// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/event.proto

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

type EventInfo struct {
	Id                   uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string              `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Created              int64               `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64               `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string              `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Parent               string              `protobuf:"bytes,6,opt,name=parent,proto3" json:"parent,omitempty"`
	Description          string              `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Date                 *DateInfo           `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo          `protobuf:"bytes,9,opt,name=place,proto3" json:"place,omitempty"`
	Creator              string              `protobuf:"bytes,10,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string              `protobuf:"bytes,11,opt,name=operator,proto3" json:"operator,omitempty"`
	Cover                string              `protobuf:"bytes,12,opt,name=cover,proto3" json:"cover,omitempty"`
	Quote                string              `protobuf:"bytes,13,opt,name=quote,proto3" json:"quote,omitempty"`
	Type                 uint32              `protobuf:"varint,14,opt,name=type,proto3" json:"type,omitempty"`
	Access               uint32              `protobuf:"varint,15,opt,name=access,proto3" json:"access,omitempty"`
	Tags                 []string            `protobuf:"bytes,16,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string            `protobuf:"bytes,17,rep,name=assets,proto3" json:"assets,omitempty"`
	Relations            []*RelationshipInfo `protobuf:"bytes,18,rep,name=relations,proto3" json:"relations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EventInfo) Reset()         { *m = EventInfo{} }
func (m *EventInfo) String() string { return proto.CompactTextString(m) }
func (*EventInfo) ProtoMessage()    {}
func (*EventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{0}
}

func (m *EventInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventInfo.Unmarshal(m, b)
}
func (m *EventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventInfo.Marshal(b, m, deterministic)
}
func (m *EventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventInfo.Merge(m, src)
}
func (m *EventInfo) XXX_Size() int {
	return xxx_messageInfo_EventInfo.Size(m)
}
func (m *EventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EventInfo proto.InternalMessageInfo

func (m *EventInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *EventInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *EventInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *EventInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventInfo) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *EventInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventInfo) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *EventInfo) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *EventInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EventInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *EventInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *EventInfo) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *EventInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *EventInfo) GetAccess() uint32 {
	if m != nil {
		return m.Access
	}
	return 0
}

func (m *EventInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *EventInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *EventInfo) GetRelations() []*RelationshipInfo {
	if m != nil {
		return m.Relations
	}
	return nil
}

type RelationshipInfo struct {
	Direction            DirectionType `protobuf:"varint,1,opt,name=direction,proto3,enum=omo.msp.vocabulary.DirectionType" json:"direction,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category             string        `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Entity               string        `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	Uid                  string        `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RelationshipInfo) Reset()         { *m = RelationshipInfo{} }
func (m *RelationshipInfo) String() string { return proto.CompactTextString(m) }
func (*RelationshipInfo) ProtoMessage()    {}
func (*RelationshipInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{1}
}

func (m *RelationshipInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipInfo.Unmarshal(m, b)
}
func (m *RelationshipInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipInfo.Marshal(b, m, deterministic)
}
func (m *RelationshipInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipInfo.Merge(m, src)
}
func (m *RelationshipInfo) XXX_Size() int {
	return xxx_messageInfo_RelationshipInfo.Size(m)
}
func (m *RelationshipInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipInfo proto.InternalMessageInfo

func (m *RelationshipInfo) GetDirection() DirectionType {
	if m != nil {
		return m.Direction
	}
	return DirectionType_Double
}

func (m *RelationshipInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RelationshipInfo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *RelationshipInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *RelationshipInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ReqEventAdd struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Parent               string              `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Description          string              `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date                 *DateInfo           `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo          `protobuf:"bytes,5,opt,name=place,proto3" json:"place,omitempty"`
	Operator             string              `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Cover                string              `protobuf:"bytes,7,opt,name=cover,proto3" json:"cover,omitempty"`
	Quote                string              `protobuf:"bytes,8,opt,name=quote,proto3" json:"quote,omitempty"`
	Type                 uint32              `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Access               uint32              `protobuf:"varint,10,opt,name=access,proto3" json:"access,omitempty"`
	Tags                 []string            `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string            `protobuf:"bytes,12,rep,name=assets,proto3" json:"assets,omitempty"`
	Relations            []*RelationshipInfo `protobuf:"bytes,13,rep,name=relations,proto3" json:"relations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ReqEventAdd) Reset()         { *m = ReqEventAdd{} }
func (m *ReqEventAdd) String() string { return proto.CompactTextString(m) }
func (*ReqEventAdd) ProtoMessage()    {}
func (*ReqEventAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{2}
}

func (m *ReqEventAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEventAdd.Unmarshal(m, b)
}
func (m *ReqEventAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEventAdd.Marshal(b, m, deterministic)
}
func (m *ReqEventAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEventAdd.Merge(m, src)
}
func (m *ReqEventAdd) XXX_Size() int {
	return xxx_messageInfo_ReqEventAdd.Size(m)
}
func (m *ReqEventAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEventAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEventAdd proto.InternalMessageInfo

func (m *ReqEventAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqEventAdd) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReqEventAdd) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReqEventAdd) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ReqEventAdd) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ReqEventAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqEventAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqEventAdd) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqEventAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqEventAdd) GetAccess() uint32 {
	if m != nil {
		return m.Access
	}
	return 0
}

func (m *ReqEventAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqEventAdd) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ReqEventAdd) GetRelations() []*RelationshipInfo {
	if m != nil {
		return m.Relations
	}
	return nil
}

type ReqEventUpdate struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date                 *DateInfo  `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo `protobuf:"bytes,5,opt,name=place,proto3" json:"place,omitempty"`
	Operator             string     `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Cover                string     `protobuf:"bytes,7,opt,name=cover,proto3" json:"cover,omitempty"`
	Quote                string     `protobuf:"bytes,8,opt,name=quote,proto3" json:"quote,omitempty"`
	Access               uint32     `protobuf:"varint,9,opt,name=access,proto3" json:"access,omitempty"`
	Assets               []string   `protobuf:"bytes,10,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqEventUpdate) Reset()         { *m = ReqEventUpdate{} }
func (m *ReqEventUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqEventUpdate) ProtoMessage()    {}
func (*ReqEventUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{3}
}

func (m *ReqEventUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEventUpdate.Unmarshal(m, b)
}
func (m *ReqEventUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEventUpdate.Marshal(b, m, deterministic)
}
func (m *ReqEventUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEventUpdate.Merge(m, src)
}
func (m *ReqEventUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqEventUpdate.Size(m)
}
func (m *ReqEventUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEventUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEventUpdate proto.InternalMessageInfo

func (m *ReqEventUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqEventUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqEventUpdate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReqEventUpdate) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ReqEventUpdate) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ReqEventUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqEventUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqEventUpdate) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqEventUpdate) GetAccess() uint32 {
	if m != nil {
		return m.Access
	}
	return 0
}

func (m *ReqEventUpdate) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqEventAccess struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Access               uint32   `protobuf:"varint,3,opt,name=access,proto3" json:"access,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEventAccess) Reset()         { *m = ReqEventAccess{} }
func (m *ReqEventAccess) String() string { return proto.CompactTextString(m) }
func (*ReqEventAccess) ProtoMessage()    {}
func (*ReqEventAccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{4}
}

func (m *ReqEventAccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEventAccess.Unmarshal(m, b)
}
func (m *ReqEventAccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEventAccess.Marshal(b, m, deterministic)
}
func (m *ReqEventAccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEventAccess.Merge(m, src)
}
func (m *ReqEventAccess) XXX_Size() int {
	return xxx_messageInfo_ReqEventAccess.Size(m)
}
func (m *ReqEventAccess) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEventAccess.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEventAccess proto.InternalMessageInfo

func (m *ReqEventAccess) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqEventAccess) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqEventAccess) GetAccess() uint32 {
	if m != nil {
		return m.Access
	}
	return 0
}

type ReqEventAsset struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Asset                string   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEventAsset) Reset()         { *m = ReqEventAsset{} }
func (m *ReqEventAsset) String() string { return proto.CompactTextString(m) }
func (*ReqEventAsset) ProtoMessage()    {}
func (*ReqEventAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{5}
}

func (m *ReqEventAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEventAsset.Unmarshal(m, b)
}
func (m *ReqEventAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEventAsset.Marshal(b, m, deterministic)
}
func (m *ReqEventAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEventAsset.Merge(m, src)
}
func (m *ReqEventAsset) XXX_Size() int {
	return xxx_messageInfo_ReqEventAsset.Size(m)
}
func (m *ReqEventAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEventAsset.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEventAsset proto.InternalMessageInfo

func (m *ReqEventAsset) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqEventAsset) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *ReqEventAsset) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyEventAssets struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Assets               []string     `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyEventAssets) Reset()         { *m = ReplyEventAssets{} }
func (m *ReplyEventAssets) String() string { return proto.CompactTextString(m) }
func (*ReplyEventAssets) ProtoMessage()    {}
func (*ReplyEventAssets) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{6}
}

func (m *ReplyEventAssets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyEventAssets.Unmarshal(m, b)
}
func (m *ReplyEventAssets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyEventAssets.Marshal(b, m, deterministic)
}
func (m *ReplyEventAssets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyEventAssets.Merge(m, src)
}
func (m *ReplyEventAssets) XXX_Size() int {
	return xxx_messageInfo_ReplyEventAssets.Size(m)
}
func (m *ReplyEventAssets) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyEventAssets.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyEventAssets proto.InternalMessageInfo

func (m *ReplyEventAssets) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyEventAssets) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ReplyEventAssets) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqEventRelation struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Relation             *RelationshipInfo `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	Operator             string            `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReqEventRelation) Reset()         { *m = ReqEventRelation{} }
func (m *ReqEventRelation) String() string { return proto.CompactTextString(m) }
func (*ReqEventRelation) ProtoMessage()    {}
func (*ReqEventRelation) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{7}
}

func (m *ReqEventRelation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEventRelation.Unmarshal(m, b)
}
func (m *ReqEventRelation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEventRelation.Marshal(b, m, deterministic)
}
func (m *ReqEventRelation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEventRelation.Merge(m, src)
}
func (m *ReqEventRelation) XXX_Size() int {
	return xxx_messageInfo_ReqEventRelation.Size(m)
}
func (m *ReqEventRelation) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEventRelation.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEventRelation proto.InternalMessageInfo

func (m *ReqEventRelation) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqEventRelation) GetRelation() *RelationshipInfo {
	if m != nil {
		return m.Relation
	}
	return nil
}

func (m *ReqEventRelation) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyEventRelations struct {
	Uid                  string              `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Relations            []*RelationshipInfo `protobuf:"bytes,2,rep,name=relations,proto3" json:"relations,omitempty"`
	Status               *ReplyStatus        `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ReplyEventRelations) Reset()         { *m = ReplyEventRelations{} }
func (m *ReplyEventRelations) String() string { return proto.CompactTextString(m) }
func (*ReplyEventRelations) ProtoMessage()    {}
func (*ReplyEventRelations) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{8}
}

func (m *ReplyEventRelations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyEventRelations.Unmarshal(m, b)
}
func (m *ReplyEventRelations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyEventRelations.Marshal(b, m, deterministic)
}
func (m *ReplyEventRelations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyEventRelations.Merge(m, src)
}
func (m *ReplyEventRelations) XXX_Size() int {
	return xxx_messageInfo_ReplyEventRelations.Size(m)
}
func (m *ReplyEventRelations) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyEventRelations.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyEventRelations proto.InternalMessageInfo

func (m *ReplyEventRelations) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyEventRelations) GetRelations() []*RelationshipInfo {
	if m != nil {
		return m.Relations
	}
	return nil
}

func (m *ReplyEventRelations) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyEventInfo struct {
	Info                 *EventInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyEventInfo) Reset()         { *m = ReplyEventInfo{} }
func (m *ReplyEventInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyEventInfo) ProtoMessage()    {}
func (*ReplyEventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{9}
}

func (m *ReplyEventInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyEventInfo.Unmarshal(m, b)
}
func (m *ReplyEventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyEventInfo.Marshal(b, m, deterministic)
}
func (m *ReplyEventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyEventInfo.Merge(m, src)
}
func (m *ReplyEventInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyEventInfo.Size(m)
}
func (m *ReplyEventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyEventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyEventInfo proto.InternalMessageInfo

func (m *ReplyEventInfo) GetInfo() *EventInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyEventInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyEventList struct {
	List                 []*EventInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyEventList) Reset()         { *m = ReplyEventList{} }
func (m *ReplyEventList) String() string { return proto.CompactTextString(m) }
func (*ReplyEventList) ProtoMessage()    {}
func (*ReplyEventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc16f7316c625f8c, []int{10}
}

func (m *ReplyEventList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyEventList.Unmarshal(m, b)
}
func (m *ReplyEventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyEventList.Marshal(b, m, deterministic)
}
func (m *ReplyEventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyEventList.Merge(m, src)
}
func (m *ReplyEventList) XXX_Size() int {
	return xxx_messageInfo_ReplyEventList.Size(m)
}
func (m *ReplyEventList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyEventList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyEventList proto.InternalMessageInfo

func (m *ReplyEventList) GetList() []*EventInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyEventList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*EventInfo)(nil), "omo.msp.vocabulary.EventInfo")
	proto.RegisterType((*RelationshipInfo)(nil), "omo.msp.vocabulary.RelationshipInfo")
	proto.RegisterType((*ReqEventAdd)(nil), "omo.msp.vocabulary.ReqEventAdd")
	proto.RegisterType((*ReqEventUpdate)(nil), "omo.msp.vocabulary.ReqEventUpdate")
	proto.RegisterType((*ReqEventAccess)(nil), "omo.msp.vocabulary.ReqEventAccess")
	proto.RegisterType((*ReqEventAsset)(nil), "omo.msp.vocabulary.ReqEventAsset")
	proto.RegisterType((*ReplyEventAssets)(nil), "omo.msp.vocabulary.ReplyEventAssets")
	proto.RegisterType((*ReqEventRelation)(nil), "omo.msp.vocabulary.ReqEventRelation")
	proto.RegisterType((*ReplyEventRelations)(nil), "omo.msp.vocabulary.ReplyEventRelations")
	proto.RegisterType((*ReplyEventInfo)(nil), "omo.msp.vocabulary.ReplyEventInfo")
	proto.RegisterType((*ReplyEventList)(nil), "omo.msp.vocabulary.ReplyEventList")
}

func init() {
	proto.RegisterFile("proto/vocabulary/event.proto", fileDescriptor_fc16f7316c625f8c)
}

var fileDescriptor_fc16f7316c625f8c = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0x1e, 0xdb, 0xf9, 0x73, 0x39, 0x09, 0xa1, 0x59, 0xa1, 0x56, 0xb4, 0x23, 0xbc, 0xd6, 0x48,
	0xe4, 0x94, 0x85, 0xec, 0x81, 0x23, 0x64, 0x00, 0xad, 0x90, 0x80, 0x59, 0x9c, 0x65, 0x41, 0x7b,
	0xf3, 0xd8, 0xbd, 0x8b, 0xa5, 0xc4, 0xed, 0xb1, 0xdb, 0x91, 0x22, 0x71, 0xe2, 0x19, 0x78, 0x04,
	0xc4, 0x8d, 0x33, 0x8f, 0xc0, 0x6b, 0xa1, 0xae, 0xf6, 0x4f, 0x32, 0xe3, 0x78, 0x32, 0x09, 0x5c,
	0xf6, 0xe6, 0xea, 0xfe, 0xfa, 0xab, 0xaf, 0xaa, 0xbf, 0x54, 0x2b, 0xf0, 0x38, 0x4e, 0xb8, 0xe0,
	0x4f, 0xd7, 0xdc, 0xf7, 0xae, 0xb3, 0xa5, 0x97, 0x6c, 0x9e, 0xb2, 0x35, 0x8b, 0xc4, 0x14, 0x97,
	0x09, 0xe1, 0x2b, 0x3e, 0x5d, 0xa5, 0xf1, 0xb4, 0xda, 0x1f, 0x9f, 0xdf, 0x39, 0xe1, 0xf3, 0xd5,
	0x8a, 0x47, 0xea, 0x88, 0xf3, 0x7b, 0x0b, 0xcc, 0xaf, 0x25, 0xc5, 0x37, 0xd1, 0x1b, 0x4e, 0x86,
	0xa0, 0x87, 0x01, 0xd5, 0x6c, 0x6d, 0xd2, 0x72, 0xf5, 0x30, 0x20, 0x23, 0x30, 0xb2, 0x30, 0xa0,
	0xba, 0xad, 0x4d, 0x4c, 0x57, 0x7e, 0x12, 0x0a, 0x5d, 0x3f, 0x61, 0x9e, 0x60, 0x01, 0x35, 0x6c,
	0x6d, 0x62, 0xb8, 0x45, 0x28, 0x77, 0xb2, 0x38, 0xc0, 0x9d, 0x96, 0xda, 0xc9, 0x43, 0x42, 0xa0,
	0x15, 0x79, 0x2b, 0x46, 0xdb, 0x48, 0x83, 0xdf, 0xe4, 0x43, 0xe8, 0xc4, 0x5e, 0xc2, 0x22, 0x41,
	0x3b, 0xb8, 0x9a, 0x47, 0xc4, 0x06, 0x2b, 0x60, 0xa9, 0x9f, 0x84, 0xb1, 0x08, 0x79, 0x44, 0xbb,
	0xb8, 0xb9, 0xbd, 0x44, 0x3e, 0x81, 0x96, 0xa4, 0xa5, 0x3d, 0x5b, 0x9b, 0x58, 0xb3, 0xc7, 0xd3,
	0xbb, 0x35, 0x4f, 0xbf, 0xf2, 0x04, 0x93, 0xf5, 0xb8, 0x88, 0x24, 0xcf, 0xa0, 0x1d, 0x2f, 0x3d,
	0x9f, 0x51, 0x13, 0x8f, 0x9c, 0xd7, 0x1d, 0x79, 0x21, 0x01, 0x78, 0x46, 0x61, 0xcb, 0x42, 0x79,
	0x42, 0x01, 0x45, 0x14, 0x21, 0x19, 0x43, 0x8f, 0xc7, 0x2c, 0xc1, 0x2d, 0x0b, 0xb7, 0xca, 0x98,
	0x3c, 0x82, 0xb6, 0xcf, 0xd7, 0x2c, 0xa1, 0x7d, 0xdc, 0x50, 0x81, 0x5c, 0xbd, 0xc9, 0xb8, 0x60,
	0x74, 0xa0, 0x56, 0x31, 0x90, 0x6d, 0x11, 0x9b, 0x98, 0xd1, 0xa1, 0xad, 0x4d, 0x06, 0x2e, 0x7e,
	0xcb, 0xb6, 0x78, 0xbe, 0xcf, 0xd2, 0x94, 0xbe, 0x87, 0xab, 0x79, 0x84, 0x58, 0xef, 0x6d, 0x4a,
	0x47, 0xb6, 0x21, 0x5b, 0x28, 0xbf, 0x11, 0x9b, 0xa6, 0x4c, 0xa4, 0xf4, 0x7d, 0x5c, 0xcd, 0x23,
	0x72, 0x09, 0x66, 0xc2, 0x96, 0x9e, 0x6c, 0x56, 0x4a, 0x89, 0x6d, 0x4c, 0xac, 0xd9, 0x45, 0x5d,
	0xc9, 0x6e, 0x01, 0xfa, 0x25, 0x8c, 0xb1, 0xf2, 0xea, 0x98, 0xf3, 0x97, 0x06, 0xa3, 0xdb, 0xfb,
	0xe4, 0x73, 0x30, 0x83, 0x30, 0x61, 0x3e, 0xde, 0x8c, 0x34, 0xc9, 0x70, 0xf6, 0xa4, 0xb6, 0xfd,
	0x05, 0xe8, 0xe5, 0x26, 0x66, 0x6e, 0x75, 0xa6, 0x34, 0x82, 0xbe, 0x65, 0x84, 0x31, 0xf4, 0x7c,
	0x4f, 0xb0, 0xb7, 0x3c, 0xd9, 0xa0, 0xa3, 0x4c, 0xb7, 0x8c, 0x65, 0x85, 0x2c, 0x12, 0xa1, 0xd8,
	0xa0, 0xa3, 0x4c, 0x37, 0x8f, 0x0a, 0x5b, 0xb6, 0x4b, 0x5b, 0x3a, 0x7f, 0x1a, 0x60, 0xb9, 0xec,
	0x06, 0x9d, 0x3c, 0x0f, 0x2a, 0xcb, 0x69, 0xb5, 0x96, 0xd3, 0x9b, 0x2c, 0x67, 0xec, 0xb7, 0x5c,
	0xeb, 0xe1, 0x96, 0x6b, 0x3f, 0xc0, 0x72, 0xdb, 0xc6, 0xea, 0xec, 0x33, 0x56, 0xb7, 0xd6, 0x58,
	0xbd, 0x3a, 0x63, 0x99, 0xb5, 0xc6, 0x82, 0x5a, 0x63, 0x59, 0xb5, 0xc6, 0xea, 0xef, 0x37, 0xd6,
	0xe0, 0x38, 0x63, 0xfd, 0xad, 0xc3, 0xb0, 0xb8, 0xa8, 0x1f, 0x71, 0x3e, 0x14, 0xb7, 0xa9, 0x55,
	0x43, 0xa6, 0xce, 0x27, 0xef, 0xdc, 0x2d, 0x55, 0x37, 0x62, 0xee, 0xdc, 0x48, 0xd5, 0x7d, 0xd8,
	0xee, 0xbe, 0xf3, 0xaa, 0x6a, 0xdc, 0x5c, 0x21, 0xef, 0x36, 0x6e, 0x5b, 0x9b, 0x7e, 0x4b, 0x5b,
	0x95, 0xcf, 0xd8, 0xce, 0xe7, 0x2c, 0x60, 0x50, 0xf2, 0xca, 0x4c, 0x35, 0xb4, 0x8f, 0xa0, 0x8d,
	0x22, 0x72, 0x4e, 0x15, 0xec, 0x24, 0x33, 0x76, 0x93, 0x39, 0x99, 0x1c, 0x1f, 0xf1, 0x72, 0x53,
	0xd1, 0xd6, 0xc9, 0xad, 0x4a, 0xd5, 0x77, 0x8c, 0xf6, 0x19, 0x74, 0x52, 0xe1, 0x89, 0x4c, 0x49,
	0xb5, 0x66, 0x1f, 0xd5, 0xbb, 0x2c, 0x5e, 0x6e, 0x16, 0x08, 0x73, 0x73, 0xb8, 0xf3, 0x1b, 0x8e,
	0x2d, 0x55, 0x4c, 0xe1, 0xc2, 0x9a, 0xbc, 0x5f, 0x40, 0xaf, 0x70, 0x24, 0x96, 0x74, 0xa8, 0x8f,
	0xcb, 0x53, 0x8d, 0xb5, 0xff, 0xa1, 0xc1, 0x07, 0x55, 0xf1, 0x25, 0x49, 0x8d, 0x8e, 0x9d, 0x1f,
	0x94, 0x7e, 0xd4, 0x0f, 0xea, 0xf8, 0x5e, 0xfd, 0x2a, 0xfd, 0x54, 0xa8, 0xc4, 0xf9, 0xfe, 0x29,
	0xb4, 0xc2, 0xe8, 0x0d, 0x47, 0x85, 0x7b, 0x7e, 0x0d, 0x25, 0xd8, 0x45, 0xe8, 0x56, 0x76, 0xfd,
	0x84, 0xec, 0xdf, 0x86, 0xa9, 0x90, 0xd9, 0x97, 0x61, 0x2a, 0xa8, 0x86, 0x7d, 0xb8, 0x2f, 0xbb,
	0x84, 0x1e, 0x9d, 0x7d, 0xf6, 0x8f, 0x09, 0x7d, 0x24, 0x5b, 0xb0, 0x64, 0x1d, 0xfa, 0x8c, 0x5c,
	0x41, 0x67, 0x1e, 0x04, 0x57, 0x11, 0x23, 0x7b, 0x38, 0xca, 0xa7, 0x65, 0xec, 0xec, 0x4d, 0x52,
	0xca, 0x73, 0xce, 0x24, 0xe1, 0x73, 0x26, 0x9a, 0x08, 0x33, 0x96, 0x22, 0xf8, 0x40, 0xc2, 0xef,
	0xc0, 0x74, 0xd9, 0x8a, 0xaf, 0xd9, 0x41, 0x9c, 0xe7, 0x7b, 0x39, 0x73, 0xba, 0x17, 0xd0, 0x7d,
	0xce, 0x54, 0xe3, 0x4f, 0x15, 0x28, 0x49, 0x9c, 0x33, 0xf2, 0x0a, 0x40, 0x0d, 0xf4, 0x4b, 0x2f,
	0x65, 0xc4, 0x69, 0x6a, 0xa3, 0xc2, 0x1d, 0x58, 0xf8, 0xf7, 0x05, 0xef, 0x4b, 0xf9, 0x36, 0x35,
	0x89, 0x95, 0x42, 0xee, 0xaf, 0xfc, 0x0a, 0x2c, 0xc5, 0xf7, 0x25, 0x0e, 0xe7, 0xd3, 0x5b, 0xb9,
	0x80, 0xbe, 0x22, 0xcc, 0xc7, 0x72, 0x63, 0xe9, 0x0a, 0xf3, 0x00, 0x95, 0x3f, 0xe0, 0x63, 0x71,
	0xba, 0xca, 0x9f, 0x4a, 0x95, 0x6a, 0xc6, 0xde, 0xdb, 0xc8, 0x8b, 0xe6, 0xdb, 0x51, 0x34, 0xce,
	0x19, 0xf9, 0x19, 0xac, 0x79, 0x1c, 0xb3, 0x28, 0x50, 0xaf, 0xc7, 0x93, 0xc6, 0xea, 0x25, 0xe4,
	0x60, 0xe6, 0xd7, 0x30, 0x58, 0x64, 0xd7, 0x22, 0xf1, 0x7c, 0xf1, 0x9f, 0x73, 0x7b, 0x30, 0x54,
	0xaa, 0xcb, 0x67, 0xe2, 0xa2, 0x89, 0xbc, 0x40, 0x8d, 0x3f, 0x6e, 0xe6, 0x2f, 0x27, 0xb5, 0x73,
	0x46, 0x7c, 0x18, 0x15, 0xf2, 0xff, 0xb7, 0x24, 0x97, 0xe4, 0xf5, 0xe8, 0xf6, 0x1f, 0xbc, 0xeb,
	0x0e, 0xae, 0x3c, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xd4, 0x2b, 0xfa, 0x2d, 0x0e, 0x00,
	0x00,
}
