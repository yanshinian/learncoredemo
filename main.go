package main

import (
	"context"
	"fmt"
	"learncoredemo/framework"
	"learncoredemo/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    "localhost:8080",
	}
	// 这个goroutine 是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的Goroutine等待信号量
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	// 调用server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	fmt.Println("UUUUUUU")
	defer cancel()
	server.RegisterOnShutdown(func() {
		fmt.Println("我要完犊子了。。。。")
	})
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
