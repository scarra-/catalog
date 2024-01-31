package middlewares

import (
	"time"

	"github.com/aadejanovs/catalog/database"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/extra/redisprometheus/v9"
	"go.uber.org/zap"
)

func PrometheusMiddleware() fiber.Handler {
	collector := redisprometheus.NewCollector("catalog", "prometheus", database.NewPrometheusRedis())
	prometheus.MustRegister(collector)

	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request durations for each endpoint",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	HttpRequests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)

	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(HttpRequests)

	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		err := c.Next()

		elapsed := time.Since(startTime)
		duration := elapsed.Seconds()

		requestDuration.WithLabelValues(c.Route().Path).Observe(duration)
		HttpRequests.WithLabelValues(c.Method(), c.Route().Path).Inc()

		logger := c.Locals("logger").(*zap.SugaredLogger)
		logger.Infow("Request",
			"method", c.Method(),
			"path", c.Path(),
			"duration", elapsed.Milliseconds(),
		)

		return err
	}
}
