// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.0
// source: fund.proto

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

type Fund struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Balance      string                 `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	CurrencyType string                 `protobuf:"bytes,3,opt,name=currency_type,json=currencyType,proto3" json:"currency_type,omitempty"`
	SsoIdentifer string                 `protobuf:"bytes,4,opt,name=sso_identifer,json=ssoIdentifer,proto3" json:"sso_identifer,omitempty"`
	UpDate       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=up_date,json=upDate,proto3" json:"up_date,omitempty"`
	CrDate       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=cr_date,json=crDate,proto3" json:"cr_date,omitempty"`
}

func (x *Fund) Reset() {
	*x = Fund{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fund) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fund) ProtoMessage() {}

func (x *Fund) ProtoReflect() protoreflect.Message {
	mi := &file_fund_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fund.ProtoReflect.Descriptor instead.
func (*Fund) Descriptor() ([]byte, []int) {
	return file_fund_proto_rawDescGZIP(), []int{0}
}

func (x *Fund) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Fund) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *Fund) GetCurrencyType() string {
	if x != nil {
		return x.CurrencyType
	}
	return ""
}

func (x *Fund) GetSsoIdentifer() string {
	if x != nil {
		return x.SsoIdentifer
	}
	return ""
}

func (x *Fund) GetUpDate() *timestamppb.Timestamp {
	if x != nil {
		return x.UpDate
	}
	return nil
}

func (x *Fund) GetCrDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CrDate
	}
	return nil
}

type RealizedGain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName  string `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Quantity     int64  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	AvgCost      string `protobuf:"bytes,3,opt,name=avg_cost,json=avgCost,proto3" json:"avg_cost,omitempty"`
	Cost         string `protobuf:"bytes,4,opt,name=cost,proto3" json:"cost,omitempty"`
	CurrentPrice string `protobuf:"bytes,5,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
	CurrentValue string `protobuf:"bytes,6,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	Gain         string `protobuf:"bytes,7,opt,name=gain,proto3" json:"gain,omitempty"`
	GainPt       string `protobuf:"bytes,8,opt,name=gain_pt,json=gainPt,proto3" json:"gain_pt,omitempty"`
}

func (x *RealizedGain) Reset() {
	*x = RealizedGain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fund_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealizedGain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealizedGain) ProtoMessage() {}

func (x *RealizedGain) ProtoReflect() protoreflect.Message {
	mi := &file_fund_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealizedGain.ProtoReflect.Descriptor instead.
func (*RealizedGain) Descriptor() ([]byte, []int) {
	return file_fund_proto_rawDescGZIP(), []int{1}
}

func (x *RealizedGain) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *RealizedGain) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *RealizedGain) GetAvgCost() string {
	if x != nil {
		return x.AvgCost
	}
	return ""
}

func (x *RealizedGain) GetCost() string {
	if x != nil {
		return x.Cost
	}
	return ""
}

func (x *RealizedGain) GetCurrentPrice() string {
	if x != nil {
		return x.CurrentPrice
	}
	return ""
}

func (x *RealizedGain) GetCurrentValue() string {
	if x != nil {
		return x.CurrentValue
	}
	return ""
}

func (x *RealizedGain) GetGain() string {
	if x != nil {
		return x.Gain
	}
	return ""
}

func (x *RealizedGain) GetGainPt() string {
	if x != nil {
		return x.GainPt
	}
	return ""
}

var File_fund_proto protoreflect.FileDescriptor

var file_fund_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xed, 0x01, 0x0a, 0x04, 0x46, 0x75, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x73, 0x6f, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x73, 0x6f, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x75, 0x70, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x75, 0x70, 0x44, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x07,
	0x63, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x63, 0x72, 0x44, 0x61, 0x74,
	0x65, 0x22, 0xf3, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x47, 0x61,
	0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x76, 0x67, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x76, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61,
	0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x61, 0x69, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x61, 0x69, 0x6e, 0x50, 0x74, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x79, 0x63, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fund_proto_rawDescOnce sync.Once
	file_fund_proto_rawDescData = file_fund_proto_rawDesc
)

func file_fund_proto_rawDescGZIP() []byte {
	file_fund_proto_rawDescOnce.Do(func() {
		file_fund_proto_rawDescData = protoimpl.X.CompressGZIP(file_fund_proto_rawDescData)
	})
	return file_fund_proto_rawDescData
}

var file_fund_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fund_proto_goTypes = []interface{}{
	(*Fund)(nil),                  // 0: pb.Fund
	(*RealizedGain)(nil),          // 1: pb.RealizedGain
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_fund_proto_depIdxs = []int32{
	2, // 0: pb.Fund.up_date:type_name -> google.protobuf.Timestamp
	2, // 1: pb.Fund.cr_date:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fund_proto_init() }
func file_fund_proto_init() {
	if File_fund_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fund_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fund); i {
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
		file_fund_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RealizedGain); i {
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
			RawDescriptor: file_fund_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fund_proto_goTypes,
		DependencyIndexes: file_fund_proto_depIdxs,
		MessageInfos:      file_fund_proto_msgTypes,
	}.Build()
	File_fund_proto = out.File
	file_fund_proto_rawDesc = nil
	file_fund_proto_goTypes = nil
	file_fund_proto_depIdxs = nil
}
