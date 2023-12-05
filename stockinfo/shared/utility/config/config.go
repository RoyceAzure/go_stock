package config

import (
	"time"

	"github.com/spf13/viper"
)

// var AppConfig *Config

// func init() {
// 	config, err := LoadConfig(".")
// 	if err != nil {
// 		// 处理错误，可能是记录日志，也可能是退出程序
// 		log.Fatalf("Failed to load configuration: %v", err)
// 	}
// 	AppConfig = &config
// }

type Config struct {
	Enviornmant          string        `mapstructure:"ENVIRONMENT"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrateURL           string        `mapstructure:"MIGRATE_URL"`
	HttpStockinfoAddress string        `mapstructure:"HTTP_STOCKINFO_ADDRESS"`
	GRPCStockinfoAddress string        `mapstructure:"GRPC_STOCKINFO_ADDRESS"`
	GRPCSchedulerAddress string        `mapstructure:"GRPC_SCHEDULER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	RedisQueueAddress    string        `mapstructure:"REDIS_ADDRESS"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

/*
path: app.env所在的資料夾
*/
func LoadConfig(path string) (config Config, err error) {
	// if AppConfig != nil {
	// 	return *AppConfig, nil
	// }
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") //JSON XML  這是指extension

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	// if err == nil {
	// 	AppConfig = &config
	// }
	return
}
