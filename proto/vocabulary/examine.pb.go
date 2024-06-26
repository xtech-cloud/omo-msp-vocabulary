// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/examine.proto

package vocabulary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExamineInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Target               string   `protobuf:"bytes,7,opt,name=target,proto3" json:"target,omitempty"`
	Status               uint32   `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Type                 uint32   `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Key                  string   `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,11,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExamineInfo) Reset()         { *m = ExamineInfo{} }
func (m *ExamineInfo) String() string { return proto.CompactTextString(m) }
func (*ExamineInfo) ProtoMessage()    {}
func (*ExamineInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be56c51a90ed54b, []int{0}
}

func (m *ExamineInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamineInfo.Unmarshal(m, b)
}
func (m *ExamineInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamineInfo.Marshal(b, m, deterministic)
}
func (m *ExamineInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamineInfo.Merge(m, src)
}
func (m *ExamineInfo) XXX_Size() int {
	return xxx_messageInfo_ExamineInfo.Size(m)
}
func (m *ExamineInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamineInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExamineInfo proto.InternalMessageInfo

func (m *ExamineInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ExamineInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExamineInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ExamineInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ExamineInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ExamineInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ExamineInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ExamineInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ExamineInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ExamineInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ExamineInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReqExamineAdd struct {
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Type                 uint32   `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqExamineAdd) Reset()         { *m = ReqExamineAdd{} }
func (m *ReqExamineAdd) String() string { return proto.CompactTextString(m) }
func (*ReqExamineAdd) ProtoMessage()    {}
func (*ReqExamineAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be56c51a90ed54b, []int{1}
}

func (m *ReqExamineAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqExamineAdd.Unmarshal(m, b)
}
func (m *ReqExamineAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqExamineAdd.Marshal(b, m, deterministic)
}
func (m *ReqExamineAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqExamineAdd.Merge(m, src)
}
func (m *ReqExamineAdd) XXX_Size() int {
	return xxx_messageInfo_ReqExamineAdd.Size(m)
}
func (m *ReqExamineAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqExamineAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqExamineAdd proto.InternalMessageInfo

func (m *ReqExamineAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqExamineAdd) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqExamineAdd) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReqExamineAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqExamineAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyExamineInfo struct {
	Info                 *ExamineInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyExamineInfo) Reset()         { *m = ReplyExamineInfo{} }
func (m *ReplyExamineInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyExamineInfo) ProtoMessage()    {}
func (*ReplyExamineInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be56c51a90ed54b, []int{2}
}

func (m *ReplyExamineInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyExamineInfo.Unmarshal(m, b)
}
func (m *ReplyExamineInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyExamineInfo.Marshal(b, m, deterministic)
}
func (m *ReplyExamineInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyExamineInfo.Merge(m, src)
}
func (m *ReplyExamineInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyExamineInfo.Size(m)
}
func (m *ReplyExamineInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyExamineInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyExamineInfo proto.InternalMessageInfo

func (m *ReplyExamineInfo) GetInfo() *ExamineInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyExamineInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyExamineList struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	List                 []*ExamineInfo `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyExamineList) Reset()         { *m = ReplyExamineList{} }
func (m *ReplyExamineList) String() string { return proto.CompactTextString(m) }
func (*ReplyExamineList) ProtoMessage()    {}
func (*ReplyExamineList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be56c51a90ed54b, []int{3}
}

func (m *ReplyExamineList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyExamineList.Unmarshal(m, b)
}
func (m *ReplyExamineList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyExamineList.Marshal(b, m, deterministic)
}
func (m *ReplyExamineList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyExamineList.Merge(m, src)
}
func (m *ReplyExamineList) XXX_Size() int {
	return xxx_messageInfo_ReplyExamineList.Size(m)
}
func (m *ReplyExamineList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyExamineList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyExamineList proto.InternalMessageInfo

func (m *ReplyExamineList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyExamineList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyExamineList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyExamineList) GetList() []*ExamineInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ExamineInfo)(nil), "omo.msp.vocabulary.ExamineInfo")
	proto.RegisterType((*ReqExamineAdd)(nil), "omo.msp.vocabulary.ReqExamineAdd")
	proto.RegisterType((*ReplyExamineInfo)(nil), "omo.msp.vocabulary.ReplyExamineInfo")
	proto.RegisterType((*ReplyExamineList)(nil), "omo.msp.vocabulary.ReplyExamineList")
}

func init() {
	proto.RegisterFile("proto/vocabulary/examine.proto", fileDescriptor_9be56c51a90ed54b)
}

var fileDescriptor_9be56c51a90ed54b = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xe3, 0x3f, 0x71, 0x9b, 0xc9, 0x9b, 0xbc, 0xd1, 0x0a, 0xa1, 0x55, 0xa4, 0x82, 0x6b,
	0x38, 0xe4, 0xe4, 0x4a, 0xed, 0x81, 0x73, 0x2b, 0x41, 0x84, 0x04, 0x42, 0x6c, 0x84, 0x90, 0x10,
	0x1c, 0x5c, 0x7b, 0x5a, 0xad, 0x6a, 0x7b, 0x5d, 0x7b, 0x1d, 0xe1, 0x13, 0x5c, 0xf8, 0x18, 0x7c,
	0x03, 0x3e, 0x24, 0xda, 0xb5, 0x63, 0x4c, 0x5b, 0x47, 0xe6, 0xe6, 0x67, 0x66, 0xf6, 0x37, 0xcf,
	0xcc, 0x6e, 0x02, 0x4f, 0xb2, 0x5c, 0x48, 0x71, 0xb2, 0x15, 0x61, 0x70, 0x59, 0xc6, 0x41, 0x5e,
	0x9d, 0xe0, 0xd7, 0x20, 0xe1, 0x29, 0xfa, 0x3a, 0x41, 0x88, 0x48, 0x84, 0x9f, 0x14, 0x99, 0xff,
	0xa7, 0x62, 0x79, 0x74, 0xef, 0x4c, 0x28, 0x92, 0x44, 0xa4, 0xf5, 0x11, 0xef, 0x87, 0x09, 0xd3,
	0x97, 0x35, 0xe4, 0x75, 0x7a, 0x25, 0xc8, 0x02, 0xac, 0x92, 0x47, 0xd4, 0x70, 0x8d, 0xd5, 0x84,
	0xa9, 0x4f, 0x32, 0x07, 0x93, 0x47, 0xd4, 0x74, 0x8d, 0x95, 0xcd, 0x4c, 0x1e, 0x11, 0x0a, 0x07,
	0x61, 0x8e, 0x81, 0xc4, 0x88, 0x5a, 0xae, 0xb1, 0xb2, 0xd8, 0x4e, 0xaa, 0x4c, 0x99, 0x45, 0x3a,
	0x63, 0xd7, 0x99, 0x46, 0xb6, 0x67, 0x44, 0x4e, 0xc7, 0x9a, 0xbc, 0x93, 0x64, 0x09, 0x87, 0x22,
	0xc3, 0x5c, 0xa7, 0x1c, 0x9d, 0x6a, 0x35, 0x79, 0x0c, 0x8e, 0x0c, 0xf2, 0x6b, 0x94, 0xf4, 0x40,
	0x67, 0x1a, 0xa5, 0xe2, 0x85, 0x0c, 0x64, 0x59, 0xd0, 0x43, 0xd7, 0x58, 0xcd, 0x58, 0xa3, 0x08,
	0x01, 0x5b, 0x56, 0x19, 0xd2, 0x89, 0x8e, 0xea, 0x6f, 0x35, 0xcf, 0x0d, 0x56, 0x14, 0xea, 0x79,
	0x6e, 0xb0, 0x22, 0x8f, 0x60, 0xbc, 0x0d, 0xe2, 0x12, 0xe9, 0x54, 0xc7, 0x6a, 0xe1, 0x7d, 0x83,
	0x19, 0xc3, 0xdb, 0x66, 0x13, 0xe7, 0x51, 0xd4, 0x69, 0x6e, 0xfc, 0xd5, 0xbc, 0x01, 0x9a, 0x0f,
	0x00, 0xad, 0x0e, 0xb0, 0x35, 0x33, 0xee, 0x98, 0xd9, 0x33, 0xac, 0xf7, 0xdd, 0x80, 0x05, 0xc3,
	0x2c, 0xae, 0xba, 0xb7, 0x71, 0x06, 0x36, 0x4f, 0xaf, 0x84, 0xb6, 0x30, 0x3d, 0x7d, 0xea, 0xdf,
	0xbf, 0x5f, 0xbf, 0x53, 0xce, 0x74, 0x31, 0x79, 0xd1, 0xae, 0xc7, 0xec, 0x3f, 0xa6, 0x5b, 0x6d,
	0x74, 0xd9, 0x6e, 0x7f, 0xde, 0xaf, 0x3b, 0x16, 0xde, 0xf0, 0x42, 0x76, 0x68, 0xc6, 0x3f, 0xd1,
	0xd4, 0x5a, 0xa4, 0x90, 0x41, 0xac, 0x5d, 0xcc, 0x58, 0x2d, 0x54, 0x34, 0x0b, 0xae, 0xb1, 0xd0,
	0xcb, 0x9a, 0xb1, 0x5a, 0xa8, 0x39, 0x63, 0x5e, 0x48, 0x3a, 0x75, 0xad, 0x41, 0x73, 0xaa, 0xe2,
	0xd3, 0x9f, 0x36, 0xcc, 0x9b, 0xe8, 0x06, 0xf3, 0x2d, 0x0f, 0x91, 0x6c, 0xc0, 0x39, 0x8f, 0xa2,
	0x77, 0x29, 0x92, 0xe3, 0x87, 0x6d, 0x76, 0x6e, 0x78, 0xf9, 0xbc, 0x77, 0x92, 0x4e, 0x2f, 0x6f,
	0x44, 0xde, 0x83, 0xb3, 0x46, 0xa9, 0xa0, 0x3d, 0xb3, 0xdf, 0x96, 0x58, 0x48, 0x55, 0x3c, 0x18,
	0xf9, 0x16, 0x26, 0x0c, 0x13, 0xb1, 0xc5, 0x41, 0xd4, 0xa3, 0x5e, 0x6a, 0x83, 0xfb, 0x02, 0xf3,
	0x0f, 0xfa, 0x97, 0x76, 0x51, 0xbd, 0xe2, 0xb1, 0xc4, 0x9c, 0x3c, 0xeb, 0x61, 0xd6, 0x65, 0x75,
	0xd1, 0x60, 0xb7, 0x9f, 0xe1, 0xff, 0x35, 0x4a, 0xf5, 0x1a, 0x5a, 0xfe, 0xf1, 0x1e, 0xcf, 0x43,
	0xe9, 0x0a, 0xe8, 0x8d, 0xc8, 0x47, 0xf8, 0x6f, 0x8d, 0x52, 0x3d, 0x1e, 0x5e, 0x48, 0x1e, 0x0e,
	0x41, 0x7b, 0x7b, 0xdf, 0xa0, 0xc6, 0x78, 0xa3, 0x0b, 0xf2, 0x69, 0x71, 0xf7, 0xbf, 0xef, 0xd2,
	0xd1, 0x91, 0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x26, 0x04, 0x95, 0x4a, 0x05, 0x00,
	0x00,
}
