package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	api "github.com/RoyceAzure/go-stockinfo-api"
	"github.com/RoyceAzure/go-stockinfo-api/gapi"
	"github.com/RoyceAzure/go-stockinfo-api/pb"
	_ "github.com/RoyceAzure/go-stockinfo-doc/statik"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	worker "github.com/RoyceAzure/go-stockinfo-worker"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := utility.LoadConfig(".") //表示讀取當前資料夾
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot load config")
	}
	if config.Enviornmant == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot connect to db:")
	}
	runDBMigration(config.MigrateURL, config.DBSource)
	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	//因為qsynq.client 是concurrent
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(redisOpt, store)
	go runGRPCGatewayServer(config, store, taskDistributor)
	runGRPCServer(config, store, taskDistributor)
}

func runGinServer(config utility.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}

	err = server.Start(config.HttpServerAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}
}

func runGRPCServer(config utility.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}
	/*
		使用 pb.RegisterStockInfoServer 函數註冊了先前創建的伺服器實例，使其能夠處理 StockInfoServer 接口的 RPC 請求。
	*/

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	//NewServer 可以接收多個grpc.ServerOption  而上面的Interceptor 就是一個grpc.ServerOption
	grpcServer := grpc.NewServer(grpcLogger)
	/*
		gRPC 中，一個 grpc.Server 可以註冊多個服務接口。
		每個服務接口通常對應於 .proto 文件中定義的一個 service。這允許單個 gRPC 伺服器同時提供多個服務，而不需要啟動多個伺服器實例。
	*/

	pb.RegisterStockInfoServer(grpcServer, server)
	//reflection.Register 允許客戶端使用反射來獲知伺服器上的服務和方法。
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
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

// runGRPCGatewayServer 啟動gRPC Gateway伺服器。此伺服器提供了一個HTTP接口，允許通過HTTP與gRPC服務進行交互。
func runGRPCGatewayServer(config utility.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	// 創建新的gRPC伺服器
	server, err := gapi.NewServer(config, store, taskDistributor)
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
	err = pb.RegisterStockInfoHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot register handler server")
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
	listener, err := net.Listen("tcp", config.HttpServerAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot create listener")
	}
	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())

	//
	handler := gapi.HttpLogger(mux)

	// 啟動HTTP伺服器
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start HTTP gateway server")
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

func runTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
