// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.0
// source: service_stockinfo.proto

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
	StockInfo_CreateUser_FullMethodName              = "/pb.StockInfo/CreateUser"
	StockInfo_GetUser_FullMethodName                 = "/pb.StockInfo/GetUser"
	StockInfo_UpdateUser_FullMethodName              = "/pb.StockInfo/UpdateUser"
	StockInfo_LoginUser_FullMethodName               = "/pb.StockInfo/LoginUser"
	StockInfo_VerifyEmail_FullMethodName             = "/pb.StockInfo/VerifyEmail"
	StockInfo_InitStock_FullMethodName               = "/pb.StockInfo/InitStock"
	StockInfo_GetUnRealizedGain_FullMethodName       = "/pb.StockInfo/GetUnRealizedGain"
	StockInfo_GetRealizedGain_FullMethodName         = "/pb.StockInfo/GetRealizedGain"
	StockInfo_GetFund_FullMethodName                 = "/pb.StockInfo/GetFund"
	StockInfo_AddFund_FullMethodName                 = "/pb.StockInfo/AddFund"
	StockInfo_GetStock_FullMethodName                = "/pb.StockInfo/GetStock"
	StockInfo_GetStocks_FullMethodName               = "/pb.StockInfo/GetStocks"
	StockInfo_TransationStock_FullMethodName         = "/pb.StockInfo/TransationStock"
	StockInfo_GetAllTransations_FullMethodName       = "/pb.StockInfo/GetAllTransations"
	StockInfo_GetUserStock_FullMethodName            = "/pb.StockInfo/GetUserStock"
	StockInfo_GetUserStockById_FullMethodName        = "/pb.StockInfo/GetUserStockById"
	StockInfo_GetRealizedProfitLoss_FullMethodName   = "/pb.StockInfo/GetRealizedProfitLoss"
	StockInfo_GetUnRealizedProfitLoss_FullMethodName = "/pb.StockInfo/GetUnRealizedProfitLoss"
)

// StockInfoClient is the client API for StockInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockInfoClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	InitStock(ctx context.Context, in *InitStockRequest, opts ...grpc.CallOption) (*InitStockResponse, error)
	GetUnRealizedGain(ctx context.Context, in *GetUnRealizedGainRequest, opts ...grpc.CallOption) (*GetUnRealizedGainResponse, error)
	GetRealizedGain(ctx context.Context, in *GetRealizedGainRequest, opts ...grpc.CallOption) (*GetRealizedGainResponse, error)
	GetFund(ctx context.Context, in *GetFundRequest, opts ...grpc.CallOption) (*GetFundResponse, error)
	AddFund(ctx context.Context, in *AddFundRequest, opts ...grpc.CallOption) (*AddFundResponse, error)
	GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error)
	GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error)
	TransationStock(ctx context.Context, in *TransationRequest, opts ...grpc.CallOption) (*TransationResponse, error)
	GetAllTransations(ctx context.Context, in *GetAllStockTransationRequest, opts ...grpc.CallOption) (*StockTransatsionResponse, error)
	GetUserStock(ctx context.Context, in *GetUserStockRequest, opts ...grpc.CallOption) (*GetUserStockResponse, error)
	GetUserStockById(ctx context.Context, in *GetUserStockByIdRequest, opts ...grpc.CallOption) (*GetUserStockBuIdResponse, error)
	GetRealizedProfitLoss(ctx context.Context, in *GetRealizedProfitLossRequest, opts ...grpc.CallOption) (*GetRealizedProfitLossResponse, error)
	GetUnRealizedProfitLoss(ctx context.Context, in *GetUnRealizedProfitLossRequest, opts ...grpc.CallOption) (*GetUnRealizedProfitLossResponse, error)
}

type stockInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewStockInfoClient(cc grpc.ClientConnInterface) StockInfoClient {
	return &stockInfoClient{cc}
}

func (c *stockInfoClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, StockInfo_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, StockInfo_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, StockInfo_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, StockInfo_VerifyEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) InitStock(ctx context.Context, in *InitStockRequest, opts ...grpc.CallOption) (*InitStockResponse, error) {
	out := new(InitStockResponse)
	err := c.cc.Invoke(ctx, StockInfo_InitStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetUnRealizedGain(ctx context.Context, in *GetUnRealizedGainRequest, opts ...grpc.CallOption) (*GetUnRealizedGainResponse, error) {
	out := new(GetUnRealizedGainResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetUnRealizedGain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetRealizedGain(ctx context.Context, in *GetRealizedGainRequest, opts ...grpc.CallOption) (*GetRealizedGainResponse, error) {
	out := new(GetRealizedGainResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetRealizedGain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetFund(ctx context.Context, in *GetFundRequest, opts ...grpc.CallOption) (*GetFundResponse, error) {
	out := new(GetFundResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetFund_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) AddFund(ctx context.Context, in *AddFundRequest, opts ...grpc.CallOption) (*AddFundResponse, error) {
	out := new(AddFundResponse)
	err := c.cc.Invoke(ctx, StockInfo_AddFund_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetStock(ctx context.Context, in *GetStockRequest, opts ...grpc.CallOption) (*GetStockResponse, error) {
	out := new(GetStockResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error) {
	out := new(GetStocksResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetStocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) TransationStock(ctx context.Context, in *TransationRequest, opts ...grpc.CallOption) (*TransationResponse, error) {
	out := new(TransationResponse)
	err := c.cc.Invoke(ctx, StockInfo_TransationStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetAllTransations(ctx context.Context, in *GetAllStockTransationRequest, opts ...grpc.CallOption) (*StockTransatsionResponse, error) {
	out := new(StockTransatsionResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetAllTransations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetUserStock(ctx context.Context, in *GetUserStockRequest, opts ...grpc.CallOption) (*GetUserStockResponse, error) {
	out := new(GetUserStockResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetUserStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetUserStockById(ctx context.Context, in *GetUserStockByIdRequest, opts ...grpc.CallOption) (*GetUserStockBuIdResponse, error) {
	out := new(GetUserStockBuIdResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetUserStockById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetRealizedProfitLoss(ctx context.Context, in *GetRealizedProfitLossRequest, opts ...grpc.CallOption) (*GetRealizedProfitLossResponse, error) {
	out := new(GetRealizedProfitLossResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetRealizedProfitLoss_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockInfoClient) GetUnRealizedProfitLoss(ctx context.Context, in *GetUnRealizedProfitLossRequest, opts ...grpc.CallOption) (*GetUnRealizedProfitLossResponse, error) {
	out := new(GetUnRealizedProfitLossResponse)
	err := c.cc.Invoke(ctx, StockInfo_GetUnRealizedProfitLoss_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockInfoServer is the server API for StockInfo service.
// All implementations must embed UnimplementedStockInfoServer
// for forward compatibility
type StockInfoServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	InitStock(context.Context, *InitStockRequest) (*InitStockResponse, error)
	GetUnRealizedGain(context.Context, *GetUnRealizedGainRequest) (*GetUnRealizedGainResponse, error)
	GetRealizedGain(context.Context, *GetRealizedGainRequest) (*GetRealizedGainResponse, error)
	GetFund(context.Context, *GetFundRequest) (*GetFundResponse, error)
	AddFund(context.Context, *AddFundRequest) (*AddFundResponse, error)
	GetStock(context.Context, *GetStockRequest) (*GetStockResponse, error)
	GetStocks(context.Context, *GetStocksRequest) (*GetStocksResponse, error)
	TransationStock(context.Context, *TransationRequest) (*TransationResponse, error)
	GetAllTransations(context.Context, *GetAllStockTransationRequest) (*StockTransatsionResponse, error)
	GetUserStock(context.Context, *GetUserStockRequest) (*GetUserStockResponse, error)
	GetUserStockById(context.Context, *GetUserStockByIdRequest) (*GetUserStockBuIdResponse, error)
	GetRealizedProfitLoss(context.Context, *GetRealizedProfitLossRequest) (*GetRealizedProfitLossResponse, error)
	GetUnRealizedProfitLoss(context.Context, *GetUnRealizedProfitLossRequest) (*GetUnRealizedProfitLossResponse, error)
	mustEmbedUnimplementedStockInfoServer()
}

// UnimplementedStockInfoServer must be embedded to have forward compatible implementations.
type UnimplementedStockInfoServer struct {
}

func (UnimplementedStockInfoServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedStockInfoServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedStockInfoServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedStockInfoServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedStockInfoServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedStockInfoServer) InitStock(context.Context, *InitStockRequest) (*InitStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitStock not implemented")
}
func (UnimplementedStockInfoServer) GetUnRealizedGain(context.Context, *GetUnRealizedGainRequest) (*GetUnRealizedGainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnRealizedGain not implemented")
}
func (UnimplementedStockInfoServer) GetRealizedGain(context.Context, *GetRealizedGainRequest) (*GetRealizedGainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealizedGain not implemented")
}
func (UnimplementedStockInfoServer) GetFund(context.Context, *GetFundRequest) (*GetFundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFund not implemented")
}
func (UnimplementedStockInfoServer) AddFund(context.Context, *AddFundRequest) (*AddFundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFund not implemented")
}
func (UnimplementedStockInfoServer) GetStock(context.Context, *GetStockRequest) (*GetStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (UnimplementedStockInfoServer) GetStocks(context.Context, *GetStocksRequest) (*GetStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStocks not implemented")
}
func (UnimplementedStockInfoServer) TransationStock(context.Context, *TransationRequest) (*TransationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransationStock not implemented")
}
func (UnimplementedStockInfoServer) GetAllTransations(context.Context, *GetAllStockTransationRequest) (*StockTransatsionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTransations not implemented")
}
func (UnimplementedStockInfoServer) GetUserStock(context.Context, *GetUserStockRequest) (*GetUserStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStock not implemented")
}
func (UnimplementedStockInfoServer) GetUserStockById(context.Context, *GetUserStockByIdRequest) (*GetUserStockBuIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStockById not implemented")
}
func (UnimplementedStockInfoServer) GetRealizedProfitLoss(context.Context, *GetRealizedProfitLossRequest) (*GetRealizedProfitLossResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealizedProfitLoss not implemented")
}
func (UnimplementedStockInfoServer) GetUnRealizedProfitLoss(context.Context, *GetUnRealizedProfitLossRequest) (*GetUnRealizedProfitLossResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnRealizedProfitLoss not implemented")
}
func (UnimplementedStockInfoServer) mustEmbedUnimplementedStockInfoServer() {}

// UnsafeStockInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockInfoServer will
// result in compilation errors.
type UnsafeStockInfoServer interface {
	mustEmbedUnimplementedStockInfoServer()
}

func RegisterStockInfoServer(s grpc.ServiceRegistrar, srv StockInfoServer) {
	s.RegisterService(&StockInfo_ServiceDesc, srv)
}

func _StockInfo_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_InitStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).InitStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_InitStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).InitStock(ctx, req.(*InitStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetUnRealizedGain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnRealizedGainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetUnRealizedGain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetUnRealizedGain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetUnRealizedGain(ctx, req.(*GetUnRealizedGainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetRealizedGain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealizedGainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetRealizedGain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetRealizedGain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetRealizedGain(ctx, req.(*GetRealizedGainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetFund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetFund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetFund_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetFund(ctx, req.(*GetFundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_AddFund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).AddFund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_AddFund_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).AddFund(ctx, req.(*AddFundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetStock(ctx, req.(*GetStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetStocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetStocks(ctx, req.(*GetStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_TransationStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).TransationStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_TransationStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).TransationStock(ctx, req.(*TransationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetAllTransations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllStockTransationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetAllTransations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetAllTransations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetAllTransations(ctx, req.(*GetAllStockTransationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetUserStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetUserStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetUserStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetUserStock(ctx, req.(*GetUserStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetUserStockById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStockByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetUserStockById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetUserStockById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetUserStockById(ctx, req.(*GetUserStockByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetRealizedProfitLoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealizedProfitLossRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetRealizedProfitLoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetRealizedProfitLoss_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetRealizedProfitLoss(ctx, req.(*GetRealizedProfitLossRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockInfo_GetUnRealizedProfitLoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnRealizedProfitLossRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockInfoServer).GetUnRealizedProfitLoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockInfo_GetUnRealizedProfitLoss_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockInfoServer).GetUnRealizedProfitLoss(ctx, req.(*GetUnRealizedProfitLossRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockInfo_ServiceDesc is the grpc.ServiceDesc for StockInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StockInfo",
	HandlerType: (*StockInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _StockInfo_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _StockInfo_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _StockInfo_UpdateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _StockInfo_LoginUser_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _StockInfo_VerifyEmail_Handler,
		},
		{
			MethodName: "InitStock",
			Handler:    _StockInfo_InitStock_Handler,
		},
		{
			MethodName: "GetUnRealizedGain",
			Handler:    _StockInfo_GetUnRealizedGain_Handler,
		},
		{
			MethodName: "GetRealizedGain",
			Handler:    _StockInfo_GetRealizedGain_Handler,
		},
		{
			MethodName: "GetFund",
			Handler:    _StockInfo_GetFund_Handler,
		},
		{
			MethodName: "AddFund",
			Handler:    _StockInfo_AddFund_Handler,
		},
		{
			MethodName: "GetStock",
			Handler:    _StockInfo_GetStock_Handler,
		},
		{
			MethodName: "GetStocks",
			Handler:    _StockInfo_GetStocks_Handler,
		},
		{
			MethodName: "TransationStock",
			Handler:    _StockInfo_TransationStock_Handler,
		},
		{
			MethodName: "GetAllTransations",
			Handler:    _StockInfo_GetAllTransations_Handler,
		},
		{
			MethodName: "GetUserStock",
			Handler:    _StockInfo_GetUserStock_Handler,
		},
		{
			MethodName: "GetUserStockById",
			Handler:    _StockInfo_GetUserStockById_Handler,
		},
		{
			MethodName: "GetRealizedProfitLoss",
			Handler:    _StockInfo_GetRealizedProfitLoss_Handler,
		},
		{
			MethodName: "GetUnRealizedProfitLoss",
			Handler:    _StockInfo_GetUnRealizedProfitLoss_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_stockinfo.proto",
}
