package router

import (
	"net/http"

	"github.com/AbiXnash/theta-api/internals/components"
	"github.com/gin-gonic/gin"
)

func setRoutes(r *gin.Engine, c *components.Components) {
	r.GET("/health", func(ctx *gin.Context) { health(ctx, c) })
	r.POST("/login", func(ctx *gin.Context) { userLogin(ctx, c) })
}

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

func userLogin(c *gin.Context, _ *components.Components) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	c.JSON(http.StatusAccepted, gin.H{
		"username": username,
		"password": password,
	})
}
