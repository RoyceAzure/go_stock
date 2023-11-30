package repository

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/random"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testClient FrontendClient
var testUserId int64

func TestMain(m *testing.M) {
	setUP()
	os.Exit(m.Run())
}

func setUP() {
	config, _ := config.LoadConfig("../../../")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, _ := pgxpool.New(ctx, config.DBSource)

	dao := NewSQLDistributorDao(conn)

	ip := random.RandomIP()
	region := "tw"
	// CreatedAt := time.Now().UTC()
	testClient, _ = dao.CreateFrontendClient(ctx, CreateFrontendClientParams{
		Ip:     ip,
		Region: region,
	})

	testUserId = int64(rand.Intn(256))
}
