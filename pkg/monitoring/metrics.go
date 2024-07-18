package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests",
		},
		[]string{"path", "method"},
	)
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests in seconds",
		},
		[]string{"path", "method"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(RequestDuration)
}

func StartMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)
}
