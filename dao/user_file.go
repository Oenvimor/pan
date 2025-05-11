package dao

import "pan/model"

type IUserFileRepository interface {
	Save(userFileInfo *model.UserFile) error
	GetUserFileByHash(fileHash string) (*model.UserFile, error)
	List() ([]*model.UserFile, error)
	Delete(fileSHA1 string) error
}

type UserFileRepository struct{}

func (d *UserFileRepository) Save(userFileInfo *model.UserFile) error {
	return DB.Save(&userFileInfo).Error
}

func (d *UserFileRepository) GetUserFileByHash(fileHash string) (*model.UserFile, error) {
	var userFile model.UserFile
	return &userFile, DB.Where("file_sha1 = ?", fileHash).First(&userFile).Error
}

func (d *UserFileRepository) List() ([]*model.UserFile, error) {
	var userFileInfo []*model.UserFile
	return userFileInfo, DB.Find(&userFileInfo).Error
}

func (d *UserFileRepository) Delete(fileSHA1 string) error {
	return DB.Where("file_sha1 = ?", fileSHA1).Delete(&model.UserFile{}).Error
}
