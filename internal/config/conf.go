package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	RunMode      string        `mapstructure:"run_mode"`
	HTTPPort     int           `mapstructure:"http_port"`
	ReadTimeOut  time.Duration `mapstructure:"read_timeout"`
	WriteTimeOut time.Duration `mapstructure:"write_timeout"`
}

type App struct {
	AppName     string `mapstructure:"app_name"`
	JwtSecret   string `mapstructure:"jwt_secret"`
	LogSavePath string `mapstructure:"log_save_path"`
	LogSaveName string `mapstructure:"log_save_name"`
	LogFileExt  string `mapstructure:"log_file_ext"`
	TimeFormat  string `mapstructure:"time_format"`
}

type Database struct {
	Type        string `mapstructure:"type"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table_prefix"`
}

type Config struct {
	Debug      bool      `mapstructure:"debug"`
	JWTExpired int       `mapstructure:"expire_jwt"`
	Server     *Server   `mapstructure:"server"`
	App        *App      `mapstructure:"app"`
	Database   *Database `mapstructure:"database"`
}

func NewConfig(configFile string) (config Config, err error) {
	now := time.Now()
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)
	return config, nil
}
