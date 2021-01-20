package handlers

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/gin-gonic/gin"
)

// MetricsHandler returns the applications version
func MetricsHandler(c *gin.Context) {

	metrics.WritePrometheus(c.Writer, true)

}
