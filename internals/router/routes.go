package router

import (
	"github.com/AbiXnash/theta-api/internals/components"
	"github.com/AbiXnash/theta-api/internals/router/auth"
	"github.com/gin-gonic/gin"
)

func setRoutes(r *gin.Engine, c *components.Components) {
	healthRoutes(r, c)
	authRoutes(r, c)
}

func healthRoutes(r *gin.Engine, c *components.Components) {
	r.GET("/health", func(ctx *gin.Context) { health(ctx, c) })
}

func authRoutes(r *gin.Engine, c *components.Components) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", func(ctx *gin.Context) { auth.UserLogin(ctx, c) })
		authGroup.POST("/register", func(ctx *gin.Context) { auth.UserRegister(ctx, c) })
	}
}
