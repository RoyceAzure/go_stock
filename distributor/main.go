package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/api"
	"github.com/RoyceAzure/go-stockinfo-distributor/api/gapi"
	"github.com/RoyceAzure/go-stockinfo-distributor/api/token"
	"github.com/RoyceAzure/go-stockinfo-distributor/cronworker"
	"github.com/RoyceAzure/go-stockinfo-distributor/jkafka"
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	logger "github.com/RoyceAzure/go-stockinfo-distributor/repository/logger_distributor"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/RoyceAzure/go-stockinfo-distributor/service"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	config, err := config.LoadConfig(".") //表示讀取當前資料夾
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot load config")
	}

	//set up mongo logger
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	}
	redisClient := asynq.NewClient(redisOpt)
	loggerDis := logger.NewLoggerDistributor(redisClient)
	err = logger.SetUpLoggerDistributor(loggerDis, config.ServiceID)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	pgxPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("err create db connect")
	}
	defer pgxPool.Close()

	runDBMigration(config.MigrateFilePath, config.DBSource)

	dbDao := sqlc.NewSQLDistributorDao(pgxPool)

	remoteDao, err := remote_repo.NewJSchdulerInfoDao(config.GrpcSchedulerAddress)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("err create schduler conn")
	}

	jwriter := jkafka.NewJKafkaWriter(config.KafkaDistributorAddress)

	service := service.NewDistributorService(remoteDao, dbDao, jwriter)

	cronWorker := cronworker.NewSchdulerWorker(service, time.Local)
	cronWorker.SetUpSchdulerWorker(context.Background())
	defer cronWorker.StopAsync()

	tokenMakerGW, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("create token maker")
	}

	tokenMakerServer, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("create token maker")
	}

	go runGoCron(context.Background(), cronWorker)
	go runGRPCGatewayServer(config, dbDao, remoteDao, tokenMakerGW)
	runGRPCServer(config, dbDao, remoteDao, tokenMakerServer)
}

func runGoCron(ctx context.Context, cronWorker cronworker.CornWorker) {
	logger.Logger.Info().Msg("start cron worker")
	cronWorker.Start()
}

func runGinServer(configs config.Config, dbDao sqlc.DistributorDao, schdulerDao remote_repo.SchdulerInfoDao) {
	server := api.NewServer(dbDao, schdulerDao)

	logger.Logger.Info().Str("server start at", configs.HttpServerAddress).Msg("server start")
	err := server.Start(configs.HttpServerAddress)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot start server")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
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

func runGRPCServer(configs config.Config, dbDao sqlc.DistributorDao, schdulerDao remote_repo.SchdulerInfoDao, tokenMaker token.Maker) {
	server := gapi.NewServer(dbDao, schdulerDao, tokenMaker)

	grpcServer := grpc.NewServer()

	pb.RegisterStockInfoDistributorServer(grpcServer, server)
	//reflection.Register 允許客戶端使用反射來獲知伺服器上的服務和方法。
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", configs.GrpcServerAddress)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot create listener")
	}
	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot start gRPC server")
	}
}

func runGRPCGatewayServer(configs config.Config, dbDao sqlc.DistributorDao, schdulerDao remote_repo.SchdulerInfoDao, tokenMaker token.Maker) {
	// 創建新的gRPC伺服器
	server := gapi.NewServer(dbDao, schdulerDao, tokenMaker)

	/*runtime.WithMarshalerOption: 這是一個設置gRPC-Gateway中的Marshaller選項的函數。Marshaller是用於將數據格式從Protobuf轉換為JSON的組件。這裡指定了對於所有的MIME類型(runtime.MIMEWildcard)要使用的Marshaller配置。

	  &runtime.JSONPb: 這是gRPC-Gateway用於Protobuf消息和JSON之間轉換的標準Marshaller。這裡將它實例化並配置其選項。

	  protojson.MarshalOptions: 這是指定序列化（Protobuf到JSON的轉換）行為的選項。使用UseProtoNames設為true意味著在JSON中使用Protobuf字段的原始名稱，而不是其轉換/標準化的JSON名稱。

	  protojson.UnmarshalOptions: 這是指定反序列化（JSON到Protobuf的轉換）行為的選項。DiscardUnknown設為true表示在反序列化過程中，如果遇到JSON中有而Protobuf模型中沒有的字段，這些字段將被忽略掉，而不會導致錯誤。*/

	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: false,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// 初始化gRPC Gateway的多路復用器
	/*runtime.NewServeMux() 創建的是一個 gRPC-Gateway 的多路復用器（multiplexer），它允許你將 HTTP/JSON 請求轉換為 gRPC 請求。

	它是一個 handler 嗎？

	是的，它實現了 http.Handler 接口，因此你可以將其用作 HTTP 伺服器的主要處理器。
	它是一個 multiplexer 嗎？

	是的，它是一個特殊的多路復用器，專為將 HTTP 請求轉換為 gRPC 請求而設計。當一個 HTTP 請求到達時，這個多路復用器會根據註冊的 gRPC 路由和方法轉換該請求，然後轉發它到對應的 gRPC 伺服器方法。
	總之，runtime.NewServeMux() 既是一個 handler，也是一個 multiplexer，但它專為 grpc-gateway 設計，用於在 gRPC 伺服器和 HTTP 客戶端之間進行轉換和路由。*/
	grpcMux := runtime.NewServeMux(jsonOpt)

	// 創建一個可取消的背景上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 註冊gRPC伺服器到Gateway的多路復用器
	/*

		你在service.proto裡面定義的路由跟function, 都會在RegisterStockInfoHandlerServer 被設置，
		當呼叫RegisterStockInfoHandlerServer時，就會把路由以及handler設定到 *runtime.ServeMux上面
		RegisterStockInfoHandlerServer 會直接call grpc function (由RegisterStockInfoHandlerServer設置) , 不會經過intercepter
		RegisterStockInfoHandlerServer會把路由根handler 設置在你傳入的grpcMux 參數
	*/
	err := pb.RegisterStockInfoDistributorHandlerServer(ctx, grpcMux, server)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot register handler server")
	}

	// 創建HTTP多路復用器並將gRPC多路復用器掛載到其上
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// 在指定地址上建立監聽
	listener, err := net.Listen("tcp", configs.HttpServerAddress)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot create listener")
	}
	logger.Logger.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	//
	// 啟動HTTP伺服器
	err = http.Serve(listener, mux)
	if err != nil {
		logger.Logger.Fatal().
			Err(err).
			Msg("cannot start HTTP gateway server")
	}
}
