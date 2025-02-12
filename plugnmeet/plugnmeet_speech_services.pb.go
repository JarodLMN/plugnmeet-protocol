// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: plugnmeet_speech_services.proto

package plugnmeet

import (
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

type SpeechServiceUserStatusTasks int32

const (
	SpeechServiceUserStatusTasks_SPEECH_TO_TEXT_SESSION_STARTED SpeechServiceUserStatusTasks = 0
	SpeechServiceUserStatusTasks_SPEECH_TO_TEXT_SESSION_ENDED   SpeechServiceUserStatusTasks = 1
	SpeechServiceUserStatusTasks_SPEECH_TO_TEXT_TOTAL_USAGE     SpeechServiceUserStatusTasks = 2
)

// Enum value maps for SpeechServiceUserStatusTasks.
var (
	SpeechServiceUserStatusTasks_name = map[int32]string{
		0: "SPEECH_TO_TEXT_SESSION_STARTED",
		1: "SPEECH_TO_TEXT_SESSION_ENDED",
		2: "SPEECH_TO_TEXT_TOTAL_USAGE",
	}
	SpeechServiceUserStatusTasks_value = map[string]int32{
		"SPEECH_TO_TEXT_SESSION_STARTED": 0,
		"SPEECH_TO_TEXT_SESSION_ENDED":   1,
		"SPEECH_TO_TEXT_TOTAL_USAGE":     2,
	}
)

func (x SpeechServiceUserStatusTasks) Enum() *SpeechServiceUserStatusTasks {
	p := new(SpeechServiceUserStatusTasks)
	*p = x
	return p
}

func (x SpeechServiceUserStatusTasks) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpeechServiceUserStatusTasks) Descriptor() protoreflect.EnumDescriptor {
	return file_plugnmeet_speech_services_proto_enumTypes[0].Descriptor()
}

func (SpeechServiceUserStatusTasks) Type() protoreflect.EnumType {
	return &file_plugnmeet_speech_services_proto_enumTypes[0]
}

func (x SpeechServiceUserStatusTasks) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpeechServiceUserStatusTasks.Descriptor instead.
func (SpeechServiceUserStatusTasks) EnumDescriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{0}
}

type SpeechToTextTranslationReq struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	RoomId               string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	IsEnabled            bool                   `protobuf:"varint,3,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
	AllowedSpeechLangs   []string               `protobuf:"bytes,4,rep,name=allowed_speech_langs,json=allowedSpeechLangs,proto3" json:"allowed_speech_langs,omitempty"`
	AllowedSpeechUsers   []string               `protobuf:"bytes,5,rep,name=allowed_speech_users,json=allowedSpeechUsers,proto3" json:"allowed_speech_users,omitempty"`
	IsEnabledTranslation bool                   `protobuf:"varint,6,opt,name=is_enabled_translation,json=isEnabledTranslation,proto3" json:"is_enabled_translation,omitempty"`
	AllowedTransLangs    []string               `protobuf:"bytes,7,rep,name=allowed_trans_langs,json=allowedTransLangs,proto3" json:"allowed_trans_langs,omitempty"`
	DefaultSubtitleLang  *string                `protobuf:"bytes,8,opt,name=default_subtitle_lang,json=defaultSubtitleLang,proto3,oneof" json:"default_subtitle_lang,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *SpeechToTextTranslationReq) Reset() {
	*x = SpeechToTextTranslationReq{}
	mi := &file_plugnmeet_speech_services_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpeechToTextTranslationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechToTextTranslationReq) ProtoMessage() {}

func (x *SpeechToTextTranslationReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechToTextTranslationReq.ProtoReflect.Descriptor instead.
func (*SpeechToTextTranslationReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{0}
}

func (x *SpeechToTextTranslationReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *SpeechToTextTranslationReq) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *SpeechToTextTranslationReq) GetAllowedSpeechLangs() []string {
	if x != nil {
		return x.AllowedSpeechLangs
	}
	return nil
}

func (x *SpeechToTextTranslationReq) GetAllowedSpeechUsers() []string {
	if x != nil {
		return x.AllowedSpeechUsers
	}
	return nil
}

func (x *SpeechToTextTranslationReq) GetIsEnabledTranslation() bool {
	if x != nil {
		return x.IsEnabledTranslation
	}
	return false
}

func (x *SpeechToTextTranslationReq) GetAllowedTransLangs() []string {
	if x != nil {
		return x.AllowedTransLangs
	}
	return nil
}

func (x *SpeechToTextTranslationReq) GetDefaultSubtitleLang() string {
	if x != nil && x.DefaultSubtitleLang != nil {
		return *x.DefaultSubtitleLang
	}
	return ""
}

type GenerateAzureTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserSid       string                 `protobuf:"bytes,2,opt,name=user_sid,json=userSid,proto3" json:"user_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateAzureTokenReq) Reset() {
	*x = GenerateAzureTokenReq{}
	mi := &file_plugnmeet_speech_services_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateAzureTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAzureTokenReq) ProtoMessage() {}

func (x *GenerateAzureTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAzureTokenReq.ProtoReflect.Descriptor instead.
func (*GenerateAzureTokenReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateAzureTokenReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GenerateAzureTokenReq) GetUserSid() string {
	if x != nil {
		return x.UserSid
	}
	return ""
}

type GenerateAzureTokenRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token         *string                `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"`
	ServiceRegion *string                `protobuf:"bytes,4,opt,name=service_region,json=serviceRegion,proto3,oneof" json:"service_region,omitempty"`
	KeyId         *string                `protobuf:"bytes,5,opt,name=key_id,json=keyId,proto3,oneof" json:"key_id,omitempty"`
	Renew         bool                   `protobuf:"varint,6,opt,name=renew,proto3" json:"renew,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateAzureTokenRes) Reset() {
	*x = GenerateAzureTokenRes{}
	mi := &file_plugnmeet_speech_services_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateAzureTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAzureTokenRes) ProtoMessage() {}

func (x *GenerateAzureTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAzureTokenRes.ProtoReflect.Descriptor instead.
func (*GenerateAzureTokenRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateAzureTokenRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GenerateAzureTokenRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GenerateAzureTokenRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *GenerateAzureTokenRes) GetServiceRegion() string {
	if x != nil && x.ServiceRegion != nil {
		return *x.ServiceRegion
	}
	return ""
}

func (x *GenerateAzureTokenRes) GetKeyId() string {
	if x != nil && x.KeyId != nil {
		return *x.KeyId
	}
	return ""
}

func (x *GenerateAzureTokenRes) GetRenew() bool {
	if x != nil {
		return x.Renew
	}
	return false
}

type AzureTokenRenewReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserSid       string                 `protobuf:"bytes,2,opt,name=user_sid,json=userSid,proto3" json:"user_sid,omitempty"`
	ServiceRegion string                 `protobuf:"bytes,3,opt,name=service_region,json=serviceRegion,proto3" json:"service_region,omitempty"`
	KeyId         string                 `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AzureTokenRenewReq) Reset() {
	*x = AzureTokenRenewReq{}
	mi := &file_plugnmeet_speech_services_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureTokenRenewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureTokenRenewReq) ProtoMessage() {}

func (x *AzureTokenRenewReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureTokenRenewReq.ProtoReflect.Descriptor instead.
func (*AzureTokenRenewReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{3}
}

func (x *AzureTokenRenewReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *AzureTokenRenewReq) GetUserSid() string {
	if x != nil {
		return x.UserSid
	}
	return ""
}

func (x *AzureTokenRenewReq) GetServiceRegion() string {
	if x != nil {
		return x.ServiceRegion
	}
	return ""
}

func (x *AzureTokenRenewReq) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type SpeechServiceUserStatusReq struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	RoomId        string                       `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid       string                       `protobuf:"bytes,2,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	UserId        string                       `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	KeyId         string                       `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Task          SpeechServiceUserStatusTasks `protobuf:"varint,5,opt,name=task,proto3,enum=plugnmeet.SpeechServiceUserStatusTasks" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SpeechServiceUserStatusReq) Reset() {
	*x = SpeechServiceUserStatusReq{}
	mi := &file_plugnmeet_speech_services_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpeechServiceUserStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechServiceUserStatusReq) ProtoMessage() {}

func (x *SpeechServiceUserStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_speech_services_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechServiceUserStatusReq.ProtoReflect.Descriptor instead.
func (*SpeechServiceUserStatusReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_speech_services_proto_rawDescGZIP(), []int{4}
}

func (x *SpeechServiceUserStatusReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *SpeechServiceUserStatusReq) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *SpeechServiceUserStatusReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SpeechServiceUserStatusReq) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *SpeechServiceUserStatusReq) GetTask() SpeechServiceUserStatusTasks {
	if x != nil {
		return x.Task
	}
	return SpeechServiceUserStatusTasks_SPEECH_TO_TEXT_SESSION_STARTED
}

var File_plugnmeet_speech_services_proto protoreflect.FileDescriptor

var file_plugnmeet_speech_services_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x22, 0xf1, 0x02, 0x0a,
	0x1a, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x54, 0x6f, 0x54, 0x65, 0x78, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68,
	0x4c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x6c,
	0x61, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x37, 0x0a,
	0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x13,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c,
	0x61, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67,
	0x22, 0x4b, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x7a, 0x75, 0x72,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53, 0x69, 0x64, 0x22, 0xe2, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53, 0x69, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x1a,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x3b,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x2a, 0x84, 0x01, 0x0a, 0x1c,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x22, 0x0a, 0x1e,
	0x53, 0x50, 0x45, 0x45, 0x43, 0x48, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x53, 0x50, 0x45, 0x45, 0x43, 0x48, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x45,
	0x58, 0x54, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x50, 0x45, 0x45, 0x43, 0x48, 0x5f, 0x54, 0x4f, 0x5f,
	0x54, 0x45, 0x58, 0x54, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x02, 0x42, 0xa5, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e,
	0x6d, 0x65, 0x65, 0x74, 0x42, 0x1c, 0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x53,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xca, 0x02, 0x09, 0x50, 0x6c, 0x75,
	0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xe2, 0x02, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65,
	0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x09, 0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_plugnmeet_speech_services_proto_rawDescOnce sync.Once
	file_plugnmeet_speech_services_proto_rawDescData []byte
)

func file_plugnmeet_speech_services_proto_rawDescGZIP() []byte {
	file_plugnmeet_speech_services_proto_rawDescOnce.Do(func() {
		file_plugnmeet_speech_services_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_plugnmeet_speech_services_proto_rawDesc), len(file_plugnmeet_speech_services_proto_rawDesc)))
	})
	return file_plugnmeet_speech_services_proto_rawDescData
}

var file_plugnmeet_speech_services_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_plugnmeet_speech_services_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_plugnmeet_speech_services_proto_goTypes = []any{
	(SpeechServiceUserStatusTasks)(0),  // 0: plugnmeet.SpeechServiceUserStatusTasks
	(*SpeechToTextTranslationReq)(nil), // 1: plugnmeet.SpeechToTextTranslationReq
	(*GenerateAzureTokenReq)(nil),      // 2: plugnmeet.GenerateAzureTokenReq
	(*GenerateAzureTokenRes)(nil),      // 3: plugnmeet.GenerateAzureTokenRes
	(*AzureTokenRenewReq)(nil),         // 4: plugnmeet.AzureTokenRenewReq
	(*SpeechServiceUserStatusReq)(nil), // 5: plugnmeet.SpeechServiceUserStatusReq
}
var file_plugnmeet_speech_services_proto_depIdxs = []int32{
	0, // 0: plugnmeet.SpeechServiceUserStatusReq.task:type_name -> plugnmeet.SpeechServiceUserStatusTasks
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_plugnmeet_speech_services_proto_init() }
func file_plugnmeet_speech_services_proto_init() {
	if File_plugnmeet_speech_services_proto != nil {
		return
	}
	file_plugnmeet_speech_services_proto_msgTypes[0].OneofWrappers = []any{}
	file_plugnmeet_speech_services_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_plugnmeet_speech_services_proto_rawDesc), len(file_plugnmeet_speech_services_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_speech_services_proto_goTypes,
		DependencyIndexes: file_plugnmeet_speech_services_proto_depIdxs,
		EnumInfos:         file_plugnmeet_speech_services_proto_enumTypes,
		MessageInfos:      file_plugnmeet_speech_services_proto_msgTypes,
	}.Build()
	File_plugnmeet_speech_services_proto = out.File
	file_plugnmeet_speech_services_proto_goTypes = nil
	file_plugnmeet_speech_services_proto_depIdxs = nil
}
