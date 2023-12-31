// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0
// source: feed/service/v1/feed.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime string `protobuf:"bytes,1,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListFeedRequest) Reset() {
	*x = ListFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_service_v1_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedRequest) ProtoMessage() {}

func (x *ListFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feed_service_v1_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedRequest.ProtoReflect.Descriptor instead.
func (*ListFeedRequest) Descriptor() ([]byte, []int) {
	return file_feed_service_v1_feed_proto_rawDescGZIP(), []int{0}
}

func (x *ListFeedRequest) GetLatestTime() string {
	if x != nil {
		return x.LatestTime
	}
	return ""
}

func (x *ListFeedRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListFeedReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,proto3" json:"status_code,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,proto3" json:"status_msg,omitempty"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,proto3" json:"video_list,omitempty"`
	NextTime   int64    `protobuf:"varint,4,opt,name=next_time,proto3" json:"next_time,omitempty"`
}

func (x *ListFeedReply) Reset() {
	*x = ListFeedReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_service_v1_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedReply) ProtoMessage() {}

func (x *ListFeedReply) ProtoReflect() protoreflect.Message {
	mi := &file_feed_service_v1_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedReply.ProtoReflect.Descriptor instead.
func (*ListFeedReply) Descriptor() ([]byte, []int) {
	return file_feed_service_v1_feed_proto_rawDescGZIP(), []int{1}
}

func (x *ListFeedReply) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListFeedReply) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListFeedReply) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *ListFeedReply) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author        *User  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	PlayUrl       string `protobuf:"bytes,3,opt,name=play_url,proto3" json:"play_url,omitempty"`
	CoverUrl      string `protobuf:"bytes,4,opt,name=cover_url,proto3" json:"cover_url,omitempty"`
	FavoriteCount uint32 `protobuf:"varint,5,opt,name=favorite_count,proto3" json:"favorite_count,omitempty"`
	CommentCount  uint32 `protobuf:"varint,6,opt,name=comment_count,proto3" json:"comment_count,omitempty"`
	IsFavorite    bool   `protobuf:"varint,7,opt,name=is_favorite,proto3" json:"is_favorite,omitempty"`
	Title         string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_service_v1_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_feed_service_v1_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_feed_service_v1_feed_proto_rawDescGZIP(), []int{2}
}

func (x *Video) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() uint32 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() uint32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FollowCount     uint32 `protobuf:"varint,3,opt,name=follow_count,proto3" json:"follow_count,omitempty"`
	FollowerCount   uint32 `protobuf:"varint,4,opt,name=follower_count,proto3" json:"follower_count,omitempty"`
	IsFollow        bool   `protobuf:"varint,5,opt,name=is_follow,proto3" json:"is_follow,omitempty"`
	Avatar          string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,proto3" json:"background_image,omitempty"`
	Signature       string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	TotalFavorited  uint32 `protobuf:"varint,9,opt,name=total_favorited,proto3" json:"total_favorited,omitempty"`
	WorkCount       uint32 `protobuf:"varint,10,opt,name=work_count,proto3" json:"work_count,omitempty"`
	FavoriteCount   uint32 `protobuf:"varint,11,opt,name=favorite_count,proto3" json:"favorite_count,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_service_v1_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_feed_service_v1_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_feed_service_v1_feed_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() uint32 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() uint32 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetTotalFavorited() uint32 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *User) GetWorkCount() uint32 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *User) GetFavoriteCount() uint32 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

var File_feed_service_v1_feed_proto protoreflect.FileDescriptor

var file_feed_service_v1_feed_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x12, 0x36, 0x0a, 0x0a, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x86, 0x02, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x75, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0x71, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x62, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x42, 0x22, 0x5a, 0x20, 0x6e, 0x6f, 0x6e, 0x74, 0x69, 0x6b,
	0x74, 0x6f, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_feed_service_v1_feed_proto_rawDescOnce sync.Once
	file_feed_service_v1_feed_proto_rawDescData = file_feed_service_v1_feed_proto_rawDesc
)

func file_feed_service_v1_feed_proto_rawDescGZIP() []byte {
	file_feed_service_v1_feed_proto_rawDescOnce.Do(func() {
		file_feed_service_v1_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_service_v1_feed_proto_rawDescData)
	})
	return file_feed_service_v1_feed_proto_rawDescData
}

var file_feed_service_v1_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_feed_service_v1_feed_proto_goTypes = []interface{}{
	(*ListFeedRequest)(nil), // 0: feed.service.v1.ListFeedRequest
	(*ListFeedReply)(nil),   // 1: feed.service.v1.ListFeedReply
	(*Video)(nil),           // 2: feed.service.v1.Video
	(*User)(nil),            // 3: feed.service.v1.User
}
var file_feed_service_v1_feed_proto_depIdxs = []int32{
	2, // 0: feed.service.v1.ListFeedReply.video_list:type_name -> feed.service.v1.Video
	3, // 1: feed.service.v1.Video.author:type_name -> feed.service.v1.User
	0, // 2: feed.service.v1.FeedService.FeedList:input_type -> feed.service.v1.ListFeedRequest
	1, // 3: feed.service.v1.FeedService.FeedList:output_type -> feed.service.v1.ListFeedReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_feed_service_v1_feed_proto_init() }
func file_feed_service_v1_feed_proto_init() {
	if File_feed_service_v1_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feed_service_v1_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedRequest); i {
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
		file_feed_service_v1_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedReply); i {
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
		file_feed_service_v1_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_feed_service_v1_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_feed_service_v1_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_service_v1_feed_proto_goTypes,
		DependencyIndexes: file_feed_service_v1_feed_proto_depIdxs,
		MessageInfos:      file_feed_service_v1_feed_proto_msgTypes,
	}.Build()
	File_feed_service_v1_feed_proto = out.File
	file_feed_service_v1_feed_proto_rawDesc = nil
	file_feed_service_v1_feed_proto_goTypes = nil
	file_feed_service_v1_feed_proto_depIdxs = nil
}
