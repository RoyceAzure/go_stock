// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.0
// source: rpc_stock_transation.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockCode      int64  `protobuf:"varint,2,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
	TransationType string `protobuf:"bytes,3,opt,name=transation_type,json=transationType,proto3" json:"transation_type,omitempty"`
	TransAmt       int64  `protobuf:"varint,4,opt,name=trans_amt,json=transAmt,proto3" json:"trans_amt,omitempty"`
}

func (x *TransationRequest) Reset() {
	*x = TransationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransationRequest) ProtoMessage() {}

func (x *TransationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransationRequest.ProtoReflect.Descriptor instead.
func (*TransationRequest) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{0}
}

func (x *TransationRequest) GetStockCode() int64 {
	if x != nil {
		return x.StockCode
	}
	return 0
}

func (x *TransationRequest) GetTransationType() string {
	if x != nil {
		return x.TransationType
	}
	return ""
}

func (x *TransationRequest) GetTransAmt() int64 {
	if x != nil {
		return x.TransAmt
	}
	return 0
}

type TransationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TransationResponse) Reset() {
	*x = TransationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransationResponse) ProtoMessage() {}

func (x *TransationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransationResponse.ProtoReflect.Descriptor instead.
func (*TransationResponse) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{1}
}

func (x *TransationResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GetTransationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	UserId   int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetTransationsRequest) Reset() {
	*x = GetTransationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransationsRequest) ProtoMessage() {}

func (x *GetTransationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransationsRequest.ProtoReflect.Descriptor instead.
func (*GetTransationsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetTransationsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetTransationsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetransationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*StockTransation `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetransationsResponse) Reset() {
	*x = GetransationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetransationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetransationsResponse) ProtoMessage() {}

func (x *GetransationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetransationsResponse.ProtoReflect.Descriptor instead.
func (*GetransationsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{3}
}

func (x *GetransationsResponse) GetData() []*StockTransation {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllStockTransationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StockId        int64  `protobuf:"varint,2,opt,name=stock_id,json=stockId,proto3" json:"stock_id,omitempty"`
	TransationType string `protobuf:"bytes,3,opt,name=transation_type,json=transationType,proto3" json:"transation_type,omitempty"`
	Page           int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize       int32  `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetAllStockTransationRequest) Reset() {
	*x = GetAllStockTransationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStockTransationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStockTransationRequest) ProtoMessage() {}

func (x *GetAllStockTransationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStockTransationRequest.ProtoReflect.Descriptor instead.
func (*GetAllStockTransationRequest) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllStockTransationRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetAllStockTransationRequest) GetStockId() int64 {
	if x != nil {
		return x.StockId
	}
	return 0
}

func (x *GetAllStockTransationRequest) GetTransationType() string {
	if x != nil {
		return x.TransationType
	}
	return ""
}

func (x *GetAllStockTransationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllStockTransationRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type StockTransatsionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*StockTransation `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *StockTransatsionResponse) Reset() {
	*x = StockTransatsionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockTransatsionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockTransatsionResponse) ProtoMessage() {}

func (x *StockTransatsionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockTransatsionResponse.ProtoReflect.Descriptor instead.
func (*StockTransatsionResponse) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{5}
}

func (x *StockTransatsionResponse) GetData() []*StockTransation {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllStockTransatioBySidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockId  int64 `protobuf:"varint,1,opt,name=stock_id,json=stockId,proto3" json:"stock_id,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetAllStockTransatioBySidRequest) Reset() {
	*x = GetAllStockTransatioBySidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStockTransatioBySidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStockTransatioBySidRequest) ProtoMessage() {}

func (x *GetAllStockTransatioBySidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStockTransatioBySidRequest.ProtoReflect.Descriptor instead.
func (*GetAllStockTransatioBySidRequest) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllStockTransatioBySidRequest) GetStockId() int64 {
	if x != nil {
		return x.StockId
	}
	return 0
}

func (x *GetAllStockTransatioBySidRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllStockTransatioBySidRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetAllStockTransatioByUidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetAllStockTransatioByUidRequest) Reset() {
	*x = GetAllStockTransatioByUidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_stock_transation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllStockTransatioByUidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStockTransatioByUidRequest) ProtoMessage() {}

func (x *GetAllStockTransatioByUidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_stock_transation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStockTransatioByUidRequest.ProtoReflect.Descriptor instead.
func (*GetAllStockTransatioByUidRequest) Descriptor() ([]byte, []int) {
	return file_rpc_stock_transation_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllStockTransatioByUidRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetAllStockTransatioByUidRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllStockTransatioByUidRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_rpc_stock_transation_proto protoreflect.FileDescriptor

var file_rpc_stock_transation_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x16, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x61,
	0x6d, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x41,
	0x6d, 0x74, 0x22, 0x2c, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x61, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xac, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x43, 0x0a, 0x18, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x74, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6e, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x42, 0x79, 0x53, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6c, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x42, 0x79, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x79, 0x63, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_stock_transation_proto_rawDescOnce sync.Once
	file_rpc_stock_transation_proto_rawDescData = file_rpc_stock_transation_proto_rawDesc
)

func file_rpc_stock_transation_proto_rawDescGZIP() []byte {
	file_rpc_stock_transation_proto_rawDescOnce.Do(func() {
		file_rpc_stock_transation_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_stock_transation_proto_rawDescData)
	})
	return file_rpc_stock_transation_proto_rawDescData
}

var file_rpc_stock_transation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_stock_transation_proto_goTypes = []interface{}{
	(*TransationRequest)(nil),                // 0: pb.TransationRequest
	(*TransationResponse)(nil),               // 1: pb.TransationResponse
	(*GetTransationsRequest)(nil),            // 2: pb.GetTransationsRequest
	(*GetransationsResponse)(nil),            // 3: pb.GetransationsResponse
	(*GetAllStockTransationRequest)(nil),     // 4: pb.GetAllStockTransationRequest
	(*StockTransatsionResponse)(nil),         // 5: pb.StockTransatsionResponse
	(*GetAllStockTransatioBySidRequest)(nil), // 6: pb.GetAllStockTransatioBySidRequest
	(*GetAllStockTransatioByUidRequest)(nil), // 7: pb.GetAllStockTransatioByUidRequest
	(*StockTransation)(nil),                  // 8: pb.StockTransation
}
var file_rpc_stock_transation_proto_depIdxs = []int32{
	8, // 0: pb.GetransationsResponse.data:type_name -> pb.StockTransation
	8, // 1: pb.StockTransatsionResponse.data:type_name -> pb.StockTransation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_stock_transation_proto_init() }
func file_rpc_stock_transation_proto_init() {
	if File_rpc_stock_transation_proto != nil {
		return
	}
	file_stock_transation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_stock_transation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_stock_transation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_stock_transation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransationsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_stock_transation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetransationsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_stock_transation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStockTransationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_stock_transation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockTransatsionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_stock_transation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStockTransatioBySidRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_stock_transation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllStockTransatioByUidRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_stock_transation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_stock_transation_proto_goTypes,
		DependencyIndexes: file_rpc_stock_transation_proto_depIdxs,
		MessageInfos:      file_rpc_stock_transation_proto_msgTypes,
	}.Build()
	File_rpc_stock_transation_proto = out.File
	file_rpc_stock_transation_proto_rawDesc = nil
	file_rpc_stock_transation_proto_goTypes = nil
	file_rpc_stock_transation_proto_depIdxs = nil
}
