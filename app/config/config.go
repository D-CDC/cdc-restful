package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	v *viper.Viper
}

var appConfig AppConfig

func LoadConfig(name string) {
	v := viper.New()

	v.SetConfigName(name)
	v.SetConfigType("toml")
	v.AddConfigPath("./conf/")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	appConfig.v = v
}

func Get(key string) interface{} {
	return appConfig.v.Get(key)
}

func GetString(key string) string {
	return appConfig.v.GetString(key)
}

func GetInt(key string) int {
	return appConfig.v.GetInt(key)
}
