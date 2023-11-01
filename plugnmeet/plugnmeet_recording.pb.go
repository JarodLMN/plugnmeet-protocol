// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: plugnmeet_recording.proto

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

type RecordingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task         RecordingTasks `protobuf:"varint,1,opt,name=task,proto3,enum=plugnmeet.RecordingTasks" json:"task,omitempty"`
	RoomId       string         `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomTableId  int64          `protobuf:"varint,3,opt,name=room_table_id,json=roomTableId,proto3" json:"room_table_id,omitempty"`
	Sid          string         `protobuf:"bytes,4,opt,name=sid,proto3" json:"sid,omitempty"`
	RtmpUrl      *string        `protobuf:"bytes,5,opt,name=rtmp_url,json=rtmpUrl,proto3,oneof" json:"rtmp_url,omitempty"`
	CustomDesign *string        `protobuf:"bytes,6,opt,name=custom_design,json=customDesign,proto3,oneof" json:"custom_design,omitempty"`
}

func (x *RecordingReq) Reset() {
	*x = RecordingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recording_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingReq) ProtoMessage() {}

func (x *RecordingReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recording_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingReq.ProtoReflect.Descriptor instead.
func (*RecordingReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recording_proto_rawDescGZIP(), []int{0}
}

func (x *RecordingReq) GetTask() RecordingTasks {
	if x != nil {
		return x.Task
	}
	return RecordingTasks_START_RECORDING
}

func (x *RecordingReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RecordingReq) GetRoomTableId() int64 {
	if x != nil {
		return x.RoomTableId
	}
	return 0
}

func (x *RecordingReq) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *RecordingReq) GetRtmpUrl() string {
	if x != nil && x.RtmpUrl != nil {
		return *x.RtmpUrl
	}
	return ""
}

func (x *RecordingReq) GetCustomDesign() string {
	if x != nil && x.CustomDesign != nil {
		return *x.CustomDesign
	}
	return ""
}

type RecordingInfoFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomTableId      int64   `protobuf:"varint,1,opt,name=room_table_id,json=roomTableId,proto3" json:"room_table_id,omitempty"`
	RoomId           string  `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomTitle        string  `protobuf:"bytes,3,opt,name=room_title,json=roomTitle,proto3" json:"room_title,omitempty"`
	RoomSid          string  `protobuf:"bytes,4,opt,name=room_sid,json=roomSid,proto3" json:"room_sid,omitempty"`
	RoomCreationTime int64   `protobuf:"varint,5,opt,name=room_creation_time,json=roomCreationTime,proto3" json:"room_creation_time,omitempty"`
	RoomEnded        int64   `protobuf:"varint,6,opt,name=room_ended,json=roomEnded,proto3" json:"room_ended,omitempty"`
	RecordingId      string  `protobuf:"bytes,7,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	RecorderId       string  `protobuf:"bytes,8,opt,name=recorder_id,json=recorderId,proto3" json:"recorder_id,omitempty"`
	FilePath         string  `protobuf:"bytes,9,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileSize         float32 `protobuf:"fixed32,10,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	CreationTime     int64   `protobuf:"varint,11,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
}

func (x *RecordingInfoFile) Reset() {
	*x = RecordingInfoFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugnmeet_recording_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordingInfoFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordingInfoFile) ProtoMessage() {}

func (x *RecordingInfoFile) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_recording_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordingInfoFile.ProtoReflect.Descriptor instead.
func (*RecordingInfoFile) Descriptor() ([]byte, []int) {
	return file_plugnmeet_recording_proto_rawDescGZIP(), []int{1}
}

func (x *RecordingInfoFile) GetRoomTableId() int64 {
	if x != nil {
		return x.RoomTableId
	}
	return 0
}

func (x *RecordingInfoFile) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RecordingInfoFile) GetRoomTitle() string {
	if x != nil {
		return x.RoomTitle
	}
	return ""
}

func (x *RecordingInfoFile) GetRoomSid() string {
	if x != nil {
		return x.RoomSid
	}
	return ""
}

func (x *RecordingInfoFile) GetRoomCreationTime() int64 {
	if x != nil {
		return x.RoomCreationTime
	}
	return 0
}

func (x *RecordingInfoFile) GetRoomEnded() int64 {
	if x != nil {
		return x.RoomEnded
	}
	return 0
}

func (x *RecordingInfoFile) GetRecordingId() string {
	if x != nil {
		return x.RecordingId
	}
	return ""
}

func (x *RecordingInfoFile) GetRecorderId() string {
	if x != nil {
		return x.RecorderId
	}
	return ""
}

func (x *RecordingInfoFile) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *RecordingInfoFile) GetFileSize() float32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *RecordingInfoFile) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

var File_plugnmeet_recording_proto protoreflect.FileDescriptor

var file_plugnmeet_recording_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75,
	0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x1a, 0x18, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65,
	0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf5, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x08, 0x72, 0x74, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x72, 0x74, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x28, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x74,
	0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x22, 0xfa, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x53, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x45, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_recording_proto_rawDescOnce sync.Once
	file_plugnmeet_recording_proto_rawDescData = file_plugnmeet_recording_proto_rawDesc
)

func file_plugnmeet_recording_proto_rawDescGZIP() []byte {
	file_plugnmeet_recording_proto_rawDescOnce.Do(func() {
		file_plugnmeet_recording_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_recording_proto_rawDescData)
	})
	return file_plugnmeet_recording_proto_rawDescData
}

var file_plugnmeet_recording_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_plugnmeet_recording_proto_goTypes = []interface{}{
	(*RecordingReq)(nil),      // 0: plugnmeet.RecordingReq
	(*RecordingInfoFile)(nil), // 1: plugnmeet.RecordingInfoFile
	(RecordingTasks)(0),       // 2: plugnmeet.RecordingTasks
}
var file_plugnmeet_recording_proto_depIdxs = []int32{
	2, // 0: plugnmeet.RecordingReq.task:type_name -> plugnmeet.RecordingTasks
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_plugnmeet_recording_proto_init() }
func file_plugnmeet_recording_proto_init() {
	if File_plugnmeet_recording_proto != nil {
		return
	}
	file_plugnmeet_recorder_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plugnmeet_recording_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingReq); i {
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
		file_plugnmeet_recording_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordingInfoFile); i {
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
	file_plugnmeet_recording_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_recording_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_recording_proto_goTypes,
		DependencyIndexes: file_plugnmeet_recording_proto_depIdxs,
		MessageInfos:      file_plugnmeet_recording_proto_msgTypes,
	}.Build()
	File_plugnmeet_recording_proto = out.File
	file_plugnmeet_recording_proto_rawDesc = nil
	file_plugnmeet_recording_proto_goTypes = nil
	file_plugnmeet_recording_proto_depIdxs = nil
}
