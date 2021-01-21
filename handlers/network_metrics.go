package handlers

import (
	"github.com/arkits/onhub-web/domain"
	"github.com/arkits/onhub-web/models"
	"github.com/gin-gonic/gin"
)

// KickOffNetworkMetricsPolling returns the applications version
func KickOffNetworkMetricsPolling(c *gin.Context) {

	domain.BeginPollingNetworkMetrics()
	c.JSON(200, models.HttpResponse{
		Message: "Started polling for Network Metrics",
	})

}

// GetNetworkMetricsHandler return the Network Metrics based on the request params
func GetNetworkMetricsHandler(c *gin.Context) {

	metricsData, err := domain.GetLastStoredNetworkMetrics()
	if err != nil {
		c.JSON(500, models.HttpErrorResponse{
			Error:       "Fatal Error in GetLastStoredNetworkMetrics",
			Description: err.Error(),
		})
		return
	}

	c.JSON(200, metricsData)

}

// GetNetworkMetricsStatusHandler returns various stats and debugging information
// about the Network Metrics related features
func GetNetworkMetricsStatusHandler(c *gin.Context) {

	networkMetricsStatus := domain.GenerateNetworkMetricsStatus()
	c.JSON(200, networkMetricsStatus)

}
