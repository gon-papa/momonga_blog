package main

import (
	"context"
	"fmt"
	"momonga_blog/api"
	"momonga_blog/config"
	"momonga_blog/consts"
	"momonga_blog/database"
	"momonga_blog/handler"
	"momonga_blog/internal/auth"
	"momonga_blog/internal/logging"
	"momonga_blog/middleware"
	"net"
	"net/http"
	"os"
)

func main() {
    // ログの初期化
    if err := logging.Init(); err != nil {
        fmt.Printf("Failed to open log file: %v\n", err)
        return
    }
    defer logging.Close()

    // DBの起動
    _, err := database.New()
    if err != nil {
        fmt.Printf("db not run: %v\n", err)
        logging.ErrorLogger.Error("DBの接続に失敗しました。", "error", err)
        os.Exit(1)
    }

    if err := database.HealthCheck(); err != nil {
        fmt.Printf("db not run: %v\n", err)
        logging.ErrorLogger.Error("DBのヘルスチェックに失敗しました。", "error", err)
        os.Exit(1)
    }

    logging.AppLogger.Info("DBの接続に成功しました。")

    defer database.Close()

    // サーバーの起動
    if err := run(context.Background()); err != nil {
		fmt.Printf("server not run: %v\n", err)
		os.Exit(1)
	}
}

func run(context context.Context) error {
    // 設定ファイルの読み込み
    cnf, err := config.GetConfig()
    if err != nil {
        fmt.Printf("server not run: %v\n", err)
        logging.ErrorLogger.Error("設定ファイルの読み込みに失敗しました。", "error", err)
		return err
    }

    os.Setenv("Timezone", cnf.TimeZone)

    l, err := net.Listen("tcp", cnf.Port)
    if err != nil {
        logging.ErrorLogger.Error("指定ポートでのサーバーの起動に失敗しました。", "port", cnf.Port, "error", err)
        return err
    }

    url := fmt.Sprintf("http://%s", l.Addr().String())
    logging.AppLogger.Info("サーバーの起動に成功しました。", "url", url)

    // ハンドラーの初期化
    mux := http.NewServeMux()
    if err := getDynamicHandler(mux); err != nil {
        logging.ErrorLogger.Error("ハンドラーの作成に失敗しました。", "error", err)
        return err
    }

    logging.ErrorLogger.Error(cnf.Env)
    logging.ErrorLogger.Error(consts.DevEnv)
    if cnf.Env == consts.DevEnv {
        getStaticHandler(mux)
    }
    
    addMiddlewareSrv := addMiddleware(mux)
    s := NewServer(l, addMiddlewareSrv)
    return s.Run(context)
}

func getDynamicHandler(mux *http.ServeMux) error {
    // API ハンドラーの初期化
    handler := &handler.Handler{}
    authHandler := auth.NewLoginUseCase()
    apiHandler, err := api.NewServer(handler, authHandler)
    if err != nil {
        logging.ErrorLogger.Error("API サーバーの作成に失敗しました。", "error", err)
        return err
    }
    mux.Handle(consts.ApiPathPrefix, http.StripPrefix("/api", apiHandler))

    return nil
}

func getStaticHandler(mux *http.ServeMux) {
    staticFiles := http.FileServer(http.Dir(consts.StaticFileDir))
    imageHandler := http.StripPrefix(consts.StaticFileEndpoint, staticFiles)

    mux.Handle(consts.StaticFileEndpoint, imageHandler)
}

func addMiddleware(h http.Handler) http.Handler {
    return middleware.ChainMiddleware(
        h,
        middleware.AccessLogMiddleware,
        middleware.RecoveryMiddleware,
    )
}