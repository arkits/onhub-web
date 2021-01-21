package handlers

import (
	"github.com/arkits/onhub-web/domain"
	"github.com/gin-gonic/gin"
)

// KickOffNetworkMetricsPolling returns the applications version
func KickOffNetworkMetricsPolling(c *gin.Context) {

	domain.BeginPollingNetworkMetrics()
	c.JSON(200, gin.H{
		"message": "Started Polling for Network Metrics",
	})

}

// GetNetworkMetricsHandler return the Network Metrics based on the request params
func GetNetworkMetricsHandler(c *gin.Context) {

	metricsData := domain.GetStoredNetworkMetrics()
	c.JSON(200, metricsData)

}

// GetNetworkMetricsStatusHandler returns various stats and debugging information
// about the Network Metrics related features
func GetNetworkMetricsStatusHandler(c *gin.Context) {

	networkMetricsStatus := domain.GenerateNetworkMetricsStatus()
	c.JSON(200, networkMetricsStatus)

}
