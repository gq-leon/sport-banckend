package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		App      *App      `mapstructure:"app"`
		Database *Database `mapstructure:"database"`
		Server   *Server   `mapstructure:"server"`
	}

	App struct {
		Name string `mapstructure:"name"`
		Env  string `mapstructure:"env"`
	}

	Token struct{}

	Redis struct{}

	Database struct {
		Type     string `mapstructure:"type"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	}

	Server struct {
		Port int    `mapstructure:"PORT"`
		Host string `mapstructure:"HOST"`
	}
)

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml") // 设置为读取 YAML 文件
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	err := v.Unmarshal(&config)
	return &config, err
}
