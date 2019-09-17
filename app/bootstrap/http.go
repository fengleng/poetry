/*
@Time : 2019/9/17 15:02
@Author : zxr
@File : http
@Software: GoLand
*/
package bootstrap

import (
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

//启动http服务
func StartHttp() (err error) {
	var (
		mux    *http.ServeMux
		server *http.Server
	)
	logrus.Info("开始启动HTTP服务,监听端口是:", G_Conf.HttpPortStr)
	mux = http.NewServeMux()
	InitRouting(mux)
	w := logrus.New().Writer()
	defer w.Close()
	server = &http.Server{
		Addr:         G_Conf.HttpPortStr,
		Handler:      mux,
		ReadTimeout:  time.Duration(G_Conf.HttpReadTimeOut) * time.Millisecond,
		WriteTimeout: time.Duration(G_Conf.HttpWriteTimeOut) * time.Millisecond,
		ErrorLog:     log.New(w, "poetry", 0),
	}
	err = server.ListenAndServe()
	return
}
