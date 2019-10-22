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
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"poetry/app/bootstrap"
	"strconv"
	"syscall"
	"time"
)

//启动http服务
func StartHttp() {
	var (
		mux      *http.ServeMux
		err      error
		server   *http.Server
		quitChan chan os.Signal
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
	WriteHttpPid(os.Getpid())
	ctx, cancel := context.WithCancel(context.Background())
	quitChan = make(chan os.Signal)
	go HttpSig(server, quitChan, ctx)
	//启动http服务
	if err = server.ListenAndServe(); err != nil {
		logrus.Infoln("ListenAndServe error:", err)
		cancel()
		os.Exit(1)
		return
	}
	return
}

//接收停止信号
func HttpSig(httpServer *http.Server, quitChan chan os.Signal, cont context.Context) {
	logrus.Infoln("HttpSig.......")
	if httpServer == nil {
		logrus.Infoln("server is nil")
		return
	}
	//接收退出信号
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	for {
		select {
		case <-cont.Done():
			logrus.Infoln("sig context cancel")
			return
		case signals := <-quitChan:
			logrus.Infoln("收到结束信号signals:", signals)
			//15秒之内必须退出
			ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancelFunc()
			if err := httpServer.Shutdown(ctx); err != nil {
				logrus.Infoln("Shutdown error:", err)
				return
			} else {
				logrus.Infoln("server Shutdown ok")
				return
			}
		}
	}
}

//将pid写入文件
func WriteHttpPid(pid int) {
	if pid <= 0 {
		return
	}
	pidStr := strconv.Itoa(pid)
	go func() {
		err := ioutil.WriteFile("server.pid", []byte(pidStr), 0777)
		if err != nil {
			logrus.Infoln("WriteFile error:", err)
		}
	}()
	return
}
