package domain

import (
	"fmt"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/gin-gonic/gin"
)

// MetricsMiddleware generates Prometheus metrics data about the request
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		timeStart := time.Now()

		logger.Debugf("Started %v - %v", c.Request.URL.Path, c.Request.Method)

		c.Next()

		// Update Prometheus metrics
		requestDuration := fmt.Sprintf(`http_requests_duration_seconds{path="%v", method="%v"}`,
			c.Request.URL.Path, c.Request.Method,
		)
		metrics.GetOrCreateSummary(requestDuration).UpdateDuration(timeStart)

		timeTaken := time.Now().Sub(timeStart).Milliseconds()
		logger.Debugf("Finshed %v - %v - timeTaken=%vms", c.Request.URL.Path, c.Request.Method, timeTaken)
	}
}
