package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"time"
)

// InitViperConfigFile 初始化配置文件
func InitViperConfigFile(configFile string, version string) error {
	tagName := "viper"
	cfgObj := viper.New()
	_, err := os.Stat(configFile)
	if err == nil {
		cfgObj.SetConfigFile(configFile)
	} else {
		cfgObj.SetConfigFile(configFile)
		cfgObj.AddConfigPath(".")
	}
	if err := cfgObj.ReadInConfig(); err != nil {
		return fmt.Errorf("read config failed: %v", err)
	}
	if err := cfgObj.Unmarshal(&coreConfig, func(c *mapstructure.DecoderConfig) {
		c.TagName = tagName
	}); err != nil {
		return fmt.Errorf("config unmarshal filed: %v", err)
	}
	cfgObj.WatchConfig()
	cfgObj.OnConfigChange(func(in fsnotify.Event) {
		if err := cfgObj.Unmarshal(&coreConfig, func(c *mapstructure.DecoderConfig) {
			c.TagName = tagName
		}); err != nil {
			fmt.Println(fmt.Errorf("config unmarshal filed: %v", err))
		}
		fmt.Printf("配置信息： %+v\n", *coreConfig)
	})
	coreConfig.Version = version
	fmt.Printf("配置信息：%+v\n", *coreConfig)
	return nil
}

func GetConfig() CoreConf {
	return *coreConfig
}

var coreConfig *CoreConf

type CoreConf struct {
	Log     Log    `viper:"log"`     // 日志配置
	Server  Server `viper:"server"`  // 服务配置
	Version string `viper:"version"` // 版本号
}

type Log struct {
	LogPath    string `viper:"log_path"`
	LogLevel   string `viper:"log_level"`
	Compress   bool   `viper:"compress"`
	MaxAge     int    `viper:"max_age"`
	MaxBackups int    `viper:"max_backups"`
	MaxSize    int    `viper:"max_size"`
	Format     string `viper:"format"`
}

type Server struct {
	Port    int    `viper:"port"`
	Mode    string `viper:"mode"`
	SqlPath string `viper:"sql_path"`
	DB      db     `viper:"db"`
}

type db struct {
	DriveName    string        `viper:"drive_name"`
	Dsn          string        `viper:"dsn"`
	MaxIdle      int           `viper:"max_idle"`
	MaxConn      int           `viper:"max_conn"`
	MaxQueryTime time.Duration `viper:"max_query_time"`
}
