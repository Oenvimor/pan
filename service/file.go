package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log/slog"
	"mime/multipart"
	"os"
	"pan/dao"
	"pan/model"
	"pan/utils"
)

type IFileService interface {
	UploadFile(c *gin.Context, file *multipart.FileHeader) (*model.Resp, error)
	ListFile(c *gin.Context) ([]*model.File, error)
}

type FileService struct {
	Repository dao.IFileRepository
}

func (d *FileService) UploadFile(c *gin.Context, file *multipart.FileHeader) (*model.Resp, error) {
	// 打开文件
	f, err := file.Open()
	if err != nil {
		slog.Error("打开文件失败", "err", err)
		return nil, err
	}
	defer f.Close()
	// 计算文件哈希
	fileHash, err := utils.GenerateSHA1(f)
	if err != nil {
		slog.Error("计算文件哈希失败", "err", err)
		return nil, err
	}
	// 创建目录
	uploadPath := "./uploads"
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		slog.Error("创建目录失败", "err", err)
		return nil, err
	}
	// 检查文件是否存在
	fileInfo, err := d.Repository.GetFileBySha1(fileHash)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error("查询文件失败", "err", err)
			return nil, err
		}
	}
	if fileInfo != nil {
		slog.Info("文件已存在，无需重复上传")
		return &model.Resp{
			Msg: "文件已存在，无需重复上传",
		}, nil
	}
	// 保存文件到本地
	err = c.SaveUploadedFile(file, uploadPath+"/"+file.Filename)
	if err != nil {
		slog.Error("保存文件失败", "err", err)
		return nil, err
	}
	// 保存文件到数据库
	FileInfo := model.File{
		FileAddr: uploadPath + "/" + file.Filename,
		FileName: file.Filename,
		FileSha1: fileHash,
		FileSize: file.Size,
		Status:   1,
	}
	err = d.Repository.Save(&FileInfo)
	if err != nil {
		slog.Error("保存文件信息进数据库失败", "err", err)
		return nil, err
	}
	return nil, nil
}

func (d *FileService) ListFile(c *gin.Context) ([]*model.File, error) {
	fileInfo, err := d.Repository.List()
	if err != nil {
		slog.Error("获取文件列表失败", "err", err)
		return nil, err
	}
	return fileInfo, nil
}
