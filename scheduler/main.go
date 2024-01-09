package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/RoyceAzure/go-stockinfo-scheduler/api"
	"github.com/RoyceAzure/go-stockinfo-scheduler/api/gapi"
	"github.com/RoyceAzure/go-stockinfo-scheduler/api/pb"
	"github.com/RoyceAzure/go-stockinfo-scheduler/cronwoeker"
	logger "github.com/RoyceAzure/go-stockinfo-scheduler/repository/logger_distributor"
	"github.com/RoyceAzure/go-stockinfo-scheduler/repository/redis"
	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	service "github.com/RoyceAzure/go-stockinfo-scheduler/service"
	"github.com/RoyceAzure/go-stockinfo-scheduler/service/redisService"
	"github.com/RoyceAzure/go-stockinfo-scheduler/util/config"
	"github.com/RoyceAzure/go-stockinfo-scheduler/worker"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hibiken/asynq"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot load config")
	}
	if config.Enviornmant == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	ctx := context.Background()

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	}
	redisClient := asynq.NewClient(redisOpt)

	loggerDis := logger.NewLoggerDistributor(redisClient)
	err = logger.SetUpLoggerDistributor(loggerDis, config.ServiceID)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}

	pgxPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	defer pgxPool.Close()
	pgxPool.Config().AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	runDBMigration(config.MigrateURL, config.DBSource)

	dao := repository.NewSQLDao(pgxPool)

	taskDistributor := worker.NewRedisTaskDistributor(redisClient)

	redisDao := jredis.NewJredis(config)

	service := service.NewService(dao, taskDistributor, redisDao)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go service.InitSDA(ctx)
	chGrpcServer := make(chan error, 1)
	chGinServer := make(chan error, 1)

	cronWorker := cronworker.NewSchdulerWorker(service)
	cronWorker.SetUpSchdulerWorker(ctx)
	defer cronWorker.StopAsync()

	jservice := redisService.NewJRedisService(redisDao)

	go runGoCron(ctx, cronWorker)
	go runGrpcServer(chGrpcServer, config, dao, service, jservice)
	go runGinServer(chGinServer, config, dao, service)

	select {
	case err = <-chGrpcServer:
		logger.Logger.Fatal().
			Err(err).
			Msg("failed to run grpc server")
	case err = <-chGinServer:
		logger.Logger.Fatal().
			Err(err).
			Msg("failed to run gin server")
	case <-ctx.Done():
		logger.Logger.Warn().Msg("Received stop signal, app will shut down after 10 second")
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		// 等待超时或者其他中断信号
		select {
		case <-timeout.Done():
			// 超时发生，程序结束
			logger.Logger.Warn().Msg("Timeout reached, shutting down.")
		case <-ctx.Done():
			// 如果在超时期间收到另一个中断信号，立即结束程序
			logger.Logger.Warn().Msg("Received another stop signal, shutting down immediately.")
		}
	}
}

/*
build Server and run
*/
func runGinServer(ch chan<- error, configs config.Config, dao repository.Dao, service service.SyncDataService) {
	server, err := api.NewServer(configs, dao, service)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot start server")
	}
	log.Info().Msgf("start gin server at %s", configs.HttpServerAddress)
	err = server.Start(configs.HttpServerAddress)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot start server")
	}
}

func runGrpcServer(ch chan<- error, configs config.Config, dao repository.Dao, service service.SyncDataService, redisService redisService.RedisService) {
	server, err := gapi.NewServer(configs, dao, service, redisService)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot start server")
	}
	/*
		使用 pb.RegisterStockInfoServer 函數註冊了先前創建的伺服器實例，使其能夠處理 StockInfoServer 接口的 RPC 請求。
	*/

	//NewServer 可以接收多個grpc.ServerOption  而上面的Interceptor 就是一個grpc.ServerOption
	/*
		gRPC 中，一個 grpc.Server 可以註冊多個服務接口。
		每個服務接口通常對應於 .proto 文件中定義的一個 service。這允許單個 gRPC 伺服器同時提供多個服務，而不需要啟動多個伺服器實例。
	*/
	grpcServer := grpc.NewServer()
	pb.RegisterStockInfoSchdulerServer(grpcServer, server)
	//reflection.Register 允許客戶端使用反射來獲知伺服器上的服務和方法。
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", configs.SCHDULER_GRPC_SERVER_ADDRESS)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot create listener")
	}
	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gRPC server")
		ch <- err
	}
}

/*
main 最好能保有組件控制權
*/
func runGoCron(ctx context.Context, cronWorker cronworker.CornWorker) {
	logger.Logger.Info().Msg("start cron worker")
	cronWorker.Start()
}

func runDBMigration(migrationURL string, dbSource string) {
	logger.Logger.Info().Msg("start db migrate")
	migrateion, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("failed to create db migrate err")
	}

	if err := migrateion.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Logger.Fatal().
			Err(err).
			Msg("failed to run db migrate err")
	}
	logger.Logger.Info().Msgf("db migrate successfully")
}