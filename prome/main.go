package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cast"
)

const (
	namespace = "QPT"
	subsystem = "GINSTUDY"
)
var metricsRequestsCost = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_cost",
		Help:      "request(ms) cost milliseconds",
	},
	[]string{"method", "path", "success", "http_code", "business_code", "cost_milliseconds", "trace_id"},
)
var metricsRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "request(ms) total",
	},
	[]string{"method", "path"},
)

func main() {
   // 定义一个Histogram类型的指标
   histogram := promauto.NewHistogram(prometheus.HistogramOpts{
      Name:    "histogram_showcase_metric",
      Buckets: []float64{5.0, 10.0, 20.0, 50.0, 100.0}, // 根据场景需求配置bucket的范围
   })

   RecordMetrics("GET","V1",true,200,300,50,"SDFASDFSDF")
   RecordMetrics("POST","V2",true,200,300,50,"SDFSDFSD")
   go func() {
      for {
         // 这里搜集一些0-100之间的随机数
         // 实际应用中，这里可以搜集系统耗时等指标
         histogram.Observe(rand.Float64() * 100.0)
         time.Sleep(1 * time.Second)
      }
   }()
   // 指标上报的路径，可以通过该路径获取实时的监控数据
   http.Handle("/metrics", promhttp.Handler())
   log.Fatal(http.ListenAndServe(":8080", nil))
}

// RecordMetrics 记录指标
func RecordMetrics(method, path string, success bool, httpCode, businessCode int, costSeconds float64, traceId string) {

	metricsRequestsCost.With(prometheus.Labels{
		"method":            method,
		"path":              path,
		"success":           cast.ToString(success),
		"http_code":         cast.ToString(httpCode),
		"business_code":     cast.ToString(businessCode),
		"cost_milliseconds": cast.ToString(costSeconds * 1000),
		"trace_id":          traceId,
	}).Observe(costSeconds)

   metricsRequestsTotal.With(prometheus.Labels{
		"method": method,
		"path":   path,
	}).Inc()

}
