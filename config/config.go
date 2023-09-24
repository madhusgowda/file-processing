package config

import (
	"github.com/spf13/viper"
	"log"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type AppConfig struct {
	Server ServerConfig
	Redis  RedisConfig
}

func LoadConfig() *AppConfig {
	var configuration *AppConfig
	var configName string
	configName = "default" // single config file
	viper.SetConfigName(configName)
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err.Error())
	}
	err := viper.MergeInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = viper.UnmarshalExact(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return configuration
}
