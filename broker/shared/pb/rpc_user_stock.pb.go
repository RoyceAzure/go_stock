// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.0
// source: rpc_user_stock.proto

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

type GetUserStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetUserStockRequest) Reset() {
	*x = GetUserStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserStockRequest) ProtoMessage() {}

func (x *GetUserStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserStockRequest.ProtoReflect.Descriptor instead.
func (*GetUserStockRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_stock_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserStockRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetUserStockRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetUserStockResponseDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StockId               int64  `protobuf:"varint,2,opt,name=stock_id,json=stockId,proto3" json:"stock_id,omitempty"`
	StockCode             string `protobuf:"bytes,3,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
	StockName             string `protobuf:"bytes,4,opt,name=stock_name,json=stockName,proto3" json:"stock_name,omitempty"`
	Quantity              int32  `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PurchasePricePerShare string `protobuf:"bytes,6,opt,name=purchase_price_per_share,json=purchasePricePerShare,proto3" json:"purchase_price_per_share,omitempty"`
}

func (x *GetUserStockResponseDTO) Reset() {
	*x = GetUserStockResponseDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserStockResponseDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserStockResponseDTO) ProtoMessage() {}

func (x *GetUserStockResponseDTO) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserStockResponseDTO.ProtoReflect.Descriptor instead.
func (*GetUserStockResponseDTO) Descriptor() ([]byte, []int) {
	return file_rpc_user_stock_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserStockResponseDTO) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserStockResponseDTO) GetStockId() int64 {
	if x != nil {
		return x.StockId
	}
	return 0
}

func (x *GetUserStockResponseDTO) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

func (x *GetUserStockResponseDTO) GetStockName() string {
	if x != nil {
		return x.StockName
	}
	return ""
}

func (x *GetUserStockResponseDTO) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *GetUserStockResponseDTO) GetPurchasePricePerShare() string {
	if x != nil {
		return x.PurchasePricePerShare
	}
	return ""
}

type GetUserStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*GetUserStockResponseDTO `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserStockResponse) Reset() {
	*x = GetUserStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserStockResponse) ProtoMessage() {}

func (x *GetUserStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserStockResponse.ProtoReflect.Descriptor instead.
func (*GetUserStockResponse) Descriptor() ([]byte, []int) {
	return file_rpc_user_stock_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserStockResponse) GetData() []*GetUserStockResponseDTO {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetUserStockByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserStockByIdRequest) Reset() {
	*x = GetUserStockByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserStockByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserStockByIdRequest) ProtoMessage() {}

func (x *GetUserStockByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserStockByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserStockByIdRequest) Descriptor() ([]byte, []int) {
	return file_rpc_user_stock_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserStockByIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserStockBuIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*UserStock `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUserStockBuIdResponse) Reset() {
	*x = GetUserStockBuIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_user_stock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserStockBuIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserStockBuIdResponse) ProtoMessage() {}

func (x *GetUserStockBuIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_user_stock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserStockBuIdResponse.ProtoReflect.Descriptor instead.
func (*GetUserStockBuIdResponse) Descriptor() ([]byte, []int) {
	return file_rpc_user_stock_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserStockBuIdResponse) GetData() []*UserStock {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_rpc_user_stock_proto protoreflect.FileDescriptor

var file_rpc_user_stock_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x54, 0x4f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x37,
	0x0a, 0x18, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x54, 0x4f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x42, 0x75, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x52, 0x6f, 0x79, 0x63, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x2d,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x2d, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_user_stock_proto_rawDescOnce sync.Once
	file_rpc_user_stock_proto_rawDescData = file_rpc_user_stock_proto_rawDesc
)

func file_rpc_user_stock_proto_rawDescGZIP() []byte {
	file_rpc_user_stock_proto_rawDescOnce.Do(func() {
		file_rpc_user_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_user_stock_proto_rawDescData)
	})
	return file_rpc_user_stock_proto_rawDescData
}

var file_rpc_user_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_user_stock_proto_goTypes = []interface{}{
	(*GetUserStockRequest)(nil),      // 0: pb.GetUserStockRequest
	(*GetUserStockResponseDTO)(nil),  // 1: pb.GetUserStockResponseDTO
	(*GetUserStockResponse)(nil),     // 2: pb.GetUserStockResponse
	(*GetUserStockByIdRequest)(nil),  // 3: pb.GetUserStockByIdRequest
	(*GetUserStockBuIdResponse)(nil), // 4: pb.GetUserStockBuIdResponse
	(*UserStock)(nil),                // 5: pb.UserStock
}
var file_rpc_user_stock_proto_depIdxs = []int32{
	1, // 0: pb.GetUserStockResponse.data:type_name -> pb.GetUserStockResponseDTO
	5, // 1: pb.GetUserStockBuIdResponse.data:type_name -> pb.UserStock
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_user_stock_proto_init() }
func file_rpc_user_stock_proto_init() {
	if File_rpc_user_stock_proto != nil {
		return
	}
	file_user_stock_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_user_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserStockRequest); i {
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
		file_rpc_user_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserStockResponseDTO); i {
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
		file_rpc_user_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserStockResponse); i {
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
		file_rpc_user_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserStockByIdRequest); i {
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
		file_rpc_user_stock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserStockBuIdResponse); i {
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
			RawDescriptor: file_rpc_user_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_user_stock_proto_goTypes,
		DependencyIndexes: file_rpc_user_stock_proto_depIdxs,
		MessageInfos:      file_rpc_user_stock_proto_msgTypes,
	}.Build()
	File_rpc_user_stock_proto = out.File
	file_rpc_user_stock_proto_rawDesc = nil
	file_rpc_user_stock_proto_goTypes = nil
	file_rpc_user_stock_proto_depIdxs = nil
}
