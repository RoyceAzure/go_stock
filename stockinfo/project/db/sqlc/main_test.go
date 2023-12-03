package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	_ "os"
	"testing"

	"github.com/RoyceAzure/go-stockinfo/shared/utility/config"
	_ "github.com/lib/pq"

)

// database/sql  只是個介面  所以測試時需要額外使用github.com/lib/pq
// 引入後就有實作了
var testQueries *Queries
var testDB *sql.DB

var testFund Fund
var testStock Stock
var testUser User
var testUserStock UserStock

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	// 你的初始化代碼

	config, err := config.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	testUser, testFund, testStock, testUserStock = CreateRandomUserStockNoTest()
}

func teardown() {
	// 你的清理代碼
	ctx := context.Background()
	testQueries.DeleteUserStock(ctx, testUserStock.UserStockID)
	testQueries.DeleteStock(ctx, testStock.StockID)
	testQueries.DeleteFund(ctx, testFund.FundID)
	testQueries.DeleteUser(ctx, testUser.UserID)
}

func TestYourFunction(t *testing.T) {
	// 你的測試代碼
}
