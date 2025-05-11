package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"pan/common/response"
	"pan/model"
	"pan/service"
)

type FileControllerType struct {
	Service service.IFileService
}

var FileController FileControllerType

func (s *FileControllerType) UploadFile(c *gin.Context) {
	fileHash := c.PostForm("file_sha1")
	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("上传文件失败", "err", err)
		response.InternalServerError(c, "上传文件失败")
		return
	}
	err = s.Service.UploadFile(c, file, fileHash)
	if err != nil {
		slog.Error("文件处理出错", "err", err)
		response.InternalServerError(c, "文件处理出错")
		return
	}
	response.Ok(c, nil)
}

func (s *FileControllerType) ListFile(c *gin.Context) {
	fileInfo, err := s.Service.ListFile()
	if err != nil {
		slog.Error("获取文件失败", "err", err)
		response.InternalServerError(c, "获取文件失败")
		return
	}
	response.Ok(c, fileInfo)
}

func (s *FileControllerType) DeleteFile(c *gin.Context) {
	var qry = struct {
		FileSha1 string `form:"file_sha1" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&qry); err != nil {
		slog.Error("参数错误", "err", err)
		response.BadRequest(c, "参数错误")
		return
	}
	if err := s.Service.DeleteFile(qry.FileSha1); err != nil {
		slog.Error("删除文件失败", "err", err)
		response.InternalServerError(c, "删除文件失败")
		return
	}
	response.Ok(c, nil)
}

func (s *FileControllerType) RapidUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("接收文件失败", "err", err)
		response.BadRequest(c, "上传文件失败")
		return
	}
	resp, err := s.Service.RapidUpload(file)
	if err != nil {
		slog.Error("文件处理出错", "err", err)
		response.InternalServerError(c, "文件处理出错")
		return
	}
	if resp.Msg == model.FileExist {
		response.OnlyMsg(c, resp.Msg)
		return
	}
	if resp.Msg == model.FailRapidUpload {
		response.Data(c, resp.Msg, resp.Data)
		return
	}
	response.Ok(c, nil)
}
