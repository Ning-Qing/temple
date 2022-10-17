package config

import (
	"github.com/Ning-Qing/temple/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var log = logger.NewLogger("config", logger.DebugLevel)

var (
	DefaultHost = ""
	DefaultPort = 8080
	DefaultMode = "debug"
)

func InitConfig(path string) {
	settings := &Settings{}
	// 配置读取文件
	loadLocalConfigs(path, settings)
	log.Debugf("update settings: %s", settings.String())
	GlobalSettings = settings

}

func loadLocalConfigs(path string, settings *Settings) {
	log.Debugf("load local config file %s", path)
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("load local config file %s faild: %s", path, err.Error())
	}
	if err := viper.Unmarshal(settings); err != nil {
		log.Errorf("unmarshal global settings faild: %s", err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(settings); err != nil {
			log.Errorf("unmarshal global settings faild: %s", err.Error())
		}
		log.Debugf("update settings: %s", settings.String())
	})
}
