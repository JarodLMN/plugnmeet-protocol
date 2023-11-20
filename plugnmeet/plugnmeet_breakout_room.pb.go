// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: plugnmeet_breakout_room.proto

package plugnmeet

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

type CreateBreakoutRoomsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId          string          `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RequestedUserId string          `protobuf:"bytes,2,opt,name=requested_user_id,json=requestedUserId,proto3" json:"requested_user_id,omitempty"`
	Duration        uint64          `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	WelcomeMsg      *string         `protobuf:"bytes,4,opt,name=welcome_msg,json=welcomeMsg,proto3,oneof" json:"welcome_msg,omitempty"`
	Rooms           []*BreakoutRoom `protobuf:"bytes,5,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *CreateBreakoutRoomsReq) Reset() {
	*x = CreateBreakoutRoomsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBreakoutRoomsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBreakoutRoomsReq) ProtoMessage() {}

func (x *CreateBreakoutRoomsReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBreakoutRoomsReq.ProtoReflect.Descriptor instead.
func (*CreateBreakoutRoomsReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBreakoutRoomsReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *CreateBreakoutRoomsReq) GetRequestedUserId() string {
	if x != nil {
		return x.RequestedUserId
	}
	return ""
}

func (x *CreateBreakoutRoomsReq) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CreateBreakoutRoomsReq) GetWelcomeMsg() string {
	if x != nil && x.WelcomeMsg != nil {
		return *x.WelcomeMsg
	}
	return ""
}

func (x *CreateBreakoutRoomsReq) GetRooms() []*BreakoutRoom {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type BreakoutRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Duration uint64              `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Started  bool                `protobuf:"varint,4,opt,name=started,proto3" json:"started,omitempty"`
	Created  uint64              `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Users    []*BreakoutRoomUser `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *BreakoutRoom) Reset() {
	*x = BreakoutRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakoutRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoom) ProtoMessage() {}

func (x *BreakoutRoom) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutRoom.ProtoReflect.Descriptor instead.
func (*BreakoutRoom) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{1}
}

func (x *BreakoutRoom) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BreakoutRoom) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BreakoutRoom) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *BreakoutRoom) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

func (x *BreakoutRoom) GetCreated() uint64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *BreakoutRoom) GetUsers() []*BreakoutRoomUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type BreakoutRoomUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Joined bool   `protobuf:"varint,3,opt,name=joined,proto3" json:"joined,omitempty"`
}

func (x *BreakoutRoomUser) Reset() {
	*x = BreakoutRoomUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakoutRoomUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoomUser) ProtoMessage() {}

func (x *BreakoutRoomUser) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutRoomUser.ProtoReflect.Descriptor instead.
func (*BreakoutRoomUser) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{2}
}

func (x *BreakoutRoomUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BreakoutRoomUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BreakoutRoomUser) GetJoined() bool {
	if x != nil {
		return x.Joined
	}
	return false
}

type IncreaseBreakoutRoomDurationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BreakoutRoomId string `protobuf:"bytes,1,opt,name=breakout_room_id,json=breakoutRoomId,proto3" json:"breakout_room_id,omitempty"`
	Duration       uint64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	RoomId         string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *IncreaseBreakoutRoomDurationReq) Reset() {
	*x = IncreaseBreakoutRoomDurationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseBreakoutRoomDurationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseBreakoutRoomDurationReq) ProtoMessage() {}

func (x *IncreaseBreakoutRoomDurationReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseBreakoutRoomDurationReq.ProtoReflect.Descriptor instead.
func (*IncreaseBreakoutRoomDurationReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{3}
}

func (x *IncreaseBreakoutRoomDurationReq) GetBreakoutRoomId() string {
	if x != nil {
		return x.BreakoutRoomId
	}
	return ""
}

func (x *IncreaseBreakoutRoomDurationReq) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *IncreaseBreakoutRoomDurationReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type BroadcastBreakoutRoomMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg    string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	RoomId string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *BroadcastBreakoutRoomMsgReq) Reset() {
	*x = BroadcastBreakoutRoomMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastBreakoutRoomMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastBreakoutRoomMsgReq) ProtoMessage() {}

func (x *BroadcastBreakoutRoomMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastBreakoutRoomMsgReq.ProtoReflect.Descriptor instead.
func (*BroadcastBreakoutRoomMsgReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastBreakoutRoomMsgReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BroadcastBreakoutRoomMsgReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type JoinBreakoutRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BreakoutRoomId string `protobuf:"bytes,1,opt,name=breakout_room_id,json=breakoutRoomId,proto3" json:"breakout_room_id,omitempty"`
	UserId         string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoomId         string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	IsAdmin        bool   `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *JoinBreakoutRoomReq) Reset() {
	*x = JoinBreakoutRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinBreakoutRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinBreakoutRoomReq) ProtoMessage() {}

func (x *JoinBreakoutRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinBreakoutRoomReq.ProtoReflect.Descriptor instead.
func (*JoinBreakoutRoomReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{5}
}

func (x *JoinBreakoutRoomReq) GetBreakoutRoomId() string {
	if x != nil {
		return x.BreakoutRoomId
	}
	return ""
}

func (x *JoinBreakoutRoomReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *JoinBreakoutRoomReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *JoinBreakoutRoomReq) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type EndBreakoutRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BreakoutRoomId string `protobuf:"bytes,1,opt,name=breakout_room_id,json=breakoutRoomId,proto3" json:"breakout_room_id,omitempty"`
	RoomId         string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *EndBreakoutRoomReq) Reset() {
	*x = EndBreakoutRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndBreakoutRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndBreakoutRoomReq) ProtoMessage() {}

func (x *EndBreakoutRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndBreakoutRoomReq.ProtoReflect.Descriptor instead.
func (*EndBreakoutRoomReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{6}
}

func (x *EndBreakoutRoomReq) GetBreakoutRoomId() string {
	if x != nil {
		return x.BreakoutRoomId
	}
	return ""
}

func (x *EndBreakoutRoomReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type BreakoutRoomRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool            `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token  *string         `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"` // join token
	Room   *BreakoutRoom   `protobuf:"bytes,4,opt,name=room,proto3,oneof" json:"room,omitempty"`   // for my breakout room
	Rooms  []*BreakoutRoom `protobuf:"bytes,5,rep,name=rooms,proto3" json:"rooms,omitempty"`       // rooms list
}

func (x *BreakoutRoomRes) Reset() {
	*x = BreakoutRoomRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_breakout_room_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BreakoutRoomRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BreakoutRoomRes) ProtoMessage() {}

func (x *BreakoutRoomRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_breakout_room_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BreakoutRoomRes.ProtoReflect.Descriptor instead.
func (*BreakoutRoomRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_breakout_room_proto_rawDescGZIP(), []int{7}
}

func (x *BreakoutRoomRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *BreakoutRoomRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *BreakoutRoomRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *BreakoutRoomRes) GetRoom() *BreakoutRoom {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *BreakoutRoomRes) GetRooms() []*BreakoutRoom {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_plugnmeet_breakout_room_proto protoreflect.FileDescriptor

var file_plugnmeet_breakout_room_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x62, 0x72, 0x65, 0x61,
	0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d,
	0x65, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x77,
	0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x05,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x22, 0xb7, 0x01, 0x0a, 0x0c,
	0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x4e, 0x0a, 0x10, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6a,
	0x6f, 0x69, 0x6e, 0x65, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x1f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61,
	0x73, 0x65, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x1b, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x13, 0x4a, 0x6f, 0x69, 0x6e, 0x42, 0x72, 0x65, 0x61, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x22, 0x57, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x72, 0x65, 0x61, 0x6b,
	0x6f, 0x75, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x0f, 0x42,
	0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x42, 0x72,
	0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x01, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74,
	0x2e, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_breakout_room_proto_rawDescOnce sync.Once
	file_plugnmeet_breakout_room_proto_rawDescData = file_plugnmeet_breakout_room_proto_rawDesc
)

func file_plugnmeet_breakout_room_proto_rawDescGZIP() []byte {
	file_plugnmeet_breakout_room_proto_rawDescOnce.Do(func() {
		file_plugnmeet_breakout_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_breakout_room_proto_rawDescData)
	})
	return file_plugnmeet_breakout_room_proto_rawDescData
}

var file_plugnmeet_breakout_room_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_plugnmeet_breakout_room_proto_goTypes = []interface{}{
	(*CreateBreakoutRoomsReq)(nil),          // 0: plugnmeet.CreateBreakoutRoomsReq
	(*BreakoutRoom)(nil),                    // 1: plugnmeet.BreakoutRoom
	(*BreakoutRoomUser)(nil),                // 2: plugnmeet.BreakoutRoomUser
	(*IncreaseBreakoutRoomDurationReq)(nil), // 3: plugnmeet.IncreaseBreakoutRoomDurationReq
	(*BroadcastBreakoutRoomMsgReq)(nil),     // 4: plugnmeet.BroadcastBreakoutRoomMsgReq
	(*JoinBreakoutRoomReq)(nil),             // 5: plugnmeet.JoinBreakoutRoomReq
	(*EndBreakoutRoomReq)(nil),              // 6: plugnmeet.EndBreakoutRoomReq
	(*BreakoutRoomRes)(nil),                 // 7: plugnmeet.BreakoutRoomRes
}
var file_plugnmeet_breakout_room_proto_depIdxs = []int32{
	1, // 0: plugnmeet.CreateBreakoutRoomsReq.rooms:type_name -> plugnmeet.BreakoutRoom
	2, // 1: plugnmeet.BreakoutRoom.users:type_name -> plugnmeet.BreakoutRoomUser
	1, // 2: plugnmeet.BreakoutRoomRes.room:type_name -> plugnmeet.BreakoutRoom
	1, // 3: plugnmeet.BreakoutRoomRes.rooms:type_name -> plugnmeet.BreakoutRoom
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_plugnmeet_breakout_room_proto_init() }
func file_plugnmeet_breakout_room_proto_init() {
	if File_plugnmeet_breakout_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_breakout_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBreakoutRoomsReq); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakoutRoom); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakoutRoomUser); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseBreakoutRoomDurationReq); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastBreakoutRoomMsgReq); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinBreakoutRoomReq); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndBreakoutRoomReq); i {
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
		file_plugnmeet_breakout_room_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BreakoutRoomRes); i {
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
	file_plugnmeet_breakout_room_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_plugnmeet_breakout_room_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_breakout_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_breakout_room_proto_goTypes,
		DependencyIndexes: file_plugnmeet_breakout_room_proto_depIdxs,
		MessageInfos:      file_plugnmeet_breakout_room_proto_msgTypes,
	}.Build()
	File_plugnmeet_breakout_room_proto = out.File
	file_plugnmeet_breakout_room_proto_rawDesc = nil
	file_plugnmeet_breakout_room_proto_goTypes = nil
	file_plugnmeet_breakout_room_proto_depIdxs = nil
}
