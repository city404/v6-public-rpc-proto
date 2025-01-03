// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: common/pub_common.proto

package common

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

type ScanListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Limit         int64                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy       []*OrderByInfo         `protobuf:"bytes,3,rep,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Version       int32                  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScanListRequest) Reset() {
	*x = ScanListRequest{}
	mi := &file_common_pub_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScanListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanListRequest) ProtoMessage() {}

func (x *ScanListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_pub_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanListRequest.ProtoReflect.Descriptor instead.
func (*ScanListRequest) Descriptor() ([]byte, []int) {
	return file_common_pub_common_proto_rawDescGZIP(), []int{0}
}

func (x *ScanListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ScanListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ScanListRequest) GetOrderBy() []*OrderByInfo {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *ScanListRequest) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type OrderByInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Field         string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Asc           bool                   `protobuf:"varint,2,opt,name=asc,proto3" json:"asc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderByInfo) Reset() {
	*x = OrderByInfo{}
	mi := &file_common_pub_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderByInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderByInfo) ProtoMessage() {}

func (x *OrderByInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_pub_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderByInfo.ProtoReflect.Descriptor instead.
func (*OrderByInfo) Descriptor() ([]byte, []int) {
	return file_common_pub_common_proto_rawDescGZIP(), []int{1}
}

func (x *OrderByInfo) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *OrderByInfo) GetAsc() bool {
	if x != nil {
		return x.Asc
	}
	return false
}

type UserNameValidateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Valid         bool                   `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserNameValidateResponse) Reset() {
	*x = UserNameValidateResponse{}
	mi := &file_common_pub_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserNameValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNameValidateResponse) ProtoMessage() {}

func (x *UserNameValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_pub_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNameValidateResponse.ProtoReflect.Descriptor instead.
func (*UserNameValidateResponse) Descriptor() ([]byte, []int) {
	return file_common_pub_common_proto_rawDescGZIP(), []int{2}
}

func (x *UserNameValidateResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserNameValidateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserNameValidateResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_common_pub_common_proto protoreflect.FileDescriptor

var file_common_pub_common_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61,
	0x73, 0x63, 0x22, 0x5e, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x34, 0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x76, 0x36, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_pub_common_proto_rawDescOnce sync.Once
	file_common_pub_common_proto_rawDescData = file_common_pub_common_proto_rawDesc
)

func file_common_pub_common_proto_rawDescGZIP() []byte {
	file_common_pub_common_proto_rawDescOnce.Do(func() {
		file_common_pub_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pub_common_proto_rawDescData)
	})
	return file_common_pub_common_proto_rawDescData
}

var file_common_pub_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_pub_common_proto_goTypes = []any{
	(*ScanListRequest)(nil),          // 0: v6.services.pub.common.ScanListRequest
	(*OrderByInfo)(nil),              // 1: v6.services.pub.common.OrderByInfo
	(*UserNameValidateResponse)(nil), // 2: v6.services.pub.common.UserNameValidateResponse
}
var file_common_pub_common_proto_depIdxs = []int32{
	1, // 0: v6.services.pub.common.ScanListRequest.order_by:type_name -> v6.services.pub.common.OrderByInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_pub_common_proto_init() }
func file_common_pub_common_proto_init() {
	if File_common_pub_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_pub_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pub_common_proto_goTypes,
		DependencyIndexes: file_common_pub_common_proto_depIdxs,
		MessageInfos:      file_common_pub_common_proto_msgTypes,
	}.Build()
	File_common_pub_common_proto = out.File
	file_common_pub_common_proto_rawDesc = nil
	file_common_pub_common_proto_goTypes = nil
	file_common_pub_common_proto_depIdxs = nil
}
