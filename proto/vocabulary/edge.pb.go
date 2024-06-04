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

type ReqVEdgeAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Center               string   `protobuf:"bytes,3,opt,name=center,proto3" json:"center,omitempty"`
	Relation             string   `protobuf:"bytes,4,opt,name=relation,proto3" json:"relation,omitempty"`
	Target               string   `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	Label                string   `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	Direction            uint32   `protobuf:"varint,7,opt,name=direction,proto3" json:"direction,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Weight               uint32   `protobuf:"varint,9,opt,name=weight,proto3" json:"weight,omitempty"`
	Thumb                string   `protobuf:"bytes,10,opt,name=thumb,proto3" json:"thumb,omitempty"`
	Remark               string   `protobuf:"bytes,11,opt,name=remark,proto3" json:"remark,omitempty"`
	Desc                 string   `protobuf:"bytes,12,opt,name=desc,proto3" json:"desc,omitempty"`
	Type                 uint32   `protobuf:"varint,13,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVEdgeAdd) Reset()         { *m = ReqVEdgeAdd{} }
func (m *ReqVEdgeAdd) String() string { return proto.CompactTextString(m) }
func (*ReqVEdgeAdd) ProtoMessage()    {}
func (*ReqVEdgeAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{0}
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

func (m *ReqVEdgeAdd) GetSource() string {
	if m != nil {
		return m.Source
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

func (m *ReqVEdgeAdd) GetThumb() string {
	if m != nil {
		return m.Thumb
	}
	return ""
}

func (m *ReqVEdgeAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqVEdgeAdd) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ReqVEdgeAdd) GetType() uint32 {
	if m != nil {
		return m.Type
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
	Thumb                string   `protobuf:"bytes,8,opt,name=thumb,proto3" json:"thumb,omitempty"`
	Remark               string   `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Desc                 string   `protobuf:"bytes,10,opt,name=desc,proto3" json:"desc,omitempty"`
	Type                 uint32   `protobuf:"varint,11,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVEdgeUpdate) Reset()         { *m = ReqVEdgeUpdate{} }
func (m *ReqVEdgeUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqVEdgeUpdate) ProtoMessage()    {}
func (*ReqVEdgeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{1}
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

func (m *ReqVEdgeUpdate) GetThumb() string {
	if m != nil {
		return m.Thumb
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
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
	return fileDescriptor_06ad0988cd483830, []int{2}
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
	return fileDescriptor_06ad0988cd483830, []int{3}
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
	proto.RegisterType((*ReqVEdgeAdd)(nil), "omo.msp.vocabulary.ReqVEdgeAdd")
	proto.RegisterType((*ReqVEdgeUpdate)(nil), "omo.msp.vocabulary.ReqVEdgeUpdate")
	proto.RegisterType((*ReplyVEdgeInfo)(nil), "omo.msp.vocabulary.ReplyVEdgeInfo")
	proto.RegisterType((*ReplyVEdgeList)(nil), "omo.msp.vocabulary.ReplyVEdgeList")
}

func init() {
	proto.RegisterFile("proto/vocabulary/edge.proto", fileDescriptor_06ad0988cd483830)
}

var fileDescriptor_06ad0988cd483830 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0xd7, 0x3f, 0xcb, 0x5a, 0xa7, 0x9d, 0x26, 0x0b, 0x21, 0xab, 0x50, 0x51, 0xc2, 0xcd,
	0xae, 0x3a, 0x51, 0x2e, 0xb8, 0xee, 0x24, 0x98, 0x90, 0x40, 0x13, 0x99, 0x18, 0x12, 0x77, 0x6e,
	0x72, 0xd6, 0x59, 0x24, 0x71, 0xb0, 0x9d, 0xa2, 0x4a, 0x3c, 0x04, 0x8f, 0xc1, 0x7b, 0xf0, 0x18,
	0xbc, 0x0c, 0xf2, 0x71, 0x96, 0x55, 0x5b, 0xd2, 0x55, 0x70, 0xe7, 0xef, 0xf4, 0x9c, 0xef, 0x3b,
	0xfe, 0xcd, 0xd1, 0xc8, 0x93, 0x5c, 0x49, 0x23, 0x4f, 0x56, 0x32, 0xe2, 0x8b, 0x22, 0xe1, 0x6a,
	0x7d, 0x02, 0xf1, 0x12, 0xa6, 0x58, 0xa5, 0x54, 0xa6, 0x72, 0x9a, 0xea, 0x7c, 0x7a, 0xfb, 0xf3,
	0x68, 0x7c, 0x6f, 0x20, 0x92, 0x69, 0x2a, 0x33, 0x37, 0x12, 0xfc, 0x6e, 0x13, 0x3f, 0x84, 0x6f,
	0x97, 0x6f, 0xe2, 0x25, 0xcc, 0xe3, 0x98, 0x52, 0xd2, 0xcd, 0x78, 0x0a, 0xac, 0x35, 0x69, 0x1d,
	0xf7, 0x43, 0x3c, 0xd3, 0xc7, 0xc4, 0xd3, 0xb2, 0x50, 0x11, 0xb0, 0x36, 0x56, 0x4b, 0x65, 0xeb,
	0x11, 0x64, 0x06, 0x14, 0xeb, 0xb8, 0xba, 0x53, 0x74, 0x44, 0x7a, 0x0a, 0x12, 0x6e, 0x84, 0xcc,
	0x58, 0x17, 0x7f, 0xa9, 0xb4, 0x9d, 0x31, 0x5c, 0x2d, 0xc1, 0xb0, 0x7d, 0x37, 0xe3, 0x14, 0x7d,
	0x44, 0xf6, 0x13, 0xbe, 0x80, 0x84, 0x79, 0x58, 0x76, 0x82, 0x3e, 0x25, 0xfd, 0x58, 0x28, 0x88,
	0xd0, 0xea, 0x60, 0xd2, 0x3a, 0x1e, 0x86, 0xb7, 0x05, 0x9b, 0x23, 0x73, 0x50, 0xdc, 0x48, 0xc5,
	0x7a, 0x2e, 0xe7, 0x46, 0xdb, 0x9c, 0xef, 0x20, 0x96, 0xd7, 0x86, 0xf5, 0x71, 0xac, 0x54, 0x36,
	0xc7, 0x5c, 0x17, 0xe9, 0x82, 0x11, 0x97, 0x83, 0xc2, 0x76, 0x2b, 0x48, 0xb9, 0xfa, 0xca, 0x7c,
	0xb7, 0x95, 0x53, 0x96, 0x46, 0x0c, 0x3a, 0x62, 0x03, 0x47, 0xc3, 0x9e, 0x6d, 0xcd, 0xac, 0x73,
	0x60, 0x43, 0xf4, 0xc5, 0x73, 0xf0, 0xb3, 0x4d, 0x0e, 0x6f, 0x28, 0x7e, 0xca, 0x63, 0x6e, 0x80,
	0x1e, 0x91, 0x4e, 0x21, 0xe2, 0x92, 0xa3, 0x3d, 0x56, 0x68, 0xdb, 0x1b, 0x68, 0x37, 0x51, 0x75,
	0x1a, 0x51, 0x75, 0xeb, 0x51, 0xed, 0x37, 0xa2, 0xf2, 0xb6, 0xa1, 0x3a, 0xb8, 0x83, 0xaa, 0x42,
	0xd2, 0xab, 0x47, 0xd2, 0xaf, 0x45, 0x42, 0x6a, 0x90, 0xf8, 0x1b, 0x48, 0x7e, 0x58, 0x22, 0x79,
	0xb2, 0x46, 0x26, 0xef, 0xb2, 0x2b, 0x49, 0x5f, 0x13, 0x4f, 0x1b, 0x6e, 0x0a, 0x8d, 0x50, 0xfc,
	0xd9, 0xb3, 0xe9, 0xfd, 0xe7, 0x3a, 0xc5, 0x99, 0x0b, 0x6c, 0x0b, 0xcb, 0x76, 0xfa, 0x92, 0x74,
	0x45, 0x76, 0x25, 0x11, 0x9c, 0x3f, 0x1b, 0xd7, 0x8d, 0x55, 0x29, 0x21, 0xb6, 0x06, 0xbf, 0x5a,
	0x9b, 0xf1, 0xef, 0x85, 0x36, 0xff, 0x1e, 0x6f, 0xf9, 0x48, 0xc3, 0x13, 0xcc, 0x1f, 0x86, 0x4e,
	0xd8, 0x6a, 0xce, 0x97, 0xa0, 0xf1, 0xcf, 0x36, 0x0c, 0x9d, 0xb0, 0xab, 0x26, 0x42, 0x1b, 0xe6,
	0x4f, 0x3a, 0x3b, 0xac, 0x6a, 0x5b, 0x67, 0x7f, 0xba, 0x64, 0x80, 0xb5, 0x0b, 0x50, 0x2b, 0x11,
	0x01, 0x3d, 0x27, 0xde, 0x3c, 0x8e, 0xcf, 0x33, 0xa0, 0x0d, 0x2b, 0x56, 0x5f, 0xeb, 0x28, 0x68,
	0xbc, 0x43, 0x95, 0x12, 0xec, 0x59, 0xc3, 0x33, 0x30, 0xdb, 0x0c, 0x0b, 0xd0, 0xc6, 0x36, 0xef,
	0x68, 0xf8, 0x81, 0xf4, 0x43, 0x48, 0xe5, 0x0a, 0x76, 0xf2, 0x1c, 0x37, 0x7a, 0x96, 0x76, 0x97,
	0x84, 0xb8, 0x8f, 0x06, 0x9f, 0x49, 0xb0, 0xed, 0xd2, 0xae, 0x6f, 0xc7, 0x35, 0x3f, 0xe2, 0xbd,
	0xe7, 0x49, 0x42, 0x9f, 0x6f, 0xd9, 0xf1, 0xad, 0x48, 0x0c, 0xa8, 0x87, 0x2c, 0xed, 0x13, 0x0a,
	0xf6, 0xe8, 0x67, 0x32, 0x38, 0x03, 0x63, 0x1f, 0x88, 0xd0, 0x46, 0x44, 0xff, 0x67, 0x5c, 0xd9,
	0x20, 0x83, 0x43, 0x77, 0xb7, 0xd3, 0xb5, 0x9b, 0xa3, 0x2f, 0x1a, 0xac, 0x5d, 0x5b, 0x69, 0xfe,
	0x10, 0xdb, 0x53, 0xfa, 0xe5, 0xe8, 0xee, 0x3f, 0x80, 0x85, 0x87, 0x95, 0x57, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x33, 0x41, 0x94, 0x89, 0x4c, 0x06, 0x00, 0x00,
}
