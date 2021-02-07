package handlers

import (
	"github.com/arkits/onhub-web/models"
	"github.com/gin-gonic/gin"
)

// ReturnHTTPError is a wrapper function to return a generic HTTP error response
func ReturnHTTPError(c *gin.Context, httpStatusCode int, errorMessage string, errorDescription string) {
	c.JSON(500, models.HTTPErrorResponse{
		Error:       errorMessage,
		Description: errorDescription,
	})
	return
}
