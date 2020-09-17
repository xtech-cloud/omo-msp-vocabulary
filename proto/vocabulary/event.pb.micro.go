// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/vocabulary/event.proto

package vocabulary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for EventService service

func NewEventServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for EventService service

type EventService interface {
	AddOne(ctx context.Context, in *ReqEventAdd, opts ...client.CallOption) (*ReplyEventInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyEventInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyEventList, error)
	UpdateBase(ctx context.Context, in *ReqEventUpdate, opts ...client.CallOption) (*ReplyEventInfo, error)
	UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyEventInfo, error)
	UpdateCover(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyEventInfo, error)
	UpdateAssets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyEventAssets, error)
	AppendAsset(ctx context.Context, in *ReqEventAsset, opts ...client.CallOption) (*ReplyEventAssets, error)
	SubtractAsset(ctx context.Context, in *ReqEventAsset, opts ...client.CallOption) (*ReplyEventAssets, error)
	AppendRelation(ctx context.Context, in *ReqEventRelation, opts ...client.CallOption) (*ReplyEventRelations, error)
	SubtractRelation(ctx context.Context, in *ReqEventRelation, opts ...client.CallOption) (*ReplyEventRelations, error)
}

type eventService struct {
	c    client.Client
	name string
}

func NewEventService(name string, c client.Client) EventService {
	return &eventService{
		c:    c,
		name: name,
	}
}

func (c *eventService) AddOne(ctx context.Context, in *ReqEventAdd, opts ...client.CallOption) (*ReplyEventInfo, error) {
	req := c.c.NewRequest(c.name, "EventService.AddOne", in)
	out := new(ReplyEventInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyEventInfo, error) {
	req := c.c.NewRequest(c.name, "EventService.GetOne", in)
	out := new(ReplyEventInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "EventService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) GetList(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyEventList, error) {
	req := c.c.NewRequest(c.name, "EventService.GetList", in)
	out := new(ReplyEventList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) UpdateBase(ctx context.Context, in *ReqEventUpdate, opts ...client.CallOption) (*ReplyEventInfo, error) {
	req := c.c.NewRequest(c.name, "EventService.UpdateBase", in)
	out := new(ReplyEventInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyEventInfo, error) {
	req := c.c.NewRequest(c.name, "EventService.UpdateTags", in)
	out := new(ReplyEventInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) UpdateCover(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyEventInfo, error) {
	req := c.c.NewRequest(c.name, "EventService.UpdateCover", in)
	out := new(ReplyEventInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) UpdateAssets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyEventAssets, error) {
	req := c.c.NewRequest(c.name, "EventService.UpdateAssets", in)
	out := new(ReplyEventAssets)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) AppendAsset(ctx context.Context, in *ReqEventAsset, opts ...client.CallOption) (*ReplyEventAssets, error) {
	req := c.c.NewRequest(c.name, "EventService.AppendAsset", in)
	out := new(ReplyEventAssets)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) SubtractAsset(ctx context.Context, in *ReqEventAsset, opts ...client.CallOption) (*ReplyEventAssets, error) {
	req := c.c.NewRequest(c.name, "EventService.SubtractAsset", in)
	out := new(ReplyEventAssets)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) AppendRelation(ctx context.Context, in *ReqEventRelation, opts ...client.CallOption) (*ReplyEventRelations, error) {
	req := c.c.NewRequest(c.name, "EventService.AppendRelation", in)
	out := new(ReplyEventRelations)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventService) SubtractRelation(ctx context.Context, in *ReqEventRelation, opts ...client.CallOption) (*ReplyEventRelations, error) {
	req := c.c.NewRequest(c.name, "EventService.SubtractRelation", in)
	out := new(ReplyEventRelations)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventService service

type EventServiceHandler interface {
	AddOne(context.Context, *ReqEventAdd, *ReplyEventInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyEventInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *RequestInfo, *ReplyEventList) error
	UpdateBase(context.Context, *ReqEventUpdate, *ReplyEventInfo) error
	UpdateTags(context.Context, *RequestList, *ReplyEventInfo) error
	UpdateCover(context.Context, *RequestInfo, *ReplyEventInfo) error
	UpdateAssets(context.Context, *RequestList, *ReplyEventAssets) error
	AppendAsset(context.Context, *ReqEventAsset, *ReplyEventAssets) error
	SubtractAsset(context.Context, *ReqEventAsset, *ReplyEventAssets) error
	AppendRelation(context.Context, *ReqEventRelation, *ReplyEventRelations) error
	SubtractRelation(context.Context, *ReqEventRelation, *ReplyEventRelations) error
}

func RegisterEventServiceHandler(s server.Server, hdlr EventServiceHandler, opts ...server.HandlerOption) error {
	type eventService interface {
		AddOne(ctx context.Context, in *ReqEventAdd, out *ReplyEventInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyEventInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *RequestInfo, out *ReplyEventList) error
		UpdateBase(ctx context.Context, in *ReqEventUpdate, out *ReplyEventInfo) error
		UpdateTags(ctx context.Context, in *RequestList, out *ReplyEventInfo) error
		UpdateCover(ctx context.Context, in *RequestInfo, out *ReplyEventInfo) error
		UpdateAssets(ctx context.Context, in *RequestList, out *ReplyEventAssets) error
		AppendAsset(ctx context.Context, in *ReqEventAsset, out *ReplyEventAssets) error
		SubtractAsset(ctx context.Context, in *ReqEventAsset, out *ReplyEventAssets) error
		AppendRelation(ctx context.Context, in *ReqEventRelation, out *ReplyEventRelations) error
		SubtractRelation(ctx context.Context, in *ReqEventRelation, out *ReplyEventRelations) error
	}
	type EventService struct {
		eventService
	}
	h := &eventServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&EventService{h}, opts...))
}

type eventServiceHandler struct {
	EventServiceHandler
}

func (h *eventServiceHandler) AddOne(ctx context.Context, in *ReqEventAdd, out *ReplyEventInfo) error {
	return h.EventServiceHandler.AddOne(ctx, in, out)
}

func (h *eventServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyEventInfo) error {
	return h.EventServiceHandler.GetOne(ctx, in, out)
}

func (h *eventServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.EventServiceHandler.RemoveOne(ctx, in, out)
}

func (h *eventServiceHandler) GetList(ctx context.Context, in *RequestInfo, out *ReplyEventList) error {
	return h.EventServiceHandler.GetList(ctx, in, out)
}

func (h *eventServiceHandler) UpdateBase(ctx context.Context, in *ReqEventUpdate, out *ReplyEventInfo) error {
	return h.EventServiceHandler.UpdateBase(ctx, in, out)
}

func (h *eventServiceHandler) UpdateTags(ctx context.Context, in *RequestList, out *ReplyEventInfo) error {
	return h.EventServiceHandler.UpdateTags(ctx, in, out)
}

func (h *eventServiceHandler) UpdateCover(ctx context.Context, in *RequestInfo, out *ReplyEventInfo) error {
	return h.EventServiceHandler.UpdateCover(ctx, in, out)
}

func (h *eventServiceHandler) UpdateAssets(ctx context.Context, in *RequestList, out *ReplyEventAssets) error {
	return h.EventServiceHandler.UpdateAssets(ctx, in, out)
}

func (h *eventServiceHandler) AppendAsset(ctx context.Context, in *ReqEventAsset, out *ReplyEventAssets) error {
	return h.EventServiceHandler.AppendAsset(ctx, in, out)
}

func (h *eventServiceHandler) SubtractAsset(ctx context.Context, in *ReqEventAsset, out *ReplyEventAssets) error {
	return h.EventServiceHandler.SubtractAsset(ctx, in, out)
}

func (h *eventServiceHandler) AppendRelation(ctx context.Context, in *ReqEventRelation, out *ReplyEventRelations) error {
	return h.EventServiceHandler.AppendRelation(ctx, in, out)
}

func (h *eventServiceHandler) SubtractRelation(ctx context.Context, in *ReqEventRelation, out *ReplyEventRelations) error {
	return h.EventServiceHandler.SubtractRelation(ctx, in, out)
}
