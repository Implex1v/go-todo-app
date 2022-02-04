package handler

import (
	"api-server/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HttpHandler struct {
	Gin *gin.Engine
}

func NewHttpHandler(l *zap.Logger) *HttpHandler {
	gin.DisableConsoleColor()
	r := gin.New()
	r.Use(logger.GinLogger(l), logger.GinRecovery(l, true))
	return &HttpHandler{Gin: r}
}

var Module = fx.Options(fx.Provide(NewHttpHandler))
