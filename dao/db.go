package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log/slog"
	"pan/model"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pwd"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
	))
	if err != nil {
		slog.Error("open mysql failed", "err", err)
		return
	}
	DB = db
	db.SingularTable(true)
	Migration()
}

func GetDB() *gorm.DB {
	if DB == nil {
		Init()
	}
	return DB
}

func Migration() {
	DB.AutoMigrate(&model.File{})
	DB.AutoMigrate(&model.UserFile{})
}

func Close() {
	DB.Close()
}
