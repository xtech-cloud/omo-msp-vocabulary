// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.2
// source: proto/vocabulary/common.proto

package vocabulary

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ResultStatus int32

const (
	ResultStatus_Success     ResultStatus = 0
	ResultStatus_MaxLimit    ResultStatus = 1
	ResultStatus_Repeated    ResultStatus = 2
	ResultStatus_NotExisted  ResultStatus = 3
	ResultStatus_DBException ResultStatus = 4
	ResultStatus_Empty       ResultStatus = 5
)

// Enum value maps for ResultStatus.
var (
	ResultStatus_name = map[int32]string{
		0: "Success",
		1: "MaxLimit",
		2: "Repeated",
		3: "NotExisted",
		4: "DBException",
		5: "Empty",
	}
	ResultStatus_value = map[string]int32{
		"Success":     0,
		"MaxLimit":    1,
		"Repeated":    2,
		"NotExisted":  3,
		"DBException": 4,
		"Empty":       5,
	}
)

func (x ResultStatus) Enum() *ResultStatus {
	p := new(ResultStatus)
	*p = x
	return p
}

func (x ResultStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vocabulary_common_proto_enumTypes[0].Descriptor()
}

func (ResultStatus) Type() protoreflect.EnumType {
	return &file_proto_vocabulary_common_proto_enumTypes[0]
}

func (x ResultStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultStatus.Descriptor instead.
func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_vocabulary_common_proto_rawDescGZIP(), []int{0}
}

type DirectionType int32

const (
	DirectionType_Double DirectionType = 0
	DirectionType_FromTo DirectionType = 1
	DirectionType_ToFrom DirectionType = 2
)

// Enum value maps for DirectionType.
var (
	DirectionType_name = map[int32]string{
		0: "Double",
		1: "FromTo",
		2: "ToFrom",
	}
	DirectionType_value = map[string]int32{
		"Double": 0,
		"FromTo": 1,
		"ToFrom": 2,
	}
)

func (x DirectionType) Enum() *DirectionType {
	p := new(DirectionType)
	*p = x
	return p
}

func (x DirectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DirectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vocabulary_common_proto_enumTypes[1].Descriptor()
}

func (DirectionType) Type() protoreflect.EnumType {
	return &file_proto_vocabulary_common_proto_enumTypes[1]
}

func (x DirectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DirectionType.Descriptor instead.
func (DirectionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_vocabulary_common_proto_rawDescGZIP(), []int{1}
}

type RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Id       uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Operator string `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (x *RequestInfo) Reset() {
	*x = RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInfo) ProtoMessage() {}

func (x *RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInfo.ProtoReflect.Descriptor instead.
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_common_proto_rawDescGZIP(), []int{0}
}

func (x *RequestInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RequestInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RequestInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RequestInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

type RequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List     []string `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *RequestList) Reset() {
	*x = RequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestList) ProtoMessage() {}

func (x *RequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestList.ProtoReflect.Descriptor instead.
func (*RequestList) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_common_proto_rawDescGZIP(), []int{1}
}

func (x *RequestList) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RequestList) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *RequestList) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type ReplyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Key    string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Id     uint64       `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Status *ReplyStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReplyInfo) Reset() {
	*x = ReplyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyInfo) ProtoMessage() {}

func (x *ReplyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyInfo.ProtoReflect.Descriptor instead.
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_common_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ReplyInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReplyInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReplyInfo) GetStatus() *ReplyStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ReplyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ReplyStatus) Reset() {
	*x = ReplyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyStatus) ProtoMessage() {}

func (x *ReplyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyStatus.ProtoReflect.Descriptor instead.
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_common_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyStatus) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ReplyStatus) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AttributeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id       uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type     uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Key      string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Begin    string `protobuf:"bytes,6,opt,name=begin,proto3" json:"begin,omitempty"`
	End      string `protobuf:"bytes,7,opt,name=end,proto3" json:"end,omitempty"`
	Remark   string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Time     int64  `protobuf:"varint,9,opt,name=time,proto3" json:"time,omitempty"`
	Updated  int64  `protobuf:"varint,10,opt,name=updated,proto3" json:"updated,omitempty"`
	Created  int64  `protobuf:"varint,11,opt,name=created,proto3" json:"created,omitempty"`
	Creator  string `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator string `protobuf:"bytes,13,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (x *AttributeInfo) Reset() {
	*x = AttributeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeInfo) ProtoMessage() {}

func (x *AttributeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeInfo.ProtoReflect.Descriptor instead.
func (*AttributeInfo) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_common_proto_rawDescGZIP(), []int{4}
}

func (x *AttributeInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AttributeInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttributeInfo) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AttributeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttributeInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AttributeInfo) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *AttributeInfo) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *AttributeInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *AttributeInfo) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *AttributeInfo) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *AttributeInfo) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *AttributeInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *AttributeInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

var File_proto_vocabulary_common_proto protoreflect.FileDescriptor

var file_proto_vocabulary_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c,
	0x61, 0x72, 0x79, 0x22, 0x5d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x22, 0x4f, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76,
	0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x33, 0x0a,
	0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0xa9, 0x02, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2a, 0x63,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x42, 0x45, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x10, 0x05, 0x2a, 0x33, 0x0a, 0x0d, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x54, 0x6f, 0x46, 0x72, 0x6f, 0x6d, 0x10, 0x02, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vocabulary_common_proto_rawDescOnce sync.Once
	file_proto_vocabulary_common_proto_rawDescData = file_proto_vocabulary_common_proto_rawDesc
)

func file_proto_vocabulary_common_proto_rawDescGZIP() []byte {
	file_proto_vocabulary_common_proto_rawDescOnce.Do(func() {
		file_proto_vocabulary_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vocabulary_common_proto_rawDescData)
	})
	return file_proto_vocabulary_common_proto_rawDescData
}

var file_proto_vocabulary_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_vocabulary_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_vocabulary_common_proto_goTypes = []interface{}{
	(ResultStatus)(0),     // 0: omo.msp.vocabulary.ResultStatus
	(DirectionType)(0),    // 1: omo.msp.vocabulary.DirectionType
	(*RequestInfo)(nil),   // 2: omo.msp.vocabulary.RequestInfo
	(*RequestList)(nil),   // 3: omo.msp.vocabulary.RequestList
	(*ReplyInfo)(nil),     // 4: omo.msp.vocabulary.ReplyInfo
	(*ReplyStatus)(nil),   // 5: omo.msp.vocabulary.ReplyStatus
	(*AttributeInfo)(nil), // 6: omo.msp.vocabulary.AttributeInfo
}
var file_proto_vocabulary_common_proto_depIdxs = []int32{
	5, // 0: omo.msp.vocabulary.ReplyInfo.status:type_name -> omo.msp.vocabulary.ReplyStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_vocabulary_common_proto_init() }
func file_proto_vocabulary_common_proto_init() {
	if File_proto_vocabulary_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vocabulary_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInfo); i {
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
		file_proto_vocabulary_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestList); i {
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
		file_proto_vocabulary_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyInfo); i {
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
		file_proto_vocabulary_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyStatus); i {
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
		file_proto_vocabulary_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeInfo); i {
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
			RawDescriptor: file_proto_vocabulary_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_vocabulary_common_proto_goTypes,
		DependencyIndexes: file_proto_vocabulary_common_proto_depIdxs,
		EnumInfos:         file_proto_vocabulary_common_proto_enumTypes,
		MessageInfos:      file_proto_vocabulary_common_proto_msgTypes,
	}.Build()
	File_proto_vocabulary_common_proto = out.File
	file_proto_vocabulary_common_proto_rawDesc = nil
	file_proto_vocabulary_common_proto_goTypes = nil
	file_proto_vocabulary_common_proto_depIdxs = nil
}
