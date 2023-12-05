package repository

import (
	"context"
	"os"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-scheduler/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var testDao Dao
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	config, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal().Err(err).Msg("err load config")
	}
	ctx := context.Background()
	testDB, err = pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	testDao = NewSQLDao(testDB)
}
