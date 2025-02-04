// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: ticket/tickets_svc.proto

package ticket

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

type Ticket struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Identity      string                 `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Subject       string                 `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Category      string                 `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	SubCategory   string                 `protobuf:"bytes,6,opt,name=sub_category,json=subCategory,proto3" json:"sub_category,omitempty"`
	Channel       string                 `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	Description   string                 `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	CreateTs      int64                  `protobuf:"varint,9,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	Attachments   []string               `protobuf:"bytes,10,rep,name=attachments,proto3" json:"attachments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	mi := &file_ticket_tickets_svc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_tickets_svc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_ticket_tickets_svc_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Ticket) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Ticket) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Ticket) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Ticket) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Ticket) GetSubCategory() string {
	if x != nil {
		return x.SubCategory
	}
	return ""
}

func (x *Ticket) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Ticket) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Ticket) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

func (x *Ticket) GetAttachments() []string {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type TicketListRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	ListInfo      *common.ScanListRequest `protobuf:"bytes,3,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TicketListRequest) Reset() {
	*x = TicketListRequest{}
	mi := &file_ticket_tickets_svc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketListRequest) ProtoMessage() {}

func (x *TicketListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_tickets_svc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketListRequest.ProtoReflect.Descriptor instead.
func (*TicketListRequest) Descriptor() ([]byte, []int) {
	return file_ticket_tickets_svc_proto_rawDescGZIP(), []int{1}
}

func (x *TicketListRequest) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type TicketListResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Tickets       []*Ticket               `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
	ListInfo      *common.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TicketListResponse) Reset() {
	*x = TicketListResponse{}
	mi := &file_ticket_tickets_svc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketListResponse) ProtoMessage() {}

func (x *TicketListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_tickets_svc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketListResponse.ProtoReflect.Descriptor instead.
func (*TicketListResponse) Descriptor() ([]byte, []int) {
	return file_ticket_tickets_svc_proto_rawDescGZIP(), []int{2}
}

func (x *TicketListResponse) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

func (x *TicketListResponse) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type ContactInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FirstName     string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContactInfo) Reset() {
	*x = ContactInfo{}
	mi := &file_ticket_tickets_svc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContactInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfo) ProtoMessage() {}

func (x *ContactInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_tickets_svc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactInfo.ProtoReflect.Descriptor instead.
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return file_ticket_tickets_svc_proto_rawDescGZIP(), []int{3}
}

func (x *ContactInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *ContactInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ContactInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ConversationListRequest struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	TicketIdentity string                  `protobuf:"bytes,1,opt,name=ticket_identity,json=ticketIdentity,proto3" json:"ticket_identity,omitempty"`
	ListInfo       *common.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ConversationListRequest) Reset() {
	*x = ConversationListRequest{}
	mi := &file_ticket_tickets_svc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationListRequest) ProtoMessage() {}

func (x *ConversationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_tickets_svc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationListRequest.ProtoReflect.Descriptor instead.
func (*ConversationListRequest) Descriptor() ([]byte, []int) {
	return file_ticket_tickets_svc_proto_rawDescGZIP(), []int{4}
}

func (x *ConversationListRequest) GetTicketIdentity() string {
	if x != nil {
		return x.TicketIdentity
	}
	return ""
}

func (x *ConversationListRequest) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type ConversationListResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Conversations []*Conversation         `protobuf:"bytes,1,rep,name=conversations,proto3" json:"conversations,omitempty"`
	ListInfo      *common.ScanListRequest `protobuf:"bytes,2,opt,name=list_info,json=listInfo,proto3" json:"list_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversationListResponse) Reset() {
	*x = ConversationListResponse{}
	mi := &file_ticket_tickets_svc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationListResponse) ProtoMessage() {}

func (x *ConversationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_tickets_svc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationListResponse.ProtoReflect.Descriptor instead.
func (*ConversationListResponse) Descriptor() ([]byte, []int) {
	return file_ticket_tickets_svc_proto_rawDescGZIP(), []int{5}
}

func (x *ConversationListResponse) GetConversations() []*Conversation {
	if x != nil {
		return x.Conversations
	}
	return nil
}

func (x *ConversationListResponse) GetListInfo() *common.ScanListRequest {
	if x != nil {
		return x.ListInfo
	}
	return nil
}

type Conversation struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	IsDescriptionThread bool                   `protobuf:"varint,1,opt,name=is_description_thread,json=isDescriptionThread,proto3" json:"is_description_thread,omitempty"`
	CanReply            bool                   `protobuf:"varint,2,opt,name=can_reply,json=canReply,proto3" json:"can_reply,omitempty"`
	Bcc                 string                 `protobuf:"bytes,3,opt,name=bcc,proto3" json:"bcc,omitempty"`
	Channel             string                 `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	Type                string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	LastRatingIconUrl   string                 `protobuf:"bytes,6,opt,name=last_rating_icon_url,json=lastRatingIconUrl,proto3" json:"last_rating_icon_url,omitempty"`
	ImpersonatedUser    string                 `protobuf:"bytes,7,opt,name=impersonated_user,json=impersonatedUser,proto3" json:"impersonated_user,omitempty"`
	Identity            string                 `protobuf:"bytes,8,opt,name=identity,proto3" json:"identity,omitempty"`
	ContentType         string                 `protobuf:"bytes,9,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Direction           string                 `protobuf:"bytes,10,opt,name=direction,proto3" json:"direction,omitempty"`
	Summary             string                 `protobuf:"bytes,11,opt,name=summary,proto3" json:"summary,omitempty"`
	Cc                  string                 `protobuf:"bytes,12,opt,name=cc,proto3" json:"cc,omitempty"`
	KeyWords            string                 `protobuf:"bytes,13,opt,name=key_words,json=keyWords,proto3" json:"key_words,omitempty"`
	Visibility          string                 `protobuf:"bytes,14,opt,name=visibility,proto3" json:"visibility,omitempty"`
	Content             string                 `protobuf:"bytes,15,opt,name=content,proto3" json:"content,omitempty"`
	EncodedContent      string                 `protobuf:"bytes,16,opt,name=encoded_content,json=encodedContent,proto3" json:"encoded_content,omitempty"`
	IsForward           bool                   `protobuf:"varint,17,opt,name=is_forward,json=isForward,proto3" json:"is_forward,omitempty"`
	HasAttach           bool                   `protobuf:"varint,18,opt,name=has_attach,json=hasAttach,proto3" json:"has_attach,omitempty"`
	AttachmentCount     string                 `protobuf:"bytes,19,opt,name=attachment_count,json=attachmentCount,proto3" json:"attachment_count,omitempty"`
	To                  string                 `protobuf:"bytes,20,opt,name=to,proto3" json:"to,omitempty"`
	FromEmailAddress    string                 `protobuf:"bytes,21,opt,name=from_email_address,json=fromEmailAddress,proto3" json:"from_email_address,omitempty"`
	Status              string                 `protobuf:"bytes,22,opt,name=status,proto3" json:"status,omitempty"`
	CreateTs            int64                  `protobuf:"varint,23,opt,name=create_ts,json=createTs,proto3" json:"create_ts,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Conversation) Reset() {
	*x = Conversation{}
	mi := &file_ticket_tickets_svc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Conversation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conversation) ProtoMessage() {}

func (x *Conversation) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_tickets_svc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conversation.ProtoReflect.Descriptor instead.
func (*Conversation) Descriptor() ([]byte, []int) {
	return file_ticket_tickets_svc_proto_rawDescGZIP(), []int{6}
}

func (x *Conversation) GetIsDescriptionThread() bool {
	if x != nil {
		return x.IsDescriptionThread
	}
	return false
}

func (x *Conversation) GetCanReply() bool {
	if x != nil {
		return x.CanReply
	}
	return false
}

func (x *Conversation) GetBcc() string {
	if x != nil {
		return x.Bcc
	}
	return ""
}

func (x *Conversation) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Conversation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Conversation) GetLastRatingIconUrl() string {
	if x != nil {
		return x.LastRatingIconUrl
	}
	return ""
}

func (x *Conversation) GetImpersonatedUser() string {
	if x != nil {
		return x.ImpersonatedUser
	}
	return ""
}

func (x *Conversation) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Conversation) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Conversation) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *Conversation) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Conversation) GetCc() string {
	if x != nil {
		return x.Cc
	}
	return ""
}

func (x *Conversation) GetKeyWords() string {
	if x != nil {
		return x.KeyWords
	}
	return ""
}

func (x *Conversation) GetVisibility() string {
	if x != nil {
		return x.Visibility
	}
	return ""
}

func (x *Conversation) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Conversation) GetEncodedContent() string {
	if x != nil {
		return x.EncodedContent
	}
	return ""
}

func (x *Conversation) GetIsForward() bool {
	if x != nil {
		return x.IsForward
	}
	return false
}

func (x *Conversation) GetHasAttach() bool {
	if x != nil {
		return x.HasAttach
	}
	return false
}

func (x *Conversation) GetAttachmentCount() string {
	if x != nil {
		return x.AttachmentCount
	}
	return ""
}

func (x *Conversation) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Conversation) GetFromEmailAddress() string {
	if x != nil {
		return x.FromEmailAddress
	}
	return ""
}

func (x *Conversation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Conversation) GetCreateTs() int64 {
	if x != nil {
		return x.CreateTs
	}
	return 0
}

var File_ticket_tickets_svc_proto protoreflect.FileDescriptor

var file_ticket_tickets_svc_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x1a, 0x17, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x59, 0x0a, 0x11, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x12, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x75, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x88, 0x01,
	0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53,
	0x63, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08,
	0x6c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76,
	0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x09, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0xe0, 0x05, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x69, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x62, 0x63, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x63, 0x6f, 0x6e,
	0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x63, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x57,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x2c, 0x0a, 0x12, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x72,
	0x6f, 0x6d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x73, 0x32, 0xa2, 0x04, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x1a, 0x17, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x76, 0x36, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1c, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1c, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x1c, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x69, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x05, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x2e, 0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x17, 0x2e,
	0x76, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x34, 0x30, 0x34, 0x2f, 0x76,
	0x36, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x36, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_ticket_tickets_svc_proto_rawDescOnce sync.Once
	file_ticket_tickets_svc_proto_rawDescData []byte
)

func file_ticket_tickets_svc_proto_rawDescGZIP() []byte {
	file_ticket_tickets_svc_proto_rawDescOnce.Do(func() {
		file_ticket_tickets_svc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ticket_tickets_svc_proto_rawDesc), len(file_ticket_tickets_svc_proto_rawDesc)))
	})
	return file_ticket_tickets_svc_proto_rawDescData
}

var file_ticket_tickets_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ticket_tickets_svc_proto_goTypes = []any{
	(*Ticket)(nil),                   // 0: v6.services.pub.Ticket
	(*TicketListRequest)(nil),        // 1: v6.services.pub.TicketListRequest
	(*TicketListResponse)(nil),       // 2: v6.services.pub.TicketListResponse
	(*ContactInfo)(nil),              // 3: v6.services.pub.ContactInfo
	(*ConversationListRequest)(nil),  // 4: v6.services.pub.ConversationListRequest
	(*ConversationListResponse)(nil), // 5: v6.services.pub.ConversationListResponse
	(*Conversation)(nil),             // 6: v6.services.pub.Conversation
	(*common.ScanListRequest)(nil),   // 7: v6.services.pub.common.ScanListRequest
}
var file_ticket_tickets_svc_proto_depIdxs = []int32{
	7,  // 0: v6.services.pub.TicketListRequest.list_info:type_name -> v6.services.pub.common.ScanListRequest
	0,  // 1: v6.services.pub.TicketListResponse.tickets:type_name -> v6.services.pub.Ticket
	7,  // 2: v6.services.pub.TicketListResponse.list_info:type_name -> v6.services.pub.common.ScanListRequest
	7,  // 3: v6.services.pub.ConversationListRequest.list_info:type_name -> v6.services.pub.common.ScanListRequest
	6,  // 4: v6.services.pub.ConversationListResponse.conversations:type_name -> v6.services.pub.Conversation
	7,  // 5: v6.services.pub.ConversationListResponse.list_info:type_name -> v6.services.pub.common.ScanListRequest
	0,  // 6: v6.services.pub.PubTicket.Get:input_type -> v6.services.pub.Ticket
	1,  // 7: v6.services.pub.PubTicket.List:input_type -> v6.services.pub.TicketListRequest
	0,  // 8: v6.services.pub.PubTicket.Create:input_type -> v6.services.pub.Ticket
	3,  // 9: v6.services.pub.PubTicket.GetContactInfo:input_type -> v6.services.pub.ContactInfo
	3,  // 10: v6.services.pub.PubTicket.UpdateContactInfo:input_type -> v6.services.pub.ContactInfo
	4,  // 11: v6.services.pub.PubTicket.ListConversation:input_type -> v6.services.pub.ConversationListRequest
	0,  // 12: v6.services.pub.PubTicket.Reply:input_type -> v6.services.pub.Ticket
	0,  // 13: v6.services.pub.PubTicket.Get:output_type -> v6.services.pub.Ticket
	2,  // 14: v6.services.pub.PubTicket.List:output_type -> v6.services.pub.TicketListResponse
	0,  // 15: v6.services.pub.PubTicket.Create:output_type -> v6.services.pub.Ticket
	3,  // 16: v6.services.pub.PubTicket.GetContactInfo:output_type -> v6.services.pub.ContactInfo
	3,  // 17: v6.services.pub.PubTicket.UpdateContactInfo:output_type -> v6.services.pub.ContactInfo
	5,  // 18: v6.services.pub.PubTicket.ListConversation:output_type -> v6.services.pub.ConversationListResponse
	0,  // 19: v6.services.pub.PubTicket.Reply:output_type -> v6.services.pub.Ticket
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_ticket_tickets_svc_proto_init() }
func file_ticket_tickets_svc_proto_init() {
	if File_ticket_tickets_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ticket_tickets_svc_proto_rawDesc), len(file_ticket_tickets_svc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_tickets_svc_proto_goTypes,
		DependencyIndexes: file_ticket_tickets_svc_proto_depIdxs,
		MessageInfos:      file_ticket_tickets_svc_proto_msgTypes,
	}.Build()
	File_ticket_tickets_svc_proto = out.File
	file_ticket_tickets_svc_proto_goTypes = nil
	file_ticket_tickets_svc_proto_depIdxs = nil
}
