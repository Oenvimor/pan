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

	v := r.Group("/api/v1/file")
	{
		v.GET("", controller.FileController.ListFile)                 // 获取文件列表
		v.POST("", controller.FileController.UploadFile)              // 上传文件
		v.DELETE("", controller.FileController.DeleteFile)            // 删除文件
		v.POST("/rapidUpload", controller.FileController.RapidUpload) // 秒传文件
	}
	r.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}
