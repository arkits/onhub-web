package handlers

import (
	"strconv"

	"github.com/arkits/onhub-web/domain"
	"github.com/arkits/onhub-web/models"
	"github.com/gin-gonic/gin"
)

// KickOffNetworkMetricsPolling returns the applications version
func KickOffNetworkMetricsPolling(c *gin.Context) {

	domain.BeginPollingNetworkMetrics()
	c.JSON(200, models.HTTPResponse{
		Message: "Started polling for Network Metrics",
	})

}

// GetNetworkMetricsHandler return the Network Metrics based on the request params
func GetNetworkMetricsHandler(c *gin.Context) {

	limitStr := c.DefaultQuery("limit", "30")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ReturnHTTPError(c, 400, "Unable to parse limit", err.Error())
		return
	}

	skipStr := c.DefaultQuery("skip", "0")
	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		ReturnHTTPError(c, 400, "Unable to parse skip", err.Error())
		return
	}

	metricsData, err := domain.GetStoredNetworkMetrics(limit, skip)
	if err != nil {
		ReturnHTTPError(c, 500, "Internal Error in GetStoredNetworkMetrics", err.Error())
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
