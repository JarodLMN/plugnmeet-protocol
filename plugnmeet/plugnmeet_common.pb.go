// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: plugnmeet_common.proto

package plugnmeet

import (
	livekit "github.com/livekit/protocol/livekit"
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

type CommonNotifyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event         *string                  `protobuf:"bytes,1,opt,name=event,proto3,oneof" json:"event,omitempty"`
	Room          *NotifyEventRoom         `protobuf:"bytes,2,opt,name=room,proto3,oneof" json:"room,omitempty"`
	Participant   *livekit.ParticipantInfo `protobuf:"bytes,3,opt,name=participant,proto3,oneof" json:"participant,omitempty"`
	RecordingInfo *RecordingInfoEvent      `protobuf:"bytes,4,opt,name=recording_info,json=recordingInfo,proto3,oneof" json:"recording_info,omitempty"`
	SpeechService *SpeechServiceEvent      `protobuf:"bytes,5,opt,name=speech_service,json=speechService,proto3,oneof" json:"speech_service,omitempty"`
	Track         *livekit.TrackInfo       `protobuf:"bytes,6,opt,name=track,proto3,oneof" json:"track,omitempty"`
	Id            *string                  `protobuf:"bytes,9,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CreatedAt     *int64                   `protobuf:"varint,10,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
}

func (x *CommonNotifyEvent) Reset() {
	*x = CommonNotifyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonNotifyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonNotifyEvent) ProtoMessage() {}

func (x *CommonNotifyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonNotifyEvent.ProtoReflect.Descriptor instead.
func (*CommonNotifyEvent) Descriptor() ([]byte, []int) {
	return file_plugnmeet_common_proto_rawDescGZIP(), []int{0}
}

func (x *CommonNotifyEvent) GetEvent() string {
	if x != nil && x.Event != nil {
		return *x.Event
	}
	return ""
}

func (x *CommonNotifyEvent) GetRoom() *NotifyEventRoom {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *CommonNotifyEvent) GetParticipant() *livekit.ParticipantInfo {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *CommonNotifyEvent) GetRecordingInfo() *RecordingInfoEvent {
	if x != nil {
		return x.RecordingInfo
	}
	return nil
}

func (x *CommonNotifyEvent) GetSpeechService() *SpeechServiceEvent {
	if x != nil {
		return x.SpeechService
	}
	return nil
}

func (x *CommonNotifyEvent) GetTrack() *livekit.TrackInfo {
	if x != nil {
		return x.Track
	}
	return nil
}

func (x *CommonNotifyEvent) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *CommonNotifyEvent) GetCreatedAt() int64 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

type NotifyEventRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid             *string          `protobuf:"bytes,1,opt,name=sid,proto3,oneof" json:"sid,omitempty"`
	RoomId          *string          `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3,oneof" json:"room_id,omitempty"`
	EmptyTimeout    *uint32          `protobuf:"varint,3,opt,name=empty_timeout,json=emptyTimeout,proto3,oneof" json:"empty_timeout,omitempty"`
	MaxParticipants *uint32          `protobuf:"varint,4,opt,name=max_participants,json=maxParticipants,proto3,oneof" json:"max_participants,omitempty"`
	CreationTime    *uint64          `protobuf:"varint,5,opt,name=creation_time,json=creationTime,proto3,oneof" json:"creation_time,omitempty"`
	EnabledCodecs   []*livekit.Codec `protobuf:"bytes,6,rep,name=enabled_codecs,json=enabledCodecs,proto3" json:"enabled_codecs,omitempty"`
	Metadata        *string          `protobuf:"bytes,7,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	NumParticipants *uint32          `protobuf:"varint,8,opt,name=num_participants,json=numParticipants,proto3,oneof" json:"num_participants,omitempty"`
}

func (x *NotifyEventRoom) Reset() {
	*x = NotifyEventRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyEventRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyEventRoom) ProtoMessage() {}

func (x *NotifyEventRoom) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyEventRoom.ProtoReflect.Descriptor instead.
func (*NotifyEventRoom) Descriptor() ([]byte, []int) {
	return file_plugnmeet_common_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyEventRoom) GetSid() string {
	if x != nil && x.Sid != nil {
		return *x.Sid
	}
	return ""
}

func (x *NotifyEventRoom) GetRoomId() string {
	if x != nil && x.RoomId != nil {
		return *x.RoomId
	}
	return ""
}

func (x *NotifyEventRoom) GetEmptyTimeout() uint32 {
	if x != nil && x.EmptyTimeout != nil {
		return *x.EmptyTimeout
	}
	return 0
}

func (x *NotifyEventRoom) GetMaxParticipants() uint32 {
	if x != nil && x.MaxParticipants != nil {
		return *x.MaxParticipants
	}
	return 0
}

func (x *NotifyEventRoom) GetCreationTime() uint64 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return 0
}

func (x *NotifyEventRoom) GetEnabledCodecs() []*livekit.Codec {
	if x != nil {
		return x.EnabledCodecs
	}
	return nil
}

func (x *NotifyEventRoom) GetMetadata() string {
	if x != nil && x.Metadata != nil {
		return *x.Metadata
	}
	return ""
}

func (x *NotifyEventRoom) GetNumParticipants() uint32 {
	if x != nil && x.NumParticipants != nil {
		return *x.NumParticipants
	}
	return 0
}

type RecordingInfoEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId    string   `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	RecorderId  string   `protobuf:"bytes,2,opt,name=recorder_id,json=recorderId,proto3" json:"recorder_id,omitempty"`
	RecorderMsg string   `protobuf:"bytes,3,opt,name=recorder_msg,json=recorderMsg,proto3" json:"recorder_msg,omitempty"`
	FilePath    *string  `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3,oneof" json:"file_path,omitempty"`
	FileSize    *float32 `protobuf:"fixed32,5,opt,name=file_size,json=fileSize,proto3,oneof" json:"file_size,omitempty"`
}

func (x *RecordingInfoEvent) Reset() {
	*x = RecordingInfoEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingInfoEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingInfoEvent) ProtoMessage() {}

func (x *RecordingInfoEvent) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingInfoEvent.ProtoReflect.Descriptor instead.
func (*RecordingInfoEvent) Descriptor() ([]byte, []int) {
	return file_plugnmeet_common_proto_rawDescGZIP(), []int{2}
}

func (x *RecordingInfoEvent) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *RecordingInfoEvent) GetRecorderId() string {
	if x != nil {
		return x.RecorderId
	}
	return ""
}

func (x *RecordingInfoEvent) GetRecorderMsg() string {
	if x != nil {
		return x.RecorderMsg
	}
	return ""
}

func (x *RecordingInfoEvent) GetFilePath() string {
	if x != nil && x.FilePath != nil {
		return *x.FilePath
	}
	return ""
}

func (x *RecordingInfoEvent) GetFileSize() float32 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

type SpeechServiceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     *string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	TotalUsage int64   `protobuf:"varint,4,opt,name=total_usage,json=totalUsage,proto3" json:"total_usage,omitempty"`
}

func (x *SpeechServiceEvent) Reset() {
	*x = SpeechServiceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeechServiceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechServiceEvent) ProtoMessage() {}

func (x *SpeechServiceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechServiceEvent.ProtoReflect.Descriptor instead.
func (*SpeechServiceEvent) Descriptor() ([]byte, []int) {
	return file_plugnmeet_common_proto_rawDescGZIP(), []int{3}
}

func (x *SpeechServiceEvent) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *SpeechServiceEvent) GetTotalUsage() int64 {
	if x != nil {
		return x.TotalUsage
	}
	return 0
}

var File_plugnmeet_common_proto protoreflect.FileDescriptor

var file_plugnmeet_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d,
	0x65, 0x65, 0x74, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x04, 0x0a, 0x11, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e,
	0x6d, 0x65, 0x65, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12,
	0x3f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x02,
	0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x49, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e,
	0x6d, 0x65, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x49, 0x0a, 0x0e, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x48, 0x04, 0x52, 0x0d, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x05, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x06, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x07,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x72, 0x6f, 0x6f,
	0x6d, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0xc1, 0x03, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x15, 0x0a, 0x03, 0x73,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x73, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x28, 0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x6d, 0x61,
	0x78, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x04, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x52, 0x0d, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10,
	0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x06, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x73, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0x5f, 0x0a, 0x12, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_plugnmeet_common_proto_rawDescOnce sync.Once
	file_plugnmeet_common_proto_rawDescData = file_plugnmeet_common_proto_rawDesc
)

func file_plugnmeet_common_proto_rawDescGZIP() []byte {
	file_plugnmeet_common_proto_rawDescOnce.Do(func() {
		file_plugnmeet_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_common_proto_rawDescData)
	})
	return file_plugnmeet_common_proto_rawDescData
}

var file_plugnmeet_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_plugnmeet_common_proto_goTypes = []interface{}{
	(*CommonNotifyEvent)(nil),       // 0: plugnmeet.CommonNotifyEvent
	(*NotifyEventRoom)(nil),         // 1: plugnmeet.NotifyEventRoom
	(*RecordingInfoEvent)(nil),      // 2: plugnmeet.RecordingInfoEvent
	(*SpeechServiceEvent)(nil),      // 3: plugnmeet.SpeechServiceEvent
	(*livekit.ParticipantInfo)(nil), // 4: livekit.ParticipantInfo
	(*livekit.TrackInfo)(nil),       // 5: livekit.TrackInfo
	(*livekit.Codec)(nil),           // 6: livekit.Codec
}
var file_plugnmeet_common_proto_depIdxs = []int32{
	1, // 0: plugnmeet.CommonNotifyEvent.room:type_name -> plugnmeet.NotifyEventRoom
	4, // 1: plugnmeet.CommonNotifyEvent.participant:type_name -> livekit.ParticipantInfo
	2, // 2: plugnmeet.CommonNotifyEvent.recording_info:type_name -> plugnmeet.RecordingInfoEvent
	3, // 3: plugnmeet.CommonNotifyEvent.speech_service:type_name -> plugnmeet.SpeechServiceEvent
	5, // 4: plugnmeet.CommonNotifyEvent.track:type_name -> livekit.TrackInfo
	6, // 5: plugnmeet.NotifyEventRoom.enabled_codecs:type_name -> livekit.Codec
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_plugnmeet_common_proto_init() }
func file_plugnmeet_common_proto_init() {
	if File_plugnmeet_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonNotifyEvent); i {
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
		file_plugnmeet_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyEventRoom); i {
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
		file_plugnmeet_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingInfoEvent); i {
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
		file_plugnmeet_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeechServiceEvent); i {
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
	file_plugnmeet_common_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_plugnmeet_common_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_plugnmeet_common_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_plugnmeet_common_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_common_proto_goTypes,
		DependencyIndexes: file_plugnmeet_common_proto_depIdxs,
		MessageInfos:      file_plugnmeet_common_proto_msgTypes,
	}.Build()
	File_plugnmeet_common_proto = out.File
	file_plugnmeet_common_proto_rawDesc = nil
	file_plugnmeet_common_proto_goTypes = nil
	file_plugnmeet_common_proto_depIdxs = nil
}
