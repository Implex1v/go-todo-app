package routes

import (
	"api-server/db"
	"api-server/handler"
	"github.com/gin-gonic/gin"
)

func registerRoutesUser(handler *handler.HttpHandler, dao db.UserDao) {
	handler.Gin.GET("/users", func(context *gin.Context) {
		err, users := dao.GetAll()
		if err == nil {
			context.JSON(200, users)
		} else {
			// TODO log
			context.JSON(404, ApiErrorOf(err))
		}
	})
}
