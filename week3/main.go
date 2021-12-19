package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, errctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	//注册一个http路由
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("new server"))
		fmt.Fprintf(w, "new server")
	})

	// windows下通过路由地址模拟单个服务退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	//goroutine1 注册服务器
	g.Go(func() error {
		return server.ListenAndServe()
	})
	//goroutine2 监听服务器退出
	g.Go(func() error {
		select {
		case <-errctx.Done():
			fmt.Println("====errgroup exit====")
		case <-serverOut:
			fmt.Println("====server will out====")
		}
		// 为了模拟windows下的优雅退出
		// 收到服务器错误退出的信号后设置context超时处理
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		//关闭服务器
		return server.Shutdown(ctx)
	})
	//goroutine3
	g.Go(func() error {
		signalCh := make(chan os.Signal)
		signal.Notify(signalCh,
			syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

		select {
		case <-errctx.Done():
			fmt.Println("====errctx cancel====")
			return errctx.Err()
		case s := <-signalCh:
			// 收到监听的信号后设置退出
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			fmt.Printf("get os signal :%v \n", s)
			//关闭服务器
			return errors.Errorf("get os signal: %v ,get errors: %v", s, server.Shutdown(ctx))
		}
	})
	// 阻塞在这里，等待所有的goroutine退出,returns the first non-nil error (if any) from them.
	if err := g.Wait(); err != nil {
		fmt.Printf("====all goroutine are dead, get errors: %+v \n", err)
	}
}
