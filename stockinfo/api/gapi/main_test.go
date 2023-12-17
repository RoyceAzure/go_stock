package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo/api/token"
	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	utility "github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/config"
	"github.com/RoyceAzure/go-stockinfo/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

// NewServer(config utility.Config, store db.Store, taskDistributor worker.TaskDistributor)
func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor, clientFactory SchedulerClientFactory) *Server {
	config := config.Config{
		TokenSymmetricKey:    utility.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}

	server, err := NewServer(config, store, taskDistributor, clientFactory)
	require.NoError(t, err)
	return server
}

func newContextWithBearerToken(t *testing.T, tokenMaker token.Maker, userRpn string, userID int64, duartion time.Duration) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(
		userRpn,
		userID,
		duartion,
	)
	require.NoError(t, err)
	brarerToken := fmt.Sprintf("%s %s", authorizationTypeBearer, accessToken)
	md := metadata.MD{
		authorizationHeaderKey: []string{
			brarerToken,
		},
	}
	return metadata.NewIncomingContext(context.Background(), md)
}
