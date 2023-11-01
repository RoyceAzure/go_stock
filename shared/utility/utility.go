package utility

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// StringToSqlNiStr 將str轉換成sql.NullInt64
func StringToSqlNiStr(str string) sql.NullString {
	return sql.NullString{String: str, Valid: true}
}

// TimeToSqlNiTime 將Time轉換成sql.NullTime
func TimeToSqlNiTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func IntToSqlNiInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

func IntToSqlNiString(i int64) sql.NullString {
	str := strconv.FormatInt(i, 10)
	return sql.NullString{String: str, Valid: true}
}

// sql.Null 就有一個valid 欄位  所以可以根據這個欄位判斷是否轉換成功  不需要額外回傳error
// StringToSqlNiInt64 將str轉換成sql.NullInt64
// 零值字串或轉換失敗將會回傳 Valid為false的sql.NullInt64
func StringToSqlNiInt64(str string) sql.NullInt64 {
	if str == "" {
		return sql.NullInt64{Valid: false}
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: i, Valid: true}
}

// StringToSqlNiTime 將str轉換成sql.NullTime
// str參數是timestamp格式，需要是毫秒級的timestamp
// 零值字串或轉換失敗將會回傳 Valid為false的sql.NullTime
func StringToSqlNiTime(str string) sql.NullTime {
	if str == "" {
		return sql.NullTime{Valid: false}
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return sql.NullTime{Valid: false}
	}
	t := time.Unix(i, 0)
	return sql.NullTime{Time: t, Valid: true}
}

type Config struct {
	Enviornmant          string        `mapstructure:"ENVIRONMENT"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrateURL           string        `mapstructure:"MIGRATE_URL"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
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
