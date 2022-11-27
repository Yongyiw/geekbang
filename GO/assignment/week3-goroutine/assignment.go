/**
* 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
* ref:
* https://blog.csdn.net/gxy_2016/article/details/116615568
* https://www.cnblogs.com/failymao/p/15677782.html
* https://github.com/Mpetrel/homework_week3
 */
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)
 
 func newAppServer() *http.Server {
	 mux := http.DefaultServeMux
	 mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		 _, _ = fmt.Fprintln(w, "App Server!")
	 })
	 return &http.Server{
		 Addr: ":8080",
		 Handler: mux,
	 }
 }
 
 func newDebugServer() *http.Server {
	 mux := http.NewServeMux()
	 mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		 _, _ = fmt.Fprintln(w, "Debug Server!")
	 })
	 return &http.Server{
		 Addr: ":8081",
		 Handler: mux,
	 }
 }
 
 func main() {
	 group, ctx := errgroup.WithContext(context.Background())
 
	 serverClose := make(chan bool)
	 // 创建两个server
	 appServer := newAppServer()
	 debugServer := newDebugServer()
 
	 // 启动http服务
	 group.Go(func() error {
		 return appServer.ListenAndServe()
	 })
	 group.Go(func() error {
		 return debugServer.ListenAndServe()
	 })
 
	 // 监听group完成事件，关闭服务
	 group.Go(func() error {
		 select {
		 case <- ctx.Done():
			 log.Printf("errgroup finished, close server..")
		 case <- serverClose:
			 log.Printf("receive custom close signal, close server...")
		 }
		 timeoutCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

		 defer cancel()
		 err := appServer.Shutdown(timeoutCtx)
		 errDebug := debugServer.Shutdown(timeoutCtx)

		 if err != nil {
			 return err
		 }

		 return errDebug
	 })
 
	 // 监听系统事件
	 group.Go(func() error {
		 c := make(chan os.Signal)
		 signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		 select {
		 case <- ctx.Done():
			 return ctx.Err()
		 case sig := <- c:
			 return errors.Errorf("receive system signal: %v", sig)
		 }
	 })

	 if err := group.Wait(); err != nil {
		 log.Fatal(err)
	 }
 }