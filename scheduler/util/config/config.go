package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Enviornmant                  string        `mapstructure:"ENVIRONMENT"`
	DBDriver                     string        `mapstructure:"DB_DRIVER"`
	DBSource                     string        `mapstructure:"DB_SOURCE"`
	MigrateURL                   string        `mapstructure:"MIGRATE_URL"`
	HttpServerAddress            string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	SCHDULER_GRPC_SERVER_ADDRESS string        `mapstructure:"SCHDULER_GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey            string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration          time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration         time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	RedisQueueAddress            string        `mapstructure:"REDIS_Q_ADDRESS"`
	RedisSchdulerAddress         string        `mapstructure:"REDIS_SCHDULER_ADDRESS"`
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
