package router

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	setRoutes(r)
	return r
}
