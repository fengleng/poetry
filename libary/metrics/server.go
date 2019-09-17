/*
@Time : 2019/9/16 19:42
@Author : zxr
@File : server
@Software: GoLand
*/
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var G_Metrics *ServerMetrics

type ServerMetrics struct {
	httpRequestTotal     *prometheus.CounterVec //请求总数
	httpRequestCodeTotal *prometheus.CounterVec //请求错误总数
	httpRequestCost      *prometheus.SummaryVec //请求耗时
	httpRequestInFlight  prometheus.Gauge       //当前正在处理的请求的数量
}

func NewServerMetrics() *ServerMetrics {
	return &ServerMetrics{
		httpRequestTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_request_total",
				Help: "Number of hello requests in total",
			}, []string{"method", "endpoint"},
		),
		httpRequestCodeTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_code_total",
				Help: "Number of hello requests code in total",
			}, []string{"method", "endpoint", "code", "msg"},
		),
		httpRequestCost: promauto.NewSummaryVec(
			prometheus.SummaryOpts{
				Name:       "http_request_cost",
				Help:       "request time ",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			}, []string{"method", "endpoint"},
		),
		httpRequestInFlight: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "http_request_in_flight",
				Help: "Current number of http requests in flight",
			},
		),
	}
}

//初始化prometheus打点统计
func InitMetrics() {
	G_Metrics = NewServerMetrics()
}

func (s *ServerMetrics) RequestInFlightInc() {
	s.httpRequestInFlight.Inc()
}

func (s *ServerMetrics) RequestInFlightDec() {
	s.httpRequestInFlight.Dec()
}

//请求总数自增
func (s *ServerMetrics) RequestTotalInc(method, endpoint string) {
	s.httpRequestTotal.WithLabelValues(method, endpoint).Inc()
}

//请求错误总数
func (s *ServerMetrics) RequestCodeTotalInc(method, endpoint string, code string, err error) {
	s.httpRequestCodeTotal.WithLabelValues(method, endpoint, code, err.Error()).Inc()
}

//请求耗时
func (s *ServerMetrics) RequestCostInc(method, endpoint string, us float64) {
	s.httpRequestCost.WithLabelValues(method, endpoint).Observe(us)
}
