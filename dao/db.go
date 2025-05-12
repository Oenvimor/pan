package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log/slog"
	"pan/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pwd"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
	)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		slog.Error("open mysql failed", "err", err)
		return nil
	}
	DB = db
	Migration()
	return DB
}

func Migration() {
	DB.AutoMigrate(&model.File{})
	DB.AutoMigrate(&model.UserFile{})
	slog.Info("表迁移成功")
}
