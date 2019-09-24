/*
@Time : 2019/9/17 15:15
@Author : zxr
@File : routing
@Software: GoLand
*/
package bootstrap

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"poetry/app/controllers"
	"poetry/libary/metrics"
	"poetry/libary/middleware"
	"poetry/tools"
	"strings"
	"time"
)

var (
	middleObj *middleware.MiddleCenter
)

//路由配置， 初始路由
func InitRouting(mux *http.ServeMux) {
	InitMiddleWare()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", CallMiddleWare(controllers.Index))
}

//初始化中间件
func InitMiddleWare() {
	middleObj = middleware.NewMiddleCenter()
	middleObj.RegisterMiddleware(middleware.NewMidMetrics())
}

//中间件处理
func CallMiddleWare(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				tools.WriteRecover(err)
				if G_Conf.ENV != "product" {
					logrus.Infoln(err)
				}
			}
		}()
		for _, middle := range middleObj.AllMiddle() {
			middle.Before(writer, request)
		}
		if request.RequestURI == "/favicon.ico" {
			return
		}
		if strings.Contains(request.RequestURI, ".") {
			return
		}
		logrus.Infoln("url:", request.RequestURI)
		now := time.Now()
		handlerFunc(writer, request)
		metrics.G_Metrics.RequestCostInc(request.Method, request.URL.Path, time.Since(now).Seconds())
		for _, middle := range middleObj.AllMiddle() {
			middle.After(writer, request)
		}
	}
}
