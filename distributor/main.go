package main

import (
	"context"
	"os"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/api"
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.LoadConfig(".") //表示讀取當前資料夾
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot load config")
	}
	if config.Enviornmant == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	pgxPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	defer pgxPool.Close()

	runDBMigration(config.MigrateFilePath, config.DBSource)

	dbDao := sqlc.NewSQLDistributorDao(pgxPool)
	remoteDao, err := remote_repo.NewJSchdulerInfoDao(config.GrpcSchedulerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("err create schduler conn")
	}

	runGinServer(config, dbDao, remoteDao)
}

func runGinServer(configs config.Config, dbDao sqlc.DistributorDao, schdulerDao remote_repo.SchdulerInfoDao) {
	server := api.NewServer(dbDao, schdulerDao)

	err := server.Start(configs.HttpServerAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migrateion, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to create db migrate err")
	}

	if err := migrateion.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().
			Err(err).
			Msg("failed to run db migrate err")
	}
	log.Info().Msgf("db migrate successfully")
}
