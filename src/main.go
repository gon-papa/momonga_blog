package main

import (
	"context"
	"fmt"
	"log"
	"momonga_blog/api"
	"momonga_blog/config"
	"momonga_blog/handler"
	"net"
	"os"
)

func main() {
    if err := run(context.Background()); err != nil {
		fmt.Printf("server not run: %v\n", err)
		os.Exit(1)
	}
}

func run(context context.Context) error {
    cnf, err := config.GetConfig()

    if err != nil {
        fmt.Printf("server not run: %v\n", err)
		os.Exit(1)
    }

    os.Setenv("Timezone", cnf.TimeZone)

    l, err := net.Listen("tcp", cnf.Port)
    if err != nil {
        log.Fatalf("指定ポートでのサーバーの起動に失敗しました。 port%s->error: %v", cnf.Port, err)
    }

    url := fmt.Sprintf("http://%s", l.Addr().String())
    log.Printf("start with: %v", url)

        // ハンドラーの初期化
        h := getHandler()
        srv, err := getServer(h)
        if err != nil {
            log.Fatalf("Failed to create ogen server: %v", err)
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