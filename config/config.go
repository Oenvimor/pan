package config

import (
	"github.com/spf13/viper"
	"log/slog"
)

var Cfg = &Config{}

type Config struct {
	MysqlConfig  *Mysql  `mapstructure:"mysql"`
	ServerConfig *Server `mapstructure:"server"`
	RedisConfig  *Redis  `mapstructure:"redis"`
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

type Redis struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Pwd         string `mapstructure:"pwd"`
	Db          int    `mapstructure:"db"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxActive   int    `mapstructure:"max_active"`
	IdleTimeout int    `mapstructure:"idle_timeout"`
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
