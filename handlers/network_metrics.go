package handlers

import (
	"github.com/arkits/onhub-web/domain"
	"github.com/gin-gonic/gin"
)

// KickOffNetworkMetricsPolling returns the applications version
func KickOffNetworkMetricsPolling(c *gin.Context) {

	domain.BeginPollingNetworkMetrics()

	c.JSON(200, "Started")
}

func GetNetworkMetricsHandler(c *gin.Context) {

	metricsData := domain.GetStoredNetworkMetrics()
	c.JSON(200, metricsData)
}
