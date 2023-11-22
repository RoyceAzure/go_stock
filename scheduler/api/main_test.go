package api

import (
	"context"
	"os"
	"testing"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/service"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

var testDao repository.Dao
var testDB *pgxpool.Pool
var testServer *Server

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal().Err(err).Msg("err load config")
	}
	ctx := context.Background()
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	testDB, err = pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	testDao = repository.NewSQLDao(testDB)
}

func NewTestServer(t *testing.T, config config.Config, dao repository.Dao, service service.SyncDataService) *Server {
	server, err := NewServer(config, dao, service)
	require.NoError(t, err)
	return server
}
