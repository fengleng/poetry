/*
@Time : 2019/9/17 11:44
@Author : zxr
@File : metrics
@Software: GoLand
*/
package bootstrap

import "poetry/libary/metrics"

//初始化prometheus打点统计
func InitMetrics() {
	metrics.InitMetrics()
}
