package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	server := http.Server{
		Handler: mux,
		Addr:    ":8000",
	}
	server2 := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// 任务1
	g.Go(func() error {
		log.Println("启动web1服务")
		return server.ListenAndServe()
	})

	// 任务2
	g.Go(func() error {
		log.Println("启动web2服务")
		return server2.ListenAndServe()
	})

	// 执行退出的协程
	g.Go(func() error {
		select {
		case <-ctx.Done():
			_ = server.Shutdown(ctx)
			log.Println("退出 web1 服务")
			_ = server2.Shutdown(ctx)
			log.Println("退出 web2 服务")
			return nil
		}
	})

	// 管理退出信号的协程
	g.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-c:
			return errors.New("exit")
		}
	})

	g.Wait()
}
