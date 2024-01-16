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
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid                  string         `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Cover                string         `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string         `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Updated              int64          `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Created              int64          `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	Creator              string         `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Concept              string         `protobuf:"bytes,9,opt,name=concept,proto3" json:"concept,omitempty"`
	Type                 uint32         `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`
	Count                uint32         `protobuf:"varint,11,opt,name=count,proto3" json:"count,omitempty"`
	Workflow             string         `protobuf:"bytes,12,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Owner                string         `protobuf:"bytes,13,opt,name=owner,proto3" json:"owner,omitempty"`
	Reviewers            []string       `protobuf:"bytes,31,rep,name=reviewers,proto3" json:"reviewers,omitempty"`
	Keywords             []string       `protobuf:"bytes,32,rep,name=keywords,proto3" json:"keywords,omitempty"`
	Users                []string       `protobuf:"bytes,33,rep,name=users,proto3" json:"users,omitempty"`
	Contents             []*ContentInfo `protobuf:"bytes,41,rep,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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

func (m *BoxInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BoxInfo) GetReviewers() []string {
	if m != nil {
		return m.Reviewers
	}
	return nil
}

func (m *BoxInfo) GetKeywords() []string {
	if m != nil {
		return m.Keywords
	}
	return nil
}

func (m *BoxInfo) GetUsers() []string {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BoxInfo) GetContents() []*ContentInfo {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ContentInfo struct {
	Keywords             string   `protobuf:"bytes,1,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               uint32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Count                uint32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentInfo) Reset()         { *m = ContentInfo{} }
func (m *ContentInfo) String() string { return proto.CompactTextString(m) }
func (*ContentInfo) ProtoMessage()    {}
func (*ContentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{1}
}

func (m *ContentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentInfo.Unmarshal(m, b)
}
func (m *ContentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentInfo.Marshal(b, m, deterministic)
}
func (m *ContentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentInfo.Merge(m, src)
}
func (m *ContentInfo) XXX_Size() int {
	return xxx_messageInfo_ContentInfo.Size(m)
}
func (m *ContentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContentInfo proto.InternalMessageInfo

func (m *ContentInfo) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *ContentInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContentInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ContentInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReqBoxAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Concept              string   `protobuf:"bytes,4,opt,name=concept,proto3" json:"concept,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Workflow             string   `protobuf:"bytes,7,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Owner                string   `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Keywords             []string `protobuf:"bytes,21,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBoxAdd) Reset()         { *m = ReqBoxAdd{} }
func (m *ReqBoxAdd) String() string { return proto.CompactTextString(m) }
func (*ReqBoxAdd) ProtoMessage()    {}
func (*ReqBoxAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{2}
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

func (m *ReqBoxAdd) GetOwner() string {
	if m != nil {
		return m.Owner
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
	Keywords             []string `protobuf:"bytes,21,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBoxUpdate) Reset()         { *m = ReqBoxUpdate{} }
func (m *ReqBoxUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqBoxUpdate) ProtoMessage()    {}
func (*ReqBoxUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{3}
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
	Reviewer             bool     `protobuf:"varint,3,opt,name=reviewer,proto3" json:"reviewer,omitempty"`
	Keywords             []string `protobuf:"bytes,21,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBoxKeywords) Reset()         { *m = ReqBoxKeywords{} }
func (m *ReqBoxKeywords) String() string { return proto.CompactTextString(m) }
func (*ReqBoxKeywords) ProtoMessage()    {}
func (*ReqBoxKeywords) Descriptor() ([]byte, []int) {
	return fileDescriptor_37d071ad8f8ae54f, []int{4}
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

func (m *ReqBoxKeywords) GetReviewer() bool {
	if m != nil {
		return m.Reviewer
	}
	return false
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
	return fileDescriptor_37d071ad8f8ae54f, []int{5}
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
	return fileDescriptor_37d071ad8f8ae54f, []int{6}
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
	proto.RegisterType((*ContentInfo)(nil), "omo.msp.vocabulary.ContentInfo")
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
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xc3, 0x34,
	0x14, 0x5e, 0x9a, 0x36, 0x6d, 0x4f, 0xd7, 0x69, 0xb2, 0x60, 0xb2, 0x0a, 0xd3, 0xb2, 0x70, 0x53,
	0x6e, 0x3a, 0x69, 0x5c, 0x70, 0xc1, 0x55, 0x8b, 0x44, 0x85, 0x06, 0x42, 0xca, 0x28, 0x43, 0xbb,
	0x4b, 0x13, 0x17, 0x45, 0x4d, 0xe3, 0xcc, 0x71, 0xfa, 0xf3, 0x0e, 0x3c, 0x0e, 0x4f, 0xc1, 0x15,
	0x2f, 0xc1, 0x7b, 0x20, 0xdb, 0x71, 0x9a, 0xfe, 0xa4, 0x14, 0xd1, 0xbb, 0x1c, 0x9f, 0xe3, 0xef,
	0x9c, 0xcf, 0xdf, 0x67, 0xb7, 0xd0, 0x4b, 0x18, 0xe5, 0xf4, 0x69, 0x49, 0x7d, 0x6f, 0x9a, 0x45,
	0x1e, 0xdb, 0x3c, 0x4d, 0xe9, 0x7a, 0x20, 0x17, 0x11, 0xa2, 0x0b, 0x3a, 0x58, 0xa4, 0xc9, 0x60,
	0x9b, 0xed, 0xdd, 0x1f, 0xd4, 0xfb, 0x74, 0xb1, 0xa0, 0xb1, 0xda, 0xe2, 0xfc, 0x69, 0x42, 0x73,
	0x44, 0xd7, 0xdf, 0xc7, 0x33, 0x8a, 0x10, 0xd4, 0x63, 0x6f, 0x41, 0xb0, 0x61, 0x1b, 0xfd, 0xb6,
	0x2b, 0xbf, 0xd1, 0x2d, 0x98, 0x59, 0x18, 0xe0, 0x9a, 0x5c, 0x12, 0x9f, 0xe8, 0x13, 0x68, 0xf8,
	0x74, 0x49, 0x18, 0x36, 0xe5, 0x9a, 0x0a, 0xd0, 0x1d, 0x58, 0x8c, 0x2c, 0x3c, 0x36, 0xc7, 0x75,
	0xb9, 0x9c, 0x47, 0x08, 0x43, 0x33, 0x4b, 0x02, 0x8f, 0x93, 0x00, 0x37, 0x6c, 0xa3, 0x6f, 0xba,
	0x3a, 0x14, 0x19, 0x9f, 0x11, 0x99, 0xb1, 0x54, 0x26, 0x0f, 0x8b, 0x0c, 0x65, 0xb8, 0x29, 0xc1,
	0x74, 0x88, 0x7a, 0xd0, 0xa2, 0x09, 0x61, 0x32, 0xd5, 0x92, 0xa9, 0x22, 0x96, 0xbb, 0x68, 0xec,
	0x93, 0x84, 0xe3, 0x76, 0xbe, 0x4b, 0x85, 0x82, 0x17, 0xdf, 0x24, 0x04, 0x83, 0x6d, 0xf4, 0xbb,
	0xae, 0xfc, 0x56, 0x2c, 0xb2, 0x98, 0xe3, 0x8e, 0x5c, 0x54, 0x81, 0xc0, 0x5f, 0x51, 0x36, 0x9f,
	0x45, 0x74, 0x85, 0xaf, 0x15, 0xbe, 0x8e, 0xc5, 0x0e, 0xba, 0x8a, 0x09, 0xc3, 0x5d, 0xc5, 0x5b,
	0x06, 0xe8, 0x73, 0x68, 0x33, 0xb2, 0x0c, 0xc9, 0x8a, 0xb0, 0x14, 0x3f, 0xd8, 0x66, 0xbf, 0xed,
	0x6e, 0x17, 0x04, 0xde, 0x9c, 0x6c, 0x56, 0x94, 0x05, 0x29, 0xb6, 0x65, 0xb2, 0x88, 0x05, 0x5e,
	0x96, 0x8a, 0x5d, 0x8f, 0x32, 0xa1, 0x02, 0xf4, 0x0d, 0xb4, 0x7c, 0x1a, 0x73, 0x12, 0xf3, 0x14,
	0x7f, 0x69, 0x9b, 0xfd, 0xce, 0xf3, 0xc3, 0xe0, 0x50, 0xd5, 0xc1, 0xb7, 0xaa, 0x46, 0xc8, 0xe6,
	0x16, 0x1b, 0x9c, 0x39, 0x74, 0x4a, 0x89, 0x9d, 0xee, 0x4a, 0xd3, 0x6d, 0x77, 0xad, 0x75, 0xad,
	0xa4, 0xf5, 0x1d, 0x58, 0x29, 0xf7, 0x78, 0x96, 0x4a, 0x69, 0xbb, 0x6e, 0x1e, 0x6d, 0xcf, 0xaa,
	0x5e, 0x3a, 0x2b, 0xe7, 0x6f, 0x03, 0xda, 0x2e, 0xf9, 0x18, 0xd1, 0xf5, 0x30, 0x08, 0x8e, 0x7a,
	0xa7, 0x70, 0x4a, 0xed, 0xb8, 0x53, 0xcc, 0x7d, 0xa7, 0x68, 0xfd, 0xea, 0xbb, 0xfa, 0x95, 0x55,
	0x6f, 0xec, 0xa9, 0xae, 0xb5, 0xb5, 0x4a, 0xda, 0x96, 0x55, 0x6c, 0x56, 0xa9, 0xd8, 0x2a, 0xab,
	0x58, 0x3e, 0xa9, 0x4f, 0x77, 0x75, 0x72, 0xfe, 0x32, 0xe0, 0x5a, 0xf1, 0x9c, 0x48, 0xe7, 0xea,
	0x2b, 0x61, 0x6c, 0xaf, 0x84, 0x26, 0x6f, 0x1e, 0x23, 0x5f, 0x3f, 0x4e, 0xbe, 0xb1, 0x43, 0xbe,
	0x4c, 0xd1, 0xaa, 0x36, 0x76, 0xf3, 0xe0, 0x60, 0x0a, 0xa2, 0xad, 0x3d, 0xa2, 0xa7, 0x28, 0x2d,
	0xe1, 0x46, 0x31, 0x7a, 0xd1, 0x76, 0x38, 0xe4, 0x54, 0x9e, 0xa8, 0xb6, 0x37, 0x51, 0x0f, 0x5a,
	0xda, 0xe3, 0x92, 0x73, 0xcb, 0x2d, 0xe2, 0x93, 0x7d, 0xd7, 0xe2, 0x24, 0x93, 0x68, 0xa3, 0x1f,
	0x9c, 0x27, 0xa8, 0x87, 0xf1, 0x8c, 0xca, 0xb6, 0x9d, 0xe7, 0xcf, 0x8e, 0x19, 0x3d, 0x2f, 0x75,
	0x65, 0x21, 0xfa, 0xba, 0x70, 0x68, 0x4d, 0x6e, 0x39, 0x7a, 0x37, 0x64, 0x8b, 0x57, 0x59, 0xa6,
	0x2d, 0xec, 0xfc, 0x6e, 0x6c, 0x5b, 0xff, 0x10, 0xa6, 0xf2, 0x4d, 0x98, 0x45, 0xde, 0x6f, 0xda,
	0xaf, 0xe2, 0x5b, 0x8c, 0x13, 0x85, 0x29, 0xc7, 0x35, 0x79, 0xef, 0x4e, 0x8f, 0x23, 0x0a, 0x4b,
	0xe3, 0x98, 0xff, 0x69, 0x9c, 0xe7, 0x3f, 0xda, 0x00, 0x23, 0xba, 0x7e, 0x25, 0x6c, 0x19, 0xfa,
	0x04, 0xbd, 0x80, 0x35, 0x0c, 0x82, 0x9f, 0x62, 0x82, 0xee, 0x8f, 0x23, 0xe4, 0xb7, 0xac, 0x67,
	0x57, 0x36, 0xc8, 0x07, 0x73, 0xae, 0xd0, 0x8f, 0x60, 0x8d, 0x09, 0x17, 0x60, 0x15, 0xe3, 0x7c,
	0x64, 0x24, 0x95, 0x0f, 0xc4, 0x99, 0x70, 0x6d, 0x97, 0x2c, 0xe8, 0x92, 0x9c, 0x85, 0x78, 0x5f,
	0x89, 0xb8, 0x33, 0xdd, 0x30, 0x8a, 0xfe, 0xe7, 0x74, 0x42, 0x44, 0xe7, 0x0a, 0xfd, 0x0c, 0xdd,
	0x31, 0xe1, 0x22, 0x18, 0x6d, 0x26, 0x29, 0x61, 0x97, 0x42, 0xed, 0x8c, 0x09, 0x1f, 0x6d, 0xbe,
	0x0b, 0x23, 0x4e, 0x18, 0x7a, 0x3c, 0x81, 0xa9, 0x4a, 0xce, 0x42, 0x7d, 0x83, 0xeb, 0x31, 0xe1,
	0xc2, 0x09, 0x61, 0xca, 0x43, 0xff, 0x1c, 0x58, 0xe7, 0xa4, 0xa1, 0x24, 0x8c, 0x73, 0x85, 0x5c,
	0x00, 0xf5, 0x34, 0x8d, 0xbc, 0x94, 0x20, 0xbb, 0xda, 0x42, 0xaa, 0xea, 0x2c, 0xd9, 0x7f, 0x85,
	0x9b, 0x61, 0x92, 0x90, 0x38, 0x28, 0x9e, 0x08, 0xa7, 0x1a, 0x57, 0xd7, 0x9c, 0x85, 0xfc, 0x0e,
	0xb7, 0xaf, 0xd9, 0x94, 0x33, 0xcf, 0xe7, 0x17, 0xc7, 0x9e, 0x40, 0x47, 0x4d, 0x3d, 0x91, 0x3f,
	0xa6, 0x97, 0x82, 0x7d, 0x83, 0xae, 0x1e, 0xf9, 0xb2, 0xc0, 0x13, 0xe8, 0x28, 0x4d, 0x2e, 0x0b,
	0xfb, 0x0b, 0xdc, 0xe4, 0x86, 0xd0, 0x16, 0xfe, 0xa2, 0x02, 0x59, 0x95, 0xe5, 0x6e, 0xfb, 0xb7,
	0xcb, 0x3b, 0x42, 0xef, 0xb7, 0xfb, 0xff, 0x26, 0xa7, 0x96, 0x5c, 0xf9, 0xea, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0xd9, 0x1a, 0x28, 0x98, 0x0a, 0x00, 0x00,
}
