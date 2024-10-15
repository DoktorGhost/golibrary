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
)

func Init() {
	// Регистрация метрик
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(DBDuration)
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
