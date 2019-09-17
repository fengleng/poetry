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
	"poetry/config"
	"runtime"
)
var (
	confFile string
	HttpEnv string
)

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	HttpEnv = os.Getenv("http_env")
	if HttpEnv == config.DevEnvStr{
		HttpEnv = config.DevEnvStr
	}else{
		HttpEnv = config.ProductEnvStr
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
	//加载日志配置
	bootstrap.InitLogger()
	//加载 prometheus
	bootstrap.InitMetrics()
	//加载 配置文件
	if err = bootstrap.InitConfig(confFile);err!=nil{
         goto PrintErr
	}
    //启动prometheus  metrics端口
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(bootstrap.G_Conf.MetricsPortStr, nil))
	}()
	//启动http端口
	if err = bootstrap.StartHttp();err!=nil{
		logrus.Debug("启动HTTP服务错误:", err)
		return
	}
	return
PrintErr:
	fmt.Println(err)
	return
}
