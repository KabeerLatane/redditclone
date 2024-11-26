// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: api/proto/generated/messages.proto

package generated

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

type UserMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserMessage) Reset() {
	*x = UserMessage{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessage) ProtoMessage() {}

func (x *UserMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessage.ProtoReflect.Descriptor instead.
func (*UserMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{0}
}

func (x *UserMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SubredditMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   string `protobuf:"bytes,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *SubredditMessage) Reset() {
	*x = SubredditMessage{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubredditMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubredditMessage) ProtoMessage() {}

func (x *SubredditMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubredditMessage.ProtoReflect.Descriptor instead.
func (*SubredditMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{1}
}

func (x *SubredditMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubredditMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubredditMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SubredditMessage) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

type PostMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubredditId string `protobuf:"bytes,2,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
	AuthorId    string `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content     string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt   int64  `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	IsRepost    bool   `protobuf:"varint,7,opt,name=is_repost,json=isRepost,proto3" json:"is_repost,omitempty"`
}

func (x *PostMessage) Reset() {
	*x = PostMessage{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMessage) ProtoMessage() {}

func (x *PostMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMessage.ProtoReflect.Descriptor instead.
func (*PostMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PostMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostMessage) GetSubredditId() string {
	if x != nil {
		return x.SubredditId
	}
	return ""
}

func (x *PostMessage) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *PostMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostMessage) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PostMessage) GetIsRepost() bool {
	if x != nil {
		return x.IsRepost
	}
	return false
}

type VoteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsUpvote bool   `protobuf:"varint,3,opt,name=is_upvote,json=isUpvote,proto3" json:"is_upvote,omitempty"`
}

func (x *VoteMessage) Reset() {
	*x = VoteMessage{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteMessage) ProtoMessage() {}

func (x *VoteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteMessage.ProtoReflect.Descriptor instead.
func (*VoteMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{3}
}

func (x *VoteMessage) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *VoteMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VoteMessage) GetIsUpvote() bool {
	if x != nil {
		return x.IsUpvote
	}
	return false
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{5}
}

func (x *SuccessResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CommentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId    string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	ParentId  string `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	AuthorId  string `protobuf:"bytes,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Content   string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CommentMessage) Reset() {
	*x = CommentMessage{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentMessage) ProtoMessage() {}

func (x *CommentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentMessage.ProtoReflect.Descriptor instead.
func (*CommentMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{6}
}

func (x *CommentMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommentMessage) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *CommentMessage) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CommentMessage) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *CommentMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentMessage) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type JoinSubredditMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubredditId string `protobuf:"bytes,1,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *JoinSubredditMessage) Reset() {
	*x = JoinSubredditMessage{}
	mi := &file_api_proto_generated_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinSubredditMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSubredditMessage) ProtoMessage() {}

func (x *JoinSubredditMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_generated_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSubredditMessage.ProtoReflect.Descriptor instead.
func (*JoinSubredditMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_generated_messages_proto_rawDescGZIP(), []int{7}
}

func (x *JoinSubredditMessage) GetSubredditId() string {
	if x != nil {
		return x.SubredditId
	}
	return ""
}

func (x *JoinSubredditMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_api_proto_generated_messages_proto protoreflect.FileDescriptor

var file_api_proto_generated_messages_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x22, 0x5e, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x77, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64,
	0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62,
	0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x74, 0x22, 0x60, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x55, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x52, 0x0a, 0x14, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x75,
	0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x22, 0x5a, 0x20, 0x72, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x2d, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_generated_messages_proto_rawDescOnce sync.Once
	file_api_proto_generated_messages_proto_rawDescData = file_api_proto_generated_messages_proto_rawDesc
)

func file_api_proto_generated_messages_proto_rawDescGZIP() []byte {
	file_api_proto_generated_messages_proto_rawDescOnce.Do(func() {
		file_api_proto_generated_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_generated_messages_proto_rawDescData)
	})
	return file_api_proto_generated_messages_proto_rawDescData
}

var file_api_proto_generated_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_generated_messages_proto_goTypes = []any{
	(*UserMessage)(nil),          // 0: reddit.UserMessage
	(*SubredditMessage)(nil),     // 1: reddit.SubredditMessage
	(*PostMessage)(nil),          // 2: reddit.PostMessage
	(*VoteMessage)(nil),          // 3: reddit.VoteMessage
	(*ErrorResponse)(nil),        // 4: reddit.ErrorResponse
	(*SuccessResponse)(nil),      // 5: reddit.SuccessResponse
	(*CommentMessage)(nil),       // 6: reddit.CommentMessage
	(*JoinSubredditMessage)(nil), // 7: reddit.JoinSubredditMessage
}
var file_api_proto_generated_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_generated_messages_proto_init() }
func file_api_proto_generated_messages_proto_init() {
	if File_api_proto_generated_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_generated_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_generated_messages_proto_goTypes,
		DependencyIndexes: file_api_proto_generated_messages_proto_depIdxs,
		MessageInfos:      file_api_proto_generated_messages_proto_msgTypes,
	}.Build()
	File_api_proto_generated_messages_proto = out.File
	file_api_proto_generated_messages_proto_rawDesc = nil
	file_api_proto_generated_messages_proto_goTypes = nil
	file_api_proto_generated_messages_proto_depIdxs = nil
}