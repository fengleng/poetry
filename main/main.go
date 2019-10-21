/*
@Time : 2019/9/16 18:46
@Author : zxr
@File : main
@Software: GoLand
*/
package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"poetry/app/bootstrap"
	"poetry/config/define"
	"poetry/libary/server"
	"runtime"
	"syscall"
	"time"
)
var (
	confFile string
	HttpEnv string
)

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	HttpEnv = os.Getenv("http_env")
	if HttpEnv == define.DevEnvStr{
		HttpEnv = define.DevEnvStr
	}else{
		HttpEnv = define.ProductEnvStr
	}
}

func initConfFile() {
	dir, _ := os.Getwd()
	confFile = dir + "/config/"+HttpEnv+"/config.json"
	logrus.Info("加载配置文件:", confFile)
}

func init()  {
	initEnv()
	initConfFile()
}

func main() {
	var (
		httpServer *http.Server
		serverStatus chan bool
		err error
	)
	if err = bootstrap.InitBootstrap(confFile);err!=nil{
		fmt.Println(err)
		return
	}
    //启动prometheus  metrics端口
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(bootstrap.G_Conf.MetricsPortStr, nil))
	}()
	  httpServer = server.StartHttp()
	  serverStatus = make(chan bool,1)
	//启动http服务
	go func() {
		if err = httpServer.ListenAndServe();err!=nil{
			serverStatus<-false
			logrus.Infoln("ListenAndServe error:",err)
			return
		}
		serverStatus<-true
		logrus.Infoln("pid is: ", os.Getpid())
	}()
	if status:= <-serverStatus;status==false{
		logrus.Infoln("启动服务失败......")
		return
	}
	//接收退出信号
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,syscall.SIGINT,syscall.SIGHUP,syscall.SIGTERM,syscall.SIGQUIT)
	go func() {
		<-quitChan
		//15秒之内必须退出
		ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFunc()
		if  err =  httpServer.Shutdown(ctx);err!=nil{
			logrus.Infoln("Shutdown error:",err)
		}
	}()
	return
}
