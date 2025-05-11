package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log/slog"
	"mime/multipart"
	"os"
	"pan/dao"
	"pan/model"
	"pan/utils"
)

type IFileService interface {
	UploadFile(c *gin.Context, file *multipart.FileHeader, fileHash string) error
	ListFile() ([]*model.File, error)
	DeleteFile(fileSHA1 string) error
	RapidUpload(file *multipart.FileHeader) (*model.Resp, error)
}

type FileService struct {
	FileRepository     dao.IFileRepository
	UserFileRepository dao.IUserFileRepository
}

func (d *FileService) UploadFile(c *gin.Context, file *multipart.FileHeader, fileHash string) error {
	// 创建目录
	uploadPath := viper.GetString("server.path")
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		slog.Error("创建目录失败", "err", err)
		return err
	}
	// 保存文件到本地
	err = c.SaveUploadedFile(file, uploadPath+"/"+file.Filename)
	if err != nil {
		slog.Error("保存文件失败", "err", err)
		return err
	}
	// 保存文件到数据库
	FileInfo := model.File{
		FileAddr: uploadPath + "/" + file.Filename,
		FileName: file.Filename,
		FileSha1: fileHash,
		FileSize: file.Size,
	}
	err = d.FileRepository.Save(&FileInfo)
	if err != nil {
		slog.Error("保存文件信息进数据库失败", "err", err)
		return err
	}
	userFileInfo := model.UserFile{
		FileName: FileInfo.FileName,
		FileSha1: FileInfo.FileSha1,
		FileSize: FileInfo.FileSize,
		UserName: "zhangcheng", //todo:从token中获取
	}
	err = d.UserFileRepository.Save(&userFileInfo)
	if err != nil {
		slog.Error("保存用户文件信息进数据库失败", "err", err)
		return err
	}
	return nil
}

func (d *FileService) ListFile() ([]*model.File, error) {
	return d.FileRepository.List()
}

func (d *FileService) DeleteFile(fileSHA1 string) error {
	return d.FileRepository.Delete(fileSHA1)
}

func (d *FileService) RapidUpload(file *multipart.FileHeader) (*model.Resp, error) {
	fileHASH, File, err := utils.GenerateSHA1(file)
	if err != nil {
		return nil, err
	}
	defer File.Close()
	userFile, err := d.UserFileRepository.GetUserFileByHash(fileHASH) // todo:从token获取id
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error("获取用户文件信息失败", "err", err)
			return nil, err
		}
	}
	if userFile.FileSha1 != "" {
		return &model.Resp{
			Msg: model.FileExist,
		}, nil
	}
	fileInfo, err := d.FileRepository.GetFileBySha1(fileHASH)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if fileInfo.FileSha1 == "" {
		// 秒传失败返回哈希值
		return &model.Resp{
			Msg:  model.FailRapidUpload,
			Data: fileHASH,
		}, nil
	}
	// 触发秒传
	var req = model.UserFile{
		FileName: fileInfo.FileName,
		FileSha1: fileInfo.FileSha1,
		FileSize: fileInfo.FileSize,
		UserName: "zhangcheng", //todo:从token中获取
	}
	err = d.UserFileRepository.Save(&req)
	return &model.Resp{
		Msg: model.SuccessRapidUpload,
	}, err
}
