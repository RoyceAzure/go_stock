package config

import (
	"time"

	"github.com/spf13/viper"
)

var AppConfig *Config

func init() {
	config, err := LoadConfig(".")
	if err == nil {
		AppConfig = &config
	}
}

type Config struct {
	ServiceID               string        `mapstructure:"SERVICEID"`
	Enviornmant             string        `mapstructure:"ENVIRONMENT"`
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSource                string        `mapstructure:"DB_SOURCE"`
	HttpServerAddress       string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress       string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	GrpcSchedulerAddress    string        `mapstructure:"GRPC_SCHEDULER_ADDRESS"`
	GrpcStockinfoAddress    string        `mapstructure:"GRPC_STOCKINFO_ADDRESS"`
	GrpcDistributorAddress  string        `mapstructure:"GRPC_DISTRIBUTOR_ADDRESS"`
	KafkaDistributorAddress string        `mapstructure:"KAFKA_DISTRIBUTOR_ADDRESS"`
	MigrateFilePath         string        `mapstructure:"MIGRATE_FILE_PATH"`
	RedisQueueAddress       string        `mapstructure:"REDIS_Q_ADDRESS"`
	RabbitMQAddress         string        `mapstructure:"RABBITMQ_ADDRESS"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration    time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
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
	if err == nil {
		AppConfig = &config
	}
	return
}
