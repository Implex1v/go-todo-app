package routes

import (
	"api-server/db"
	"api-server/handler"
	"api-server/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func registerRoutesUser(handler *handler.HttpHandler, dao db.UserDao, logger *zap.SugaredLogger) {
	handler.Gin.GET("/users", func(context *gin.Context) {
		err, users := dao.GetAll()
		if err == nil {
			context.JSON(200, users)
		} else {
			context.JSON(http.StatusInternalServerError, ApiErrorOf(err))
		}
	})

	handler.Gin.GET("/users/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, NewApiError("id is invalid"))
			return
		}

		err, users := dao.Get(id)
		if err == nil {
			context.JSON(http.StatusOK, users)
		} else {
			context.JSON(http.StatusNotFound, ApiErrorOf(err))
		}
	})

	handler.Gin.POST("/users/", func(context *gin.Context) {
		var user types.User
		if err := context.BindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, NewApiError("User payload is invalid"))
			return
		}

		err, createdUser := dao.Create(&user)
		if err == nil {
			context.JSON(http.StatusCreated, createdUser)
		} else {
			context.JSON(http.StatusInternalServerError, NewApiError("Internal server error"))
		}
	})

	handler.Gin.PUT("/users/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, NewApiError("id is invalid"))
			return
		}

		var user types.User
		if err := context.BindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, NewApiError("User payload is invalid"))
			return
		}
		user.ID = uint(id)

		err, createdUser := dao.Update(&user)
		if err == nil {
			context.JSON(http.StatusOK, createdUser)
		} else {
			context.JSON(http.StatusInternalServerError, NewApiError("Internal server error"))
		}
	})
}
