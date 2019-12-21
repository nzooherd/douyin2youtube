package config

import (
	"github.com/spf13/viper"
	"runtime"
	"time"
)

var configBaseName, configType = "config_base", "toml"
var configPath, configName = "", ""

var BaseConfig *viper.Viper
var SpecifyConfig *viper.Viper


func init() {
	var platform = runtime.GOOS
	if platform == "darwin"{
		configPath = "/Users/shiwenhao/Documents/Project/My/douyin2youtube/src/project/static"
		configName = "config_darwin"
	}else if platform == "linux"{
		configPath = "/home/shiwenhao/Documents//Project/My/douyin2youtube/src/project/static"
		configName = "config_linux"
	}

	BaseConfig = viper.New()
	BaseConfig.SetConfigName(configBaseName)
	BaseConfig.AddConfigPath(configPath)
	BaseConfig.SetConfigType(configType)
	if err := BaseConfig.ReadInConfig(); err != nil {
		panic(err)
	}

	SpecifyConfig = viper.New()
	SpecifyConfig.SetConfigName(configName)
	SpecifyConfig.AddConfigPath(configPath)
	SpecifyConfig.SetConfigType(configType)

	if err := SpecifyConfig.ReadInConfig(); err != nil {
		panic(err)
	}
}


func GetString(key string) string{
	res := SpecifyConfig.GetString(key)
	if res != ""{
		return res
	}
	return BaseConfig.GetString(key)
}

func GetInt(key string) int{
	res := SpecifyConfig.GetInt(key)
	if res != 0{
		return res
	}
	return BaseConfig.GetInt(key)
}

func GetDuration(key string) time.Duration{
	res := SpecifyConfig.GetDuration(key)
	if res != 0{
		return res
	}
	return BaseConfig.GetDuration(key)
}
