package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init function initialises the yaml config
// Can be run with different environments
func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("./")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot load the configuration file ", env)
	}
}

// GetConfig gets the currently loaded config
func GetConfig() *viper.Viper {
	return config
}
