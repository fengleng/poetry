/*
@Time : 2019/9/16 18:46
@Author : zxr
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"poetry/app/bootstrap"
	"poetry/config/define"
	"poetry/libary/server"
	"runtime"
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
	var err error
	if err = bootstrap.InitBootstrap(confFile);err!=nil{
         goto PrintErr
	}
    //启动prometheus  metrics端口
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(bootstrap.G_Conf.MetricsPortStr, nil))
	}()
	//启动http端口
	if err = server.StartHttp();err!=nil{
		logrus.Debug("启动HTTP服务错误:", err)
		return
	}
	return
PrintErr:
	fmt.Println(err)
	return
}
