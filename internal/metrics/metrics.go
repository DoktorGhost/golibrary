package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Метрики для времени запросов
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Метрики для количества запросов
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_count",
			Help: "Number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	// Метрики для времени обращения к БД
	DBDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_duration_seconds",
			Help:    "Duration of DB requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func Init() {
	// Регистрация метрик
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(DBDuration)
}
