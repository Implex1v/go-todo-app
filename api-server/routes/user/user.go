package user

import (
	"api-server/handler"
	"api-server/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func registerRoutes(handler *handler.HttpHandler, db *gorm.DB) {
	handler.Gin.GET("/users", func(context *gin.Context) {
		var users []types.User
		db.Find(&users)
		context.JSON(200, &users)
	})
}

var Module = fx.Options(fx.Invoke(registerRoutes))
