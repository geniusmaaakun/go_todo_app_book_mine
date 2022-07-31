package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/budougumi0617/go_todo_app/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

//責務を分離したのでコードもリファクタリングする
func run(ctx context.Context) error {
	//コンフィグ
	cfg, err := config.New()
	if err != nil {
		return err
	}
	//listener作成
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	//ハンドラ作成
	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
