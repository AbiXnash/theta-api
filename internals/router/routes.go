package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"server": "running",
		})
	})
}
