// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/box.proto

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

type BoxInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Created              int64    `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	Creator              string   `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Concept              string   `protobuf:"bytes,9,opt,name=concept,proto3" json:"concept,omitempty"`
	Type                 uint32   `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`
	Keywords             []string `protobuf:"bytes,11,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoxInfo) Reset()         { *m = BoxInfo{} }
func (m *BoxInfo) String() string { return proto.CompactTextString(m) }
func (*BoxInfo) ProtoMessage()    {}
func (*BoxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{0}
}

func (m *BoxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoxInfo.Unmarshal(m, b)
}
func (m *BoxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoxInfo.Marshal(b, m, deterministic)
}
func (m *BoxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoxInfo.Merge(m, src)
}
func (m *BoxInfo) XXX_Size() int {
	return xxx_messageInfo_BoxInfo.Size(m)
}
func (m *BoxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BoxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BoxInfo proto.InternalMessageInfo

func (m *BoxInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BoxInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *BoxInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *BoxInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *BoxInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *BoxInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *BoxInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BoxInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *BoxInfo) GetConcept() string {
	if m != nil {
		return m.Concept
	}
	return ""
}

func (m *BoxInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *BoxInfo) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type ReqBoxAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Concept              string   `protobuf:"bytes,4,opt,name=concept,proto3" json:"concept,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Keywords             []string `protobuf:"bytes,7,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBoxAdd) Reset()         { *m = ReqBoxAdd{} }
func (m *ReqBoxAdd) String() string { return proto.CompactTextString(m) }
func (*ReqBoxAdd) ProtoMessage()    {}
func (*ReqBoxAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{1}
}

func (m *ReqBoxAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBoxAdd.Unmarshal(m, b)
}
func (m *ReqBoxAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBoxAdd.Marshal(b, m, deterministic)
}
func (m *ReqBoxAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBoxAdd.Merge(m, src)
}
func (m *ReqBoxAdd) XXX_Size() int {
	return xxx_messageInfo_ReqBoxAdd.Size(m)
}
func (m *ReqBoxAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBoxAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBoxAdd proto.InternalMessageInfo

func (m *ReqBoxAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqBoxAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqBoxAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqBoxAdd) GetConcept() string {
	if m != nil {
		return m.Concept
	}
	return ""
}

func (m *ReqBoxAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqBoxAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqBoxAdd) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type ReqBoxUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Concept              string   `protobuf:"bytes,7,opt,name=concept,proto3" json:"concept,omitempty"`
	Keywords             []string `protobuf:"bytes,8,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBoxUpdate) Reset()         { *m = ReqBoxUpdate{} }
func (m *ReqBoxUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqBoxUpdate) ProtoMessage()    {}
func (*ReqBoxUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{2}
}

func (m *ReqBoxUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBoxUpdate.Unmarshal(m, b)
}
func (m *ReqBoxUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBoxUpdate.Marshal(b, m, deterministic)
}
func (m *ReqBoxUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBoxUpdate.Merge(m, src)
}
func (m *ReqBoxUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqBoxUpdate.Size(m)
}
func (m *ReqBoxUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBoxUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBoxUpdate proto.InternalMessageInfo

func (m *ReqBoxUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqBoxUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqBoxUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqBoxUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqBoxUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqBoxUpdate) GetConcept() string {
	if m != nil {
		return m.Concept
	}
	return ""
}

func (m *ReqBoxUpdate) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type ReqBoxKeywords struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Keywords             []string `protobuf:"bytes,3,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBoxKeywords) Reset()         { *m = ReqBoxKeywords{} }
func (m *ReqBoxKeywords) String() string { return proto.CompactTextString(m) }
func (*ReqBoxKeywords) ProtoMessage()    {}
func (*ReqBoxKeywords) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{3}
}

func (m *ReqBoxKeywords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBoxKeywords.Unmarshal(m, b)
}
func (m *ReqBoxKeywords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBoxKeywords.Marshal(b, m, deterministic)
}
func (m *ReqBoxKeywords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBoxKeywords.Merge(m, src)
}
func (m *ReqBoxKeywords) XXX_Size() int {
	return xxx_messageInfo_ReqBoxKeywords.Size(m)
}
func (m *ReqBoxKeywords) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBoxKeywords.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBoxKeywords proto.InternalMessageInfo

func (m *ReqBoxKeywords) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqBoxKeywords) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqBoxKeywords) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

type ReplyBoxInfo struct {
	Info                 *BoxInfo     `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyBoxInfo) Reset()         { *m = ReplyBoxInfo{} }
func (m *ReplyBoxInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyBoxInfo) ProtoMessage()    {}
func (*ReplyBoxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{4}
}

func (m *ReplyBoxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyBoxInfo.Unmarshal(m, b)
}
func (m *ReplyBoxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyBoxInfo.Marshal(b, m, deterministic)
}
func (m *ReplyBoxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyBoxInfo.Merge(m, src)
}
func (m *ReplyBoxInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyBoxInfo.Size(m)
}
func (m *ReplyBoxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyBoxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyBoxInfo proto.InternalMessageInfo

func (m *ReplyBoxInfo) GetInfo() *BoxInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyBoxInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyBoxList struct {
	Flag                 string       `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	List                 []*BoxInfo   `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyBoxList) Reset()         { *m = ReplyBoxList{} }
func (m *ReplyBoxList) String() string { return proto.CompactTextString(m) }
func (*ReplyBoxList) ProtoMessage()    {}
func (*ReplyBoxList) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{5}
}

func (m *ReplyBoxList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyBoxList.Unmarshal(m, b)
}
func (m *ReplyBoxList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyBoxList.Marshal(b, m, deterministic)
}
func (m *ReplyBoxList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyBoxList.Merge(m, src)
}
func (m *ReplyBoxList) XXX_Size() int {
	return xxx_messageInfo_ReplyBoxList.Size(m)
}
func (m *ReplyBoxList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyBoxList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyBoxList proto.InternalMessageInfo

func (m *ReplyBoxList) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *ReplyBoxList) GetList() []*BoxInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyBoxList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*BoxInfo)(nil), "omo.msp.vocabulary.BoxInfo")
	proto.RegisterType((*ReqBoxAdd)(nil), "omo.msp.vocabulary.ReqBoxAdd")
	proto.RegisterType((*ReqBoxUpdate)(nil), "omo.msp.vocabulary.ReqBoxUpdate")
	proto.RegisterType((*ReqBoxKeywords)(nil), "omo.msp.vocabulary.ReqBoxKeywords")
	proto.RegisterType((*ReplyBoxInfo)(nil), "omo.msp.vocabulary.ReplyBoxInfo")
	proto.RegisterType((*ReplyBoxList)(nil), "omo.msp.vocabulary.ReplyBoxList")
}

func init() {
	proto.RegisterFile("proto/vocabulary/box.proto", fileDescriptor_37d071ad8f8ae54f)
}

var fileDescriptor_37d071ad8f8ae54f = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4d, 0xaf, 0x93, 0x4c,
	0x14, 0xc7, 0x2f, 0x85, 0x42, 0x7b, 0xfa, 0x3c, 0xe6, 0x66, 0x62, 0xcc, 0xa4, 0xa6, 0x91, 0xb0,
	0xea, 0x8a, 0x26, 0x75, 0xe1, 0xba, 0xdd, 0x18, 0x73, 0x7d, 0x89, 0x34, 0x6e, 0xee, 0x8e, 0xc2,
	0xd4, 0x90, 0x0b, 0x0c, 0x77, 0x18, 0x2a, 0xfd, 0x0e, 0x6e, 0xfc, 0x2c, 0xba, 0xf6, 0xb3, 0x99,
	0x19, 0x98, 0x96, 0x52, 0x88, 0x35, 0xba, 0x9b, 0xf3, 0x32, 0xff, 0xf3, 0x3b, 0x67, 0x0e, 0x01,
	0xa6, 0x19, 0xa3, 0x9c, 0x2e, 0xf6, 0x34, 0xf0, 0xb7, 0x45, 0xec, 0xb3, 0xc3, 0x62, 0x4b, 0x4b,
	0x57, 0x3a, 0x11, 0xa2, 0x09, 0x75, 0x93, 0x3c, 0x73, 0x4f, 0xd1, 0xe9, 0xec, 0x22, 0x3f, 0xa0,
	0x49, 0x42, 0xd3, 0xea, 0x8a, 0xf3, 0x6d, 0x00, 0xd6, 0x9a, 0x96, 0x6f, 0xd2, 0x1d, 0x45, 0x08,
	0x8c, 0xd4, 0x4f, 0x08, 0xd6, 0x6c, 0x6d, 0x3e, 0xf6, 0xe4, 0x19, 0xdd, 0x82, 0x5e, 0x44, 0x21,
	0x1e, 0x48, 0x97, 0x38, 0xa2, 0xa7, 0x30, 0x0c, 0xe8, 0x9e, 0x30, 0xac, 0x4b, 0x5f, 0x65, 0xa0,
	0x67, 0x60, 0x32, 0x92, 0xf8, 0xec, 0x01, 0x1b, 0xd2, 0x5d, 0x5b, 0x08, 0x83, 0x55, 0x64, 0xa1,
	0xcf, 0x49, 0x88, 0x87, 0xb6, 0x36, 0xd7, 0x3d, 0x65, 0x8a, 0x48, 0xc0, 0x88, 0x8c, 0x98, 0x55,
	0xa4, 0x36, 0x8f, 0x11, 0xca, 0xb0, 0x25, 0xc5, 0x94, 0x89, 0xa6, 0x30, 0xa2, 0x19, 0x61, 0x32,
	0x34, 0x92, 0xa1, 0xa3, 0x2d, 0x6f, 0xd1, 0x34, 0x20, 0x19, 0xc7, 0xe3, 0xfa, 0x56, 0x65, 0x8a,
	0xbe, 0xf8, 0x21, 0x23, 0x18, 0x6c, 0x6d, 0xfe, 0xbf, 0x27, 0xcf, 0x42, 0xe9, 0x81, 0x1c, 0xbe,
	0x50, 0x16, 0xe6, 0x78, 0x62, 0xeb, 0x42, 0x49, 0xd9, 0xce, 0x77, 0x0d, 0xc6, 0x1e, 0x79, 0x5c,
	0xd3, 0x72, 0x15, 0x86, 0x9d, 0x53, 0x39, 0xce, 0x60, 0xd0, 0x3d, 0x03, 0xbd, 0x3d, 0x03, 0x45,
	0x66, 0x9c, 0x93, 0x35, 0xfb, 0x19, 0xb6, 0xfa, 0x51, 0xd4, 0x66, 0x0f, 0xb5, 0xd5, 0xa2, 0xfe,
	0xa1, 0xc1, 0x7f, 0x15, 0xf5, 0x27, 0x39, 0x61, 0xf5, 0x74, 0xda, 0xe9, 0xe9, 0x54, 0x2b, 0x7a,
	0x57, 0x2b, 0x46, 0x77, 0x2b, 0xc3, 0xb3, 0x56, 0x9a, 0xc0, 0x66, 0xff, 0x03, 0x58, 0x17, 0x6d,
	0x1e, 0xb1, 0x47, 0x2d, 0xec, 0x7b, 0x78, 0x52, 0x51, 0xdf, 0xd5, 0x9e, 0x0e, 0xee, 0x66, 0xd5,
	0x41, 0xab, 0x6a, 0x53, 0x5b, 0x6f, 0x69, 0x97, 0x62, 0x22, 0x59, 0x7c, 0x50, 0x0b, 0xbe, 0x00,
	0x23, 0x4a, 0x77, 0x54, 0x4a, 0x4f, 0x96, 0xcf, 0xdd, 0xcb, 0xcf, 0xc5, 0xad, 0x53, 0x3d, 0x99,
	0x88, 0x5e, 0x81, 0x99, 0x73, 0x9f, 0x17, 0xb9, 0x2c, 0x3b, 0x59, 0xbe, 0xe8, 0xba, 0x22, 0x4b,
	0x6c, 0x64, 0x9a, 0x57, 0xa7, 0x3b, 0x5f, 0xb5, 0x53, 0xe9, 0xb7, 0x51, 0x2e, 0x77, 0x70, 0x17,
	0xfb, 0x9f, 0xd5, 0x16, 0x89, 0xb3, 0xc0, 0x89, 0xa3, 0x9c, 0xe3, 0x81, 0xad, 0xff, 0x16, 0x47,
	0x24, 0x36, 0x70, 0xf4, 0x3f, 0xc2, 0x59, 0xfe, 0x34, 0x00, 0xd6, 0xb4, 0xdc, 0x10, 0xb6, 0x8f,
	0x02, 0x82, 0xee, 0xc0, 0x5c, 0x85, 0xe1, 0x87, 0x94, 0xa0, 0x59, 0xb7, 0x42, 0xbd, 0xfb, 0x53,
	0xbb, 0xb7, 0x40, 0x0d, 0xe6, 0xdc, 0xa0, 0x77, 0x60, 0xbe, 0x26, 0x5c, 0x88, 0xf5, 0xe0, 0x3c,
	0x16, 0x24, 0xe7, 0x22, 0xf9, 0x4a, 0xb9, 0xb1, 0x47, 0x12, 0xba, 0x27, 0x57, 0x29, 0xce, 0x7a,
	0x15, 0xcf, 0xe8, 0x56, 0x71, 0xfc, 0x97, 0x74, 0xe2, 0x11, 0x9d, 0x1b, 0xf4, 0x1e, 0xcc, 0xfa,
	0xeb, 0xb2, 0xfb, 0x27, 0x57, 0x65, 0x5c, 0xd5, 0xed, 0x47, 0xb0, 0x56, 0x59, 0x46, 0xd2, 0x30,
	0x47, 0x4e, 0xbf, 0xa0, 0xfa, 0x34, 0xae, 0x92, 0xdc, 0xc0, 0x78, 0x53, 0x6c, 0x39, 0xf3, 0x03,
	0xfe, 0xcf, 0x44, 0xd7, 0xe8, 0xfe, 0xb6, 0xfd, 0x1f, 0xd9, 0x9a, 0xd2, 0xf3, 0xf2, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb8, 0x9b, 0xf4, 0x5c, 0x92, 0x06, 0x00, 0x00,
}