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
	Count                uint32   `protobuf:"varint,11,opt,name=count,proto3" json:"count,omitempty"`
	Workflow             string   `protobuf:"bytes,12,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Keywords             []string `protobuf:"bytes,13,rep,name=keywords,proto3" json:"keywords,omitempty"`
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

func (m *BoxInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *BoxInfo) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
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
	Workflow             string   `protobuf:"bytes,7,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Keywords             []string `protobuf:"bytes,8,rep,name=keywords,proto3" json:"keywords,omitempty"`
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

func (m *ReqBoxAdd) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
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
	Workflow             string   `protobuf:"bytes,8,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Keywords             []string `protobuf:"bytes,9,rep,name=keywords,proto3" json:"keywords,omitempty"`
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

func (m *ReqBoxUpdate) GetWorkflow() string {
	if m != nil {
		return m.Workflow
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
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6e, 0xdb, 0x3c,
	0x14, 0x8c, 0x2c, 0x5b, 0xb6, 0x9f, 0xe3, 0x20, 0x20, 0x3e, 0x7c, 0x20, 0x5c, 0x04, 0x15, 0xb4,
	0xca, 0xca, 0x01, 0xd2, 0x45, 0xd7, 0xd6, 0x26, 0x28, 0xd2, 0xa2, 0x80, 0xd2, 0xa0, 0x45, 0x76,
	0xb2, 0x44, 0x17, 0x86, 0x25, 0x51, 0xa1, 0x28, 0xff, 0xdc, 0xa1, 0xb7, 0xea, 0x01, 0xda, 0x55,
	0xcf, 0x53, 0x90, 0x14, 0x6d, 0x59, 0x96, 0x52, 0x17, 0xf5, 0x4e, 0xf3, 0xde, 0xe3, 0x70, 0x66,
	0x48, 0x42, 0x30, 0x4a, 0x19, 0xe5, 0xf4, 0x66, 0x49, 0x03, 0x7f, 0x9a, 0x47, 0x3e, 0xdb, 0xdc,
	0x4c, 0xe9, 0x7a, 0x2c, 0x8b, 0x08, 0xd1, 0x98, 0x8e, 0xe3, 0x2c, 0x1d, 0xef, 0xba, 0xa3, 0xab,
	0x83, 0xf9, 0x80, 0xc6, 0x31, 0x4d, 0xd4, 0x12, 0xe7, 0x7b, 0x0b, 0xba, 0x2e, 0x5d, 0xbf, 0x4b,
	0x66, 0x14, 0x21, 0x68, 0x27, 0x7e, 0x4c, 0xb0, 0x61, 0x1b, 0xd7, 0x7d, 0x4f, 0x7e, 0xa3, 0x4b,
	0x30, 0xf3, 0x79, 0x88, 0x5b, 0xb2, 0x24, 0x3e, 0xd1, 0x7f, 0xd0, 0x09, 0xe8, 0x92, 0x30, 0x6c,
	0xca, 0x9a, 0x02, 0xe8, 0x7f, 0xb0, 0x18, 0x89, 0x7d, 0xb6, 0xc0, 0x6d, 0x59, 0x2e, 0x10, 0xc2,
	0xd0, 0xcd, 0xd3, 0xd0, 0xe7, 0x24, 0xc4, 0x1d, 0xdb, 0xb8, 0x36, 0x3d, 0x0d, 0x45, 0x27, 0x60,
	0x44, 0x76, 0x2c, 0xd5, 0x29, 0xe0, 0xb6, 0x43, 0x19, 0xee, 0x4a, 0x32, 0x0d, 0xd1, 0x08, 0x7a,
	0x34, 0x25, 0x4c, 0xb6, 0x7a, 0xb2, 0xb5, 0xc5, 0x72, 0x15, 0x4d, 0x02, 0x92, 0x72, 0xdc, 0x2f,
	0x56, 0x29, 0x28, 0x7c, 0xf1, 0x4d, 0x4a, 0x30, 0xd8, 0xc6, 0xf5, 0xd0, 0x93, 0xdf, 0xca, 0x45,
	0x9e, 0x70, 0x3c, 0x90, 0x45, 0x05, 0x04, 0xff, 0x8a, 0xb2, 0xc5, 0x2c, 0xa2, 0x2b, 0x7c, 0xae,
	0xf8, 0x35, 0x16, 0xbd, 0x05, 0xd9, 0xac, 0x28, 0x0b, 0x33, 0x3c, 0xb4, 0x4d, 0xd1, 0xd3, 0xd8,
	0xf9, 0x61, 0x40, 0xdf, 0x23, 0xcf, 0x2e, 0x5d, 0x4f, 0xc2, 0xb0, 0x36, 0xc7, 0x6d, 0x6a, 0xad,
	0xfa, 0xd4, 0xcc, 0x6a, 0x6a, 0xda, 0x4b, 0x7b, 0xdf, 0x4b, 0x39, 0x81, 0x4e, 0x25, 0x01, 0xed,
	0xd3, 0x2a, 0xf9, 0x2c, 0x3b, 0xea, 0xbe, 0xe0, 0xa8, 0x57, 0x71, 0xf4, 0xd3, 0x80, 0x73, 0xe5,
	0xe8, 0x51, 0x9e, 0x97, 0xbe, 0x08, 0xc6, 0xee, 0x22, 0x68, 0x9b, 0x66, 0x9d, 0xcd, 0x76, 0xbd,
	0xcd, 0xce, 0x9e, 0xcd, 0xb2, 0x19, 0xab, 0xf9, 0x38, 0xbb, 0x07, 0x11, 0x6c, 0x2d, 0xf5, 0x5e,
	0xb0, 0xd4, 0xaf, 0x58, 0x7a, 0x82, 0x0b, 0xe5, 0xe8, 0xbe, 0xa8, 0xd4, 0x78, 0x2a, 0x2b, 0x6a,
	0x55, 0x14, 0x95, 0xb9, 0xcd, 0x0a, 0xf7, 0x5a, 0xa4, 0x95, 0x46, 0x1b, 0xfd, 0x94, 0x6e, 0xa0,
	0x3d, 0x4f, 0x66, 0x54, 0x52, 0x0f, 0x6e, 0x5f, 0x8d, 0x0f, 0x1f, 0xe6, 0xb8, 0x18, 0xf5, 0xe4,
	0x20, 0x7a, 0x0b, 0x56, 0xc6, 0x7d, 0x9e, 0x67, 0x72, 0xdb, 0xc1, 0xed, 0xeb, 0xba, 0x25, 0x72,
	0x8b, 0x07, 0x39, 0xe6, 0x15, 0xe3, 0xce, 0x37, 0x63, 0xb7, 0xf5, 0xfb, 0x79, 0x26, 0x6f, 0xfb,
	0x2c, 0xf2, 0xbf, 0xea, 0xdb, 0x27, 0xbe, 0x85, 0x9c, 0x68, 0x9e, 0x71, 0xdc, 0xb2, 0xcd, 0x3f,
	0xca, 0x11, 0x83, 0x25, 0x39, 0xe6, 0x5f, 0xc9, 0xb9, 0xfd, 0x65, 0x01, 0xb8, 0x74, 0xfd, 0x40,
	0xd8, 0x72, 0x1e, 0x10, 0x74, 0x0f, 0xd6, 0x24, 0x0c, 0x3f, 0x26, 0x04, 0x5d, 0xd5, 0x33, 0x14,
	0x6f, 0x66, 0x64, 0x37, 0x6e, 0x50, 0x08, 0x73, 0xce, 0xd0, 0x07, 0xb0, 0xee, 0x08, 0x17, 0x64,
	0x0d, 0x72, 0x9e, 0x73, 0x92, 0x71, 0x31, 0x7c, 0x24, 0x5d, 0xdf, 0x23, 0x31, 0x5d, 0x92, 0xa3,
	0x18, 0xaf, 0x1a, 0x19, 0xf7, 0xd4, 0x4d, 0xa2, 0xe8, 0x1f, 0xd5, 0x89, 0x43, 0x74, 0xce, 0xd0,
	0x27, 0x18, 0xde, 0x11, 0x2e, 0x80, 0xbb, 0x79, 0xcc, 0x08, 0x3b, 0x0d, 0xab, 0x07, 0xa0, 0xde,
	0xb3, 0xeb, 0x67, 0x04, 0xd9, 0xcd, 0x67, 0xa2, 0xa6, 0x8e, 0xca, 0xf1, 0x0b, 0x5c, 0x4c, 0xd2,
	0x94, 0x24, 0xe1, 0xf6, 0x5d, 0x39, 0xcd, 0xbc, 0x7a, 0xe6, 0x28, 0xe6, 0x27, 0xb8, 0x7c, 0xc8,
	0xa7, 0x9c, 0xf9, 0x01, 0x3f, 0x39, 0xf7, 0x23, 0x0c, 0x94, 0x6a, 0x11, 0xee, 0xe9, 0x68, 0x3f,
	0xc3, 0x50, 0x4b, 0x3e, 0x29, 0xb1, 0x8b, 0x9e, 0x2e, 0xab, 0x7f, 0xf2, 0xa9, 0x25, 0x2b, 0x6f,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x01, 0xa7, 0x96, 0x14, 0x08, 0x00, 0x00,
}
