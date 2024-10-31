package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
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

	// Метрики для времени выполнения запросов к внешним API
	ExternalAPIDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "external_api_duration_seconds",
			Help:    "Duration of external API requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"api_name", "endpoint"},
	)

	// Метрика для количества запросов к внешним API
	ExternalAPICount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "external_api_request_count",
			Help: "Number of requests to external APIs",
		},
		[]string{"api_name", "endpoint"},
	)
)

func Init() {
	// Регистрация метрик
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(DBDuration)
	prometheus.MustRegister(ExternalAPIDuration)
	prometheus.MustRegister(ExternalAPICount)
}

// TrackExternalAPIDuration - функция для отслеживания времени выполнения запроса к внешним API
func TrackExternalAPIDuration(apiName, endpoint string, duration float64) {
	ExternalAPIDuration.WithLabelValues(apiName, endpoint).Observe(duration)
	ExternalAPICount.WithLabelValues(apiName, endpoint).Inc()
}

// TrackDBDuration - функция для отслеживания времени выполнения запроса к БД
func TrackDBDuration(method string, duration float64) {
	DBDuration.WithLabelValues(method).Observe(duration)
}

// RequestMetricsMiddleware - middleware для измерения времени выполнения запросов
func RequestMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Запоминаем время начала

		// Выполняем следующий хендлер
		next.ServeHTTP(w, r)

		// Вычисляем продолжительность запроса
		duration := time.Since(start).Seconds()

		// Регистрируем метрику
		RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)

		// Увеличиваем счетчик запросов
		RequestCount.WithLabelValues(r.Method, r.URL.Path).Inc()
	})
}
