package fake

import (
	"context"
	"math/rand"
	"testing"

	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestGenerateFakeData(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, "postgres://royce:royce@localhost:5432/stock_info_scheduler?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	testDao := repository.NewSQLDao(conn)

	fakedataService, err := NewFakeSPRDataService(testDao)
	require.NoError(t, err)
	require.NotEmpty(t, fakedataService)

	fakesprs, err := fakedataService.GenerateFakeData(true)

	require.NoError(t, err)
	require.NotEmpty(t, fakesprs)

	n := 100

	var intList []int

	length := len(fakesprs)
	for i := 0; i < n; i++ {
		intList = append(intList, rand.Intn(length))
	}

	for _, index := range intList {
		require.NotZero(t, fakesprs[index].ClosingPrice)
		require.NotZero(t, fakesprs[index].OpeningPrice)
	}

}
