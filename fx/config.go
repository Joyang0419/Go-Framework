package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
}

type Server struct {
	Port int
}

func NewServerConfig() (conf Config) {
	v := viper.New()
	v.SetConfigName("config") // 配置文件名称(无扩展名)
	v.SetConfigType("yaml")   // 如果配置文件没有扩展名，需要配置此项
	v.AddConfigPath(".")      // 查找配置文件所在的路径

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("ReadInConfig error: %v", err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}

	return conf
}
