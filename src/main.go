package main

import (
	"context"
	"fmt"
	"log"
	"momonga_blog/config"
	"momonga_blog/router"
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

    mux := router.Router()
    s := NewServer(l, mux)
    return s.Run(context)
}