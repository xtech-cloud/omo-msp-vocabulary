// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/vocabulary/concept.proto

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

// Api Endpoints for ConceptService service

func NewConceptServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ConceptService service

type ConceptService interface {
	AddOne(ctx context.Context, in *ReqConceptAdd, opts ...client.CallOption) (*ReplyConceptInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyConceptInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetAll(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyConceptList, error)
	Update(ctx context.Context, in *ReqConceptUpdate, opts ...client.CallOption) (*ReplyConceptInfo, error)
	UpdateAttributes(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyConceptAttrs, error)
}

type conceptService struct {
	c    client.Client
	name string
}

func NewConceptService(name string, c client.Client) ConceptService {
	return &conceptService{
		c:    c,
		name: name,
	}
}

func (c *conceptService) AddOne(ctx context.Context, in *ReqConceptAdd, opts ...client.CallOption) (*ReplyConceptInfo, error) {
	req := c.c.NewRequest(c.name, "ConceptService.AddOne", in)
	out := new(ReplyConceptInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conceptService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyConceptInfo, error) {
	req := c.c.NewRequest(c.name, "ConceptService.GetOne", in)
	out := new(ReplyConceptInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conceptService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ConceptService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conceptService) GetAll(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyConceptList, error) {
	req := c.c.NewRequest(c.name, "ConceptService.GetAll", in)
	out := new(ReplyConceptList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conceptService) Update(ctx context.Context, in *ReqConceptUpdate, opts ...client.CallOption) (*ReplyConceptInfo, error) {
	req := c.c.NewRequest(c.name, "ConceptService.Update", in)
	out := new(ReplyConceptInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conceptService) UpdateAttributes(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyConceptAttrs, error) {
	req := c.c.NewRequest(c.name, "ConceptService.UpdateAttributes", in)
	out := new(ReplyConceptAttrs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConceptService service

type ConceptServiceHandler interface {
	AddOne(context.Context, *ReqConceptAdd, *ReplyConceptInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyConceptInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetAll(context.Context, *RequestInfo, *ReplyConceptList) error
	Update(context.Context, *ReqConceptUpdate, *ReplyConceptInfo) error
	UpdateAttributes(context.Context, *RequestList, *ReplyConceptAttrs) error
}

func RegisterConceptServiceHandler(s server.Server, hdlr ConceptServiceHandler, opts ...server.HandlerOption) error {
	type conceptService interface {
		AddOne(ctx context.Context, in *ReqConceptAdd, out *ReplyConceptInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyConceptInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetAll(ctx context.Context, in *RequestInfo, out *ReplyConceptList) error
		Update(ctx context.Context, in *ReqConceptUpdate, out *ReplyConceptInfo) error
		UpdateAttributes(ctx context.Context, in *RequestList, out *ReplyConceptAttrs) error
	}
	type ConceptService struct {
		conceptService
	}
	h := &conceptServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConceptService{h}, opts...))
}

type conceptServiceHandler struct {
	ConceptServiceHandler
}

func (h *conceptServiceHandler) AddOne(ctx context.Context, in *ReqConceptAdd, out *ReplyConceptInfo) error {
	return h.ConceptServiceHandler.AddOne(ctx, in, out)
}

func (h *conceptServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyConceptInfo) error {
	return h.ConceptServiceHandler.GetOne(ctx, in, out)
}

func (h *conceptServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.ConceptServiceHandler.RemoveOne(ctx, in, out)
}

func (h *conceptServiceHandler) GetAll(ctx context.Context, in *RequestInfo, out *ReplyConceptList) error {
	return h.ConceptServiceHandler.GetAll(ctx, in, out)
}

func (h *conceptServiceHandler) Update(ctx context.Context, in *ReqConceptUpdate, out *ReplyConceptInfo) error {
	return h.ConceptServiceHandler.Update(ctx, in, out)
}

func (h *conceptServiceHandler) UpdateAttributes(ctx context.Context, in *RequestList, out *ReplyConceptAttrs) error {
	return h.ConceptServiceHandler.UpdateAttributes(ctx, in, out)
}
