package main

import (
	"context"
	"fmt"
	"momonga_blog/api"
	"momonga_blog/config"
	"momonga_blog/handler"
	"momonga_blog/pkg/logging"
	"net"
	"os"
)

func main() {
    // ログの初期化
    if err := logging.Init(); err != nil {
        fmt.Printf("Failed to open log file: %v\n", err)
        os.Exit(1)
    }

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
		os.Exit(1)
    }

    os.Setenv("Timezone", cnf.TimeZone)

    l, err := net.Listen("tcp", cnf.Port)
    if err != nil {
        logging.ErrorLogger.Error("指定ポートでのサーバーの起動に失敗しました。", "port", cnf.Port, "error", err)
        os.Exit(1)
    }

    url := fmt.Sprintf("http://%s", l.Addr().String())
    logging.AppLogger.Info("サーバーの起動に成功しました。", "url", url)

        // ハンドラーの初期化
        h := getHandler()
        srv, err := getServer(h)
        if err != nil {
            logging.ErrorLogger.Error("serverの作成に失敗しました。", "error", err)
            os.Exit(1)
        }

    s := NewServer(l, srv)
    return s.Run(context)
}

func getHandler() *handler.Handler {
    return &handler.Handler{}
}

func getServer(handler *handler.Handler) (*api.Server, error) {
    return api.NewServer(handler, nil)
}