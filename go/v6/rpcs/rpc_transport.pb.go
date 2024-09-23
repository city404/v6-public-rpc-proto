// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: rpcs/rpc_transport.proto

package rpcs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RpcTransport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command  string       `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Identity string       `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	Type     int32        `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Data     *anypb.Any   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Status   int32        `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	Message  string       `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Meta     *RpcMetaData `protobuf:"bytes,7,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *RpcTransport) Reset() {
	*x = RpcTransport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcTransport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcTransport) ProtoMessage() {}

func (x *RpcTransport) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcTransport.ProtoReflect.Descriptor instead.
func (*RpcTransport) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_transport_proto_rawDescGZIP(), []int{0}
}

func (x *RpcTransport) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *RpcTransport) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *RpcTransport) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *RpcTransport) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RpcTransport) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RpcTransport) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RpcTransport) GetMeta() *RpcMetaData {
	if x != nil {
		return x.Meta
	}
	return nil
}

type RpcHeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *RpcHeaderValue) Reset() {
	*x = RpcHeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcHeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcHeaderValue) ProtoMessage() {}

func (x *RpcHeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcHeaderValue.ProtoReflect.Descriptor instead.
func (*RpcHeaderValue) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_transport_proto_rawDescGZIP(), []int{1}
}

func (x *RpcHeaderValue) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type RpcMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64                      `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	From      string                     `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Version   int32                      `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Header    map[string]*RpcHeaderValue `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RpcMetaData) Reset() {
	*x = RpcMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpcs_rpc_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcMetaData) ProtoMessage() {}

func (x *RpcMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_rpcs_rpc_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcMetaData.ProtoReflect.Descriptor instead.
func (*RpcMetaData) Descriptor() ([]byte, []int) {
	return file_rpcs_rpc_transport_proto_rawDescGZIP(), []int{2}
}

func (x *RpcMetaData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RpcMetaData) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *RpcMetaData) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *RpcMetaData) GetHeader() map[string]*RpcHeaderValue {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_rpcs_rpc_transport_proto protoreflect.FileDescriptor

var file_rpcs_rpc_transport_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x36, 0x2e, 0x72,
	0x70, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01,
	0x0a, 0x0c, 0x52, 0x70, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x36, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x4d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x26, 0x0a,
	0x0e, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x0b, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x74,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x36, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x51, 0x0a, 0x0b, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x36, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x70, 0x63, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x79,
	0x34, 0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x72, 0x70,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x36, 0x2f, 0x72, 0x70,
	0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpcs_rpc_transport_proto_rawDescOnce sync.Once
	file_rpcs_rpc_transport_proto_rawDescData = file_rpcs_rpc_transport_proto_rawDesc
)

func file_rpcs_rpc_transport_proto_rawDescGZIP() []byte {
	file_rpcs_rpc_transport_proto_rawDescOnce.Do(func() {
		file_rpcs_rpc_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpcs_rpc_transport_proto_rawDescData)
	})
	return file_rpcs_rpc_transport_proto_rawDescData
}

var file_rpcs_rpc_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpcs_rpc_transport_proto_goTypes = []any{
	(*RpcTransport)(nil),   // 0: v6.rpc.RpcTransport
	(*RpcHeaderValue)(nil), // 1: v6.rpc.RpcHeaderValue
	(*RpcMetaData)(nil),    // 2: v6.rpc.RpcMetaData
	nil,                    // 3: v6.rpc.RpcMetaData.HeaderEntry
	(*anypb.Any)(nil),      // 4: google.protobuf.Any
}
var file_rpcs_rpc_transport_proto_depIdxs = []int32{
	4, // 0: v6.rpc.RpcTransport.data:type_name -> google.protobuf.Any
	2, // 1: v6.rpc.RpcTransport.meta:type_name -> v6.rpc.RpcMetaData
	3, // 2: v6.rpc.RpcMetaData.header:type_name -> v6.rpc.RpcMetaData.HeaderEntry
	1, // 3: v6.rpc.RpcMetaData.HeaderEntry.value:type_name -> v6.rpc.RpcHeaderValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpcs_rpc_transport_proto_init() }
func file_rpcs_rpc_transport_proto_init() {
	if File_rpcs_rpc_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpcs_rpc_transport_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RpcTransport); i {
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
		file_rpcs_rpc_transport_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RpcHeaderValue); i {
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
		file_rpcs_rpc_transport_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RpcMetaData); i {
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
			RawDescriptor: file_rpcs_rpc_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpcs_rpc_transport_proto_goTypes,
		DependencyIndexes: file_rpcs_rpc_transport_proto_depIdxs,
		MessageInfos:      file_rpcs_rpc_transport_proto_msgTypes,
	}.Build()
	File_rpcs_rpc_transport_proto = out.File
	file_rpcs_rpc_transport_proto_rawDesc = nil
	file_rpcs_rpc_transport_proto_goTypes = nil
	file_rpcs_rpc_transport_proto_depIdxs = nil
}
