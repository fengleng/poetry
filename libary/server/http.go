/*
@Time : 2019/9/17 15:02
@Author : zxr
@File : http
@Software: GoLand
*/
package server

import (
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"poetry/app/bootstrap"
	"time"
)

//启动http服务
func StartHttp() (server *http.Server) {
	var (
		mux *http.ServeMux
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
	return
}
