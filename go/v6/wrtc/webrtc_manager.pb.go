// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: wrtc/webrtc_manager.proto

package wrtc

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

type StartDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdentity  string `protobuf:"bytes,1,opt,name=client_identity,json=clientIdentity,proto3" json:"client_identity,omitempty"`
	ContentIdentity string `protobuf:"bytes,2,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
	PeerLimit       int64  `protobuf:"varint,3,opt,name=peer_limit,json=peerLimit,proto3" json:"peer_limit,omitempty"`
}

func (x *StartDownloadRequest) Reset() {
	*x = StartDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDownloadRequest) ProtoMessage() {}

func (x *StartDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDownloadRequest.ProtoReflect.Descriptor instead.
func (*StartDownloadRequest) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{0}
}

func (x *StartDownloadRequest) GetClientIdentity() string {
	if x != nil {
		return x.ClientIdentity
	}
	return ""
}

func (x *StartDownloadRequest) GetContentIdentity() string {
	if x != nil {
		return x.ContentIdentity
	}
	return ""
}

func (x *StartDownloadRequest) GetPeerLimit() int64 {
	if x != nil {
		return x.PeerLimit
	}
	return 0
}

type RtcPeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdentity  string `protobuf:"bytes,1,opt,name=client_identity,json=clientIdentity,proto3" json:"client_identity,omitempty"`
	ContentIdentity string `protobuf:"bytes,2,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
	Sdp             string `protobuf:"bytes,3,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *RtcPeerInfo) Reset() {
	*x = RtcPeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtcPeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtcPeerInfo) ProtoMessage() {}

func (x *RtcPeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtcPeerInfo.ProtoReflect.Descriptor instead.
func (*RtcPeerInfo) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RtcPeerInfo) GetClientIdentity() string {
	if x != nil {
		return x.ClientIdentity
	}
	return ""
}

func (x *RtcPeerInfo) GetContentIdentity() string {
	if x != nil {
		return x.ContentIdentity
	}
	return ""
}

func (x *RtcPeerInfo) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

type StartDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentIdentity string         `protobuf:"bytes,1,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
	Peers           []*RtcPeerInfo `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	BackendPeers    []*RtcPeerInfo `protobuf:"bytes,3,rep,name=backend_peers,json=backendPeers,proto3" json:"backend_peers,omitempty"`
}

func (x *StartDownloadResponse) Reset() {
	*x = StartDownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDownloadResponse) ProtoMessage() {}

func (x *StartDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDownloadResponse.ProtoReflect.Descriptor instead.
func (*StartDownloadResponse) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{2}
}

func (x *StartDownloadResponse) GetContentIdentity() string {
	if x != nil {
		return x.ContentIdentity
	}
	return ""
}

func (x *StartDownloadResponse) GetPeers() []*RtcPeerInfo {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *StartDownloadResponse) GetBackendPeers() []*RtcPeerInfo {
	if x != nil {
		return x.BackendPeers
	}
	return nil
}

type StopDownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdentity  string `protobuf:"bytes,1,opt,name=client_identity,json=clientIdentity,proto3" json:"client_identity,omitempty"`
	ContentIdentity string `protobuf:"bytes,2,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
}

func (x *StopDownloadRequest) Reset() {
	*x = StopDownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopDownloadRequest) ProtoMessage() {}

func (x *StopDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopDownloadRequest.ProtoReflect.Descriptor instead.
func (*StopDownloadRequest) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{3}
}

func (x *StopDownloadRequest) GetClientIdentity() string {
	if x != nil {
		return x.ClientIdentity
	}
	return ""
}

func (x *StopDownloadRequest) GetContentIdentity() string {
	if x != nil {
		return x.ContentIdentity
	}
	return ""
}

type StopDownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentIdentity string `protobuf:"bytes,1,opt,name=content_identity,json=contentIdentity,proto3" json:"content_identity,omitempty"`
}

func (x *StopDownloadResponse) Reset() {
	*x = StopDownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopDownloadResponse) ProtoMessage() {}

func (x *StopDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopDownloadResponse.ProtoReflect.Descriptor instead.
func (*StopDownloadResponse) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{4}
}

func (x *StopDownloadResponse) GetContentIdentity() string {
	if x != nil {
		return x.ContentIdentity
	}
	return ""
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdentity string `protobuf:"bytes,1,opt,name=client_identity,json=clientIdentity,proto3" json:"client_identity,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{5}
}

func (x *DisconnectRequest) GetClientIdentity() string {
	if x != nil {
		return x.ClientIdentity
	}
	return ""
}

type DisconnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdentity string `protobuf:"bytes,1,opt,name=client_identity,json=clientIdentity,proto3" json:"client_identity,omitempty"`
}

func (x *DisconnectResponse) Reset() {
	*x = DisconnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectResponse) ProtoMessage() {}

func (x *DisconnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectResponse.ProtoReflect.Descriptor instead.
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{6}
}

func (x *DisconnectResponse) GetClientIdentity() string {
	if x != nil {
		return x.ClientIdentity
	}
	return ""
}

type SendPeerIceCandidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdentity string `protobuf:"bytes,1,opt,name=client_identity,json=clientIdentity,proto3" json:"client_identity,omitempty"`
	Candidate      string `protobuf:"bytes,3,opt,name=candidate,proto3" json:"candidate,omitempty"`
}

func (x *SendPeerIceCandidateRequest) Reset() {
	*x = SendPeerIceCandidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPeerIceCandidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPeerIceCandidateRequest) ProtoMessage() {}

func (x *SendPeerIceCandidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPeerIceCandidateRequest.ProtoReflect.Descriptor instead.
func (*SendPeerIceCandidateRequest) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{7}
}

func (x *SendPeerIceCandidateRequest) GetClientIdentity() string {
	if x != nil {
		return x.ClientIdentity
	}
	return ""
}

func (x *SendPeerIceCandidateRequest) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

type SendPeerIceCandidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIdentity string `protobuf:"bytes,1,opt,name=client_identity,json=clientIdentity,proto3" json:"client_identity,omitempty"`
}

func (x *SendPeerIceCandidateResponse) Reset() {
	*x = SendPeerIceCandidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPeerIceCandidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPeerIceCandidateResponse) ProtoMessage() {}

func (x *SendPeerIceCandidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPeerIceCandidateResponse.ProtoReflect.Descriptor instead.
func (*SendPeerIceCandidateResponse) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{8}
}

func (x *SendPeerIceCandidateResponse) GetClientIdentity() string {
	if x != nil {
		return x.ClientIdentity
	}
	return ""
}

type ReceiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReceiveRequest) Reset() {
	*x = ReceiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveRequest) ProtoMessage() {}

func (x *ReceiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveRequest.ProtoReflect.Descriptor instead.
func (*ReceiveRequest) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{9}
}

func (x *ReceiveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReceiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReceiveResponse) Reset() {
	*x = ReceiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wrtc_webrtc_manager_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveResponse) ProtoMessage() {}

func (x *ReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wrtc_webrtc_manager_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveResponse.ProtoReflect.Descriptor instead.
func (*ReceiveResponse) Descriptor() ([]byte, []int) {
	return file_wrtc_webrtc_manager_proto_rawDescGZIP(), []int{10}
}

func (x *ReceiveResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_wrtc_webrtc_manager_proto protoreflect.FileDescriptor

var file_wrtc_webrtc_manager_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x72, 0x74, 0x63, 0x2f, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x22, 0x89, 0x01, 0x0a,
	0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x65, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x73, 0x0a, 0x0b, 0x52, 0x74, 0x63, 0x50,
	0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x64, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x22, 0xb9, 0x01,
	0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x32, 0x70, 0x2e, 0x52, 0x74, 0x63, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e,
	0x52, 0x74, 0x63, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73, 0x22, 0x69, 0x0a, 0x13, 0x53, 0x74, 0x6f,
	0x70, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x41, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3c, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3d, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x64, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x22, 0x47, 0x0a, 0x1c, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x65, 0x65, 0x72, 0x49, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xf8, 0x03, 0x0a,
	0x13, 0x57, 0x65, 0x62, 0x52, 0x54, 0x43, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x24, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75,
	0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x65, 0x65, 0x72, 0x49, 0x63, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x65, 0x65, 0x72, 0x49,
	0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x12, 0x1f, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x32, 0x70, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x32, 0x70, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x34, 0x30, 0x34, 0x2f, 0x76, 0x36,
	0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x36, 0x2f, 0x77, 0x72, 0x74, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wrtc_webrtc_manager_proto_rawDescOnce sync.Once
	file_wrtc_webrtc_manager_proto_rawDescData = file_wrtc_webrtc_manager_proto_rawDesc
)

func file_wrtc_webrtc_manager_proto_rawDescGZIP() []byte {
	file_wrtc_webrtc_manager_proto_rawDescOnce.Do(func() {
		file_wrtc_webrtc_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_wrtc_webrtc_manager_proto_rawDescData)
	})
	return file_wrtc_webrtc_manager_proto_rawDescData
}

var file_wrtc_webrtc_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_wrtc_webrtc_manager_proto_goTypes = []interface{}{
	(*StartDownloadRequest)(nil),         // 0: v6.services.p2p.StartDownloadRequest
	(*RtcPeerInfo)(nil),                  // 1: v6.services.p2p.RtcPeerInfo
	(*StartDownloadResponse)(nil),        // 2: v6.services.p2p.StartDownloadResponse
	(*StopDownloadRequest)(nil),          // 3: v6.services.p2p.StopDownloadRequest
	(*StopDownloadResponse)(nil),         // 4: v6.services.p2p.StopDownloadResponse
	(*DisconnectRequest)(nil),            // 5: v6.services.p2p.DisconnectRequest
	(*DisconnectResponse)(nil),           // 6: v6.services.p2p.DisconnectResponse
	(*SendPeerIceCandidateRequest)(nil),  // 7: v6.services.p2p.SendPeerIceCandidateRequest
	(*SendPeerIceCandidateResponse)(nil), // 8: v6.services.p2p.SendPeerIceCandidateResponse
	(*ReceiveRequest)(nil),               // 9: v6.services.p2p.ReceiveRequest
	(*ReceiveResponse)(nil),              // 10: v6.services.p2p.ReceiveResponse
}
var file_wrtc_webrtc_manager_proto_depIdxs = []int32{
	1,  // 0: v6.services.p2p.StartDownloadResponse.peers:type_name -> v6.services.p2p.RtcPeerInfo
	1,  // 1: v6.services.p2p.StartDownloadResponse.backend_peers:type_name -> v6.services.p2p.RtcPeerInfo
	0,  // 2: v6.services.p2p.WebRTCManageService.StartDownload:input_type -> v6.services.p2p.StartDownloadRequest
	3,  // 3: v6.services.p2p.WebRTCManageService.StopDownload:input_type -> v6.services.p2p.StopDownloadRequest
	5,  // 4: v6.services.p2p.WebRTCManageService.Disconnect:input_type -> v6.services.p2p.DisconnectRequest
	7,  // 5: v6.services.p2p.WebRTCManageService.SendPeerIceCandidate:input_type -> v6.services.p2p.SendPeerIceCandidateRequest
	9,  // 6: v6.services.p2p.WebRTCManageService.Receive:input_type -> v6.services.p2p.ReceiveRequest
	2,  // 7: v6.services.p2p.WebRTCManageService.StartDownload:output_type -> v6.services.p2p.StartDownloadResponse
	4,  // 8: v6.services.p2p.WebRTCManageService.StopDownload:output_type -> v6.services.p2p.StopDownloadResponse
	6,  // 9: v6.services.p2p.WebRTCManageService.Disconnect:output_type -> v6.services.p2p.DisconnectResponse
	8,  // 10: v6.services.p2p.WebRTCManageService.SendPeerIceCandidate:output_type -> v6.services.p2p.SendPeerIceCandidateResponse
	10, // 11: v6.services.p2p.WebRTCManageService.Receive:output_type -> v6.services.p2p.ReceiveResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_wrtc_webrtc_manager_proto_init() }
func file_wrtc_webrtc_manager_proto_init() {
	if File_wrtc_webrtc_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wrtc_webrtc_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDownloadRequest); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RtcPeerInfo); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDownloadResponse); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopDownloadRequest); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopDownloadResponse); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectResponse); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPeerIceCandidateRequest); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPeerIceCandidateResponse); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveRequest); i {
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
		file_wrtc_webrtc_manager_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveResponse); i {
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
			RawDescriptor: file_wrtc_webrtc_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wrtc_webrtc_manager_proto_goTypes,
		DependencyIndexes: file_wrtc_webrtc_manager_proto_depIdxs,
		MessageInfos:      file_wrtc_webrtc_manager_proto_msgTypes,
	}.Build()
	File_wrtc_webrtc_manager_proto = out.File
	file_wrtc_webrtc_manager_proto_rawDesc = nil
	file_wrtc_webrtc_manager_proto_goTypes = nil
	file_wrtc_webrtc_manager_proto_depIdxs = nil
}
