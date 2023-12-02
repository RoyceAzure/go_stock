package main

import (
	"context"
	"os"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/api"
	"github.com/RoyceAzure/go-stockinfo-distributor/cronworker"
	"github.com/RoyceAzure/go-stockinfo-distributor/jkafka"
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	logger "github.com/RoyceAzure/go-stockinfo-distributor/repository/logger_distributor"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/RoyceAzure/go-stockinfo-distributor/service"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hibiken/asynq"
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

	//set up mongo logger
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	}
	redisClient := asynq.NewClient(redisOpt)
	loggerDis := logger.NewLoggerDistributor(redisClient)
	err = logger.SetUpLoggerDistributor(loggerDis)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
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

	jwriter := jkafka.NewJKafkaWriter(config.KafkaDistributorAddress)

	service := service.NewDistributorService(remoteDao, dbDao, jwriter)

	cronWorker := cronworker.NewSchdulerWorker(service, time.Local)
	cronWorker.SetUpSchdulerWorker(context.Background())
	defer cronWorker.StopAsync()

	go runGoCron(context.Background(), cronWorker)
	runGinServer(config, dbDao, remoteDao)
}

func runGoCron(ctx context.Context, cronWorker cronworker.CornWorker) {
	log.Info().Msg("start cron worker")
	cronWorker.Start()
}

func runGinServer(configs config.Config, dbDao sqlc.DistributorDao, schdulerDao remote_repo.SchdulerInfoDao) {
	server := api.NewServer(dbDao, schdulerDao)

	log.Info().Str("server start at", configs.HttpServerAddress).Msg("server start")
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
