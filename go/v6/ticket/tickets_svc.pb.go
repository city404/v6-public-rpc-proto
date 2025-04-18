// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	EntityType    string                 `protobuf:"bytes,11,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	ContentType   string                 `protobuf:"bytes,12,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Status        string                 `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	Read          bool                   `protobuf:"varint,14,opt,name=read,proto3" json:"read,omitempty"`
	CommentCount  int64                  `protobuf:"varint,15,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	ThreadCount   int64                  `protobuf:"varint,16,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
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

func (x *Ticket) GetEntityType() string {
	if x != nil {
		return x.EntityType
	}
	return ""
}

func (x *Ticket) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Ticket) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Ticket) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *Ticket) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Ticket) GetThreadCount() int64 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
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
	UserThread          bool                   `protobuf:"varint,24,opt,name=user_thread,json=userThread,proto3" json:"user_thread,omitempty"`
	CleanContent        string                 `protobuf:"bytes,25,opt,name=clean_content,json=cleanContent,proto3" json:"clean_content,omitempty"`
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

func (x *Conversation) GetUserThread() bool {
	if x != nil {
		return x.UserThread
	}
	return false
}

func (x *Conversation) GetCleanContent() string {
	if x != nil {
		return x.CleanContent
	}
	return ""
}

var File_ticket_tickets_svc_proto protoreflect.FileDescriptor

const file_ticket_tickets_svc_proto_rawDesc = "" +
	"\n" +
	"\x18ticket/tickets_svc.proto\x12\x0fv6.services.pub\x1a\x17common/pub_common.proto\"\xdc\x03\n" +
	"\x06Ticket\x12\x1a\n" +
	"\bidentity\x18\x01 \x01(\tR\bidentity\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x18\n" +
	"\asubject\x18\x04 \x01(\tR\asubject\x12\x1a\n" +
	"\bcategory\x18\x05 \x01(\tR\bcategory\x12!\n" +
	"\fsub_category\x18\x06 \x01(\tR\vsubCategory\x12\x18\n" +
	"\achannel\x18\a \x01(\tR\achannel\x12 \n" +
	"\vdescription\x18\b \x01(\tR\vdescription\x12\x1b\n" +
	"\tcreate_ts\x18\t \x01(\x03R\bcreateTs\x12 \n" +
	"\vattachments\x18\n" +
	" \x03(\tR\vattachments\x12\x1f\n" +
	"\ventity_type\x18\v \x01(\tR\n" +
	"entityType\x12!\n" +
	"\fcontent_type\x18\f \x01(\tR\vcontentType\x12\x16\n" +
	"\x06status\x18\r \x01(\tR\x06status\x12\x12\n" +
	"\x04read\x18\x0e \x01(\bR\x04read\x12#\n" +
	"\rcomment_count\x18\x0f \x01(\x03R\fcommentCount\x12!\n" +
	"\fthread_count\x18\x10 \x01(\x03R\vthreadCount\"Y\n" +
	"\x11TicketListRequest\x12D\n" +
	"\tlist_info\x18\x03 \x01(\v2'.v6.services.pub.common.ScanListRequestR\blistInfo\"\x8d\x01\n" +
	"\x12TicketListResponse\x121\n" +
	"\atickets\x18\x01 \x03(\v2\x17.v6.services.pub.TicketR\atickets\x12D\n" +
	"\tlist_info\x18\x02 \x01(\v2'.v6.services.pub.common.ScanListRequestR\blistInfo\"u\n" +
	"\vContactInfo\x12\x1d\n" +
	"\n" +
	"first_name\x18\x01 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x02 \x01(\tR\blastName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\"\x88\x01\n" +
	"\x17ConversationListRequest\x12'\n" +
	"\x0fticket_identity\x18\x01 \x01(\tR\x0eticketIdentity\x12D\n" +
	"\tlist_info\x18\x02 \x01(\v2'.v6.services.pub.common.ScanListRequestR\blistInfo\"\xa5\x01\n" +
	"\x18ConversationListResponse\x12C\n" +
	"\rconversations\x18\x01 \x03(\v2\x1d.v6.services.pub.ConversationR\rconversations\x12D\n" +
	"\tlist_info\x18\x02 \x01(\v2'.v6.services.pub.common.ScanListRequestR\blistInfo\"\xa6\x06\n" +
	"\fConversation\x122\n" +
	"\x15is_description_thread\x18\x01 \x01(\bR\x13isDescriptionThread\x12\x1b\n" +
	"\tcan_reply\x18\x02 \x01(\bR\bcanReply\x12\x10\n" +
	"\x03bcc\x18\x03 \x01(\tR\x03bcc\x12\x18\n" +
	"\achannel\x18\x04 \x01(\tR\achannel\x12\x12\n" +
	"\x04type\x18\x05 \x01(\tR\x04type\x12/\n" +
	"\x14last_rating_icon_url\x18\x06 \x01(\tR\x11lastRatingIconUrl\x12+\n" +
	"\x11impersonated_user\x18\a \x01(\tR\x10impersonatedUser\x12\x1a\n" +
	"\bidentity\x18\b \x01(\tR\bidentity\x12!\n" +
	"\fcontent_type\x18\t \x01(\tR\vcontentType\x12\x1c\n" +
	"\tdirection\x18\n" +
	" \x01(\tR\tdirection\x12\x18\n" +
	"\asummary\x18\v \x01(\tR\asummary\x12\x0e\n" +
	"\x02cc\x18\f \x01(\tR\x02cc\x12\x1b\n" +
	"\tkey_words\x18\r \x01(\tR\bkeyWords\x12\x1e\n" +
	"\n" +
	"visibility\x18\x0e \x01(\tR\n" +
	"visibility\x12\x18\n" +
	"\acontent\x18\x0f \x01(\tR\acontent\x12'\n" +
	"\x0fencoded_content\x18\x10 \x01(\tR\x0eencodedContent\x12\x1d\n" +
	"\n" +
	"is_forward\x18\x11 \x01(\bR\tisForward\x12\x1d\n" +
	"\n" +
	"has_attach\x18\x12 \x01(\bR\thasAttach\x12)\n" +
	"\x10attachment_count\x18\x13 \x01(\tR\x0fattachmentCount\x12\x0e\n" +
	"\x02to\x18\x14 \x01(\tR\x02to\x12,\n" +
	"\x12from_email_address\x18\x15 \x01(\tR\x10fromEmailAddress\x12\x16\n" +
	"\x06status\x18\x16 \x01(\tR\x06status\x12\x1b\n" +
	"\tcreate_ts\x18\x17 \x01(\x03R\bcreateTs\x12\x1f\n" +
	"\vuser_thread\x18\x18 \x01(\bR\n" +
	"userThread\x12#\n" +
	"\rclean_content\x18\x19 \x01(\tR\fcleanContent2\xcb\x05\n" +
	"\tPubTicket\x129\n" +
	"\x03Get\x12\x17.v6.services.pub.Ticket\x1a\x17.v6.services.pub.Ticket\"\x00\x12Q\n" +
	"\x04List\x12\".v6.services.pub.TicketListRequest\x1a#.v6.services.pub.TicketListResponse\"\x00\x12<\n" +
	"\x06Create\x12\x17.v6.services.pub.Ticket\x1a\x17.v6.services.pub.Ticket\"\x00\x12N\n" +
	"\x0eGetContactInfo\x12\x1c.v6.services.pub.ContactInfo\x1a\x1c.v6.services.pub.ContactInfo\"\x00\x12Q\n" +
	"\x11UpdateContactInfo\x12\x1c.v6.services.pub.ContactInfo\x1a\x1c.v6.services.pub.ContactInfo\"\x00\x12i\n" +
	"\x10ListConversation\x12(.v6.services.pub.ConversationListRequest\x1a).v6.services.pub.ConversationListResponse\"\x00\x12=\n" +
	"\aComment\x12\x17.v6.services.pub.Ticket\x1a\x17.v6.services.pub.Ticket\"\x00\x12R\n" +
	"\x06Delete\x12\".v6.services.pub.common.StringList\x1a\".v6.services.pub.common.StringList\"\x00\x12Q\n" +
	"\x05Close\x12\".v6.services.pub.common.StringList\x1a\".v6.services.pub.common.StringList\"\x00B5Z3github.com/city404/v6-public-rpc-proto/go/v6/ticketb\x06proto3"

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
	(*common.StringList)(nil),        // 8: v6.services.pub.common.StringList
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
	0,  // 12: v6.services.pub.PubTicket.Comment:input_type -> v6.services.pub.Ticket
	8,  // 13: v6.services.pub.PubTicket.Delete:input_type -> v6.services.pub.common.StringList
	8,  // 14: v6.services.pub.PubTicket.Close:input_type -> v6.services.pub.common.StringList
	0,  // 15: v6.services.pub.PubTicket.Get:output_type -> v6.services.pub.Ticket
	2,  // 16: v6.services.pub.PubTicket.List:output_type -> v6.services.pub.TicketListResponse
	0,  // 17: v6.services.pub.PubTicket.Create:output_type -> v6.services.pub.Ticket
	3,  // 18: v6.services.pub.PubTicket.GetContactInfo:output_type -> v6.services.pub.ContactInfo
	3,  // 19: v6.services.pub.PubTicket.UpdateContactInfo:output_type -> v6.services.pub.ContactInfo
	5,  // 20: v6.services.pub.PubTicket.ListConversation:output_type -> v6.services.pub.ConversationListResponse
	0,  // 21: v6.services.pub.PubTicket.Comment:output_type -> v6.services.pub.Ticket
	8,  // 22: v6.services.pub.PubTicket.Delete:output_type -> v6.services.pub.common.StringList
	8,  // 23: v6.services.pub.PubTicket.Close:output_type -> v6.services.pub.common.StringList
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
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
