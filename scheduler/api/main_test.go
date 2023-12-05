package api

import (
	"testing"

	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-scheduler/service"
	"github.com/RoyceAzure/go-stockinfo-scheduler/util/config"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, config config.Config, dao repository.Dao, service service.SyncDataService) *Server {
	server, err := NewServer(config, dao, service)
	require.NoError(t, err)
	return server
}
