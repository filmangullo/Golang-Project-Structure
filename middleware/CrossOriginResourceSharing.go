package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AllowCORS() gin.HandlerFunc {
	cfg := cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	cfg.AllowAllOrigins = true
	return cors.New(cfg)
}
