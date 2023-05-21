// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: plugnmeet_auth_recording.proto

package plugnmeet

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type FetchRecordingsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomIds []string `protobuf:"bytes,1,rep,name=room_ids,json=roomIds,proto3" json:"room_ids,omitempty"`
	From    uint32   `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit   uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy string   `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *FetchRecordingsReq) Reset() {
	*x = FetchRecordingsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRecordingsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRecordingsReq) ProtoMessage() {}

func (x *FetchRecordingsReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRecordingsReq.ProtoReflect.Descriptor instead.
func (*FetchRecordingsReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRecordingsReq) GetRoomIds() []string {
	if x != nil {
		return x.RoomIds
	}
	return nil
}

func (x *FetchRecordingsReq) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *FetchRecordingsReq) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchRecordingsReq) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type RecordingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId         string  `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	RoomId           string  `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomSid          string  `protobuf:"bytes,3,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	FilePath         string  `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileSize         float32 `protobuf:"fixed32,5,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	CreationTime     int64   `protobuf:"varint,6,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	RoomCreationTime int64   `protobuf:"varint,7,opt,name=room_creation_time,json=roomCreationTime,proto3" json:"room_creation_time,omitempty"`
}

func (x *RecordingInfo) Reset() {
	*x = RecordingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingInfo) ProtoMessage() {}

func (x *RecordingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingInfo.ProtoReflect.Descriptor instead.
func (*RecordingInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{1}
}

func (x *RecordingInfo) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

func (x *RecordingInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RecordingInfo) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *RecordingInfo) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *RecordingInfo) GetFileSize() float32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *RecordingInfo) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *RecordingInfo) GetRoomCreationTime() int64 {
	if x != nil {
		return x.RoomCreationTime
	}
	return 0
}

type FetchRecordingsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalRecordings int64            `protobuf:"varint,1,opt,name=total_recordings,json=totalRecordings,proto3" json:"total_recordings,omitempty"`
	From            uint32           `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit           uint32           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	OrderBy         string           `protobuf:"bytes,4,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	RecordingsList  []*RecordingInfo `protobuf:"bytes,5,rep,name=recordings_list,json=recordingsList,proto3" json:"recordings_list,omitempty"`
}

func (x *FetchRecordingsResult) Reset() {
	*x = FetchRecordingsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRecordingsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRecordingsResult) ProtoMessage() {}

func (x *FetchRecordingsResult) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRecordingsResult.ProtoReflect.Descriptor instead.
func (*FetchRecordingsResult) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{2}
}

func (x *FetchRecordingsResult) GetTotalRecordings() int64 {
	if x != nil {
		return x.TotalRecordings
	}
	return 0
}

func (x *FetchRecordingsResult) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *FetchRecordingsResult) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FetchRecordingsResult) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *FetchRecordingsResult) GetRecordingsList() []*RecordingInfo {
	if x != nil {
		return x.RecordingsList
	}
	return nil
}

type FetchRecordingsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Result *FetchRecordingsResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FetchRecordingsRes) Reset() {
	*x = FetchRecordingsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRecordingsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRecordingsRes) ProtoMessage() {}

func (x *FetchRecordingsRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRecordingsRes.ProtoReflect.Descriptor instead.
func (*FetchRecordingsRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{3}
}

func (x *FetchRecordingsRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *FetchRecordingsRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *FetchRecordingsRes) GetResult() *FetchRecordingsResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteRecordingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId string `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
}

func (x *DeleteRecordingReq) Reset() {
	*x = DeleteRecordingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRecordingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecordingReq) ProtoMessage() {}

func (x *DeleteRecordingReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecordingReq.ProtoReflect.Descriptor instead.
func (*DeleteRecordingReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRecordingReq) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

type DeleteRecordingRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteRecordingRes) Reset() {
	*x = DeleteRecordingRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRecordingRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRecordingRes) ProtoMessage() {}

func (x *DeleteRecordingRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRecordingRes.ProtoReflect.Descriptor instead.
func (*DeleteRecordingRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRecordingRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *DeleteRecordingRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetDownloadTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId string `protobuf:"bytes,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
}

func (x *GetDownloadTokenReq) Reset() {
	*x = GetDownloadTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDownloadTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadTokenReq) ProtoMessage() {}

func (x *GetDownloadTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadTokenReq.ProtoReflect.Descriptor instead.
func (*GetDownloadTokenReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{6}
}

func (x *GetDownloadTokenReq) GetRecordId() string {
	if x != nil {
		return x.RecordId
	}
	return ""
}

type GetDownloadTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token  *string `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *GetDownloadTokenRes) Reset() {
	*x = GetDownloadTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_auth_recording_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDownloadTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDownloadTokenRes) ProtoMessage() {}

func (x *GetDownloadTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_auth_recording_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDownloadTokenRes.ProtoReflect.Descriptor instead.
func (*GetDownloadTokenRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_auth_recording_proto_rawDescGZIP(), []int{7}
}

func (x *GetDownloadTokenRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetDownloadTokenRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetDownloadTokenRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

var File_plugnmeet_auth_recording_proto protoreflect.FileDescriptor

var file_plugnmeet_auth_recording_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0xed, 0x01, 0x0a, 0x0d, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x15, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x12, 0x41, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d,
	0x65, 0x65, 0x74, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x3a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x3e, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3b, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d,
	0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_auth_recording_proto_rawDescOnce sync.Once
	file_plugnmeet_auth_recording_proto_rawDescData = file_plugnmeet_auth_recording_proto_rawDesc
)

func file_plugnmeet_auth_recording_proto_rawDescGZIP() []byte {
	file_plugnmeet_auth_recording_proto_rawDescOnce.Do(func() {
		file_plugnmeet_auth_recording_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_auth_recording_proto_rawDescData)
	})
	return file_plugnmeet_auth_recording_proto_rawDescData
}

var file_plugnmeet_auth_recording_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_plugnmeet_auth_recording_proto_goTypes = []interface{}{
	(*FetchRecordingsReq)(nil),    // 0: plugnmeet.FetchRecordingsReq
	(*RecordingInfo)(nil),         // 1: plugnmeet.RecordingInfo
	(*FetchRecordingsResult)(nil), // 2: plugnmeet.FetchRecordingsResult
	(*FetchRecordingsRes)(nil),    // 3: plugnmeet.FetchRecordingsRes
	(*DeleteRecordingReq)(nil),    // 4: plugnmeet.DeleteRecordingReq
	(*DeleteRecordingRes)(nil),    // 5: plugnmeet.DeleteRecordingRes
	(*GetDownloadTokenReq)(nil),   // 6: plugnmeet.GetDownloadTokenReq
	(*GetDownloadTokenRes)(nil),   // 7: plugnmeet.GetDownloadTokenRes
}
var file_plugnmeet_auth_recording_proto_depIdxs = []int32{
	1, // 0: plugnmeet.FetchRecordingsResult.recordings_list:type_name -> plugnmeet.RecordingInfo
	2, // 1: plugnmeet.FetchRecordingsRes.result:type_name -> plugnmeet.FetchRecordingsResult
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plugnmeet_auth_recording_proto_init() }
func file_plugnmeet_auth_recording_proto_init() {
	if File_plugnmeet_auth_recording_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_auth_recording_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRecordingsReq); i {
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
		file_plugnmeet_auth_recording_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingInfo); i {
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
		file_plugnmeet_auth_recording_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRecordingsResult); i {
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
		file_plugnmeet_auth_recording_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRecordingsRes); i {
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
		file_plugnmeet_auth_recording_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRecordingReq); i {
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
		file_plugnmeet_auth_recording_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRecordingRes); i {
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
		file_plugnmeet_auth_recording_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDownloadTokenReq); i {
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
		file_plugnmeet_auth_recording_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDownloadTokenRes); i {
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
	file_plugnmeet_auth_recording_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_auth_recording_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_auth_recording_proto_goTypes,
		DependencyIndexes: file_plugnmeet_auth_recording_proto_depIdxs,
		MessageInfos:      file_plugnmeet_auth_recording_proto_msgTypes,
	}.Build()
	File_plugnmeet_auth_recording_proto = out.File
	file_plugnmeet_auth_recording_proto_rawDesc = nil
	file_plugnmeet_auth_recording_proto_goTypes = nil
	file_plugnmeet_auth_recording_proto_depIdxs = nil
}
