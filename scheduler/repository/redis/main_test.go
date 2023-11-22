package jredis

import (
	"context"
	"os"
	"testing"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var testDao repository.Dao

func NewTestSevice() {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal().Err(err).Msg("err load config")
	}
	ctx := context.Background()
	pgxPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	testDao = repository.NewSQLDao(pgxPool)
}

func TestMain(m *testing.M) {
	NewTestSevice()
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}