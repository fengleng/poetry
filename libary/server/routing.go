/*
@Time : 2019/9/17 15:15
@Author : zxr
@File : routing
@Software: GoLand
*/
package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"poetry/app/bootstrap"
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
	mux.HandleFunc("/", CallMiddleWare(controllers.Index)) //首页
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/shiwen/ajaxshiwencont", CallMiddleWare(controllers.AjaxShiWenCont))   //ajax 根据诗词URL crc32值获取注释和译文详情html
	mux.HandleFunc("/shiwen/ajaxshiwennotes", CallMiddleWare(controllers.AjaxShiWenNotes)) // ajax 根据赏析或译文id获取注释和译文详情html
	mux.HandleFunc("/shiwen/", CallMiddleWare(controllers.ShiWenIndex))                    //诗词详情页,上线后在nginx上做的转发

	mux.HandleFunc("/dict/fancha", CallMiddleWare(controllers.FanCha)) //词典接口

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
				if bootstrap.G_Conf.ENV != "product" {
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
