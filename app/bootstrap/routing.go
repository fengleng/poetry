/*
@Time : 2019/9/17 15:15
@Author : zxr
@File : routing
@Software: GoLand
*/
package bootstrap

import (
	"net/http"
	"poetry/libary/metrics"
	"poetry/libary/middleware"
	"time"
)

var (
	middleObj *middleware.MiddleCenter
)

//路由配置， 初始路由
func InitRouting(mux *http.ServeMux) {
	InitMiddleWare()
	mux.HandleFunc("/a", CallMiddleWare(func(writer http.ResponseWriter, request *http.Request) {

	}))
}

//初始化中间件
func InitMiddleWare() {
	middleObj = middleware.NewMiddleCenter()
	middleObj.RegisterMiddleware(middleware.NewMidMetrics())
}

//中间件处理
func CallMiddleWare(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		for _, middle := range middleObj.AllMiddle() {
			middle.Before(writer, request)
		}
		now := time.Now()
		handlerFunc(writer, request)
		metrics.G_Metrics.RequestCostInc(request.Method, request.URL.Path, time.Since(now).Seconds())
		for _, middle := range middleObj.AllMiddle() {
			middle.After(writer, request)
		}
	}
}
