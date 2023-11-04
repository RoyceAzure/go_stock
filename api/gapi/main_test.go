package gapi

import (
	"testing"
	"time"

	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	worker "github.com/RoyceAzure/go-stockinfo-worker"
	"github.com/stretchr/testify/require"
)

// NewServer(config utility.Config, store db.Store, taskDistributor worker.TaskDistributor)
func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utility.Config{
		TokenSymmetricKey:    utility.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)
	return server
}
