package inits

import (
	"github.com/spf13/viper"
	"init/config"
	"log"
)

func InitViper() {
	viper.SetConfigFile("../dev.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&config.Config)
	log.Println("配置文件获取成功", config.Config)
}
