// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: sftp/public_sftp.proto

package sftpconfig

import (
	common "github.com/city404/v6-public-rpc-proto/go/v6/common"
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

type SftpConfig struct {
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
	SshKey        string                 `protobuf:"bytes,15,opt,name=ssh_key,json=sshKey,proto3" json:"ssh_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SftpConfig) Reset() {
	*x = SftpConfig{}
	mi := &file_sftp_public_sftp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SftpConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SftpConfig) ProtoMessage() {}

func (x *SftpConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sftp_public_sftp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SftpConfig.ProtoReflect.Descriptor instead.
func (*SftpConfig) Descriptor() ([]byte, []int) {
	return file_sftp_public_sftp_proto_rawDescGZIP(), []int{0}
}

func (x *SftpConfig) GetUserIdentity() string {
	if x != nil {
		return x.UserIdentity
	}
	return ""
}

func (x *SftpConfig) GetEnableFlag() int32 {
	if x != nil {
		return x.EnableFlag
	}
	return 0
}

func (x *SftpConfig) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *SftpConfig) GetWriteFlag() int32 {
	if x != nil {
		return x.WriteFlag
	}
	return 0
}

func (x *SftpConfig) GetLockFlag() int32 {
	if x != nil {
		return x.LockFlag
	}
	return 0
}

func (x *SftpConfig) GetChunkedFlag() int32 {
	if x != nil {
		return x.ChunkedFlag
	}
	return 0
}

func (x *SftpConfig) GetPathStyle() int32 {
	if x != nil {
		return x.PathStyle
	}
	return 0
}

func (x *SftpConfig) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SftpConfig) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *SftpConfig) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SftpConfig) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SftpConfig) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SftpConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SftpConfig) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

func (x *SftpConfig) GetSshKey() string {
	if x != nil {
		return x.SshKey
	}
	return ""
}

var File_sftp_public_sftp_proto protoreflect.FileDescriptor

var file_sftp_public_sftp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x66, 0x74, 0x70, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x66,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x03, 0x0a, 0x0a, 0x53, 0x66, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x6b, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x65, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x61, 0x74, 0x68, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x73, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x73, 0x68, 0x4b, 0x65, 0x79, 0x32, 0xd4, 0x03, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x53, 0x66, 0x74,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x53, 0x66, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x66,
	0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x66, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x1b, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x53, 0x66, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x76, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x66, 0x74,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x66, 0x74, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1b, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x53, 0x66, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x53, 0x66, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1b, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x66,
	0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1b, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x66, 0x74, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x66,
	0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x30, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x34,
	0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x72, 0x70, 0x63,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x36, 0x2f, 0x73, 0x66, 0x74,
	0x70, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sftp_public_sftp_proto_rawDescOnce sync.Once
	file_sftp_public_sftp_proto_rawDescData = file_sftp_public_sftp_proto_rawDesc
)

func file_sftp_public_sftp_proto_rawDescGZIP() []byte {
	file_sftp_public_sftp_proto_rawDescOnce.Do(func() {
		file_sftp_public_sftp_proto_rawDescData = protoimpl.X.CompressGZIP(file_sftp_public_sftp_proto_rawDescData)
	})
	return file_sftp_public_sftp_proto_rawDescData
}

var file_sftp_public_sftp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sftp_public_sftp_proto_goTypes = []any{
	(*SftpConfig)(nil),                      // 0: v6.services.pub.SftpConfig
	(*common.UserNameValidateResponse)(nil), // 1: v6.services.pub.common.UserNameValidateResponse
}
var file_sftp_public_sftp_proto_depIdxs = []int32{
	0, // 0: v6.services.pub.PubSftpConfig.Get:input_type -> v6.services.pub.SftpConfig
	0, // 1: v6.services.pub.PubSftpConfig.Update:input_type -> v6.services.pub.SftpConfig
	0, // 2: v6.services.pub.PubSftpConfig.Enable:input_type -> v6.services.pub.SftpConfig
	0, // 3: v6.services.pub.PubSftpConfig.Disable:input_type -> v6.services.pub.SftpConfig
	0, // 4: v6.services.pub.PubSftpConfig.UpdateKeys:input_type -> v6.services.pub.SftpConfig
	0, // 5: v6.services.pub.PubSftpConfig.ValidateUserName:input_type -> v6.services.pub.SftpConfig
	0, // 6: v6.services.pub.PubSftpConfig.Get:output_type -> v6.services.pub.SftpConfig
	0, // 7: v6.services.pub.PubSftpConfig.Update:output_type -> v6.services.pub.SftpConfig
	0, // 8: v6.services.pub.PubSftpConfig.Enable:output_type -> v6.services.pub.SftpConfig
	0, // 9: v6.services.pub.PubSftpConfig.Disable:output_type -> v6.services.pub.SftpConfig
	0, // 10: v6.services.pub.PubSftpConfig.UpdateKeys:output_type -> v6.services.pub.SftpConfig
	1, // 11: v6.services.pub.PubSftpConfig.ValidateUserName:output_type -> v6.services.pub.common.UserNameValidateResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sftp_public_sftp_proto_init() }
func file_sftp_public_sftp_proto_init() {
	if File_sftp_public_sftp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sftp_public_sftp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sftp_public_sftp_proto_goTypes,
		DependencyIndexes: file_sftp_public_sftp_proto_depIdxs,
		MessageInfos:      file_sftp_public_sftp_proto_msgTypes,
	}.Build()
	File_sftp_public_sftp_proto = out.File
	file_sftp_public_sftp_proto_rawDesc = nil
	file_sftp_public_sftp_proto_goTypes = nil
	file_sftp_public_sftp_proto_depIdxs = nil
}
