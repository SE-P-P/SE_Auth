// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: signIn.proto

package signIn

import (
	_ "SEproject/biz/application/dto/http"
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

type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email      string  `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`
	Password   *string `protobuf:"bytes,2,opt,name=password,proto3,oneof" json:"password,omitempty" form:"password" query:"password"`
	Code       *string `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty" form:"code" query:"code"`
	SignInType string  `protobuf:"bytes,4,opt,name=signInType,proto3" json:"signInType,omitempty" form:"signInType" query:"signInType"`
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{0}
}

func (x *SignInReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInReq) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *SignInReq) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *SignInReq) GetSignInType() string {
	if x != nil {
		return x.SignInType
	}
	return ""
}

type SignInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty" form:"userId" query:"userId"`
}

func (x *SignInResp) Reset() {
	*x = SignInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResp) ProtoMessage() {}

func (x *SignInResp) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResp.ProtoReflect.Descriptor instead.
func (*SignInResp) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{1}
}

func (x *SignInResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" form:"password" query:"password"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty" form:"userId" query:"userId"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ResetPwReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" form:"password" query:"password"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
}

func (x *ResetPwReq) Reset() {
	*x = ResetPwReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPwReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPwReq) ProtoMessage() {}

func (x *ResetPwReq) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPwReq.ProtoReflect.Descriptor instead.
func (*ResetPwReq) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{4}
}

func (x *ResetPwReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResetPwReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ResetPwReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ResetPwResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetPwResp) Reset() {
	*x = ResetPwResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPwResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPwResp) ProtoMessage() {}

func (x *ResetPwResp) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPwResp.ProtoReflect.Descriptor instead.
func (*ResetPwResp) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{5}
}

type SendCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`
}

func (x *SendCodeReq) Reset() {
	*x = SendCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCodeReq) ProtoMessage() {}

func (x *SendCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCodeReq.ProtoReflect.Descriptor instead.
func (*SendCodeReq) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{6}
}

func (x *SendCodeReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendCodeResp) Reset() {
	*x = SendCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCodeResp) ProtoMessage() {}

func (x *SendCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCodeResp.ProtoReflect.Descriptor instead.
func (*SendCodeResp) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{7}
}

type SaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db      string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty" form:"db" query:"db"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty" form:"content" query:"content"`
}

func (x *SaveReq) Reset() {
	*x = SaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveReq) ProtoMessage() {}

func (x *SaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveReq.ProtoReflect.Descriptor instead.
func (*SaveReq) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{8}
}

func (x *SaveReq) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *SaveReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SaveReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SaveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveResp) Reset() {
	*x = SaveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResp) ProtoMessage() {}

func (x *SaveResp) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResp.ProtoReflect.Descriptor instead.
func (*SaveResp) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{9}
}

type LoadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
}

func (x *LoadReq) Reset() {
	*x = LoadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadReq) ProtoMessage() {}

func (x *LoadReq) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadReq.ProtoReflect.Descriptor instead.
func (*LoadReq) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{10}
}

func (x *LoadReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LoadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []string `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty" form:"content" query:"content"`
}

func (x *LoadResp) Reset() {
	*x = LoadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signIn_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadResp) ProtoMessage() {}

func (x *LoadResp) ProtoReflect() protoreflect.Message {
	mi := &file_signIn_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadResp.ProtoReflect.Descriptor instead.
func (*LoadResp) Descriptor() ([]byte, []int) {
	return file_signIn_proto_rawDescGZIP(), []int{11}
}

func (x *LoadResp) GetContent() []string {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_signIn_proto protoreflect.FileDescriptor

var file_signIn_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x1a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x0b,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x0a, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x77, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0d, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x77, 0x52, 0x65, 0x73, 0x70, 0x22, 0x23, 0x0a, 0x0b,
	0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x0e, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x43, 0x0a, 0x07, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x0a, 0x0a, 0x08, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x19, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a,
	0x08, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x32, 0x8a, 0x03, 0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12,
	0x11, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x73, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0c, 0xca, 0xc1, 0x18, 0x08, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x69, 0x6e, 0x12, 0x44, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x13, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0xca, 0xc1, 0x18,
	0x09, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x77, 0x12, 0x12, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x50, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x73, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x77, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d,
	0xca, 0xc1, 0x18, 0x09, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x77, 0x12, 0x45, 0x0a,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x0e, 0xca, 0xc1, 0x18, 0x0a, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x61, 0x76, 0x65, 0x12, 0x0f, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x09, 0xd2, 0xc1, 0x18, 0x05, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x0f, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x09, 0xca, 0xc1, 0x18, 0x05, 0x2f, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x15, 0x5a, 0x13, 0x53, 0x65, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signIn_proto_rawDescOnce sync.Once
	file_signIn_proto_rawDescData = file_signIn_proto_rawDesc
)

func file_signIn_proto_rawDescGZIP() []byte {
	file_signIn_proto_rawDescOnce.Do(func() {
		file_signIn_proto_rawDescData = protoimpl.X.CompressGZIP(file_signIn_proto_rawDescData)
	})
	return file_signIn_proto_rawDescData
}

var file_signIn_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_signIn_proto_goTypes = []interface{}{
	(*SignInReq)(nil),    // 0: signIn.signInReq
	(*SignInResp)(nil),   // 1: signIn.signInResp
	(*RegisterReq)(nil),  // 2: signIn.registerReq
	(*RegisterResp)(nil), // 3: signIn.registerResp
	(*ResetPwReq)(nil),   // 4: signIn.resetPwReq
	(*ResetPwResp)(nil),  // 5: signIn.resetPwResp
	(*SendCodeReq)(nil),  // 6: signIn.sendCodeReq
	(*SendCodeResp)(nil), // 7: signIn.sendCodeResp
	(*SaveReq)(nil),      // 8: signIn.saveReq
	(*SaveResp)(nil),     // 9: signIn.saveResp
	(*LoadReq)(nil),      // 10: signIn.loadReq
	(*LoadResp)(nil),     // 11: signIn.loadResp
}
var file_signIn_proto_depIdxs = []int32{
	0,  // 0: signIn.signInService.signIn:input_type -> signIn.signInReq
	2,  // 1: signIn.signInService.register:input_type -> signIn.registerReq
	4,  // 2: signIn.signInService.resetPw:input_type -> signIn.resetPwReq
	6,  // 3: signIn.signInService.sendCode:input_type -> signIn.sendCodeReq
	8,  // 4: signIn.signInService.save:input_type -> signIn.saveReq
	10, // 5: signIn.signInService.load:input_type -> signIn.loadReq
	1,  // 6: signIn.signInService.signIn:output_type -> signIn.signInResp
	3,  // 7: signIn.signInService.register:output_type -> signIn.registerResp
	5,  // 8: signIn.signInService.resetPw:output_type -> signIn.resetPwResp
	7,  // 9: signIn.signInService.sendCode:output_type -> signIn.sendCodeResp
	9,  // 10: signIn.signInService.save:output_type -> signIn.saveResp
	11, // 11: signIn.signInService.load:output_type -> signIn.loadResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_signIn_proto_init() }
func file_signIn_proto_init() {
	if File_signIn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signIn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInReq); i {
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
		file_signIn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResp); i {
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
		file_signIn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_signIn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_signIn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPwReq); i {
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
		file_signIn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPwResp); i {
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
		file_signIn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCodeReq); i {
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
		file_signIn_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCodeResp); i {
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
		file_signIn_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveReq); i {
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
		file_signIn_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResp); i {
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
		file_signIn_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadReq); i {
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
		file_signIn_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadResp); i {
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
	file_signIn_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_signIn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_signIn_proto_goTypes,
		DependencyIndexes: file_signIn_proto_depIdxs,
		MessageInfos:      file_signIn_proto_msgTypes,
	}.Build()
	File_signIn_proto = out.File
	file_signIn_proto_rawDesc = nil
	file_signIn_proto_goTypes = nil
	file_signIn_proto_depIdxs = nil
}
