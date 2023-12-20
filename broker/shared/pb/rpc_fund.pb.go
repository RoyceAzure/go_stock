// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.0
// source: rpc_fund.proto

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

type GetUnRealizedGainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUnRealizedGainRequest) Reset() {
	*x = GetUnRealizedGainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnRealizedGainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnRealizedGainRequest) ProtoMessage() {}

func (x *GetUnRealizedGainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnRealizedGainRequest.ProtoReflect.Descriptor instead.
func (*GetUnRealizedGainRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{0}
}

type GetUnRealizedGainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*RealizedGain `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetUnRealizedGainResponse) Reset() {
	*x = GetUnRealizedGainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnRealizedGainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnRealizedGainResponse) ProtoMessage() {}

func (x *GetUnRealizedGainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnRealizedGainResponse.ProtoReflect.Descriptor instead.
func (*GetUnRealizedGainResponse) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{1}
}

func (x *GetUnRealizedGainResponse) GetData() []*RealizedGain {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetRealizedGainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRealizedGainRequest) Reset() {
	*x = GetRealizedGainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRealizedGainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealizedGainRequest) ProtoMessage() {}

func (x *GetRealizedGainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealizedGainRequest.ProtoReflect.Descriptor instead.
func (*GetRealizedGainRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{2}
}

type GetRealizedGainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*RealizedGain `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetRealizedGainResponse) Reset() {
	*x = GetRealizedGainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRealizedGainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealizedGainResponse) ProtoMessage() {}

func (x *GetRealizedGainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealizedGainResponse.ProtoReflect.Descriptor instead.
func (*GetRealizedGainResponse) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{3}
}

func (x *GetRealizedGainResponse) GetData() []*RealizedGain {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetFundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFundRequest) Reset() {
	*x = GetFundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundRequest) ProtoMessage() {}

func (x *GetFundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundRequest.ProtoReflect.Descriptor instead.
func (*GetFundRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{4}
}

type GetFundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Fund `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetFundResponse) Reset() {
	*x = GetFundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundResponse) ProtoMessage() {}

func (x *GetFundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundResponse.ProtoReflect.Descriptor instead.
func (*GetFundResponse) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{5}
}

func (x *GetFundResponse) GetData() []*Fund {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddFundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount       string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CurrencyType string `protobuf:"bytes,2,opt,name=currency_type,json=currencyType,proto3" json:"currency_type,omitempty"`
}

func (x *AddFundRequest) Reset() {
	*x = AddFundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFundRequest) ProtoMessage() {}

func (x *AddFundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFundRequest.ProtoReflect.Descriptor instead.
func (*AddFundRequest) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{6}
}

func (x *AddFundRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *AddFundRequest) GetCurrencyType() string {
	if x != nil {
		return x.CurrencyType
	}
	return ""
}

type AddFundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   *Fund  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddFundResponse) Reset() {
	*x = AddFundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_fund_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFundResponse) ProtoMessage() {}

func (x *AddFundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_fund_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFundResponse.ProtoReflect.Descriptor instead.
func (*AddFundResponse) Descriptor() ([]byte, []int) {
	return file_rpc_fund_proto_rawDescGZIP(), []int{7}
}

func (x *AddFundResponse) GetData() *Fund {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AddFundResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_rpc_fund_proto protoreflect.FileDescriptor

var file_rpc_fund_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x47, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x47, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x47, 0x61, 0x69, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x47, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x47, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x47, 0x61, 0x69, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x0f,
	0x41, 0x64, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x79, 0x63, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x2d, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_fund_proto_rawDescOnce sync.Once
	file_rpc_fund_proto_rawDescData = file_rpc_fund_proto_rawDesc
)

func file_rpc_fund_proto_rawDescGZIP() []byte {
	file_rpc_fund_proto_rawDescOnce.Do(func() {
		file_rpc_fund_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_fund_proto_rawDescData)
	})
	return file_rpc_fund_proto_rawDescData
}

var file_rpc_fund_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_fund_proto_goTypes = []interface{}{
	(*GetUnRealizedGainRequest)(nil),  // 0: pb.GetUnRealizedGainRequest
	(*GetUnRealizedGainResponse)(nil), // 1: pb.GetUnRealizedGainResponse
	(*GetRealizedGainRequest)(nil),    // 2: pb.GetRealizedGainRequest
	(*GetRealizedGainResponse)(nil),   // 3: pb.GetRealizedGainResponse
	(*GetFundRequest)(nil),            // 4: pb.GetFundRequest
	(*GetFundResponse)(nil),           // 5: pb.GetFundResponse
	(*AddFundRequest)(nil),            // 6: pb.AddFundRequest
	(*AddFundResponse)(nil),           // 7: pb.AddFundResponse
	(*RealizedGain)(nil),              // 8: pb.RealizedGain
	(*Fund)(nil),                      // 9: pb.Fund
}
var file_rpc_fund_proto_depIdxs = []int32{
	8, // 0: pb.GetUnRealizedGainResponse.data:type_name -> pb.RealizedGain
	8, // 1: pb.GetRealizedGainResponse.data:type_name -> pb.RealizedGain
	9, // 2: pb.GetFundResponse.data:type_name -> pb.Fund
	9, // 3: pb.AddFundResponse.data:type_name -> pb.Fund
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_fund_proto_init() }
func file_rpc_fund_proto_init() {
	if File_rpc_fund_proto != nil {
		return
	}
	file_fund_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_fund_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnRealizedGainRequest); i {
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
		file_rpc_fund_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnRealizedGainResponse); i {
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
		file_rpc_fund_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRealizedGainRequest); i {
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
		file_rpc_fund_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRealizedGainResponse); i {
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
		file_rpc_fund_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundRequest); i {
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
		file_rpc_fund_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundResponse); i {
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
		file_rpc_fund_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFundRequest); i {
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
		file_rpc_fund_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFundResponse); i {
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
			RawDescriptor: file_rpc_fund_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_fund_proto_goTypes,
		DependencyIndexes: file_rpc_fund_proto_depIdxs,
		MessageInfos:      file_rpc_fund_proto_msgTypes,
	}.Build()
	File_rpc_fund_proto = out.File
	file_rpc_fund_proto_rawDesc = nil
	file_rpc_fund_proto_goTypes = nil
	file_rpc_fund_proto_depIdxs = nil
}