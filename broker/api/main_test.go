package api

import (
	"os"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/config"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/random"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

)

func newTestServer(t *testing.T) *Server {
	config := config.Config{
		TokenSymmetricKey:    random.RandomString(32),
		AccessTokenDuration:  time.Minute,
		RefreshTokenDuration: time.Hour,
	}

	server, err := NewServer(config)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
