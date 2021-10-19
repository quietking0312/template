package mylog

import (
	"fmt"
	"os"
	"server/common/log"
	"server/core/config"
)

func Init() {
	logCfg := config.GetConfig().Log
	err := log.InitLog(
		log.Path(logCfg.LogPath),
		log.Level(logCfg.LogLevel),
		log.Compress(logCfg.Compress),
		log.MaxSize(logCfg.MaxSize),
		log.MaxBackups(logCfg.MaxBackups),
		log.MaxAge(logCfg.MaxAge),
		log.Format(logCfg.Format),
	)
	if err != nil {
		fmt.Printf("Initlog failed %v\n", err)
		os.Exit(1)
	}
}
