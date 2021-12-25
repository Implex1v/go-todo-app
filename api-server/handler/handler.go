package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"io"
	"os"
)

type HttpHandler struct {
	Gin *gin.Engine
}

func NewHttpHandler() *HttpHandler {
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DisableConsoleColor()
	return &HttpHandler{Gin: gin.Default()}
}

var Module = fx.Options(fx.Provide(NewHttpHandler))
