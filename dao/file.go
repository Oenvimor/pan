package dao

import "pan/model"

type IFileRepository interface {
	Save(fileInfo *model.File) error
	GetFileBySha1(fileSha1 string) (*model.File, error)
	List() ([]*model.File, error)
	Delete(fileSHA1 string) error
}

type FileRepository struct{}

func (d *FileRepository) Save(fileInfo *model.File) error {
	return DB.Save(&fileInfo).Error
}

func (d *FileRepository) GetFileBySha1(fileSha1 string) (*model.File, error) {
	var file model.File
	return &file, DB.Where("file_sha1 = ?", fileSha1).First(&file).Error
}

func (d *FileRepository) List() ([]*model.File, error) {
	var fileInfo []*model.File
	return fileInfo, DB.Find(&fileInfo).Error
}

func (d *FileRepository) Delete(fileSHA1 string) error {
	return DB.Where("file_sha1 = ?", fileSHA1).Delete(&model.File{}).Error
}
