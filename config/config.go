package config

import (
	// "github.com/fsnotify/fsnotify"
	"fmt"

	"github.com/spf13/viper"
)

const DEFAULT_CONFIG_PATH = "./config/config.yml"

func Init(filePath string) {
	if filePath == "" {
		filePath = DEFAULT_CONFIG_PATH
	}
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load config Fail: %s\n", err))
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
