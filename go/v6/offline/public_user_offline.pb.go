// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: offline/public_user_offline.proto

package offline

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

type TaskParseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	File string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *TaskParseRequest) Reset() {
	*x = TaskParseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offline_public_user_offline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskParseRequest) ProtoMessage() {}

func (x *TaskParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskParseRequest.ProtoReflect.Descriptor instead.
func (*TaskParseRequest) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{0}
}

func (x *TaskParseRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *TaskParseRequest) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type TaskMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity       string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	Type           int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status         int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Locked         bool   `protobuf:"varint,4,opt,name=locked,proto3" json:"locked,omitempty"`
	UpdateTs       int64  `protobuf:"varint,5,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	File           string `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	CreateTs       int64  `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Url            string `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Size           uint64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Name           string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	LockBy         string `protobuf:"bytes,11,opt,name=lock_by,json=lockBy,proto3" json:"lock_by,omitempty"`
	Code           int32  `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`
	Message        string `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
	Addon          string `protobuf:"bytes,14,opt,name=addon,proto3" json:"addon,omitempty"`
	Retries        uint64 `protobuf:"varint,15,opt,name=retries,proto3" json:"retries,omitempty"`
	Progress       uint64 `protobuf:"varint,16,opt,name=progress,proto3" json:"progress,omitempty"`
	BytesTotal     uint64 `protobuf:"varint,17,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty"`
	BytesProcessed uint64 `protobuf:"varint,18,opt,name=bytes_processed,json=bytesProcessed,proto3" json:"bytes_processed,omitempty"`
	Flag           int32  `protobuf:"varint,19,opt,name=flag,proto3" json:"flag,omitempty"`
	RetryStatus    int32  `protobuf:"varint,20,opt,name=retry_status,json=retryStatus,proto3" json:"retry_status,omitempty"`
}

func (x *TaskMeta) Reset() {
	*x = TaskMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offline_public_user_offline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskMeta) ProtoMessage() {}

func (x *TaskMeta) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskMeta.ProtoReflect.Descriptor instead.
func (*TaskMeta) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{1}
}

func (x *TaskMeta) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *TaskMeta) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TaskMeta) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskMeta) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *TaskMeta) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

func (x *TaskMeta) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *TaskMeta) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *TaskMeta) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *TaskMeta) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TaskMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskMeta) GetLockBy() string {
	if x != nil {
		return x.LockBy
	}
	return ""
}

func (x *TaskMeta) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TaskMeta) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TaskMeta) GetAddon() string {
	if x != nil {
		return x.Addon
	}
	return ""
}

func (x *TaskMeta) GetRetries() uint64 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *TaskMeta) GetProgress() uint64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *TaskMeta) GetBytesTotal() uint64 {
	if x != nil {
		return x.BytesTotal
	}
	return 0
}

func (x *TaskMeta) GetBytesProcessed() uint64 {
	if x != nil {
		return x.BytesProcessed
	}
	return 0
}

func (x *TaskMeta) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *TaskMeta) GetRetryStatus() int32 {
	if x != nil {
		return x.RetryStatus
	}
	return 0
}

type TaskFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity        string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	FileIdentity    string `protobuf:"bytes,2,opt,name=file_identity,json=fileIdentity,proto3" json:"file_identity,omitempty"`
	TaskIdentity    string `protobuf:"bytes,3,opt,name=task_identity,json=taskIdentity,proto3" json:"task_identity,omitempty"`
	Path            string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status          int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	UpdateTs        int64  `protobuf:"varint,6,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	CreateTs        int64  `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Name            string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Size            uint64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	ContentIdentity string `protobuf:"bytes,10,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
	Code            int32  `protobuf:"varint,11,opt,name=code,proto3" json:"code,omitempty"`
	Message         string `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
	BytesTotal      uint64 `protobuf:"varint,13,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty"`
	BytesProcessed  uint64 `protobuf:"varint,14,opt,name=bytes_processed,json=bytesProcessed,proto3" json:"bytes_processed,omitempty"`
	Index           int32  `protobuf:"varint,15,opt,name=index,proto3" json:"index,omitempty"`
	Directory       bool   `protobuf:"varint,16,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *TaskFile) Reset() {
	*x = TaskFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offline_public_user_offline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFile) ProtoMessage() {}

func (x *TaskFile) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskFile.ProtoReflect.Descriptor instead.
func (*TaskFile) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{2}
}

func (x *TaskFile) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *TaskFile) GetFileIdentity() string {
	if x != nil {
		return x.FileIdentity
	}
	return ""
}

func (x *TaskFile) GetTaskIdentity() string {
	if x != nil {
		return x.TaskIdentity
	}
	return ""
}

func (x *TaskFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TaskFile) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskFile) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

func (x *TaskFile) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *TaskFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskFile) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *TaskFile) GetContentIdentity() string {
	if x != nil {
		return x.ContentIdentity
	}
	return ""
}

func (x *TaskFile) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TaskFile) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TaskFile) GetBytesTotal() uint64 {
	if x != nil {
		return x.BytesTotal
	}
	return 0
}

func (x *TaskFile) GetBytesProcessed() uint64 {
	if x != nil {
		return x.BytesProcessed
	}
	return 0
}

func (x *TaskFile) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *TaskFile) GetDirectory() bool {
	if x != nil {
		return x.Directory
	}
	return false
}

type TaskParseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *TaskMeta               `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	TaskFiles []*TaskFile             `protobuf:"bytes,2,rep,name=task_files,json=taskFiles,proto3" json:"task_files,omitempty"`
	ListInfo  *common.ScanListRequest `protobuf:"bytes,3,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

func (x *TaskParseResponse) Reset() {
	*x = TaskParseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_offline_public_user_offline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskParseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskParseResponse) ProtoMessage() {}

func (x *TaskParseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskParseResponse.ProtoReflect.Descriptor instead.
func (*TaskParseResponse) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{3}
}

func (x *TaskParseResponse) GetMeta() *TaskMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *TaskParseResponse) GetTaskFiles() []*TaskFile {
	if x != nil {
		return x.TaskFiles
	}
	return nil
}

func (x *TaskParseResponse) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

var File_offline_public_user_offline_proto protoreflect.FileDescriptor

var file_offline_public_user_offline_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x75, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a,
	0x10, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x86, 0x04, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x62, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x6f, 0x63, 0x6b, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xd5, 0x03, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xc2, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x73,
	0x6b, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a,
	0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x74, 0x61,
	0x73, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x62, 0x0a,
	0x0e, 0x50, 0x75, 0x62, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x50, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x21, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x69, 0x74, 0x79, 0x34, 0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x36, 0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_offline_public_user_offline_proto_rawDescOnce sync.Once
	file_offline_public_user_offline_proto_rawDescData = file_offline_public_user_offline_proto_rawDesc
)

func file_offline_public_user_offline_proto_rawDescGZIP() []byte {
	file_offline_public_user_offline_proto_rawDescOnce.Do(func() {
		file_offline_public_user_offline_proto_rawDescData = protoimpl.X.CompressGZIP(file_offline_public_user_offline_proto_rawDescData)
	})
	return file_offline_public_user_offline_proto_rawDescData
}

var file_offline_public_user_offline_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_offline_public_user_offline_proto_goTypes = []interface{}{
	(*TaskParseRequest)(nil),       // 0: v6.services.pub.TaskParseRequest
	(*TaskMeta)(nil),               // 1: v6.services.pub.TaskMeta
	(*TaskFile)(nil),               // 2: v6.services.pub.TaskFile
	(*TaskParseResponse)(nil),      // 3: v6.services.pub.TaskParseResponse
	(*common.ScanListRequest)(nil), // 4: v6.services.pub.common.ScanListRequest
}
var file_offline_public_user_offline_proto_depIdxs = []int32{
	1, // 0: v6.services.pub.TaskParseResponse.meta:type_name -> v6.services.pub.TaskMeta
	2, // 1: v6.services.pub.TaskParseResponse.task_files:type_name -> v6.services.pub.TaskFile
	4, // 2: v6.services.pub.TaskParseResponse.list_info:type_name -> v6.services.pub.common.ScanListRequest
	0, // 3: v6.services.pub.PubOfflineTask.Parse:input_type -> v6.services.pub.TaskParseRequest
	3, // 4: v6.services.pub.PubOfflineTask.Parse:output_type -> v6.services.pub.TaskParseResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_offline_public_user_offline_proto_init() }
func file_offline_public_user_offline_proto_init() {
	if File_offline_public_user_offline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_offline_public_user_offline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskParseRequest); i {
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
		file_offline_public_user_offline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskMeta); i {
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
		file_offline_public_user_offline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskFile); i {
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
		file_offline_public_user_offline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskParseResponse); i {
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
			RawDescriptor: file_offline_public_user_offline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_offline_public_user_offline_proto_goTypes,
		DependencyIndexes: file_offline_public_user_offline_proto_depIdxs,
		MessageInfos:      file_offline_public_user_offline_proto_msgTypes,
	}.Build()
	File_offline_public_user_offline_proto = out.File
	file_offline_public_user_offline_proto_rawDesc = nil
	file_offline_public_user_offline_proto_goTypes = nil
	file_offline_public_user_offline_proto_depIdxs = nil
}