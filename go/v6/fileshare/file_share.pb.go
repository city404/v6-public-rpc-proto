// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.3
// source: share/file_share.proto

package fileshare

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

type FileShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity     string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	UserIdentity string   `protobuf:"bytes,2,opt,name=user_identity,json=userIdentity,proto3" json:"user_identity,omitempty"`
	Type         int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	MimeType     string   `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Name         string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	PathList     []string `protobuf:"bytes,6,rep,name=path_list,json=pathList,proto3" json:"path_list,omitempty"`
	HasPassword  bool     `protobuf:"varint,7,opt,name=has_password,json=hasPassword,proto3" json:"has_password,omitempty"`
	Password     string   `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	Lifetime     int64    `protobuf:"varint,9,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	SaveLimit    int64    `protobuf:"varint,10,opt,name=save_limit,json=saveLimit,proto3" json:"save_limit,omitempty"`
	SaveCount    int64    `protobuf:"varint,11,opt,name=save_count,json=saveCount,proto3" json:"save_count,omitempty"`
	Like         int64    `protobuf:"varint,12,opt,name=like,proto3" json:"like,omitempty"`
	Dislike      int64    `protobuf:"varint,13,opt,name=dislike,proto3" json:"dislike,omitempty"`
	CreateTs     int64    `protobuf:"varint,14,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	UpdateTs     int64    `protobuf:"varint,15,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	ExpireTs     int64    `protobuf:"varint,16,opt,name=expire_ts,json=expireTs,proto3" json:"expire_ts,omitempty"`
	FileSize     int64    `protobuf:"varint,17,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	SavePath     string   `protobuf:"bytes,18,opt,name=save_path,json=savePath,proto3" json:"save_path,omitempty"`
}

func (x *FileShare) Reset() {
	*x = FileShare{}
	mi := &file_share_file_share_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileShare) ProtoMessage() {}

func (x *FileShare) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileShare.ProtoReflect.Descriptor instead.
func (*FileShare) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{0}
}

func (x *FileShare) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *FileShare) GetUserIdentity() string {
	if x != nil {
		return x.UserIdentity
	}
	return ""
}

func (x *FileShare) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FileShare) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *FileShare) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileShare) GetPathList() []string {
	if x != nil {
		return x.PathList
	}
	return nil
}

func (x *FileShare) GetHasPassword() bool {
	if x != nil {
		return x.HasPassword
	}
	return false
}

func (x *FileShare) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *FileShare) GetLifetime() int64 {
	if x != nil {
		return x.Lifetime
	}
	return 0
}

func (x *FileShare) GetSaveLimit() int64 {
	if x != nil {
		return x.SaveLimit
	}
	return 0
}

func (x *FileShare) GetSaveCount() int64 {
	if x != nil {
		return x.SaveCount
	}
	return 0
}

func (x *FileShare) GetLike() int64 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *FileShare) GetDislike() int64 {
	if x != nil {
		return x.Dislike
	}
	return 0
}

func (x *FileShare) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *FileShare) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

func (x *FileShare) GetExpireTs() int64 {
	if x != nil {
		return x.ExpireTs
	}
	return 0
}

func (x *FileShare) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileShare) GetSavePath() string {
	if x != nil {
		return x.SavePath
	}
	return ""
}

type FileShareListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIdentity string                  `protobuf:"bytes,1,opt,name=user_identity,json=userIdentity,proto3" json:"user_identity,omitempty"`
	ListInfo     *common.ScanListRequest `protobuf:"bytes,3,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

func (x *FileShareListRequest) Reset() {
	*x = FileShareListRequest{}
	mi := &file_share_file_share_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileShareListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileShareListRequest) ProtoMessage() {}

func (x *FileShareListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileShareListRequest.ProtoReflect.Descriptor instead.
func (*FileShareListRequest) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{1}
}

func (x *FileShareListRequest) GetUserIdentity() string {
	if x != nil {
		return x.UserIdentity
	}
	return ""
}

func (x *FileShareListRequest) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type FileShareListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shares   []*FileShare            `protobuf:"bytes,1,rep,name=shares,proto3" json:"shares,omitempty"`
	ListInfo *common.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

func (x *FileShareListResponse) Reset() {
	*x = FileShareListResponse{}
	mi := &file_share_file_share_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileShareListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileShareListResponse) ProtoMessage() {}

func (x *FileShareListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileShareListResponse.ProtoReflect.Descriptor instead.
func (*FileShareListResponse) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{2}
}

func (x *FileShareListResponse) GetShares() []*FileShare {
	if x != nil {
		return x.Shares
	}
	return nil
}

func (x *FileShareListResponse) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type FileShareDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity []string `protobuf:"bytes,1,rep,name=identity,proto3" json:"identity,omitempty"`
}

func (x *FileShareDeleteRequest) Reset() {
	*x = FileShareDeleteRequest{}
	mi := &file_share_file_share_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileShareDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileShareDeleteRequest) ProtoMessage() {}

func (x *FileShareDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileShareDeleteRequest.ProtoReflect.Descriptor instead.
func (*FileShareDeleteRequest) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{3}
}

func (x *FileShareDeleteRequest) GetIdentity() []string {
	if x != nil {
		return x.Identity
	}
	return nil
}

type FileShareDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FileShareDeleteResponse) Reset() {
	*x = FileShareDeleteResponse{}
	mi := &file_share_file_share_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileShareDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileShareDeleteResponse) ProtoMessage() {}

func (x *FileShareDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileShareDeleteResponse.ProtoReflect.Descriptor instead.
func (*FileShareDeleteResponse) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{4}
}

func (x *FileShareDeleteResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Complaint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity              string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	UserIdentity          string `protobuf:"bytes,2,opt,name=user_identity,json=userIdentity,proto3" json:"user_identity,omitempty"`
	FileShareIdentity     string `protobuf:"bytes,3,opt,name=file_share_identity,json=fileShareIdentity,proto3" json:"file_share_identity,omitempty"`
	FileShareUserIdentity string `protobuf:"bytes,4,opt,name=file_share_user_identity,json=fileShareUserIdentity,proto3" json:"file_share_user_identity,omitempty"`
	Content               string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreateTs              int64  `protobuf:"varint,6,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	UpdateTs              int64  `protobuf:"varint,7,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	Status                int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Flag                  int32  `protobuf:"varint,9,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *Complaint) Reset() {
	*x = Complaint{}
	mi := &file_share_file_share_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Complaint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Complaint) ProtoMessage() {}

func (x *Complaint) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Complaint.ProtoReflect.Descriptor instead.
func (*Complaint) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{5}
}

func (x *Complaint) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Complaint) GetUserIdentity() string {
	if x != nil {
		return x.UserIdentity
	}
	return ""
}

func (x *Complaint) GetFileShareIdentity() string {
	if x != nil {
		return x.FileShareIdentity
	}
	return ""
}

func (x *Complaint) GetFileShareUserIdentity() string {
	if x != nil {
		return x.FileShareUserIdentity
	}
	return ""
}

func (x *Complaint) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Complaint) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *Complaint) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

func (x *Complaint) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Complaint) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

type ComplaintListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string user_identity = 1;
	ListInfo *common.ScanListRequest `protobuf:"bytes,3,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

func (x *ComplaintListRequest) Reset() {
	*x = ComplaintListRequest{}
	mi := &file_share_file_share_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplaintListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplaintListRequest) ProtoMessage() {}

func (x *ComplaintListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplaintListRequest.ProtoReflect.Descriptor instead.
func (*ComplaintListRequest) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{6}
}

func (x *ComplaintListRequest) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type ComplaintListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Complaints []*Complaint            `protobuf:"bytes,1,rep,name=complaints,proto3" json:"complaints,omitempty"`
	ListInfo   *common.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
}

func (x *ComplaintListResponse) Reset() {
	*x = ComplaintListResponse{}
	mi := &file_share_file_share_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplaintListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplaintListResponse) ProtoMessage() {}

func (x *ComplaintListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplaintListResponse.ProtoReflect.Descriptor instead.
func (*ComplaintListResponse) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{7}
}

func (x *ComplaintListResponse) GetComplaints() []*Complaint {
	if x != nil {
		return x.Complaints
	}
	return nil
}

func (x *ComplaintListResponse) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type ComplaintDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity []string `protobuf:"bytes,1,rep,name=identity,proto3" json:"identity,omitempty"`
}

func (x *ComplaintDeleteRequest) Reset() {
	*x = ComplaintDeleteRequest{}
	mi := &file_share_file_share_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplaintDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplaintDeleteRequest) ProtoMessage() {}

func (x *ComplaintDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplaintDeleteRequest.ProtoReflect.Descriptor instead.
func (*ComplaintDeleteRequest) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{8}
}

func (x *ComplaintDeleteRequest) GetIdentity() []string {
	if x != nil {
		return x.Identity
	}
	return nil
}

type ComplaintDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ComplaintDeleteResponse) Reset() {
	*x = ComplaintDeleteResponse{}
	mi := &file_share_file_share_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplaintDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplaintDeleteResponse) ProtoMessage() {}

func (x *ComplaintDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_share_file_share_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplaintDeleteResponse.ProtoReflect.Descriptor instead.
func (*ComplaintDeleteResponse) Descriptor() ([]byte, []int) {
	return file_share_file_share_proto_rawDescGZIP(), []int{9}
}

func (x *ComplaintDeleteResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_share_file_share_proto protoreflect.FileDescriptor

var file_share_file_share_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x74, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6b, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74,
	0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x61, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x81, 0x01, 0x0a, 0x14,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x91, 0x01, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x44, 0x0a,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x34, 0x0a, 0x16, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x17, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb5, 0x02, 0x0a, 0x09, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x18, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x22, 0x5c, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x99, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x34, 0x0a, 0x16,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x2f, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0xaa, 0x06, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x76, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x6b,
	0x65, 0x12, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x1a, 0x1a, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x44,
	0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x1a, 0x1a,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x27, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x1a, 0x1a, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x00,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x69, 0x74, 0x79, 0x34, 0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2d, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x36,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_share_file_share_proto_rawDescOnce sync.Once
	file_share_file_share_proto_rawDescData = file_share_file_share_proto_rawDesc
)

func file_share_file_share_proto_rawDescGZIP() []byte {
	file_share_file_share_proto_rawDescOnce.Do(func() {
		file_share_file_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_share_file_share_proto_rawDescData)
	})
	return file_share_file_share_proto_rawDescData
}

var file_share_file_share_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_share_file_share_proto_goTypes = []any{
	(*FileShare)(nil),               // 0: v6.services.pub.FileShare
	(*FileShareListRequest)(nil),    // 1: v6.services.pub.FileShareListRequest
	(*FileShareListResponse)(nil),   // 2: v6.services.pub.FileShareListResponse
	(*FileShareDeleteRequest)(nil),  // 3: v6.services.pub.FileShareDeleteRequest
	(*FileShareDeleteResponse)(nil), // 4: v6.services.pub.FileShareDeleteResponse
	(*Complaint)(nil),               // 5: v6.services.pub.Complaint
	(*ComplaintListRequest)(nil),    // 6: v6.services.pub.ComplaintListRequest
	(*ComplaintListResponse)(nil),   // 7: v6.services.pub.ComplaintListResponse
	(*ComplaintDeleteRequest)(nil),  // 8: v6.services.pub.ComplaintDeleteRequest
	(*ComplaintDeleteResponse)(nil), // 9: v6.services.pub.ComplaintDeleteResponse
	(*common.ScanListRequest)(nil),  // 10: v6.services.pub.common.ScanListRequest
}
var file_share_file_share_proto_depIdxs = []int32{
	10, // 0: v6.services.pub.FileShareListRequest.list_info:type_name -> v6.services.pub.common.ScanListRequest
	0,  // 1: v6.services.pub.FileShareListResponse.shares:type_name -> v6.services.pub.FileShare
	10, // 2: v6.services.pub.FileShareListResponse.list_info:type_name -> v6.services.pub.common.ScanListRequest
	10, // 3: v6.services.pub.ComplaintListRequest.list_info:type_name -> v6.services.pub.common.ScanListRequest
	5,  // 4: v6.services.pub.ComplaintListResponse.complaints:type_name -> v6.services.pub.Complaint
	10, // 5: v6.services.pub.ComplaintListResponse.list_info:type_name -> v6.services.pub.common.ScanListRequest
	0,  // 6: v6.services.pub.PubFileShare.Create:input_type -> v6.services.pub.FileShare
	0,  // 7: v6.services.pub.PubFileShare.Get:input_type -> v6.services.pub.FileShare
	0,  // 8: v6.services.pub.PubFileShare.Like:input_type -> v6.services.pub.FileShare
	0,  // 9: v6.services.pub.PubFileShare.Dislike:input_type -> v6.services.pub.FileShare
	5,  // 10: v6.services.pub.PubFileShare.ComplaintShare:input_type -> v6.services.pub.Complaint
	1,  // 11: v6.services.pub.PubFileShare.List:input_type -> v6.services.pub.FileShareListRequest
	6,  // 12: v6.services.pub.PubFileShare.ListComplaint:input_type -> v6.services.pub.ComplaintListRequest
	3,  // 13: v6.services.pub.PubFileShare.Delete:input_type -> v6.services.pub.FileShareDeleteRequest
	8,  // 14: v6.services.pub.PubFileShare.DeleteComplaint:input_type -> v6.services.pub.ComplaintDeleteRequest
	0,  // 15: v6.services.pub.PubFileShare.Save:input_type -> v6.services.pub.FileShare
	0,  // 16: v6.services.pub.PubFileShare.Create:output_type -> v6.services.pub.FileShare
	0,  // 17: v6.services.pub.PubFileShare.Get:output_type -> v6.services.pub.FileShare
	0,  // 18: v6.services.pub.PubFileShare.Like:output_type -> v6.services.pub.FileShare
	0,  // 19: v6.services.pub.PubFileShare.Dislike:output_type -> v6.services.pub.FileShare
	5,  // 20: v6.services.pub.PubFileShare.ComplaintShare:output_type -> v6.services.pub.Complaint
	2,  // 21: v6.services.pub.PubFileShare.List:output_type -> v6.services.pub.FileShareListResponse
	7,  // 22: v6.services.pub.PubFileShare.ListComplaint:output_type -> v6.services.pub.ComplaintListResponse
	4,  // 23: v6.services.pub.PubFileShare.Delete:output_type -> v6.services.pub.FileShareDeleteResponse
	9,  // 24: v6.services.pub.PubFileShare.DeleteComplaint:output_type -> v6.services.pub.ComplaintDeleteResponse
	0,  // 25: v6.services.pub.PubFileShare.Save:output_type -> v6.services.pub.FileShare
	16, // [16:26] is the sub-list for method output_type
	6,  // [6:16] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_share_file_share_proto_init() }
func file_share_file_share_proto_init() {
	if File_share_file_share_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_share_file_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_share_file_share_proto_goTypes,
		DependencyIndexes: file_share_file_share_proto_depIdxs,
		MessageInfos:      file_share_file_share_proto_msgTypes,
	}.Build()
	File_share_file_share_proto = out.File
	file_share_file_share_proto_rawDesc = nil
	file_share_file_share_proto_goTypes = nil
	file_share_file_share_proto_depIdxs = nil
}
