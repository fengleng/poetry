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

	//搜索相关
	mux.HandleFunc("/search/shiwen/", CallMiddleWare(controllers.ShiWenSearch)) //诗词搜索页

	//诗文内容相关
	mux.HandleFunc("/shiwen/list/", CallMiddleWare(controllers.ShiWenList))                //诗词列表页-》根据分类名显示诗词列表
	mux.HandleFunc("/shiwen/ajaxshiwencont", CallMiddleWare(controllers.AjaxShiWenCont))   //ajax 根据诗词URL crc32值获取注释和译文详情html
	mux.HandleFunc("/shiwen/ajaxshiwennotes", CallMiddleWare(controllers.AjaxShiWenNotes)) // ajax 根据赏析或译文id获取注释和译文详情html
	mux.HandleFunc("/shiwen/ajaxshiwenplay", CallMiddleWare(controllers.AjaxShiWenPlay))   //ajax 根据赏析或译文id获取注释和译文的MP3文件
	mux.HandleFunc("/shiwen/", CallMiddleWare(controllers.ShiWenDetail))                   //诗词详情页,上线后在nginx上做的转发

	//作者相关
	mux.HandleFunc("/author/list/", CallMiddleWare(controllers.AuthorList))      //作者列表页
	mux.HandleFunc("/author/detail", CallMiddleWare(controllers.AuthorDetail))   //作者详情页
	mux.HandleFunc("/author/poetryList", CallMiddleWare(controllers.PoetryList)) //作者诗词列表页

	//古文相关
	mux.HandleFunc("/guwen/", CallMiddleWare(controllers.GuWenIndex))               //古文首页
	mux.HandleFunc("/guwen/bookplay", CallMiddleWare(controllers.BookPlay))         //古文播放声音
	mux.HandleFunc("/guwen/detail/", CallMiddleWare(controllers.GuWenDetail))       //古文详情页
	mux.HandleFunc("/guwen/book/", CallMiddleWare(controllers.GuWenBook))           //目录详情页
	mux.HandleFunc("/guwen/showyizhu/", CallMiddleWare(controllers.GuWenShowYizhu)) //ajax获取古文正文译注信息

	//名句相关
	mux.HandleFunc("/famous/", CallMiddleWare(controllers.FamousIndex)) //名句首页

	//词典查询
	mux.HandleFunc("/dict/fancha", CallMiddleWare(controllers.FanCha)) //词典接口

	//完善
	mux.HandleFunc("/perfect", CallMiddleWare(controllers.Correction))         //完善页面
	mux.HandleFunc("/perfect/", CallMiddleWare(controllers.Correction))        //完善页面
	mux.HandleFunc("/perfect/submit/", CallMiddleWare(controllers.CorrSubmit)) //完善页面提交处理
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
		now := time.Now()
		handlerFunc(writer, request)
		metrics.G_Metrics.RequestCostInc(request.Method, request.URL.Path, time.Since(now).Seconds())
		for _, middle := range middleObj.AllMiddle() {
			middle.After(writer, request)
		}
	}
}
