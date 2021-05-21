package routers

import (
	"auth/app/controllers/user"
	"auth/app/controllers/wechat"
	"auth/app/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r :=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"state":2000,
			"message":"success",
		})
	})

	r.Any("/auth/login", user.Login)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JWT())
	{
		apiv1.Any("/wechat/config",wechat.GetConfig)
		apiv1.GET("/user/list", user.GetUsers)
	}

	return r
}
