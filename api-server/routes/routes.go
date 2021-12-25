package routes

import (
	"api-server/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func registerRoutes(handler *handler.HttpHandler) {
	handler.Gin.GET("/health", func(context *gin.Context) {
		context.JSONP(200, gin.H{"message": "healthy"})
	})
}

var Module = fx.Options(
	fx.Invoke(registerRoutes, registerRoutesUser),
)
