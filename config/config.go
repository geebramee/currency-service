package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port          uint16 `mapstructure:"port"`
	RedisUrl      string `mapstructure:"redis_url"`
	RedisPassword string `mapstructure:"redis_password"`
	// values from 0 to 100
	API1Percent uint8 `mapstructure:"api1_percent"`
	API2Percent uint8 `mapstructure:"api2_percent"`
}

func GetConfig() *Config {
	viper.AddConfigPath("config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	return &config
}
