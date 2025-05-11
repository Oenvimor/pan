package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

var Cfg = &Config{}

type Config struct {
	MysqlConfig  *Mysql  `mapstructure:"mysql"`
	ServerConfig *Server `mapstructure:"server"`
}

type Mysql struct {
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"pwd"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Server struct {
	Port int    `mapstructure:"port"`
	Path string `mapstructure:"path"`
}

func InitConfig() {
	viper.SetConfigFile("yaml/config.yaml")
	viper.AddConfigPath("./yaml/")
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("read config failed", "err", err)
		return
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		slog.Error("Unmarshal to config", "err", err)
		return
	}
}
