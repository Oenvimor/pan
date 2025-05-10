package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"pan/common/response"
	"pan/service"
)

type FileControllerType struct {
	Service service.IFileService
}

var FileController FileControllerType

func (s *FileControllerType) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("上传文件失败", "err", err)
		response.InternalServerError(c, "上传文件失败")
		return
	}
	resp, err := s.Service.UploadFile(c, file)
	if err != nil {
		slog.Error("文件处理出错", "err", err)
		response.InternalServerError(c, "文件处理出错")
		return
	}
	if resp != nil {
		response.OnlyMsg(c, resp.Msg)
		return
	}
	response.Ok(c, nil)
}

func (s *FileControllerType) ListFile(c *gin.Context) {
	resp, err := s.Service.ListFile(c)
	if err != nil {
		slog.Error("获取文件失败", "err", err)
		response.InternalServerError(c, "获取文件失败")
		return
	}
	response.Ok(c, resp)
}
