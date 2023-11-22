// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.0
// source: service_stockinfo_schduler.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StockInfoSchduler_GetStockDayAvg_FullMethodName        = "/pb.StockInfoSchduler/GetStockDayAvg"
	StockInfoSchduler_GetStockPriceRealTime_FullMethodName = "/pb.StockInfoSchduler/GetStockPriceRealTime"
)

// StockInfoSchdulerClient is the client API for StockInfoSchduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockInfoSchdulerClient interface {
	GetStockDayAvg(ctx context.Context, in *StockDayAvgRequest, opts ...grpc.CallOption) (*StockDayAvgResponse, error)
	GetStockPriceRealTime(ctx context.Context, in *StockPriceRealTimeRequest, opts ...grpc.CallOption) (*StockPriceRealTimeResponse, error)
}

type stockInfoSchdulerClient struct {
	cc grpc.ClientConnInterface
}

func NewStockInfoSchdulerClient(cc grpc.ClientConnInterface) StockInfoSchdulerClient {
	return &stockInfoSchdulerClient{cc}
}

func (c *stockInfoSchdulerClient) GetStockDayAvg(ctx context.Context, in *StockDayAvgRequest, opts ...grpc.CallOption) (*StockDayAvgResponse, error) {
	out := new(StockDayAvgResponse)
	err := c.cc.Invoke(ctx, StockInfoSchduler_GetStockDayAvg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoSchdulerClient) GetStockPriceRealTime(ctx context.Context, in *StockPriceRealTimeRequest, opts ...grpc.CallOption) (*StockPriceRealTimeResponse, error) {
	out := new(StockPriceRealTimeResponse)
	err := c.cc.Invoke(ctx, StockInfoSchduler_GetStockPriceRealTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockInfoSchdulerServer is the server API for StockInfoSchduler service.
// All implementations must embed UnimplementedStockInfoSchdulerServer
// for forward compatibility
type StockInfoSchdulerServer interface {
	GetStockDayAvg(context.Context, *StockDayAvgRequest) (*StockDayAvgResponse, error)
	GetStockPriceRealTime(context.Context, *StockPriceRealTimeRequest) (*StockPriceRealTimeResponse, error)
	mustEmbedUnimplementedStockInfoSchdulerServer()
}

// UnimplementedStockInfoSchdulerServer must be embedded to have forward compatible implementations.
type UnimplementedStockInfoSchdulerServer struct {
}

func (UnimplementedStockInfoSchdulerServer) GetStockDayAvg(context.Context, *StockDayAvgRequest) (*StockDayAvgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockDayAvg not implemented")
}
func (UnimplementedStockInfoSchdulerServer) GetStockPriceRealTime(context.Context, *StockPriceRealTimeRequest) (*StockPriceRealTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockPriceRealTime not implemented")
}
func (UnimplementedStockInfoSchdulerServer) mustEmbedUnimplementedStockInfoSchdulerServer() {}

// UnsafeStockInfoSchdulerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockInfoSchdulerServer will
// result in compilation errors.
type UnsafeStockInfoSchdulerServer interface {
	mustEmbedUnimplementedStockInfoSchdulerServer()
}

func RegisterStockInfoSchdulerServer(s grpc.ServiceRegistrar, srv StockInfoSchdulerServer) {
	s.RegisterService(&StockInfoSchduler_ServiceDesc, srv)
}

func _StockInfoSchduler_GetStockDayAvg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockDayAvgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoSchdulerServer).GetStockDayAvg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfoSchduler_GetStockDayAvg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoSchdulerServer).GetStockDayAvg(ctx, req.(*StockDayAvgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfoSchduler_GetStockPriceRealTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockPriceRealTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoSchdulerServer).GetStockPriceRealTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfoSchduler_GetStockPriceRealTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoSchdulerServer).GetStockPriceRealTime(ctx, req.(*StockPriceRealTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockInfoSchduler_ServiceDesc is the grpc.ServiceDesc for StockInfoSchduler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockInfoSchduler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StockInfoSchduler",
	HandlerType: (*StockInfoSchdulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStockDayAvg",
			Handler:    _StockInfoSchduler_GetStockDayAvg_Handler,
		},
		{
			MethodName: "GetStockPriceRealTime",
			Handler:    _StockInfoSchduler_GetStockPriceRealTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_stockinfo_schduler.proto",
}
