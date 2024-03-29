// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.0
// source: rpc_client_register.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateClientRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	StockCode string `protobuf:"bytes,2,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
}

func (x *CreateClientRegisterRequest) Reset() {
	*x = CreateClientRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRegisterRequest) ProtoMessage() {}

func (x *CreateClientRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRegisterRequest.ProtoReflect.Descriptor instead.
func (*CreateClientRegisterRequest) Descriptor() ([]byte, []int) {
	return file_rpc_client_register_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClientRegisterRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateClientRegisterRequest) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

type CreateClientRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	StockCode string                 `protobuf:"bytes,2,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
	CreateAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *CreateClientRegisterResponse) Reset() {
	*x = CreateClientRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRegisterResponse) ProtoMessage() {}

func (x *CreateClientRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRegisterResponse.ProtoReflect.Descriptor instead.
func (*CreateClientRegisterResponse) Descriptor() ([]byte, []int) {
	return file_rpc_client_register_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClientRegisterResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateClientRegisterResponse) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

func (x *CreateClientRegisterResponse) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

type DeleteClientRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	StockCode string `protobuf:"bytes,2,opt,name=stock_code,json=stockCode,proto3" json:"stock_code,omitempty"`
}

func (x *DeleteClientRegisterRequest) Reset() {
	*x = DeleteClientRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientRegisterRequest) ProtoMessage() {}

func (x *DeleteClientRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientRegisterRequest.ProtoReflect.Descriptor instead.
func (*DeleteClientRegisterRequest) Descriptor() ([]byte, []int) {
	return file_rpc_client_register_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteClientRegisterRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *DeleteClientRegisterRequest) GetStockCode() string {
	if x != nil {
		return x.StockCode
	}
	return ""
}

type DeleteClientRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteClientRegisterResponse) Reset() {
	*x = DeleteClientRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientRegisterResponse) ProtoMessage() {}

func (x *DeleteClientRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_register_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientRegisterResponse.ProtoReflect.Descriptor instead.
func (*DeleteClientRegisterResponse) Descriptor() ([]byte, []int) {
	return file_rpc_client_register_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteClientRegisterResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GetClientRegisterByClientUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *GetClientRegisterByClientUIDRequest) Reset() {
	*x = GetClientRegisterByClientUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_register_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientRegisterByClientUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientRegisterByClientUIDRequest) ProtoMessage() {}

func (x *GetClientRegisterByClientUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_register_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientRegisterByClientUIDRequest.ProtoReflect.Descriptor instead.
func (*GetClientRegisterByClientUIDRequest) Descriptor() ([]byte, []int) {
	return file_rpc_client_register_proto_rawDescGZIP(), []int{4}
}

func (x *GetClientRegisterByClientUIDRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type GetClientRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ClientRegister `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetClientRegisterResponse) Reset() {
	*x = GetClientRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_client_register_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientRegisterResponse) ProtoMessage() {}

func (x *GetClientRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_client_register_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientRegisterResponse.ProtoReflect.Descriptor instead.
func (*GetClientRegisterResponse) Descriptor() ([]byte, []int) {
	return file_rpc_client_register_proto_rawDescGZIP(), []int{5}
}

func (x *GetClientRegisterResponse) GetData() []*ClientRegister {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_rpc_client_register_proto protoreflect.FileDescriptor

var file_rpc_client_register_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x37, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x59, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x36, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x42, 0x0a, 0x23, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x42, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x43, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x79, 0x63, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x2d, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_client_register_proto_rawDescOnce sync.Once
	file_rpc_client_register_proto_rawDescData = file_rpc_client_register_proto_rawDesc
)

func file_rpc_client_register_proto_rawDescGZIP() []byte {
	file_rpc_client_register_proto_rawDescOnce.Do(func() {
		file_rpc_client_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_client_register_proto_rawDescData)
	})
	return file_rpc_client_register_proto_rawDescData
}

var file_rpc_client_register_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rpc_client_register_proto_goTypes = []interface{}{
	(*CreateClientRegisterRequest)(nil),         // 0: pb.CreateClientRegisterRequest
	(*CreateClientRegisterResponse)(nil),        // 1: pb.CreateClientRegisterResponse
	(*DeleteClientRegisterRequest)(nil),         // 2: pb.DeleteClientRegisterRequest
	(*DeleteClientRegisterResponse)(nil),        // 3: pb.DeleteClientRegisterResponse
	(*GetClientRegisterByClientUIDRequest)(nil), // 4: pb.GetClientRegisterByClientUIDRequest
	(*GetClientRegisterResponse)(nil),           // 5: pb.GetClientRegisterResponse
	(*timestamppb.Timestamp)(nil),               // 6: google.protobuf.Timestamp
	(*ClientRegister)(nil),                      // 7: pb.ClientRegister
}
var file_rpc_client_register_proto_depIdxs = []int32{
	6, // 0: pb.CreateClientRegisterResponse.create_at:type_name -> google.protobuf.Timestamp
	7, // 1: pb.GetClientRegisterResponse.data:type_name -> pb.ClientRegister
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_client_register_proto_init() }
func file_rpc_client_register_proto_init() {
	if File_rpc_client_register_proto != nil {
		return
	}
	file_client_register_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_client_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientRegisterRequest); i {
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
		file_rpc_client_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientRegisterResponse); i {
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
		file_rpc_client_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientRegisterRequest); i {
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
		file_rpc_client_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientRegisterResponse); i {
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
		file_rpc_client_register_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientRegisterByClientUIDRequest); i {
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
		file_rpc_client_register_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientRegisterResponse); i {
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
			RawDescriptor: file_rpc_client_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_client_register_proto_goTypes,
		DependencyIndexes: file_rpc_client_register_proto_depIdxs,
		MessageInfos:      file_rpc_client_register_proto_msgTypes,
	}.Build()
	File_rpc_client_register_proto = out.File
	file_rpc_client_register_proto_rawDesc = nil
	file_rpc_client_register_proto_goTypes = nil
	file_rpc_client_register_proto_depIdxs = nil
}
