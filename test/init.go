/*
@Time : 2019/9/20 17:27
@Author : zxr
@File : init
@Software: GoLand
*/
package test

import (
	"github.com/sirupsen/logrus"
	"os"
	"poetry/app/bootstrap"
	"strings"
)

func init() {
	dir, _ := os.Getwd()
	dir = strings.TrimRight(dir, "/test")
	confFile := dir + "/config/product/config.json"
	if err := bootstrap.InitBootstrap(confFile); err != nil {
		logrus.Infoln("InitBootstrap err:", err)
	}
}
