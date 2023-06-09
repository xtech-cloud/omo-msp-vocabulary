// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/edge.proto

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

type VEdgeInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Entity               string   `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	Relation             string   `protobuf:"bytes,9,opt,name=relation,proto3" json:"relation,omitempty"`
	Direction            uint32   `protobuf:"varint,10,opt,name=direction,proto3" json:"direction,omitempty"`
	Target               string   `protobuf:"bytes,11,opt,name=target,proto3" json:"target,omitempty"`
	Label                string   `protobuf:"bytes,12,opt,name=label,proto3" json:"label,omitempty"`
	Center               string   `protobuf:"bytes,13,opt,name=center,proto3" json:"center,omitempty"`
	Weight               uint32   `protobuf:"varint,14,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VEdgeInfo) Reset()         { *m = VEdgeInfo{} }
func (m *VEdgeInfo) String() string { return proto.CompactTextString(m) }
func (*VEdgeInfo) ProtoMessage()    {}
func (*VEdgeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{0}
}

func (m *VEdgeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VEdgeInfo.Unmarshal(m, b)
}
func (m *VEdgeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VEdgeInfo.Marshal(b, m, deterministic)
}
func (m *VEdgeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VEdgeInfo.Merge(m, src)
}
func (m *VEdgeInfo) XXX_Size() int {
	return xxx_messageInfo_VEdgeInfo.Size(m)
}
func (m *VEdgeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VEdgeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VEdgeInfo proto.InternalMessageInfo

func (m *VEdgeInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VEdgeInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *VEdgeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VEdgeInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *VEdgeInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *VEdgeInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *VEdgeInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *VEdgeInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *VEdgeInfo) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *VEdgeInfo) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *VEdgeInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *VEdgeInfo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *VEdgeInfo) GetCenter() string {
	if m != nil {
		return m.Center
	}
	return ""
}

func (m *VEdgeInfo) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type ReqVEdgeAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Entity               string   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Center               string   `protobuf:"bytes,3,opt,name=center,proto3" json:"center,omitempty"`
	Relation             string   `protobuf:"bytes,4,opt,name=relation,proto3" json:"relation,omitempty"`
	Target               string   `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	Label                string   `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	Direction            uint32   `protobuf:"varint,7,opt,name=direction,proto3" json:"direction,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Weight               uint32   `protobuf:"varint,9,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVEdgeAdd) Reset()         { *m = ReqVEdgeAdd{} }
func (m *ReqVEdgeAdd) String() string { return proto.CompactTextString(m) }
func (*ReqVEdgeAdd) ProtoMessage()    {}
func (*ReqVEdgeAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{1}
}

func (m *ReqVEdgeAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqVEdgeAdd.Unmarshal(m, b)
}
func (m *ReqVEdgeAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqVEdgeAdd.Marshal(b, m, deterministic)
}
func (m *ReqVEdgeAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqVEdgeAdd.Merge(m, src)
}
func (m *ReqVEdgeAdd) XXX_Size() int {
	return xxx_messageInfo_ReqVEdgeAdd.Size(m)
}
func (m *ReqVEdgeAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqVEdgeAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqVEdgeAdd proto.InternalMessageInfo

func (m *ReqVEdgeAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqVEdgeAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqVEdgeAdd) GetCenter() string {
	if m != nil {
		return m.Center
	}
	return ""
}

func (m *ReqVEdgeAdd) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *ReqVEdgeAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqVEdgeAdd) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ReqVEdgeAdd) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqVEdgeAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqVEdgeAdd) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type ReqVEdgeUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Relation             string   `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	Target               string   `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	Label                string   `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Direction            uint32   `protobuf:"varint,6,opt,name=direction,proto3" json:"direction,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVEdgeUpdate) Reset()         { *m = ReqVEdgeUpdate{} }
func (m *ReqVEdgeUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqVEdgeUpdate) ProtoMessage()    {}
func (*ReqVEdgeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{2}
}

func (m *ReqVEdgeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqVEdgeUpdate.Unmarshal(m, b)
}
func (m *ReqVEdgeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqVEdgeUpdate.Marshal(b, m, deterministic)
}
func (m *ReqVEdgeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqVEdgeUpdate.Merge(m, src)
}
func (m *ReqVEdgeUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqVEdgeUpdate.Size(m)
}
func (m *ReqVEdgeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqVEdgeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqVEdgeUpdate proto.InternalMessageInfo

func (m *ReqVEdgeUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqVEdgeUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyVEdgeInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *VEdgeInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyVEdgeInfo) Reset()         { *m = ReplyVEdgeInfo{} }
func (m *ReplyVEdgeInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyVEdgeInfo) ProtoMessage()    {}
func (*ReplyVEdgeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{3}
}

func (m *ReplyVEdgeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyVEdgeInfo.Unmarshal(m, b)
}
func (m *ReplyVEdgeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyVEdgeInfo.Marshal(b, m, deterministic)
}
func (m *ReplyVEdgeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyVEdgeInfo.Merge(m, src)
}
func (m *ReplyVEdgeInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyVEdgeInfo.Size(m)
}
func (m *ReplyVEdgeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyVEdgeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyVEdgeInfo proto.InternalMessageInfo

func (m *ReplyVEdgeInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyVEdgeInfo) GetInfo() *VEdgeInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyVEdgeList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	List                 []*VEdgeInfo `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyVEdgeList) Reset()         { *m = ReplyVEdgeList{} }
func (m *ReplyVEdgeList) String() string { return proto.CompactTextString(m) }
func (*ReplyVEdgeList) ProtoMessage()    {}
func (*ReplyVEdgeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{4}
}

func (m *ReplyVEdgeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyVEdgeList.Unmarshal(m, b)
}
func (m *ReplyVEdgeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyVEdgeList.Marshal(b, m, deterministic)
}
func (m *ReplyVEdgeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyVEdgeList.Merge(m, src)
}
func (m *ReplyVEdgeList) XXX_Size() int {
	return xxx_messageInfo_ReplyVEdgeList.Size(m)
}
func (m *ReplyVEdgeList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyVEdgeList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyVEdgeList proto.InternalMessageInfo

func (m *ReplyVEdgeList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyVEdgeList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyVEdgeList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyVEdgeList) GetList() []*VEdgeInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*VEdgeInfo)(nil), "omo.msp.vocabulary.VEdgeInfo")
	proto.RegisterType((*ReqVEdgeAdd)(nil), "omo.msp.vocabulary.ReqVEdgeAdd")
	proto.RegisterType((*ReqVEdgeUpdate)(nil), "omo.msp.vocabulary.ReqVEdgeUpdate")
	proto.RegisterType((*ReplyVEdgeInfo)(nil), "omo.msp.vocabulary.ReplyVEdgeInfo")
	proto.RegisterType((*ReplyVEdgeList)(nil), "omo.msp.vocabulary.ReplyVEdgeList")
}

func init() {
	proto.RegisterFile("proto/vocabulary/edge.proto", fileDescriptor_06ad0988cd483830)
}

var fileDescriptor_06ad0988cd483830 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xc7, 0xeb, 0x4b, 0xdc, 0x7a, 0xd2, 0x44, 0xd5, 0xa8, 0xfa, 0x34, 0xea, 0x47, 0x45, 0x30,
	0x9b, 0xac, 0x52, 0x11, 0x16, 0xac, 0x53, 0x09, 0x2a, 0x24, 0x50, 0x25, 0x57, 0x14, 0x89, 0x9d,
	0x63, 0x9f, 0x86, 0x91, 0x6c, 0x8f, 0xb1, 0x27, 0x41, 0x91, 0x78, 0x20, 0x9e, 0x81, 0x37, 0x41,
	0x3c, 0x03, 0xef, 0x80, 0xe6, 0x8c, 0x63, 0x3b, 0x17, 0x87, 0x00, 0x3b, 0x9f, 0xdb, 0xff, 0xcc,
	0xef, 0xcc, 0x19, 0x99, 0xfc, 0x9f, 0xe5, 0x42, 0x8a, 0xab, 0x85, 0x08, 0x83, 0xe9, 0x3c, 0x0e,
	0xf2, 0xe5, 0x15, 0x44, 0x33, 0x18, 0xa1, 0x97, 0x52, 0x91, 0x88, 0x51, 0x52, 0x64, 0xa3, 0x3a,
	0x7c, 0x71, 0xb9, 0x55, 0x10, 0x8a, 0x24, 0x11, 0xa9, 0x2e, 0xf1, 0x7e, 0x98, 0xc4, 0xbd, 0x7f,
	0x19, 0xcd, 0xe0, 0x75, 0xfa, 0x20, 0x68, 0x9f, 0x98, 0x3c, 0x62, 0xc6, 0xc0, 0x18, 0x5a, 0xbe,
	0xc9, 0x23, 0x7a, 0x46, 0xac, 0x39, 0x8f, 0x98, 0x39, 0x30, 0x86, 0xae, 0xaf, 0x3e, 0x29, 0x25,
	0x76, 0x1a, 0x24, 0xc0, 0x2c, 0x74, 0xe1, 0x37, 0x65, 0xe4, 0x38, 0xcc, 0x21, 0x90, 0x10, 0x31,
	0x1b, 0x4b, 0x57, 0xa6, 0x8a, 0xcc, 0xb3, 0x08, 0x23, 0x1d, 0x1d, 0x29, 0xcd, 0xaa, 0x46, 0xe4,
	0xcc, 0x41, 0xa9, 0x95, 0x49, 0x2f, 0xc8, 0x89, 0xc8, 0x20, 0xc7, 0xd0, 0x31, 0x86, 0x2a, 0x9b,
	0xfe, 0x47, 0x1c, 0x48, 0x25, 0x97, 0x4b, 0x76, 0x82, 0x91, 0xd2, 0x52, 0x35, 0x39, 0xc4, 0x81,
	0xe4, 0x22, 0x65, 0xae, 0xae, 0x59, 0xd9, 0xf4, 0x11, 0x71, 0x23, 0x9e, 0x43, 0x88, 0x41, 0x32,
	0x30, 0x86, 0x3d, 0xbf, 0x76, 0x28, 0x45, 0x19, 0xe4, 0x33, 0x90, 0xac, 0xab, 0x15, 0xb5, 0x45,
	0xcf, 0x49, 0x27, 0x0e, 0xa6, 0x10, 0xb3, 0x53, 0x74, 0x6b, 0x43, 0x65, 0x87, 0x90, 0x4a, 0xc8,
	0x59, 0x4f, 0x67, 0x6b, 0x4b, 0xf9, 0x3f, 0x03, 0x9f, 0x7d, 0x94, 0xac, 0x8f, 0x0d, 0x4a, 0xcb,
	0xfb, 0x69, 0x90, 0xae, 0x0f, 0x9f, 0x70, 0xc0, 0x93, 0xa8, 0x9e, 0x9e, 0xd1, 0x98, 0x5e, 0xcd,
	0x64, 0xae, 0x31, 0xd5, 0xbd, 0xac, 0xb5, 0x5e, 0x4d, 0x56, 0x7b, 0x83, 0xb5, 0xa6, 0xe9, 0xec,
	0xa6, 0x71, 0x9a, 0x34, 0x6b, 0x93, 0x39, 0xde, 0x9c, 0x4c, 0xf3, 0x1e, 0x4e, 0xb6, 0xef, 0xa1,
	0xe4, 0x75, 0xd7, 0x78, 0xbf, 0x19, 0xa4, 0xbf, 0xe2, 0x7d, 0x87, 0x37, 0xbd, 0x5a, 0x21, 0x63,
	0x7b, 0x85, 0xcc, 0xc6, 0x10, 0x9a, 0x50, 0x56, 0x2b, 0x94, 0xbd, 0x1b, 0xaa, 0xd3, 0x0a, 0xe5,
	0xec, 0x83, 0xda, 0x58, 0x2e, 0xef, 0x8b, 0x3a, 0x7b, 0x16, 0x2f, 0xeb, 0xe7, 0xf0, 0x82, 0x38,
	0x85, 0x0c, 0xe4, 0xbc, 0xc0, 0xe3, 0x77, 0xc7, 0x8f, 0x47, 0xdb, 0x0f, 0x6c, 0x84, 0x35, 0x77,
	0x98, 0xe6, 0x97, 0xe9, 0xf4, 0x19, 0xb1, 0x79, 0xfa, 0x20, 0x10, 0xb1, 0x3b, 0xbe, 0xdc, 0x55,
	0x56, 0x75, 0xf1, 0x31, 0xd5, 0xfb, 0x6a, 0x34, 0xdb, 0xbf, 0xe1, 0x85, 0xfc, 0xfb, 0xf6, 0xe7,
	0xa4, 0x23, 0x85, 0x0c, 0x62, 0xec, 0xdf, 0xf3, 0xb5, 0xa1, 0xbc, 0x59, 0x30, 0x83, 0x02, 0x07,
	0xdc, 0xf3, 0xb5, 0xa1, 0x8e, 0x1a, 0xf3, 0x42, 0xad, 0xbf, 0x75, 0xc0, 0x51, 0x55, 0xea, 0xf8,
	0xbb, 0x4d, 0x4e, 0xd1, 0x77, 0x07, 0xf9, 0x82, 0x87, 0x40, 0x6f, 0x89, 0x33, 0x89, 0xa2, 0xdb,
	0x14, 0x68, 0xcb, 0x11, 0xab, 0x17, 0x70, 0xe1, 0xb5, 0x32, 0x54, 0x5d, 0xbc, 0x23, 0x25, 0x78,
	0x03, 0x72, 0x9f, 0xe0, 0x1c, 0x0a, 0xa9, 0x92, 0x0f, 0x14, 0x7c, 0x4b, 0x5c, 0x1f, 0x12, 0xb1,
	0x80, 0x83, 0x34, 0x2f, 0x5b, 0x35, 0x4b, 0xb9, 0x7b, 0x42, 0xf4, 0x7a, 0xe3, 0x9a, 0x78, 0xfb,
	0xa0, 0x75, 0xde, 0x1f, 0x71, 0x4f, 0xe2, 0xf8, 0x9f, 0xb9, 0xd5, 0x02, 0x79, 0x47, 0xf4, 0x3d,
	0x39, 0xbd, 0x01, 0xa9, 0xd6, 0x83, 0x17, 0x92, 0x87, 0xf4, 0xc9, 0x1e, 0xd9, 0x57, 0x3c, 0x96,
	0x90, 0xef, 0x11, 0xae, 0x64, 0x70, 0x02, 0x7d, 0x4d, 0x76, 0xbd, 0xd4, 0x75, 0xf4, 0x69, 0x8b,
	0xb4, 0x4e, 0x2b, 0xc5, 0x7f, 0x37, 0xd9, 0x6b, 0xfa, 0xe1, 0x6c, 0xf3, 0x87, 0x35, 0x75, 0xd0,
	0xf3, 0xfc, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x37, 0x74, 0x45, 0xfc, 0x06, 0x00, 0x00,
}