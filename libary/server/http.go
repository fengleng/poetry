/*
@Time : 2019/9/17 15:02
@Author : zxr
@File : http
@Software: GoLand
*/
package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"poetry/app/bootstrap"
	"strings"
	"syscall"
	"time"
)

//启动http服务
func StartHttp() {
	var (
		mux          *http.ServeMux
		err          error
		server       *http.Server
		serverStatus chan bool
		quitChan     chan os.Signal
	)
	logrus.Info("开始启动HTTP服务,监听端口是:", bootstrap.G_Conf.HttpPortStr)
	mux = http.NewServeMux()
	InitRouting(mux)
	w := logrus.New().Writer()
	defer w.Close()
	server = &http.Server{
		Addr:         bootstrap.G_Conf.HttpPortStr,
		Handler:      mux,
		ReadTimeout:  time.Duration(bootstrap.G_Conf.HttpReadTimeOut) * time.Millisecond,
		WriteTimeout: time.Duration(bootstrap.G_Conf.HttpWriteTimeOut) * time.Millisecond,
		ErrorLog:     log.New(w, "poetry", 0),
	}
	serverStatus = make(chan bool, 1)
	//启动http服务
	go func() {
		if err = server.ListenAndServe(); err != nil {
			serverStatus <- false
			logrus.Infoln("ListenAndServe error:", err)
			return
		}
		serverStatus <- true
		logrus.Infoln("pid is: ", os.Getpid())
	}()
	if status := <-serverStatus; status == false {
		logrus.Infoln("启动服务失败......")
		return
	}
	quitChan = make(chan os.Signal)
	HttpSig(server, quitChan)
	HttpFlagParams(quitChan)
	return
}

//接收信号
func HttpSig(httpServer *http.Server, quitChan chan os.Signal) {
	//接收退出信号
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-quitChan
		if httpServer == nil {
			logrus.Infoln("server is nil")
			return
		}
		//15秒之内必须退出
		ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFunc()
		if err := httpServer.Shutdown(ctx); err != nil {
			logrus.Infoln("Shutdown error:", err)
		}
	}()
	return
}

//如果命令行参数有stop，则停止服务
func HttpFlagParams(quitChan chan os.Signal) {
	args := os.Args
	if len(args) > 1 {
		sig := strings.ToLower(os.Args[1])
		if sig == "stop" {
			quitChan <- syscall.SIGINT
		}
	}
	return
}
