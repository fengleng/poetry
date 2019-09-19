/*
@Time : 2019/9/17 14:29
@Author : zxr
@File : config
@Software: GoLand
*/
package bootstrap

import (
	"github.com/json-iterator/go"
	"io/ioutil"
	"poetry/tools"
)

type Config struct {
	MetricsPortStr   string `json:"MetricsPortStr"`
	HttpPortStr      string `json:"HttpPortStr"`
	HttpReadTimeOut  int    `json:"httpReadTimeOut"`
	HttpWriteTimeOut int    `json:"httpWriteTimeOut"`
	DataSource       string `json:"dataSource"`
}

var (
	G_Json jsoniter.API
	G_Conf *Config
)

func InitConfig(confFile string) (err error) {
	var (
		conf    Config
		content []byte
	)
	G_Json = jsoniter.ConfigCompatibleWithStandardLibrary
	if ret, err := tools.FileExists(confFile); err != nil || ret == false {
		goto Return
	}
	if content, err = ioutil.ReadFile(confFile); err != nil {
		goto Return
	}
	if err = G_Json.Unmarshal(content, &conf); err != nil {
		goto Return
	}
	G_Conf = &conf
	return
Return:
	return
}
