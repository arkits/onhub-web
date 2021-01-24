package domain

import (
	"fmt"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/gin-gonic/gin"
)

var ignoredHTTPEndpoints []string = []string{"/ohw/api/metrics"}

// MetricsMiddleware generates Prometheus metrics data about the request
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		timeStart := time.Now()

		_, requestToIgnoredEndpoint := Find(ignoredHTTPEndpoints, c.Request.URL.Path)

		if !requestToIgnoredEndpoint {
			logger.Debugf("Started %v - %v", c.Request.URL.Path, c.Request.Method)
		}

		c.Next()

		// Update Prometheus metrics
		requestDuration := fmt.Sprintf(`http_requests_duration_seconds{path="%v", method="%v", status="%v"}`,
			c.Request.URL.Path, c.Request.Method, c.Writer.Status(),
		)
		metrics.GetOrCreateSummary(requestDuration).UpdateDuration(timeStart)

		timeTaken := time.Now().Sub(timeStart).Milliseconds()

		if !requestToIgnoredEndpoint {
			logger.Debugf("Finshed %v - %v - timeTaken=%vms", c.Request.URL.Path, c.Request.Method, timeTaken)
		}

	}
}
