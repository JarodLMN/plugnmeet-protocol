// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: plugnmeet_gen_token.proto

package plugnmeet

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type GenerateTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserInfo      *UserInfo              `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenReq) Reset() {
	*x = GenerateTokenReq{}
	mi := &file_plugnmeet_gen_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenReq) ProtoMessage() {}

func (x *GenerateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_gen_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenReq.ProtoReflect.Descriptor instead.
func (*GenerateTokenReq) Descriptor() ([]byte, []int) {
	return file_plugnmeet_gen_token_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenReq) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GenerateTokenReq) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsAdmin       bool                   `protobuf:"varint,3,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	IsHidden      bool                   `protobuf:"varint,4,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	UserMetadata  *UserMetadata          `protobuf:"bytes,5,opt,name=user_metadata,json=userMetadata,proto3" json:"user_metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_plugnmeet_gen_token_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_gen_token_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_plugnmeet_gen_token_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfo) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *UserInfo) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

func (x *UserInfo) GetUserMetadata() *UserMetadata {
	if x != nil {
		return x.UserMetadata
	}
	return nil
}

type UserMetadata struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ProfilePic      *string                `protobuf:"bytes,1,opt,name=profile_pic,json=profilePic,proto3,oneof" json:"profile_pic,omitempty"`
	IsAdmin         bool                   `protobuf:"varint,2,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	RecordWebcam    *bool                  `protobuf:"varint,3,opt,name=record_webcam,json=recordWebcam,proto3,oneof" json:"record_webcam,omitempty"`
	IsPresenter     bool                   `protobuf:"varint,4,opt,name=is_presenter,json=isPresenter,proto3" json:"is_presenter,omitempty"`
	RaisedHand      bool                   `protobuf:"varint,5,opt,name=raised_hand,json=raisedHand,proto3" json:"raised_hand,omitempty"`
	WaitForApproval bool                   `protobuf:"varint,6,opt,name=wait_for_approval,json=waitForApproval,proto3" json:"wait_for_approval,omitempty"`
	PreferredLang   *string                `protobuf:"bytes,7,opt,name=preferred_lang,json=preferredLang,proto3,oneof" json:"preferred_lang,omitempty"`
	LockSettings    *LockSettings          `protobuf:"bytes,8,opt,name=lock_settings,json=lockSettings,proto3" json:"lock_settings,omitempty"`
	// we'll use this as version control
	MetadataId *string `protobuf:"bytes,9,opt,name=metadata_id,json=metadataId,proto3,oneof" json:"metadata_id,omitempty"`
	ExtraData  *string `protobuf:"bytes,10,opt,name=extra_data,json=extraData,proto3,oneof" json:"extra_data,omitempty"`
	// if remain empty then will use value from user_id
	ExUserId      *string `protobuf:"bytes,11,opt,name=ex_user_id,json=exUserId,proto3,oneof" json:"ex_user_id,omitempty"` // Next ID: 12
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserMetadata) Reset() {
	*x = UserMetadata{}
	mi := &file_plugnmeet_gen_token_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMetadata) ProtoMessage() {}

func (x *UserMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_gen_token_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMetadata.ProtoReflect.Descriptor instead.
func (*UserMetadata) Descriptor() ([]byte, []int) {
	return file_plugnmeet_gen_token_proto_rawDescGZIP(), []int{2}
}

func (x *UserMetadata) GetProfilePic() string {
	if x != nil && x.ProfilePic != nil {
		return *x.ProfilePic
	}
	return ""
}

func (x *UserMetadata) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *UserMetadata) GetRecordWebcam() bool {
	if x != nil && x.RecordWebcam != nil {
		return *x.RecordWebcam
	}
	return false
}

func (x *UserMetadata) GetIsPresenter() bool {
	if x != nil {
		return x.IsPresenter
	}
	return false
}

func (x *UserMetadata) GetRaisedHand() bool {
	if x != nil {
		return x.RaisedHand
	}
	return false
}

func (x *UserMetadata) GetWaitForApproval() bool {
	if x != nil {
		return x.WaitForApproval
	}
	return false
}

func (x *UserMetadata) GetPreferredLang() string {
	if x != nil && x.PreferredLang != nil {
		return *x.PreferredLang
	}
	return ""
}

func (x *UserMetadata) GetLockSettings() *LockSettings {
	if x != nil {
		return x.LockSettings
	}
	return nil
}

func (x *UserMetadata) GetMetadataId() string {
	if x != nil && x.MetadataId != nil {
		return *x.MetadataId
	}
	return ""
}

func (x *UserMetadata) GetExtraData() string {
	if x != nil && x.ExtraData != nil {
		return *x.ExtraData
	}
	return ""
}

func (x *UserMetadata) GetExUserId() string {
	if x != nil && x.ExUserId != nil {
		return *x.ExUserId
	}
	return ""
}

type PlugNmeetTokenClaims struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoomId        string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	IsAdmin       bool                   `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	IsHidden      bool                   `protobuf:"varint,5,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlugNmeetTokenClaims) Reset() {
	*x = PlugNmeetTokenClaims{}
	mi := &file_plugnmeet_gen_token_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlugNmeetTokenClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlugNmeetTokenClaims) ProtoMessage() {}

func (x *PlugNmeetTokenClaims) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_gen_token_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlugNmeetTokenClaims.ProtoReflect.Descriptor instead.
func (*PlugNmeetTokenClaims) Descriptor() ([]byte, []int) {
	return file_plugnmeet_gen_token_proto_rawDescGZIP(), []int{3}
}

func (x *PlugNmeetTokenClaims) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlugNmeetTokenClaims) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PlugNmeetTokenClaims) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *PlugNmeetTokenClaims) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *PlugNmeetTokenClaims) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

type LockSettings struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	LockMicrophone      *bool                  `protobuf:"varint,1,opt,name=lock_microphone,json=lockMicrophone,proto3,oneof" json:"lock_microphone,omitempty"`
	LockWebcam          *bool                  `protobuf:"varint,2,opt,name=lock_webcam,json=lockWebcam,proto3,oneof" json:"lock_webcam,omitempty"`
	LockScreenSharing   *bool                  `protobuf:"varint,3,opt,name=lock_screen_sharing,json=lockScreenSharing,proto3,oneof" json:"lock_screen_sharing,omitempty"`
	LockChat            *bool                  `protobuf:"varint,4,opt,name=lock_chat,json=lockChat,proto3,oneof" json:"lock_chat,omitempty"`
	LockChatSendMessage *bool                  `protobuf:"varint,5,opt,name=lock_chat_send_message,json=lockChatSendMessage,proto3,oneof" json:"lock_chat_send_message,omitempty"`
	LockChatFileShare   *bool                  `protobuf:"varint,6,opt,name=lock_chat_file_share,json=lockChatFileShare,proto3,oneof" json:"lock_chat_file_share,omitempty"`
	LockPrivateChat     *bool                  `protobuf:"varint,7,opt,name=lock_private_chat,json=lockPrivateChat,proto3,oneof" json:"lock_private_chat,omitempty"`
	LockWhiteboard      *bool                  `protobuf:"varint,8,opt,name=lock_whiteboard,json=lockWhiteboard,proto3,oneof" json:"lock_whiteboard,omitempty"`
	LockSharedNotepad   *bool                  `protobuf:"varint,9,opt,name=lock_shared_notepad,json=lockSharedNotepad,proto3,oneof" json:"lock_shared_notepad,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *LockSettings) Reset() {
	*x = LockSettings{}
	mi := &file_plugnmeet_gen_token_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LockSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockSettings) ProtoMessage() {}

func (x *LockSettings) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_gen_token_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockSettings.ProtoReflect.Descriptor instead.
func (*LockSettings) Descriptor() ([]byte, []int) {
	return file_plugnmeet_gen_token_proto_rawDescGZIP(), []int{4}
}

func (x *LockSettings) GetLockMicrophone() bool {
	if x != nil && x.LockMicrophone != nil {
		return *x.LockMicrophone
	}
	return false
}

func (x *LockSettings) GetLockWebcam() bool {
	if x != nil && x.LockWebcam != nil {
		return *x.LockWebcam
	}
	return false
}

func (x *LockSettings) GetLockScreenSharing() bool {
	if x != nil && x.LockScreenSharing != nil {
		return *x.LockScreenSharing
	}
	return false
}

func (x *LockSettings) GetLockChat() bool {
	if x != nil && x.LockChat != nil {
		return *x.LockChat
	}
	return false
}

func (x *LockSettings) GetLockChatSendMessage() bool {
	if x != nil && x.LockChatSendMessage != nil {
		return *x.LockChatSendMessage
	}
	return false
}

func (x *LockSettings) GetLockChatFileShare() bool {
	if x != nil && x.LockChatFileShare != nil {
		return *x.LockChatFileShare
	}
	return false
}

func (x *LockSettings) GetLockPrivateChat() bool {
	if x != nil && x.LockPrivateChat != nil {
		return *x.LockPrivateChat
	}
	return false
}

func (x *LockSettings) GetLockWhiteboard() bool {
	if x != nil && x.LockWhiteboard != nil {
		return *x.LockWhiteboard
	}
	return false
}

func (x *LockSettings) GetLockSharedNotepad() bool {
	if x != nil && x.LockSharedNotepad != nil {
		return *x.LockSharedNotepad
	}
	return false
}

type GenerateTokenRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token         *string                `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateTokenRes) Reset() {
	*x = GenerateTokenRes{}
	mi := &file_plugnmeet_gen_token_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRes) ProtoMessage() {}

func (x *GenerateTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_gen_token_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRes.ProtoReflect.Descriptor instead.
func (*GenerateTokenRes) Descriptor() ([]byte, []int) {
	return file_plugnmeet_gen_token_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateTokenRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GenerateTokenRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GenerateTokenRes) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type CustomDesignParams struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	PrimaryColor     *string                `protobuf:"bytes,1,opt,name=primary_color,json=primaryColor,proto3,oneof" json:"primary_color,omitempty"`
	SecondaryColor   *string                `protobuf:"bytes,2,opt,name=secondary_color,json=secondaryColor,proto3,oneof" json:"secondary_color,omitempty"`
	BackgroundColor  *string                `protobuf:"bytes,3,opt,name=background_color,json=backgroundColor,proto3,oneof" json:"background_color,omitempty"`
	BackgroundImage  *string                `protobuf:"bytes,4,opt,name=background_image,json=backgroundImage,proto3,oneof" json:"background_image,omitempty"`
	HeaderBgColor    *string                `protobuf:"bytes,5,opt,name=header_bg_color,json=headerBgColor,proto3,oneof" json:"header_bg_color,omitempty"`
	FooterBgColor    *string                `protobuf:"bytes,6,opt,name=footer_bg_color,json=footerBgColor,proto3,oneof" json:"footer_bg_color,omitempty"`
	LeftSideBgColor  *string                `protobuf:"bytes,7,opt,name=left_side_bg_color,json=leftSideBgColor,proto3,oneof" json:"left_side_bg_color,omitempty"`
	RightSideBgColor *string                `protobuf:"bytes,8,opt,name=right_side_bg_color,json=rightSideBgColor,proto3,oneof" json:"right_side_bg_color,omitempty"`
	CustomCssUrl     *string                `protobuf:"bytes,9,opt,name=custom_css_url,json=customCssUrl,proto3,oneof" json:"custom_css_url,omitempty"`
	CustomLogo       *string                `protobuf:"bytes,10,opt,name=custom_logo,json=customLogo,proto3,oneof" json:"custom_logo,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CustomDesignParams) Reset() {
	*x = CustomDesignParams{}
	mi := &file_plugnmeet_gen_token_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomDesignParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDesignParams) ProtoMessage() {}

func (x *CustomDesignParams) ProtoReflect() protoreflect.Message {
	mi := &file_plugnmeet_gen_token_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDesignParams.ProtoReflect.Descriptor instead.
func (*CustomDesignParams) Descriptor() ([]byte, []int) {
	return file_plugnmeet_gen_token_proto_rawDescGZIP(), []int{6}
}

func (x *CustomDesignParams) GetPrimaryColor() string {
	if x != nil && x.PrimaryColor != nil {
		return *x.PrimaryColor
	}
	return ""
}

func (x *CustomDesignParams) GetSecondaryColor() string {
	if x != nil && x.SecondaryColor != nil {
		return *x.SecondaryColor
	}
	return ""
}

func (x *CustomDesignParams) GetBackgroundColor() string {
	if x != nil && x.BackgroundColor != nil {
		return *x.BackgroundColor
	}
	return ""
}

func (x *CustomDesignParams) GetBackgroundImage() string {
	if x != nil && x.BackgroundImage != nil {
		return *x.BackgroundImage
	}
	return ""
}

func (x *CustomDesignParams) GetHeaderBgColor() string {
	if x != nil && x.HeaderBgColor != nil {
		return *x.HeaderBgColor
	}
	return ""
}

func (x *CustomDesignParams) GetFooterBgColor() string {
	if x != nil && x.FooterBgColor != nil {
		return *x.FooterBgColor
	}
	return ""
}

func (x *CustomDesignParams) GetLeftSideBgColor() string {
	if x != nil && x.LeftSideBgColor != nil {
		return *x.LeftSideBgColor
	}
	return ""
}

func (x *CustomDesignParams) GetRightSideBgColor() string {
	if x != nil && x.RightSideBgColor != nil {
		return *x.RightSideBgColor
	}
	return ""
}

func (x *CustomDesignParams) GetCustomCssUrl() string {
	if x != nil && x.CustomCssUrl != nil {
		return *x.CustomCssUrl
	}
	return ""
}

func (x *CustomDesignParams) GetCustomLogo() string {
	if x != nil && x.CustomLogo != nil {
		return *x.CustomLogo
	}
	return ""
}

var File_plugnmeet_gen_token_proto protoreflect.FileDescriptor

var file_plugnmeet_gen_token_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x75,
	0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x94, 0x01, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x7b, 0xba, 0x48, 0x78, 0xba,
	0x01, 0x75, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x41, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x20, 0x73, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x20, 0x28, 0x61, 0x2d, 0x7a, 0x20, 0x41, 0x2d, 0x5a,
	0x29, 0x2c, 0x20, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x20, 0x28, 0x30, 0x2d, 0x39, 0x29, 0x20,
	0x6f, 0x72, 0x20, 0x2d, 0x5f, 0x1a, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39,
	0x2d, 0x5f, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x12, 0x3c, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xa2, 0x05, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2e, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x88, 0x01,
	0x01, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x28, 0x0a,
	0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x77, 0x65, 0x62, 0x63, 0x61, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x57, 0x65,
	0x62, 0x63, 0x61, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x6a, 0x02, 0x08, 0x00, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0b, 0x72, 0x61, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x68, 0x61,
	0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xba, 0x48, 0x04, 0x6a, 0x02, 0x08,
	0x00, 0x52, 0x0a, 0x72, 0x61, 0x69, 0x73, 0x65, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x33, 0x0a,
	0x11, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xba, 0x48, 0x04, 0x6a, 0x02, 0x08,
	0x00, 0x52, 0x0f, 0x77, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f,
	0x6c, 0x61, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0d, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x3c,
	0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65,
	0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0c,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x7e, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x58, 0xba, 0x48, 0x55, 0xba, 0x01, 0x52, 0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x28, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c,
	0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x61, 0x6e,
	0x79, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x24, 0x27, 0x29, 0x48, 0x03, 0x52, 0x0a, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x04, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x21, 0x0a, 0x0a, 0x65, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x65, 0x78, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x69, 0x63, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x77,
	0x65, 0x62, 0x63, 0x61, 0x6d, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x50, 0x6c, 0x75, 0x67, 0x4e,
	0x6d, 0x65, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0xfd, 0x04,
	0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2c,
	0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x6b, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x77, 0x65, 0x62, 0x63, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x65, 0x62, 0x63, 0x61, 0x6d, 0x88,
	0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x68, 0x61,
	0x72, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x63, 0x68, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x6b, 0x43, 0x68, 0x61, 0x74, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x16, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x13, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x14, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x05, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x0f, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x07, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x6b, 0x57, 0x68, 0x69, 0x74, 0x65,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x70, 0x61, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x08, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x65, 0x70, 0x61, 0x64, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x77, 0x65, 0x62, 0x63, 0x61,
	0x6d, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x70, 0x61, 0x64, 0x22, 0x61, 0x0a,
	0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xa7, 0x05, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x2e, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0f, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x2e, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0f, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x42, 0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f,
	0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x5f, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0d, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x42,
	0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x12, 0x6c, 0x65, 0x66,
	0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0f, 0x6c, 0x65, 0x66, 0x74, 0x53, 0x69, 0x64,
	0x65, 0x42, 0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x13, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x10, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x53, 0x69, 0x64, 0x65, 0x42, 0x67, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x29, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x73, 0x73, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x43, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x6f, 0x67, 0x6f, 0x88, 0x01, 0x01,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x67, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x5f,
	0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x6c, 0x65, 0x66,
	0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x62, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42,
	0x16, 0x0a, 0x14, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x62,
	0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x63, 0x73, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x42, 0x9f, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x42, 0x16, 0x50, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x79, 0x6e, 0x61, 0x70, 0x61, 0x72, 0x72, 0x6f, 0x74, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xca, 0x02, 0x09, 0x50,
	0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0xe2, 0x02, 0x15, 0x50, 0x6c, 0x75, 0x67, 0x6e,
	0x6d, 0x65, 0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x09, 0x50, 0x6c, 0x75, 0x67, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugnmeet_gen_token_proto_rawDescOnce sync.Once
	file_plugnmeet_gen_token_proto_rawDescData = file_plugnmeet_gen_token_proto_rawDesc
)

func file_plugnmeet_gen_token_proto_rawDescGZIP() []byte {
	file_plugnmeet_gen_token_proto_rawDescOnce.Do(func() {
		file_plugnmeet_gen_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugnmeet_gen_token_proto_rawDescData)
	})
	return file_plugnmeet_gen_token_proto_rawDescData
}

var file_plugnmeet_gen_token_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugnmeet_gen_token_proto_goTypes = []any{
	(*GenerateTokenReq)(nil),     // 0: plugnmeet.GenerateTokenReq
	(*UserInfo)(nil),             // 1: plugnmeet.UserInfo
	(*UserMetadata)(nil),         // 2: plugnmeet.UserMetadata
	(*PlugNmeetTokenClaims)(nil), // 3: plugnmeet.PlugNmeetTokenClaims
	(*LockSettings)(nil),         // 4: plugnmeet.LockSettings
	(*GenerateTokenRes)(nil),     // 5: plugnmeet.GenerateTokenRes
	(*CustomDesignParams)(nil),   // 6: plugnmeet.CustomDesignParams
}
var file_plugnmeet_gen_token_proto_depIdxs = []int32{
	1, // 0: plugnmeet.GenerateTokenReq.user_info:type_name -> plugnmeet.UserInfo
	2, // 1: plugnmeet.UserInfo.user_metadata:type_name -> plugnmeet.UserMetadata
	4, // 2: plugnmeet.UserMetadata.lock_settings:type_name -> plugnmeet.LockSettings
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_plugnmeet_gen_token_proto_init() }
func file_plugnmeet_gen_token_proto_init() {
	if File_plugnmeet_gen_token_proto != nil {
		return
	}
	file_plugnmeet_gen_token_proto_msgTypes[2].OneofWrappers = []any{}
	file_plugnmeet_gen_token_proto_msgTypes[4].OneofWrappers = []any{}
	file_plugnmeet_gen_token_proto_msgTypes[5].OneofWrappers = []any{}
	file_plugnmeet_gen_token_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugnmeet_gen_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugnmeet_gen_token_proto_goTypes,
		DependencyIndexes: file_plugnmeet_gen_token_proto_depIdxs,
		MessageInfos:      file_plugnmeet_gen_token_proto_msgTypes,
	}.Build()
	File_plugnmeet_gen_token_proto = out.File
	file_plugnmeet_gen_token_proto_rawDesc = nil
	file_plugnmeet_gen_token_proto_goTypes = nil
	file_plugnmeet_gen_token_proto_depIdxs = nil
}
