package gapi

import (
	"context"
	"net/http"
	"strings"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
	// logger "github.com/RoyceAzure/go-stockinfo-broker/repository/remote_dao/logger_distributor"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func IdMiddleWare(ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {

	reqID := uuid.New().String()
	ctx = context.WithValue(ctx, util.RequestIDKey, reqID)
	return handler(ctx, req)
}

func IdMiddleWareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// log := logger.Logger.Info()
		// log.Msg("in IdMiddleWareHandler")
		reqID := uuid.New().String()

		// 將 UUID 添加到 HTTP header 中
		req.Header.Set(string(util.RequestIDKey), reqID)
		handler.ServeHTTP(res, req)
	})
}

func CustomMatcher(ctx context.Context, req *http.Request) metadata.MD {
	// 創建一個空的 metadata.MD 對象
	md := metadata.MD{}

	// 從 HTTP 請求中提取 header 並添加到 metadata.MD
	if val := req.Header.Get(string(util.RequestIDKey)); val != "" {
		md[strings.ToLower(string(util.RequestIDKey))] = []string{val}
	}
	return md
}
