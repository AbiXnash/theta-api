package router

import (
	"net/http"

	"github.com/AbiXnash/theta-api/internals/components"
	"github.com/gin-gonic/gin"
)

func health(c *gin.Context, comps *components.Components) {
	if comps.Status.IsReady() {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"db":     "connected",
		})
		return
	}

	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status": "unhealthy",
		"db":     "disconnected",
	})
}
