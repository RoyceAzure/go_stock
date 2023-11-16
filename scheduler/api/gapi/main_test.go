package gapi

import (
	"testing"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/service"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/stretchr/testify/require"
)

/*
可以在這裡設定test所需的參數
*/
func newTestServer(t *testing.T, config config.Config, dao repository.Dao, service service.SyncDataService) *Server {
	server, err := NewServer(config, dao, service)
	require.NoError(t, err)
	return server
}
