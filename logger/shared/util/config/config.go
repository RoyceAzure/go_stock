package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

func init() {
	config, err := LoadConfig(".")
	if err != nil {
		// 处理错误，可能是记录日志，也可能是退出程序
		log.Fatalf("Failed to load configuration: %v", err)
	}
	AppConfig = &config
}

type Config struct {
	ServiceID         string `mapstructure:"SERVICEID"`
	MongodbAddress    string `mapstructure:"MONGODB_ADDRESS"`
	RedisQueueAddress string `mapstructure:"REDIS_Q_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	if AppConfig != nil {
		return *AppConfig, nil
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //JSON XML  這是指extension

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
