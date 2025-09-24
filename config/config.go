package config

import "github.com/spf13/viper"

var Config *viper.Viper

func GetConfig(key string) any {
	if Config == nil {
		v := viper.New()
		v.SetConfigFile("config.yaml")
		err := v.ReadInConfig()
		if err != nil {
			panic("读取配置失败:" + err.Error())
		}
		Config = v
	}
	return Config.Get(key)
}
