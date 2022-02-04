package routes

import (
	"api-server/db"
	"api-server/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func registerRoutesUser(handler *handler.HttpHandler, dao db.UserDao, logger *zap.SugaredLogger) {
	handler.Gin.GET("/users", func(context *gin.Context) {
		err, users := dao.GetAll()
		if err == nil {
			context.JSON(200, users)
		} else {
			context.JSON(404, ApiErrorOf(err))
		}
	})
}
