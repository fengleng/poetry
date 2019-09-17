/*
@Time : 2019/9/17 15:18
@Author : zxr
@File : metrics
@Software: GoLand
*/
package middleware

import (
	"net/http"
	"poetry/libary/metrics"
	"time"
)

type midMetrics struct {
	Now time.Time
}

func NewMidMetrics() *midMetrics {
	return &midMetrics{}
}

//prometheus中间件
func (m *midMetrics) Before(writer http.ResponseWriter, request *http.Request) {
	metrics.G_Metrics.RequestTotalInc(request.Method, request.URL.Path)
	metrics.G_Metrics.RequestInFlightInc()
}

func (m *midMetrics) Run(writer http.ResponseWriter, request *http.Request) {

}

func (m *midMetrics) After(writer http.ResponseWriter, request *http.Request) {
	metrics.G_Metrics.RequestInFlightDec()
}
