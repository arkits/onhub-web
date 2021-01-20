package handlers

import (
	"github.com/arkits/onhub-web/domain"
	"github.com/gin-gonic/gin"
)

// GetAllDevicesHandler returns all the devices
func GetAllDevicesHandler(c *gin.Context) {
	allDevices := domain.GetAllDevices()
	c.JSON(200, allDevices)
}
