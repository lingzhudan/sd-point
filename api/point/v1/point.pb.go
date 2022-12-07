// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: point/v1/point.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PointInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 点数编号
	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	// 点数数量
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	// 新增点数的描述
	Desc string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	// 点数新增时间
	ClickedAt int64 `protobuf:"varint,3,opt,name=clicked_at,json=clickedAt,proto3" json:"clicked_at,omitempty"`
	// 创建时间
	CreatedAt int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt int64 `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt int64 `protobuf:"varint,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *PointInfo) Reset() {
	*x = PointInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointInfo) ProtoMessage() {}

func (x *PointInfo) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointInfo.ProtoReflect.Descriptor instead.
func (*PointInfo) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{0}
}

func (x *PointInfo) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *PointInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *PointInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PointInfo) GetClickedAt() int64 {
	if x != nil {
		return x.ClickedAt
	}
	return 0
}

func (x *PointInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PointInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *PointInfo) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type CreatePointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 幂等令牌
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 新增点数信息
	Point []*PointInfo `protobuf:"bytes,2,rep,name=point,proto3" json:"point,omitempty"`
}

func (x *CreatePointsRequest) Reset() {
	*x = CreatePointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePointsRequest) ProtoMessage() {}

func (x *CreatePointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePointsRequest.ProtoReflect.Descriptor instead.
func (*CreatePointsRequest) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePointsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreatePointsRequest) GetPoint() []*PointInfo {
	if x != nil {
		return x.Point
	}
	return nil
}

type UpdatePointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 幂等令牌
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 待更新点数信息
	Point *PointInfo `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *UpdatePointRequest) Reset() {
	*x = UpdatePointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePointRequest) ProtoMessage() {}

func (x *UpdatePointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePointRequest.ProtoReflect.Descriptor instead.
func (*UpdatePointRequest) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePointRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdatePointRequest) GetPoint() *PointInfo {
	if x != nil {
		return x.Point
	}
	return nil
}

type UpdatePointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePointReply) Reset() {
	*x = UpdatePointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePointReply) ProtoMessage() {}

func (x *UpdatePointReply) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePointReply.ProtoReflect.Descriptor instead.
func (*UpdatePointReply) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{3}
}

type DeletePointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 幂等令牌
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 待删除的点数编号
	Pid int32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *DeletePointRequest) Reset() {
	*x = DeletePointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePointRequest) ProtoMessage() {}

func (x *DeletePointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePointRequest.ProtoReflect.Descriptor instead.
func (*DeletePointRequest) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePointRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeletePointRequest) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type DeletePointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePointReply) Reset() {
	*x = DeletePointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePointReply) ProtoMessage() {}

func (x *DeletePointReply) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePointReply.ProtoReflect.Descriptor instead.
func (*DeletePointReply) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{5}
}

type GetPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 幂等令牌
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 点数编号
	Pid int32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GetPointRequest) Reset() {
	*x = GetPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointRequest) ProtoMessage() {}

func (x *GetPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointRequest.ProtoReflect.Descriptor instead.
func (*GetPointRequest) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{6}
}

func (x *GetPointRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPointRequest) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type GetPointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point *PointInfo `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *GetPointReply) Reset() {
	*x = GetPointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointReply) ProtoMessage() {}

func (x *GetPointReply) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointReply.ProtoReflect.Descriptor instead.
func (*GetPointReply) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{7}
}

func (x *GetPointReply) GetPoint() *PointInfo {
	if x != nil {
		return x.Point
	}
	return nil
}

type ListPointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPointRequest) Reset() {
	*x = ListPointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPointRequest) ProtoMessage() {}

func (x *ListPointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPointRequest.ProtoReflect.Descriptor instead.
func (*ListPointRequest) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{8}
}

type ListPointReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point []*PointInfo `protobuf:"bytes,1,rep,name=point,proto3" json:"point,omitempty"`
}

func (x *ListPointReply) Reset() {
	*x = ListPointReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPointReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPointReply) ProtoMessage() {}

func (x *ListPointReply) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPointReply.ProtoReflect.Descriptor instead.
func (*ListPointReply) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{9}
}

func (x *ListPointReply) GetPoint() []*PointInfo {
	if x != nil {
		return x.Point
	}
	return nil
}

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{10}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_v1_point_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_point_v1_point_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_point_v1_point_proto_rawDescGZIP(), []int{11}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_point_v1_point_proto protoreflect.FileDescriptor

var file_point_v1_point_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x09, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x14, 0xfa, 0x42, 0x11, 0x1a, 0x0f, 0x18, 0xff, 0xff, 0x03,
	0x28, 0x80, 0x80, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x12, 0x1c, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x64, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x63, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3c, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x39, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x87, 0x04, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x5f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x1a, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x78, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x12, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x5a, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f,
	0x73, 0x61, 0x79, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x47, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x1a, 0x5a, 0x18, 0x73, 0x64, 0x2d, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_point_v1_point_proto_rawDescOnce sync.Once
	file_point_v1_point_proto_rawDescData = file_point_v1_point_proto_rawDesc
)

func file_point_v1_point_proto_rawDescGZIP() []byte {
	file_point_v1_point_proto_rawDescOnce.Do(func() {
		file_point_v1_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_point_v1_point_proto_rawDescData)
	})
	return file_point_v1_point_proto_rawDescData
}

var file_point_v1_point_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_point_v1_point_proto_goTypes = []interface{}{
	(*PointInfo)(nil),           // 0: api.point.v1.PointInfo
	(*CreatePointsRequest)(nil), // 1: api.point.v1.CreatePointsRequest
	(*UpdatePointRequest)(nil),  // 2: api.point.v1.UpdatePointRequest
	(*UpdatePointReply)(nil),    // 3: api.point.v1.UpdatePointReply
	(*DeletePointRequest)(nil),  // 4: api.point.v1.DeletePointRequest
	(*DeletePointReply)(nil),    // 5: api.point.v1.DeletePointReply
	(*GetPointRequest)(nil),     // 6: api.point.v1.GetPointRequest
	(*GetPointReply)(nil),       // 7: api.point.v1.GetPointReply
	(*ListPointRequest)(nil),    // 8: api.point.v1.ListPointRequest
	(*ListPointReply)(nil),      // 9: api.point.v1.ListPointReply
	(*HelloRequest)(nil),        // 10: api.point.v1.HelloRequest
	(*HelloReply)(nil),          // 11: api.point.v1.HelloReply
	(*emptypb.Empty)(nil),       // 12: google.protobuf.Empty
}
var file_point_v1_point_proto_depIdxs = []int32{
	0,  // 0: api.point.v1.CreatePointsRequest.point:type_name -> api.point.v1.PointInfo
	0,  // 1: api.point.v1.UpdatePointRequest.point:type_name -> api.point.v1.PointInfo
	0,  // 2: api.point.v1.GetPointReply.point:type_name -> api.point.v1.PointInfo
	0,  // 3: api.point.v1.ListPointReply.point:type_name -> api.point.v1.PointInfo
	1,  // 4: api.point.v1.Point.CreatePoints:input_type -> api.point.v1.CreatePointsRequest
	10, // 5: api.point.v1.Point.SayHello:input_type -> api.point.v1.HelloRequest
	2,  // 6: api.point.v1.Point.UpdatePoint:input_type -> api.point.v1.UpdatePointRequest
	4,  // 7: api.point.v1.Point.DeletePoint:input_type -> api.point.v1.DeletePointRequest
	6,  // 8: api.point.v1.Point.GetPoint:input_type -> api.point.v1.GetPointRequest
	8,  // 9: api.point.v1.Point.ListPoint:input_type -> api.point.v1.ListPointRequest
	12, // 10: api.point.v1.Point.CreatePoints:output_type -> google.protobuf.Empty
	11, // 11: api.point.v1.Point.SayHello:output_type -> api.point.v1.HelloReply
	12, // 12: api.point.v1.Point.UpdatePoint:output_type -> google.protobuf.Empty
	12, // 13: api.point.v1.Point.DeletePoint:output_type -> google.protobuf.Empty
	7,  // 14: api.point.v1.Point.GetPoint:output_type -> api.point.v1.GetPointReply
	9,  // 15: api.point.v1.Point.ListPoint:output_type -> api.point.v1.ListPointReply
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_point_v1_point_proto_init() }
func file_point_v1_point_proto_init() {
	if File_point_v1_point_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_point_v1_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointInfo); i {
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
		file_point_v1_point_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePointsRequest); i {
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
		file_point_v1_point_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePointRequest); i {
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
		file_point_v1_point_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePointReply); i {
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
		file_point_v1_point_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePointRequest); i {
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
		file_point_v1_point_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePointReply); i {
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
		file_point_v1_point_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPointRequest); i {
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
		file_point_v1_point_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPointReply); i {
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
		file_point_v1_point_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPointRequest); i {
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
		file_point_v1_point_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPointReply); i {
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
		file_point_v1_point_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_point_v1_point_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
			RawDescriptor: file_point_v1_point_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_point_v1_point_proto_goTypes,
		DependencyIndexes: file_point_v1_point_proto_depIdxs,
		MessageInfos:      file_point_v1_point_proto_msgTypes,
	}.Build()
	File_point_v1_point_proto = out.File
	file_point_v1_point_proto_rawDesc = nil
	file_point_v1_point_proto_goTypes = nil
	file_point_v1_point_proto_depIdxs = nil
}
