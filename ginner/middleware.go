package ginner

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Amz-Date",
		"X-Api-Key",
		"X-Amz-Security-Token",
	}
	corsConf.AllowCredentials = true
	return cors.New(corsConf)
}
