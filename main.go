package main

import (
	"database/sql"
	"log"

	"github.com/RoyceAzure/go-stockinfo-api"
	"github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utility.LoadConfig(".") //表示讀取當前資料夾
	if err != nil {
		log.Fatal("cannot load config :", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
