package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// InitViperConfigFile 初始化配置文件
func InitViperConfigFile(configFile string) error {
	_, err := os.Stat(configFile)
	if err == nil {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigFile(configFile)
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config failed: %v", err)
	}
	viper.WatchConfig()
	if err := viper.Unmarshal(&CoreConfig); err != nil {
		return fmt.Errorf("config unmarshal filed: %v", err)
	}
	return nil
}

var CoreConfig *coreConf

type coreConf struct {
	Log     Log    // 日志配置
	Server  Server // 服务配置
	Version string // 版本号
}

type Log struct {
	LogPath    string
	LogLevel   string
	Compress   bool
	MaxAge     int
	MaxBackups int
	MaxSize    int
	Format     string
}

type Server struct {
	Port    int
	Mode    string
	SqlPath string
}
