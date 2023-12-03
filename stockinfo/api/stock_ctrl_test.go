package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/RoyceAzure/go-stockinfo/project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/utility/config"
	"github.com/stretchr/testify/require"

)

func TestInitSyncStock(t *testing.T) {
	//使用Gin  *gin.Engine建立server
	//new Server已經把所有的路由都設定好
	config, err := config.LoadConfig("../") //表示讀取當前資料夾
	require.NoError(t, err)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	require.NoError(t, err)
	store := db.NewStore(conn)
	server := newTestServer(t, store)
	//Response Recoder是做甚麼的?
	//ResponseRecorder 是一个实现了 http.ResponseWriter 接口的类型
	recoder := httptest.NewRecorder()
	url := "/stock/init"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recoder, request)
	require.Equal(t, http.StatusAccepted, recoder.Code)
}
