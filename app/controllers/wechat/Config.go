package wechat

import (
	"auth/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetConfig(c *gin.Context) {
	code := c.DefaultQuery("code","")
	if code == ""{
		code = c.DefaultPostForm("code","")
	}
	if code == "" {
		c.JSON(http.StatusOK,gin.H{
			"state":3001,
			"message":"缺失参数",
		})
		c.Abort()
		return
	}

	result,err := services.FindWechatConfigOne(code)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":4001,
			"message":err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"state":2000,
		"message":"success",
		"data":result,
	})
}
