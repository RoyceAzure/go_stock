package api

import (
	"database/sql"
	"os"
	"testing"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/rs/zerolog/log"
)

var testQueries *repository.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal().Err(err).Msg("err load config")
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	testQueries = repository.New(testDB)
}
