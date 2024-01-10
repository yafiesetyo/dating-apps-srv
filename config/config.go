package config

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg Config

func Init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to initialize config, error: %v \n", err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal config, error: %v \n", err)
	}
}
