// metrics.go
package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	requestsTotal *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "api_requests_total",
				Help: "Total number of API requests",
			},
			[]string{"method"},
		),
	}
}

func (m *Metrics) Register() {
	prometheus.MustRegister(m.requestsTotal)
}

func (m *Metrics) IncrementRequest(method string) {
	m.requestsTotal.WithLabelValues(method).Inc()
}

func (m *Metrics) Handler() http.Handler {
	return promhttp.Handler()
}
