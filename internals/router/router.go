package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func GetRouter() *gin.Engine {
	if os.Getenv("ENV") == "release" {
		gin.SetMode(gin.ReleaseMode)
		slog.Info("Running in production mode")
	} else {
		slog.Info("Running in development mode")
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)

	setRoutes(r)
	return r
}
