// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: offline/public_user_offline.proto

package offline

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

type TaskParseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	File          string                 `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	Identity      string                 `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Addon         string                 `protobuf:"bytes,4,opt,name=addon,proto3" json:"addon,omitempty"`
	Flag          int32                  `protobuf:"varint,5,opt,name=flag,proto3" json:"flag,omitempty"`
	From          string                 `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskParseRequest) Reset() {
	*x = TaskParseRequest{}
	mi := &file_offline_public_user_offline_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskParseRequest) ProtoMessage() {}

func (x *TaskParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[0]
	if x != nil {
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

func (x *TaskParseRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *TaskParseRequest) GetAddon() string {
	if x != nil {
		return x.Addon
	}
	return ""
}

func (x *TaskParseRequest) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *TaskParseRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

type TaskMeta struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Identity string                 `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	Type     int32                  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status   int32                  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Locked   bool                   `protobuf:"varint,4,opt,name=locked,proto3" json:"locked,omitempty"`
	UpdateTs int64                  `protobuf:"varint,5,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	File     string                 `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	CreateTs int64                  `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Url      string                 `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Size     uint64                 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Name     string                 `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// string lock_by = 11;
	Code           int32  `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`
	Message        string `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
	Addon          string `protobuf:"bytes,14,opt,name=addon,proto3" json:"addon,omitempty"`
	Retries        uint64 `protobuf:"varint,15,opt,name=retries,proto3" json:"retries,omitempty"`
	Progress       uint64 `protobuf:"varint,16,opt,name=progress,proto3" json:"progress,omitempty"`
	BytesTotal     uint64 `protobuf:"varint,17,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty"`
	BytesProcessed uint64 `protobuf:"varint,18,opt,name=bytes_processed,json=bytesProcessed,proto3" json:"bytes_processed,omitempty"`
	Flag           int32  `protobuf:"varint,19,opt,name=flag,proto3" json:"flag,omitempty"` // int32 retry_status = 20;
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TaskMeta) Reset() {
	*x = TaskMeta{}
	mi := &file_offline_public_user_offline_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskMeta) ProtoMessage() {}

func (x *TaskMeta) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[1]
	if x != nil {
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

type TaskFile struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Identity string                 `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	// string file_identity = 2;
	// string task_identity = 3;
	Path     string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Status   int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	UpdateTs int64  `protobuf:"varint,6,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	CreateTs int64  `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Name     string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Size     uint64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// string content_identity = 10;
	// int32 code = 11;
	// string message = 12;
	BytesTotal uint64 `protobuf:"varint,13,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty"`
	// uint64 bytes_processed = 14;
	Index         int32 `protobuf:"varint,15,opt,name=index,proto3" json:"index,omitempty"`
	Directory     bool  `protobuf:"varint,16,opt,name=directory,proto3" json:"directory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskFile) Reset() {
	*x = TaskFile{}
	mi := &file_offline_public_user_offline_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFile) ProtoMessage() {}

func (x *TaskFile) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[2]
	if x != nil {
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

func (x *TaskFile) GetBytesTotal() uint64 {
	if x != nil {
		return x.BytesTotal
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *TaskMeta              `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	TaskFiles     []*TaskFile            `protobuf:"bytes,2,rep,name=task_files,json=taskFiles,proto3" json:"task_files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskParseResponse) Reset() {
	*x = TaskParseResponse{}
	mi := &file_offline_public_user_offline_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskParseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskParseResponse) ProtoMessage() {}

func (x *TaskParseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[3]
	if x != nil {
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

type OfflineTaskListRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	ListInfo      *common.ScanListRequest `protobuf:"bytes,3,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfflineTaskListRequest) Reset() {
	*x = OfflineTaskListRequest{}
	mi := &file_offline_public_user_offline_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfflineTaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineTaskListRequest) ProtoMessage() {}

func (x *OfflineTaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineTaskListRequest.ProtoReflect.Descriptor instead.
func (*OfflineTaskListRequest) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{4}
}

func (x *OfflineTaskListRequest) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type OfflineTaskListResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Tasks         []*UserTask             `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	ListInfo      *common.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfflineTaskListResponse) Reset() {
	*x = OfflineTaskListResponse{}
	mi := &file_offline_public_user_offline_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfflineTaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineTaskListResponse) ProtoMessage() {}

func (x *OfflineTaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineTaskListResponse.ProtoReflect.Descriptor instead.
func (*OfflineTaskListResponse) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{5}
}

func (x *OfflineTaskListResponse) GetTasks() []*UserTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *OfflineTaskListResponse) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type OfflineTaskDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Identity      []string               `protobuf:"bytes,1,rep,name=identity,proto3" json:"identity,omitempty"`
	DeleteFiles   bool                   `protobuf:"varint,2,opt,name=delete_files,json=deleteFiles,proto3" json:"delete_files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfflineTaskDeleteRequest) Reset() {
	*x = OfflineTaskDeleteRequest{}
	mi := &file_offline_public_user_offline_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfflineTaskDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineTaskDeleteRequest) ProtoMessage() {}

func (x *OfflineTaskDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineTaskDeleteRequest.ProtoReflect.Descriptor instead.
func (*OfflineTaskDeleteRequest) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{6}
}

func (x *OfflineTaskDeleteRequest) GetIdentity() []string {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *OfflineTaskDeleteRequest) GetDeleteFiles() bool {
	if x != nil {
		return x.DeleteFiles
	}
	return false
}

type OfflineTaskDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int64                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OfflineTaskDeleteResponse) Reset() {
	*x = OfflineTaskDeleteResponse{}
	mi := &file_offline_public_user_offline_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OfflineTaskDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineTaskDeleteResponse) ProtoMessage() {}

func (x *OfflineTaskDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineTaskDeleteResponse.ProtoReflect.Descriptor instead.
func (*OfflineTaskDeleteResponse) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{7}
}

func (x *OfflineTaskDeleteResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UserTask struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Identity string                 `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	Type     int32                  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status   int32                  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// string user_identity = 4;
	UpdateTs       int64    `protobuf:"varint,5,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	File           string   `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	CreateTs       int64    `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Url            string   `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Size           uint64   `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Name           string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	TaskIdentity   string   `protobuf:"bytes,11,opt,name=task_identity,json=taskIdentity,proto3" json:"task_identity,omitempty"`
	Code           int32    `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`
	Message        string   `protobuf:"bytes,13,opt,name=message,proto3" json:"message,omitempty"`
	Addon          string   `protobuf:"bytes,14,opt,name=addon,proto3" json:"addon,omitempty"`
	Progress       uint64   `protobuf:"varint,16,opt,name=progress,proto3" json:"progress,omitempty"`
	BytesTotal     uint64   `protobuf:"varint,17,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty"`
	BytesProcessed uint64   `protobuf:"varint,18,opt,name=bytes_processed,json=bytesProcessed,proto3" json:"bytes_processed,omitempty"`
	Flag           int32    `protobuf:"varint,19,opt,name=flag,proto3" json:"flag,omitempty"`
	SavePath       string   `protobuf:"bytes,20,opt,name=save_path,json=savePath,proto3" json:"save_path,omitempty"`
	Callbacks      []string `protobuf:"bytes,21,rep,name=callbacks,proto3" json:"callbacks,omitempty"`
	IgnoreFiles    []string `protobuf:"bytes,22,rep,name=ignore_files,json=ignoreFiles,proto3" json:"ignore_files,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UserTask) Reset() {
	*x = UserTask{}
	mi := &file_offline_public_user_offline_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTask) ProtoMessage() {}

func (x *UserTask) ProtoReflect() protoreflect.Message {
	mi := &file_offline_public_user_offline_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTask.ProtoReflect.Descriptor instead.
func (*UserTask) Descriptor() ([]byte, []int) {
	return file_offline_public_user_offline_proto_rawDescGZIP(), []int{8}
}

func (x *UserTask) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *UserTask) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *UserTask) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserTask) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

func (x *UserTask) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *UserTask) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *UserTask) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UserTask) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UserTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserTask) GetTaskIdentity() string {
	if x != nil {
		return x.TaskIdentity
	}
	return ""
}

func (x *UserTask) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserTask) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserTask) GetAddon() string {
	if x != nil {
		return x.Addon
	}
	return ""
}

func (x *UserTask) GetProgress() uint64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *UserTask) GetBytesTotal() uint64 {
	if x != nil {
		return x.BytesTotal
	}
	return 0
}

func (x *UserTask) GetBytesProcessed() uint64 {
	if x != nil {
		return x.BytesProcessed
	}
	return 0
}

func (x *UserTask) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *UserTask) GetSavePath() string {
	if x != nil {
		return x.SavePath
	}
	return ""
}

func (x *UserTask) GetCallbacks() []string {
	if x != nil {
		return x.Callbacks
	}
	return nil
}

func (x *UserTask) GetIgnoreFiles() []string {
	if x != nil {
		return x.IgnoreFiles
	}
	return nil
}

var File_offline_public_user_offline_proto protoreflect.FileDescriptor

const file_offline_public_user_offline_proto_rawDesc = "" +
	"\n" +
	"!offline/public_user_offline.proto\x12\x0fv6.services.pub\x1a\x17common/pub_common.proto\"\x92\x01\n" +
	"\x10TaskParseRequest\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12\x12\n" +
	"\x04file\x18\x02 \x01(\tR\x04file\x12\x1a\n" +
	"\bidentity\x18\x03 \x01(\tR\bidentity\x12\x14\n" +
	"\x05addon\x18\x04 \x01(\tR\x05addon\x12\x12\n" +
	"\x04flag\x18\x05 \x01(\x05R\x04flag\x12\x12\n" +
	"\x04from\x18\x06 \x01(\tR\x04from\"\xca\x03\n" +
	"\bTaskMeta\x12\x1a\n" +
	"\bidentity\x18\x01 \x01(\tR\bidentity\x12\x12\n" +
	"\x04type\x18\x02 \x01(\x05R\x04type\x12\x16\n" +
	"\x06status\x18\x03 \x01(\x05R\x06status\x12\x16\n" +
	"\x06locked\x18\x04 \x01(\bR\x06locked\x12\x1b\n" +
	"\tupdate_ts\x18\x05 \x01(\x03R\bupdateTs\x12\x12\n" +
	"\x04file\x18\x06 \x01(\tR\x04file\x12\x1b\n" +
	"\tcreate_ts\x18\a \x01(\x03R\bcreateTs\x12\x10\n" +
	"\x03url\x18\b \x01(\tR\x03url\x12\x12\n" +
	"\x04size\x18\t \x01(\x04R\x04size\x12\x12\n" +
	"\x04name\x18\n" +
	" \x01(\tR\x04name\x12\x12\n" +
	"\x04code\x18\f \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\r \x01(\tR\amessage\x12\x14\n" +
	"\x05addon\x18\x0e \x01(\tR\x05addon\x12\x18\n" +
	"\aretries\x18\x0f \x01(\x04R\aretries\x12\x1a\n" +
	"\bprogress\x18\x10 \x01(\x04R\bprogress\x12\x1f\n" +
	"\vbytes_total\x18\x11 \x01(\x04R\n" +
	"bytesTotal\x12'\n" +
	"\x0fbytes_processed\x18\x12 \x01(\x04R\x0ebytesProcessed\x12\x12\n" +
	"\x04flag\x18\x13 \x01(\x05R\x04flag\"\x89\x02\n" +
	"\bTaskFile\x12\x1a\n" +
	"\bidentity\x18\x01 \x01(\tR\bidentity\x12\x12\n" +
	"\x04path\x18\x04 \x01(\tR\x04path\x12\x16\n" +
	"\x06status\x18\x05 \x01(\x05R\x06status\x12\x1b\n" +
	"\tupdate_ts\x18\x06 \x01(\x03R\bupdateTs\x12\x1b\n" +
	"\tcreate_ts\x18\a \x01(\x03R\bcreateTs\x12\x12\n" +
	"\x04name\x18\b \x01(\tR\x04name\x12\x12\n" +
	"\x04size\x18\t \x01(\x04R\x04size\x12\x1f\n" +
	"\vbytes_total\x18\r \x01(\x04R\n" +
	"bytesTotal\x12\x14\n" +
	"\x05index\x18\x0f \x01(\x05R\x05index\x12\x1c\n" +
	"\tdirectory\x18\x10 \x01(\bR\tdirectory\"|\n" +
	"\x11TaskParseResponse\x12-\n" +
	"\x04meta\x18\x01 \x01(\v2\x19.v6.services.pub.TaskMetaR\x04meta\x128\n" +
	"\n" +
	"task_files\x18\x02 \x03(\v2\x19.v6.services.pub.TaskFileR\ttaskFiles\"^\n" +
	"\x16OfflineTaskListRequest\x12D\n" +
	"\tlist_info\x18\x03 \x01(\v2'.v6.services.pub.common.ScanListRequestR\blistInfo\"\x90\x01\n" +
	"\x17OfflineTaskListResponse\x12/\n" +
	"\x05tasks\x18\x01 \x03(\v2\x19.v6.services.pub.UserTaskR\x05tasks\x12D\n" +
	"\tlist_info\x18\x02 \x01(\v2'.v6.services.pub.common.ScanListRequestR\blistInfo\"Y\n" +
	"\x18OfflineTaskDeleteRequest\x12\x1a\n" +
	"\bidentity\x18\x01 \x03(\tR\bidentity\x12!\n" +
	"\fdelete_files\x18\x02 \x01(\bR\vdeleteFiles\"1\n" +
	"\x19OfflineTaskDeleteResponse\x12\x14\n" +
	"\x05count\x18\x01 \x01(\x03R\x05count\"\x9b\x04\n" +
	"\bUserTask\x12\x1a\n" +
	"\bidentity\x18\x01 \x01(\tR\bidentity\x12\x12\n" +
	"\x04type\x18\x02 \x01(\x05R\x04type\x12\x16\n" +
	"\x06status\x18\x03 \x01(\x05R\x06status\x12\x1b\n" +
	"\tupdate_ts\x18\x05 \x01(\x03R\bupdateTs\x12\x12\n" +
	"\x04file\x18\x06 \x01(\tR\x04file\x12\x1b\n" +
	"\tcreate_ts\x18\a \x01(\x03R\bcreateTs\x12\x10\n" +
	"\x03url\x18\b \x01(\tR\x03url\x12\x12\n" +
	"\x04size\x18\t \x01(\x04R\x04size\x12\x12\n" +
	"\x04name\x18\n" +
	" \x01(\tR\x04name\x12#\n" +
	"\rtask_identity\x18\v \x01(\tR\ftaskIdentity\x12\x12\n" +
	"\x04code\x18\f \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\r \x01(\tR\amessage\x12\x14\n" +
	"\x05addon\x18\x0e \x01(\tR\x05addon\x12\x1a\n" +
	"\bprogress\x18\x10 \x01(\x04R\bprogress\x12\x1f\n" +
	"\vbytes_total\x18\x11 \x01(\x04R\n" +
	"bytesTotal\x12'\n" +
	"\x0fbytes_processed\x18\x12 \x01(\x04R\x0ebytesProcessed\x12\x12\n" +
	"\x04flag\x18\x13 \x01(\x05R\x04flag\x12\x1b\n" +
	"\tsave_path\x18\x14 \x01(\tR\bsavePath\x12\x1c\n" +
	"\tcallbacks\x18\x15 \x03(\tR\tcallbacks\x12!\n" +
	"\fignore_files\x18\x16 \x03(\tR\vignoreFiles2\xe1\x02\n" +
	"\x0ePubOfflineTask\x12P\n" +
	"\x05Parse\x12!.v6.services.pub.TaskParseRequest\x1a\".v6.services.pub.TaskParseResponse\"\x00\x12=\n" +
	"\x03Add\x12\x19.v6.services.pub.UserTask\x1a\x19.v6.services.pub.UserTask\"\x00\x12[\n" +
	"\x04List\x12'.v6.services.pub.OfflineTaskListRequest\x1a(.v6.services.pub.OfflineTaskListResponse\"\x00\x12a\n" +
	"\x06Delete\x12).v6.services.pub.OfflineTaskDeleteRequest\x1a*.v6.services.pub.OfflineTaskDeleteResponse\"\x00B6Z4github.com/city404/v6-public-rpc-proto/go/v6/offlineb\x06proto3"

var (
	file_offline_public_user_offline_proto_rawDescOnce sync.Once
	file_offline_public_user_offline_proto_rawDescData []byte
)

func file_offline_public_user_offline_proto_rawDescGZIP() []byte {
	file_offline_public_user_offline_proto_rawDescOnce.Do(func() {
		file_offline_public_user_offline_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_offline_public_user_offline_proto_rawDesc), len(file_offline_public_user_offline_proto_rawDesc)))
	})
	return file_offline_public_user_offline_proto_rawDescData
}

var file_offline_public_user_offline_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_offline_public_user_offline_proto_goTypes = []any{
	(*TaskParseRequest)(nil),          // 0: v6.services.pub.TaskParseRequest
	(*TaskMeta)(nil),                  // 1: v6.services.pub.TaskMeta
	(*TaskFile)(nil),                  // 2: v6.services.pub.TaskFile
	(*TaskParseResponse)(nil),         // 3: v6.services.pub.TaskParseResponse
	(*OfflineTaskListRequest)(nil),    // 4: v6.services.pub.OfflineTaskListRequest
	(*OfflineTaskListResponse)(nil),   // 5: v6.services.pub.OfflineTaskListResponse
	(*OfflineTaskDeleteRequest)(nil),  // 6: v6.services.pub.OfflineTaskDeleteRequest
	(*OfflineTaskDeleteResponse)(nil), // 7: v6.services.pub.OfflineTaskDeleteResponse
	(*UserTask)(nil),                  // 8: v6.services.pub.UserTask
	(*common.ScanListRequest)(nil),    // 9: v6.services.pub.common.ScanListRequest
}
var file_offline_public_user_offline_proto_depIdxs = []int32{
	1, // 0: v6.services.pub.TaskParseResponse.meta:type_name -> v6.services.pub.TaskMeta
	2, // 1: v6.services.pub.TaskParseResponse.task_files:type_name -> v6.services.pub.TaskFile
	9, // 2: v6.services.pub.OfflineTaskListRequest.list_info:type_name -> v6.services.pub.common.ScanListRequest
	8, // 3: v6.services.pub.OfflineTaskListResponse.tasks:type_name -> v6.services.pub.UserTask
	9, // 4: v6.services.pub.OfflineTaskListResponse.list_info:type_name -> v6.services.pub.common.ScanListRequest
	0, // 5: v6.services.pub.PubOfflineTask.Parse:input_type -> v6.services.pub.TaskParseRequest
	8, // 6: v6.services.pub.PubOfflineTask.Add:input_type -> v6.services.pub.UserTask
	4, // 7: v6.services.pub.PubOfflineTask.List:input_type -> v6.services.pub.OfflineTaskListRequest
	6, // 8: v6.services.pub.PubOfflineTask.Delete:input_type -> v6.services.pub.OfflineTaskDeleteRequest
	3, // 9: v6.services.pub.PubOfflineTask.Parse:output_type -> v6.services.pub.TaskParseResponse
	8, // 10: v6.services.pub.PubOfflineTask.Add:output_type -> v6.services.pub.UserTask
	5, // 11: v6.services.pub.PubOfflineTask.List:output_type -> v6.services.pub.OfflineTaskListResponse
	7, // 12: v6.services.pub.PubOfflineTask.Delete:output_type -> v6.services.pub.OfflineTaskDeleteResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_offline_public_user_offline_proto_init() }
func file_offline_public_user_offline_proto_init() {
	if File_offline_public_user_offline_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_offline_public_user_offline_proto_rawDesc), len(file_offline_public_user_offline_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_offline_public_user_offline_proto_goTypes,
		DependencyIndexes: file_offline_public_user_offline_proto_depIdxs,
		MessageInfos:      file_offline_public_user_offline_proto_msgTypes,
	}.Build()
	File_offline_public_user_offline_proto = out.File
	file_offline_public_user_offline_proto_goTypes = nil
	file_offline_public_user_offline_proto_depIdxs = nil
}
