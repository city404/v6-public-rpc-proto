// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: webdav/public_webdav.proto

package webdavconfig

import (
	common "github.com/city404/v6-public-rpc-proto/go/v6/common"
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

type DavConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserIdentity  string                 `protobuf:"bytes,1,opt,name=user_identity,json=userIdentity,proto3" json:"user_identity,omitempty"`
	EnableFlag    int32                  `protobuf:"varint,2,opt,name=enable_flag,json=enableFlag,proto3" json:"enable_flag,omitempty"`
	Root          string                 `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	WriteFlag     int32                  `protobuf:"varint,4,opt,name=write_flag,json=writeFlag,proto3" json:"write_flag,omitempty"`
	LockFlag      int32                  `protobuf:"varint,5,opt,name=lock_flag,json=lockFlag,proto3" json:"lock_flag,omitempty"`
	ChunkedFlag   int32                  `protobuf:"varint,6,opt,name=chunked_flag,json=chunkedFlag,proto3" json:"chunked_flag,omitempty"`
	PathStyle     int32                  `protobuf:"varint,7,opt,name=path_style,json=pathStyle,proto3" json:"path_style,omitempty"`
	Status        int32                  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Prefix        string                 `protobuf:"bytes,9,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Code          int32                  `protobuf:"varint,10,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,11,opt,name=message,proto3" json:"message,omitempty"`
	Username      string                 `protobuf:"bytes,12,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,13,opt,name=password,proto3" json:"password,omitempty"`
	UpdateTs      int64                  `protobuf:"varint,14,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DavConfig) Reset() {
	*x = DavConfig{}
	mi := &file_webdav_public_webdav_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DavConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DavConfig) ProtoMessage() {}

func (x *DavConfig) ProtoReflect() protoreflect.Message {
	mi := &file_webdav_public_webdav_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DavConfig.ProtoReflect.Descriptor instead.
func (*DavConfig) Descriptor() ([]byte, []int) {
	return file_webdav_public_webdav_proto_rawDescGZIP(), []int{0}
}

func (x *DavConfig) GetUserIdentity() string {
	if x != nil {
		return x.UserIdentity
	}
	return ""
}

func (x *DavConfig) GetEnableFlag() int32 {
	if x != nil {
		return x.EnableFlag
	}
	return 0
}

func (x *DavConfig) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *DavConfig) GetWriteFlag() int32 {
	if x != nil {
		return x.WriteFlag
	}
	return 0
}

func (x *DavConfig) GetLockFlag() int32 {
	if x != nil {
		return x.LockFlag
	}
	return 0
}

func (x *DavConfig) GetChunkedFlag() int32 {
	if x != nil {
		return x.ChunkedFlag
	}
	return 0
}

func (x *DavConfig) GetPathStyle() int32 {
	if x != nil {
		return x.PathStyle
	}
	return 0
}

func (x *DavConfig) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DavConfig) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *DavConfig) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DavConfig) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DavConfig) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DavConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DavConfig) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

var File_webdav_public_webdav_proto protoreflect.FileDescriptor

var file_webdav_public_webdav_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x1a, 0x17, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x09, 0x44, 0x61, 0x76, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x73, 0x32,
	0x80, 0x03, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x44, 0x61,
	0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1a, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x44, 0x61,
	0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x62,
	0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x44, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x30,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x34, 0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x76, 0x36, 0x2f, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_webdav_public_webdav_proto_rawDescOnce sync.Once
	file_webdav_public_webdav_proto_rawDescData []byte
)

func file_webdav_public_webdav_proto_rawDescGZIP() []byte {
	file_webdav_public_webdav_proto_rawDescOnce.Do(func() {
		file_webdav_public_webdav_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_webdav_public_webdav_proto_rawDesc), len(file_webdav_public_webdav_proto_rawDesc)))
	})
	return file_webdav_public_webdav_proto_rawDescData
}

var file_webdav_public_webdav_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_webdav_public_webdav_proto_goTypes = []any{
	(*DavConfig)(nil),                       // 0: v6.services.pub.DavConfig
	(*common.UserNameValidateResponse)(nil), // 1: v6.services.pub.common.UserNameValidateResponse
}
var file_webdav_public_webdav_proto_depIdxs = []int32{
	0, // 0: v6.services.pub.PubDavConfig.Get:input_type -> v6.services.pub.DavConfig
	0, // 1: v6.services.pub.PubDavConfig.Update:input_type -> v6.services.pub.DavConfig
	0, // 2: v6.services.pub.PubDavConfig.Enable:input_type -> v6.services.pub.DavConfig
	0, // 3: v6.services.pub.PubDavConfig.Disable:input_type -> v6.services.pub.DavConfig
	0, // 4: v6.services.pub.PubDavConfig.ValidateUserName:input_type -> v6.services.pub.DavConfig
	0, // 5: v6.services.pub.PubDavConfig.Get:output_type -> v6.services.pub.DavConfig
	0, // 6: v6.services.pub.PubDavConfig.Update:output_type -> v6.services.pub.DavConfig
	0, // 7: v6.services.pub.PubDavConfig.Enable:output_type -> v6.services.pub.DavConfig
	0, // 8: v6.services.pub.PubDavConfig.Disable:output_type -> v6.services.pub.DavConfig
	1, // 9: v6.services.pub.PubDavConfig.ValidateUserName:output_type -> v6.services.pub.common.UserNameValidateResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_webdav_public_webdav_proto_init() }
func file_webdav_public_webdav_proto_init() {
	if File_webdav_public_webdav_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_webdav_public_webdav_proto_rawDesc), len(file_webdav_public_webdav_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webdav_public_webdav_proto_goTypes,
		DependencyIndexes: file_webdav_public_webdav_proto_depIdxs,
		MessageInfos:      file_webdav_public_webdav_proto_msgTypes,
	}.Build()
	File_webdav_public_webdav_proto = out.File
	file_webdav_public_webdav_proto_goTypes = nil
	file_webdav_public_webdav_proto_depIdxs = nil
}
