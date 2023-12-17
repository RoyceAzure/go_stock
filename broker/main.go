package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/RoyceAzure/go-stockinfo-broker/api"
	"github.com/RoyceAzure/go-stockinfo-broker/api/gapi"
	"github.com/RoyceAzure/go-stockinfo-broker/api/token"
	_ "github.com/RoyceAzure/go-stockinfo-broker/doc/statik"
	logger "github.com/RoyceAzure/go-stockinfo-broker/repository/remote_dao/logger_distributor"
	scheduler_dao "github.com/RoyceAzure/go-stockinfo-broker/repository/remote_dao/schedulerDao"
	stockinfo_dao "github.com/RoyceAzure/go-stockinfo-broker/repository/remote_dao/stockinfoDao"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/config"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}
	// conn, err := amqp.Dial(config.RabbitMQAddress)
	// if err != nil {
	// 	log.Fatal().
	// 		Err(err).
	// 		Msg("Failed to connect to RabbitMQ")
	// }
	// defer conn.Close()

	// e := rabbitmq.NewRabbitMqEmmiter(conn)
	// log.Print(e.EmmitEvent("1", "1").Error())

	// server, err := api.NewServer()
	// if err != nil {
	// 	log.Fatal().
	// 		Err(err).
	// 		Msg("Failed to connect to RabbitMQ")
	// }
	// server.Start(config.HttpServerAddress)

	stockinfoDao, err := stockinfo_dao.NewStockInfoDao(config.GrpcStockinfoAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot connect stockinfo grpc server")
	}

	schedulerDao, err := scheduler_dao.NewSchedulerDao(config.GrpcSchedulerAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot connect scheduler grpc server")
	}

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	}

	//set up mongo logger
	redisClient := asynq.NewClient(redisOpt)
	loggerDis := logger.NewLoggerDistributor(redisClient)
	err = logger.SetUpLoggerDistributor(loggerDis, config.ServiceID)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}

	go runGRPCGatewayServer(config, schedulerDao, stockinfoDao, config.TokenSymmetricKey)
	runGRPCServer(config, schedulerDao, stockinfoDao, config.TokenSymmetricKey)
}

func runGRPCGatewayServer(configs config.Config, schedulerDao scheduler_dao.ISchedulerDao, stockinfoDao stockinfo_dao.IStockInfoDao, symmerickey string) {

	tokenMakerStockinfo, err := token.NewPasetoMaker(symmerickey)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}

	tokenMakerScheduler, err := token.NewPasetoMaker(symmerickey)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}

	authorizeStockinfo := api.NewAuthorizor(tokenMakerStockinfo)

	authorizeScheduler := api.NewAuthorizor(tokenMakerScheduler)

	// 創建新的gRPC伺服器

	serverStockinfo, err := gapi.NewStockInfoServer(stockinfoDao, authorizeStockinfo)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}

	serverScheduler, err := gapi.NewSchedulerServer(schedulerDao, authorizeScheduler)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}

	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
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
	grpcMux := runtime.NewServeMux(jsonOpt, runtime.WithMetadata(gapi.CustomMatcher))

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
	err = pb.RegisterStockInfoHandlerServer(ctx, grpcMux, serverStockinfo)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot register stockinfo handler server")
	}

	err = pb.RegisterStockInfoSchdulerHandlerServer(ctx, grpcMux, serverScheduler)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot register scheduler handler server")
	}

	// 創建HTTP多路復用器並將gRPC多路復用器掛載到其上
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	/*
			如果路由結尾有斜線 (/)，它將匹配任何以該前綴開始的 URL。因此，/swagger/ 將匹配 /swagger/、/swagger/file1.html、/swagger/subdir/file2.html 等。
		如果路由沒有結尾斜線，它只會匹配該具體路徑。
		所以，如果你使用 http.Handle("swagger/")（注意，缺少前置的斜線）：

		它將不會如預期地工作，因為在 net/http 中，路由通常應該以斜線 (/) 開始。這可能會導致未定義的行為或不匹配任何路徑。
		正確的做法是：

		使用 http.Handle("/swagger/") 以匹配所有以 /swagger/ 開始的路徑。
		使用 http.Handle("/swagger")（沒有結尾的斜線）只匹配 /swagger 這一具體的路徑，不匹配 /swagger/abc 或其他子路徑。
		總之，要確保路由以斜線 (/) 開始，並根據你的需求決定是否在結尾添加斜線。*/

	// fs := http.FileServer(http.Dir("./doc/swagger"))

	//這個FileSystem內容是zip content data, 剛好跟statik使用的依樣

	//在statik.go init()裡面已經註冊了使用statik編譯好的data
	//這裡New就是把他轉成fileSystem, 然後再搭配http.FileServer 即可
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("can't create statik fs err")
	}

	//http.StripPrefix 會回傳handler
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	// 在指定地址上建立監聽
	listener, err := net.Listen("tcp", configs.HttpServerAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot create listener")
	}
	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())

	//
	loggerHandler := gapi.HttpLogger(mux)
	handler := gapi.IdMiddleWareHandler(loggerHandler)
	// handler1 := gapi.IdMiddleWareHandler(mux)
	// handler := gapi.HttpLogger(handler1)

	// 啟動HTTP伺服器
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start HTTP gateway server")
	}
}

func runGRPCServer(configs config.Config, schedulerDao scheduler_dao.ISchedulerDao, stockinfoDao stockinfo_dao.IStockInfoDao, symmerickey string) {

	tokenMakerStockinfo, err := token.NewPasetoMaker(symmerickey)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}

	tokenMakerScheduler, err := token.NewPasetoMaker(symmerickey)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}

	authorizeStockinfo := api.NewAuthorizor(tokenMakerStockinfo)

	authorizeScheduler := api.NewAuthorizor(tokenMakerScheduler)

	// 創建新的gRPC伺服器

	serverStockinfo, err := gapi.NewStockInfoServer(stockinfoDao, authorizeStockinfo)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}

	serverScheduler, err := gapi.NewSchedulerServer(schedulerDao, authorizeScheduler)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start gateway http server")
	}
	/*
		使用 pb.RegisterStockInfoServer 函數註冊了先前創建的伺服器實例，使其能夠處理 StockInfoServer 接口的 RPC 請求。
	*/

	unaryInterceptor := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(gapi.IdMiddleWare, gapi.GrpcLogger))

	//NewServer 可以接收多個grpc.ServerOption  而上面的Interceptor 就是一個grpc.ServerOption
	grpcServer := grpc.NewServer(unaryInterceptor)
	/*
		gRPC 中，一個 grpc.Server 可以註冊多個服務接口。
		每個服務接口通常對應於 .proto 文件中定義的一個 service。這允許單個 gRPC 伺服器同時提供多個服務，而不需要啟動多個伺服器實例。
	*/

	pb.RegisterStockInfoServer(grpcServer, serverStockinfo)
	pb.RegisterStockInfoSchdulerServer(grpcServer, serverScheduler)
	//reflection.Register 允許客戶端使用反射來獲知伺服器上的服務和方法。
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", configs.GrpcServerAddress)
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
