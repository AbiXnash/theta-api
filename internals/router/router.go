package router

import (
	"os"

	"github.com/AbiXnash/theta-api/internals/components"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func setUpEnv() {
	if os.Getenv("ENV") == "release" {
		gin.SetMode(gin.ReleaseMode)
		slog.Info("Running in production mode")
	} else {
		slog.Info("Running in development mode")
	}
}

func GetRouter(c *components.Components) *gin.Engine {
	setUpEnv()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.SetTrustedProxies(nil)

	setRoutes(r, c)
	return r
}
