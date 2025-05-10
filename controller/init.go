package controller

import (
	"pan/dao"
	"pan/service"
)

func InitService() {
	FileController = FileControllerType{
		Service: &service.FileService{
			Repository: &dao.FileRepository{},
		},
	}
}
