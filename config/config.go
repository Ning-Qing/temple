package config

import (
	"sync"

	"github.com/Ning-Qing/temple/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var log = logger.NewLogger("config", logger.DebugLevel)

func InitConfig(path string) *Settings {
	settings := &Settings{}
	settings.lock = new(sync.RWMutex)
	// 配置读取文件
	loadLocalConfigs(path, settings)
	log.Debugf("update settings: %s", settings.String())
	return settings
}

func loadLocalConfigs(path string, settings *Settings) {
	log.Debugf("load local config file %s", path)
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("load local config file %s faild: %s", path, err.Error())
	}
	settings.lock.Lock()
	defer settings.lock.Unlock()
	if err := viper.Unmarshal(settings); err != nil {
		log.Errorf("unmarshal global settings faild: %s", err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		settings.lock.Lock()
		defer settings.lock.Unlock()
		if err := viper.Unmarshal(settings); err != nil {
			log.Errorf("unmarshal global settings faild: %s", err.Error())
		}
		log.Debugf("update settings: %s", settings.String())
	})
}
