package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"pan/controller"
)

func SetUpRouter() {
	r := gin.Default()

	controller.InitService()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/file", controller.FileController.UploadFile) // 上传文件
		v1.DELETE("/file")
		v1.PUT("/file")
	}
	r.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}
