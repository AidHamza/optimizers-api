package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Service appConfig
	Messaging nsqConfig
	Storage s3Config
}

type s3Config struct {
	Host string
	Port int
	AccessKey string
	SecretKey string
	TLS bool
}

type nsqConfig struct {
	Host string
	Port int
	Topic string
}

type appConfig struct {
	Port string
	CrossOrigin string
}

var App Config

func (config *Config) Init() error {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	config.Service.Port = viper.GetString("Service.Port")
	config.Service.CrossOrigin = viper.GetString("Service.CrossOrigin")

	config.Messaging.Host = viper.GetString("Messaging.Host")
	config.Messaging.Port = viper.GetInt("Messaging.Port")
	config.Messaging.Topic = viper.GetString("Messaging.Topic")

	config.Storage.Host = viper.GetString("Storage.Host")
	config.Storage.Port = viper.GetInt("Storage.Port")
	config.Storage.AccessKey = viper.GetString("Storage.AccessKey")
	config.Storage.SecretKey = viper.GetString("Storage.SecretKey")
	config.Storage.TLS = viper.GetBool("Storage.TLS")

	return nil
}