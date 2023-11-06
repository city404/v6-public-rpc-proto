// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: userfile/user_file.proto

package userfile

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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity        string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Parent          string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Name            string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Path            string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Mime            string `protobuf:"bytes,6,opt,name=mime,proto3" json:"mime,omitempty"`
	Size            int64  `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Type            int32  `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt       int64  `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       int64  `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       int64  `protobuf:"varint,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Deleted         bool   `protobuf:"varint,12,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Dir             bool   `protobuf:"varint,13,opt,name=dir,proto3" json:"dir,omitempty"`
	Hidden          bool   `protobuf:"varint,14,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Locked          bool   `protobuf:"varint,15,opt,name=locked,proto3" json:"locked,omitempty"`
	Shared          bool   `protobuf:"varint,16,opt,name=shared,proto3" json:"shared,omitempty"`
	Starred         bool   `protobuf:"varint,17,opt,name=starred,proto3" json:"starred,omitempty"`
	Trashed         bool   `protobuf:"varint,18,opt,name=trashed,proto3" json:"trashed,omitempty"`
	LockedAt        int64  `protobuf:"varint,19,opt,name=locked_at,json=lockedAt,proto3" json:"locked_at,omitempty"`
	LockedBy        string `protobuf:"bytes,20,opt,name=locked_by,json=lockedBy,proto3" json:"locked_by,omitempty"`
	SharedAt        int64  `protobuf:"varint,21,opt,name=shared_at,json=sharedAt,proto3" json:"shared_at,omitempty"`
	Flag            int32  `protobuf:"varint,22,opt,name=flag,proto3" json:"flag,omitempty"`
	Unique          string `protobuf:"bytes,23,opt,name=unique,proto3" json:"unique,omitempty"`
	ContentIdentity string `protobuf:"bytes,24,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userfile_user_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_userfile_user_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_userfile_user_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *File) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *File) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *File) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *File) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *File) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *File) GetDir() bool {
	if x != nil {
		return x.Dir
	}
	return false
}

func (x *File) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *File) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *File) GetShared() bool {
	if x != nil {
		return x.Shared
	}
	return false
}

func (x *File) GetStarred() bool {
	if x != nil {
		return x.Starred
	}
	return false
}

func (x *File) GetTrashed() bool {
	if x != nil {
		return x.Trashed
	}
	return false
}

func (x *File) GetLockedAt() int64 {
	if x != nil {
		return x.LockedAt
	}
	return 0
}

func (x *File) GetLockedBy() string {
	if x != nil {
		return x.LockedBy
	}
	return ""
}

func (x *File) GetSharedAt() int64 {
	if x != nil {
		return x.SharedAt
	}
	return 0
}

func (x *File) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *File) GetUnique() string {
	if x != nil {
		return x.Unique
	}
	return ""
}

func (x *File) GetContentIdentity() string {
	if x != nil {
		return x.ContentIdentity
	}
	return ""
}

var File_userfile_user_file_proto protoreflect.FileDescriptor

var file_userfile_user_file_proto_rawDesc = []byte{
	0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x22, 0xd1, 0x04, 0x0a, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x73, 0x68,
	0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x72, 0x61, 0x73, 0x68, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32,
	0x7e, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x38,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a,
	0x15, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x15, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x15, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69,
	0x74, 0x79, 0x34, 0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d,
	0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x36, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userfile_user_file_proto_rawDescOnce sync.Once
	file_userfile_user_file_proto_rawDescData = file_userfile_user_file_proto_rawDesc
)

func file_userfile_user_file_proto_rawDescGZIP() []byte {
	file_userfile_user_file_proto_rawDescOnce.Do(func() {
		file_userfile_user_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_userfile_user_file_proto_rawDescData)
	})
	return file_userfile_user_file_proto_rawDescData
}

var file_userfile_user_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_userfile_user_file_proto_goTypes = []interface{}{
	(*File)(nil), // 0: v6.services.pub.File
}
var file_userfile_user_file_proto_depIdxs = []int32{
	0, // 0: v6.services.pub.PubUserFile.Create:input_type -> v6.services.pub.File
	0, // 1: v6.services.pub.PubUserFile.Get:input_type -> v6.services.pub.File
	0, // 2: v6.services.pub.PubUserFile.Create:output_type -> v6.services.pub.File
	0, // 3: v6.services.pub.PubUserFile.Get:output_type -> v6.services.pub.File
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userfile_user_file_proto_init() }
func file_userfile_user_file_proto_init() {
	if File_userfile_user_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userfile_user_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_userfile_user_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userfile_user_file_proto_goTypes,
		DependencyIndexes: file_userfile_user_file_proto_depIdxs,
		MessageInfos:      file_userfile_user_file_proto_msgTypes,
	}.Build()
	File_userfile_user_file_proto = out.File
	file_userfile_user_file_proto_rawDesc = nil
	file_userfile_user_file_proto_goTypes = nil
	file_userfile_user_file_proto_depIdxs = nil
}
