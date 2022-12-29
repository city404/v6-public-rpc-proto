// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: file/file_svc.proto

package file

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

type Slice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	Type     int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status   int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	UpdateTs int64  `protobuf:"varint,5,opt,name=update_ts,json=updateTs,proto3" json:"update_ts,omitempty"`
	Peer     string `protobuf:"bytes,6,opt,name=peer,proto3" json:"peer,omitempty"`
	CreateTs int64  `protobuf:"varint,7,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Server   string `protobuf:"bytes,8,opt,name=server,proto3" json:"server,omitempty"`
	NeedData bool   `protobuf:"varint,9,opt,name=need_data,json=needData,proto3" json:"need_data,omitempty"`
	Data     []byte `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Slice) Reset() {
	*x = Slice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slice) ProtoMessage() {}

func (x *Slice) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slice.ProtoReflect.Descriptor instead.
func (*Slice) Descriptor() ([]byte, []int) {
	return file_file_file_svc_proto_rawDescGZIP(), []int{0}
}

func (x *Slice) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Slice) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Slice) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Slice) GetUpdateTs() int64 {
	if x != nil {
		return x.UpdateTs
	}
	return 0
}

func (x *Slice) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

func (x *Slice) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *Slice) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *Slice) GetNeedData() bool {
	if x != nil {
		return x.NeedData
	}
	return false
}

func (x *Slice) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"` // identity
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_file_file_svc_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Meta) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FastLookup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity    string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	WcsEtag     string `protobuf:"bytes,2,opt,name=wcs_etag,json=wcsEtag,proto3" json:"wcs_etag,omitempty"`
	Sha1        string `protobuf:"bytes,3,opt,name=sha1,proto3" json:"sha1,omitempty"`
	HeadSha1    string `protobuf:"bytes,4,opt,name=head_sha1,json=headSha1,proto3" json:"head_sha1,omitempty"`
	FileSize    int64  `protobuf:"varint,5,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	SpecialSha1 string `protobuf:"bytes,6,opt,name=special_sha1,json=specialSha1,proto3" json:"special_sha1,omitempty"`
	Name        string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FastLookup) Reset() {
	*x = FastLookup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FastLookup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FastLookup) ProtoMessage() {}

func (x *FastLookup) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FastLookup.ProtoReflect.Descriptor instead.
func (*FastLookup) Descriptor() ([]byte, []int) {
	return file_file_file_svc_proto_rawDescGZIP(), []int{2}
}

func (x *FastLookup) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *FastLookup) GetWcsEtag() string {
	if x != nil {
		return x.WcsEtag
	}
	return ""
}

func (x *FastLookup) GetSha1() string {
	if x != nil {
		return x.Sha1
	}
	return ""
}

func (x *FastLookup) GetHeadSha1() string {
	if x != nil {
		return x.HeadSha1
	}
	return ""
}

func (x *FastLookup) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FastLookup) GetSpecialSha1() string {
	if x != nil {
		return x.SpecialSha1
	}
	return ""
}

func (x *FastLookup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FastLookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32       `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Info *FastLookup `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *FastLookupRequest) Reset() {
	*x = FastLookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FastLookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FastLookupRequest) ProtoMessage() {}

func (x *FastLookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FastLookupRequest.ProtoReflect.Descriptor instead.
func (*FastLookupRequest) Descriptor() ([]byte, []int) {
	return file_file_file_svc_proto_rawDescGZIP(), []int{3}
}

func (x *FastLookupRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FastLookupRequest) GetInfo() *FastLookup {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_file_file_svc_proto protoreflect.FileDescriptor

var file_file_file_svc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x76, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x22, 0xe6, 0x01, 0x0a, 0x05, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x6e, 0x65, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x36, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x46, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x63, 0x73, 0x5f, 0x65, 0x74, 0x61, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x63, 0x73, 0x45, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x68, 0x61, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x61,
	0x31, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x53, 0x68, 0x61, 0x31, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x31, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x58, 0x0a, 0x11, 0x46, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x9f, 0x03, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x12, 0x16, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x76, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x15, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x15, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x0d, 0x41, 0x64, 0x64, 0x46, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x1b,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x46, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x1a, 0x1b, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x61,
	0x73, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x22, 0x2e, 0x76, 0x36,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x46, 0x61,
	0x73, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x46, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x22, 0x00, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74,
	0x79, 0x34, 0x30, 0x34, 0x2f, 0x76, 0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x72,
	0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x36, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_file_svc_proto_rawDescOnce sync.Once
	file_file_file_svc_proto_rawDescData = file_file_file_svc_proto_rawDesc
)

func file_file_file_svc_proto_rawDescGZIP() []byte {
	file_file_file_svc_proto_rawDescOnce.Do(func() {
		file_file_file_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_file_svc_proto_rawDescData)
	})
	return file_file_file_svc_proto_rawDescData
}

var file_file_file_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_file_svc_proto_goTypes = []interface{}{
	(*Slice)(nil),             // 0: v6.services.pub.Slice
	(*Meta)(nil),              // 1: v6.services.pub.Meta
	(*FastLookup)(nil),        // 2: v6.services.pub.FastLookup
	(*FastLookupRequest)(nil), // 3: v6.services.pub.FastLookupRequest
}
var file_file_file_svc_proto_depIdxs = []int32{
	2, // 0: v6.services.pub.FastLookupRequest.info:type_name -> v6.services.pub.FastLookup
	0, // 1: v6.services.pub.File.GetSlice:input_type -> v6.services.pub.Slice
	0, // 2: v6.services.pub.File.CreateSlice:input_type -> v6.services.pub.Slice
	1, // 3: v6.services.pub.File.GetMeta:input_type -> v6.services.pub.Meta
	1, // 4: v6.services.pub.File.CreateMeta:input_type -> v6.services.pub.Meta
	2, // 5: v6.services.pub.File.AddFastLookup:input_type -> v6.services.pub.FastLookup
	3, // 6: v6.services.pub.File.GetFastLookup:input_type -> v6.services.pub.FastLookupRequest
	0, // 7: v6.services.pub.File.GetSlice:output_type -> v6.services.pub.Slice
	0, // 8: v6.services.pub.File.CreateSlice:output_type -> v6.services.pub.Slice
	1, // 9: v6.services.pub.File.GetMeta:output_type -> v6.services.pub.Meta
	1, // 10: v6.services.pub.File.CreateMeta:output_type -> v6.services.pub.Meta
	2, // 11: v6.services.pub.File.AddFastLookup:output_type -> v6.services.pub.FastLookup
	2, // 12: v6.services.pub.File.GetFastLookup:output_type -> v6.services.pub.FastLookup
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_file_file_svc_proto_init() }
func file_file_file_svc_proto_init() {
	if File_file_file_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_file_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slice); i {
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
		file_file_file_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_file_file_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FastLookup); i {
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
		file_file_file_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FastLookupRequest); i {
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
			RawDescriptor: file_file_file_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_file_svc_proto_goTypes,
		DependencyIndexes: file_file_file_svc_proto_depIdxs,
		MessageInfos:      file_file_file_svc_proto_msgTypes,
	}.Build()
	File_file_file_svc_proto = out.File
	file_file_file_svc_proto_rawDesc = nil
	file_file_file_svc_proto_goTypes = nil
	file_file_file_svc_proto_depIdxs = nil
}
