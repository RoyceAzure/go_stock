package main

import (
	"context"
	"net"
	"os"

	"github.com/RoyceAzure/go-stockinfo-schduler/api"
	"github.com/RoyceAzure/go-stockinfo-schduler/api/gapi"
	"github.com/RoyceAzure/go-stockinfo-schduler/api/pb"
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	service "github.com/RoyceAzure/go-stockinfo-schduler/service"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/RoyceAzure/go-stockinfo-schduler/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot load config")
	}
	if config.Enviornmant == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.Background()
	pgxPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	defer pgxPool.Close()
	dao := repository.NewSQLDao(pgxPool)
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	service := service.NewService(dao, taskDistributor)
	go runGrpcServer(config, dao, service)
	runGinServer(config, dao, service)
}

/*
build Server and run
*/
func runGinServer(configs config.Config, dao repository.Dao, service service.SyncDataService) {
	server, err := api.NewServer(configs, dao, service)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}
	err = server.Start(configs.HttpServerAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}
}

func runGrpcServer(configs config.Config, dao repository.Dao, service service.SyncDataService) {
	server, err := gapi.NewServer(configs, dao, service)
	if err != nil {
		log.Fatal().
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
	}
}

func runGoCron() {
	// s := gocron.NewScheduler(time.Local)

	// Every starts the job immediately and then runs at the
	// specified interval
	// job, err := s.Every(5).Seconds().Do(task)
	// if err != nil {
	// 	// handle the error related to setting up the job
	// }

	// to wait for the interval to pass before running the first job
	// use WaitForSchedule or WaitForScheduleAll
	// s.Every(5).Second().WaitForSchedule().Do(task)

	// s.WaitForScheduleAll()
	// s.Every(5).Second().Do(task) // waits for schedule
	// s.Every(5).Second().Do(task) // waits for schedule

	// // strings parse to duration
	// s.Every("5m").Do(task)

	// s.Every(5).Days().Do(task)

	// s.Every(1).Month(1, 2, 3).Do(task)

	// // set time
	// s.Every(1).Day().At("10:30").Do(task)

	// // set multiple times
	// s.Every(1).Day().At("10:30;08:00").Do(task)

	// s.Every(1).Day().At("10:30").At("08:00").Do(task)

	// // Schedule each last day of the month
	// s.Every(1).MonthLastDay().Do(task)

	// // Or each last day of every other month
	// s.Every(2).MonthLastDay().Do(task)

	// cron expressions supported
	// s.Cron("* * * * *").Do(task) // every minute

	// cron second-level expressions supported
	// s.CronWithSeconds("*/3 * 18 * * *").Do(task) // every second

	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	// s.StartAsync()
	// starts the scheduler and blocks current execution path
	// s.StartBlocking()

	// stop the running scheduler in two different ways:
	// stop the scheduler
	// s.Stop()

	// stop the scheduler and notify the `StartBlocking()` to exit
	// s.StopBlockingChan()
}
