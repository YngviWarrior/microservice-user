// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: user_stretegy.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserStrategy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserStrategy  uint64                 `protobuf:"varint,1,opt,name=user_strategy,json=userStrategy,proto3" json:"user_strategy,omitempty"`
	User          uint64                 `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	TradeConfig   string                 `protobuf:"bytes,3,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserStrategy) Reset() {
	*x = UserStrategy{}
	mi := &file_user_stretegy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStrategy) ProtoMessage() {}

func (x *UserStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_user_stretegy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStrategy.ProtoReflect.Descriptor instead.
func (*UserStrategy) Descriptor() ([]byte, []int) {
	return file_user_stretegy_proto_rawDescGZIP(), []int{0}
}

func (x *UserStrategy) GetUserStrategy() uint64 {
	if x != nil {
		return x.UserStrategy
	}
	return 0
}

func (x *UserStrategy) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *UserStrategy) GetTradeConfig() string {
	if x != nil {
		return x.TradeConfig
	}
	return ""
}

type UserStrategyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserStrategy  *UserStrategy          `protobuf:"bytes,1,opt,name=user_strategy,json=userStrategy,proto3" json:"user_strategy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserStrategyResponse) Reset() {
	*x = UserStrategyResponse{}
	mi := &file_user_stretegy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStrategyResponse) ProtoMessage() {}

func (x *UserStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_stretegy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStrategyResponse.ProtoReflect.Descriptor instead.
func (*UserStrategyResponse) Descriptor() ([]byte, []int) {
	return file_user_stretegy_proto_rawDescGZIP(), []int{1}
}

func (x *UserStrategyResponse) GetUserStrategy() *UserStrategy {
	if x != nil {
		return x.UserStrategy
	}
	return nil
}

type CreateUserStrategyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          uint64                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	TradeConfig   string                 `protobuf:"bytes,2,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserStrategyRequest) Reset() {
	*x = CreateUserStrategyRequest{}
	mi := &file_user_stretegy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserStrategyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserStrategyRequest) ProtoMessage() {}

func (x *CreateUserStrategyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_stretegy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserStrategyRequest.ProtoReflect.Descriptor instead.
func (*CreateUserStrategyRequest) Descriptor() ([]byte, []int) {
	return file_user_stretegy_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserStrategyRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreateUserStrategyRequest) GetTradeConfig() string {
	if x != nil {
		return x.TradeConfig
	}
	return ""
}

var File_user_stretegy_proto protoreflect.FileDescriptor

var file_user_stretegy_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x67, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x6a, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4d, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x22, 0x52, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x66, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_user_stretegy_proto_rawDescOnce sync.Once
	file_user_stretegy_proto_rawDescData []byte
)

func file_user_stretegy_proto_rawDescGZIP() []byte {
	file_user_stretegy_proto_rawDescOnce.Do(func() {
		file_user_stretegy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_stretegy_proto_rawDesc), len(file_user_stretegy_proto_rawDesc)))
	})
	return file_user_stretegy_proto_rawDescData
}

var file_user_stretegy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_stretegy_proto_goTypes = []any{
	(*UserStrategy)(nil),              // 0: pb.UserStrategy
	(*UserStrategyResponse)(nil),      // 1: pb.UserStrategyResponse
	(*CreateUserStrategyRequest)(nil), // 2: pb.CreateUserStrategyRequest
}
var file_user_stretegy_proto_depIdxs = []int32{
	0, // 0: pb.UserStrategyResponse.user_strategy:type_name -> pb.UserStrategy
	2, // 1: pb.UserStrategyService.CreateUserStrategy:input_type -> pb.CreateUserStrategyRequest
	1, // 2: pb.UserStrategyService.CreateUserStrategy:output_type -> pb.UserStrategyResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_stretegy_proto_init() }
func file_user_stretegy_proto_init() {
	if File_user_stretegy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_stretegy_proto_rawDesc), len(file_user_stretegy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_stretegy_proto_goTypes,
		DependencyIndexes: file_user_stretegy_proto_depIdxs,
		MessageInfos:      file_user_stretegy_proto_msgTypes,
	}.Build()
	File_user_stretegy_proto = out.File
	file_user_stretegy_proto_goTypes = nil
	file_user_stretegy_proto_depIdxs = nil
}
