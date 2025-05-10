package model

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	FileSha1 string `gorm:"column:file_sha1;type:varchar(64);not null;uniqueIndex" json:"file_sha1"` // 文件的哈希值，用于唯一标识文件，实现秒传
	FileName string `gorm:"column:file_name;type:varchar(255);not null" json:"file_name"`            // 文件名
	FileAddr string `gorm:"column:file_addr;type:varchar(512);not null" json:"file_addr"`            // 文件存储位置
	FileSize int64  `gorm:"column:file_size;not null" json:"file_size"`                              // 文件大小，单位为字节
	Status   int    `gorm:"column:status;default:1" json:"status"`                                   // 文件状态，1 表示正常，其他值可以用于标识删除或禁用
}

type UserFile struct {
	gorm.Model
	UserName string `gorm:"column:user_name;type:varchar(64);not null" json:"user_name"`  // 用户名，用于标识上传该文件的用户
	FileSha1 string `gorm:"column:file_sha1;type:varchar(64);not null" json:"file_sha1"`  // 文件的哈希值，用于唯一标识文件，实现秒传
	FileName string `gorm:"column:file_name;type:varchar(255);not null" json:"file_name"` // 文件名
	FileSize int64  `gorm:"column:file_size;not null" json:"file_size"`                   // 文件大小，单位为字节
	Status   int    `gorm:"column:status;default:1" json:"status"`                        // 文件状态，1 表示正常，其他值可以用于标识删除或禁用
}
