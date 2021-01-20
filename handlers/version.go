package handlers

import (
	"github.com/arkits/onhub-web/domain"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// VersionHandler returns the applications version
func VersionHandler(c *gin.Context) {

	version := domain.Version{
		Name:    viper.GetString("server.name"),
		Version: viper.GetString("server.version"),
	}

	c.JSON(200, version)
}
