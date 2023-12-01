package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongodbAddress    string `mapstructure:"MONGODB_ADDRESS"`
	RedisQueueAddress string `mapstructure:"REDIS_Q_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
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
