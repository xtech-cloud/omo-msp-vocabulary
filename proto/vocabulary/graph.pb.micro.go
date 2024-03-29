// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/vocabulary/graph.proto

package vocabulary

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GraphService service

type GraphService interface {
	AddNode(ctx context.Context, in *ReqNodeAdd, opts ...client.CallOption) (*ReplyNodeInfo, error)
	AddLink(ctx context.Context, in *ReqLinkAdd, opts ...client.CallOption) (*ReplyLinkInfo, error)
	RemoveNode(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveLink(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetNode(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyNodeInfo, error)
	GetLink(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyLinkInfo, error)
	FindNodes(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGraphInfo, error)
	FindPath(ctx context.Context, in *ReqGraphPath, opts ...client.CallOption) (*ReplyGraphInfo, error)
	FindGraph(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGraphInfo, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
}

type graphService struct {
	c    client.Client
	name string
}

func NewGraphService(name string, c client.Client) GraphService {
	return &graphService{
		c:    c,
		name: name,
	}
}

func (c *graphService) AddNode(ctx context.Context, in *ReqNodeAdd, opts ...client.CallOption) (*ReplyNodeInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.AddNode", in)
	out := new(ReplyNodeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) AddLink(ctx context.Context, in *ReqLinkAdd, opts ...client.CallOption) (*ReplyLinkInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.AddLink", in)
	out := new(ReplyLinkInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) RemoveNode(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.RemoveNode", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) RemoveLink(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.RemoveLink", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) GetNode(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyNodeInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.GetNode", in)
	out := new(ReplyNodeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) GetLink(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyLinkInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.GetLink", in)
	out := new(ReplyLinkInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) FindNodes(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGraphInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.FindNodes", in)
	out := new(ReplyGraphInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) FindPath(ctx context.Context, in *ReqGraphPath, opts ...client.CallOption) (*ReplyGraphInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.FindPath", in)
	out := new(ReplyGraphInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) FindGraph(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyGraphInfo, error) {
	req := c.c.NewRequest(c.name, "GraphService.FindGraph", in)
	out := new(ReplyGraphInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graphService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "GraphService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GraphService service

type GraphServiceHandler interface {
	AddNode(context.Context, *ReqNodeAdd, *ReplyNodeInfo) error
	AddLink(context.Context, *ReqLinkAdd, *ReplyLinkInfo) error
	RemoveNode(context.Context, *RequestInfo, *ReplyInfo) error
	RemoveLink(context.Context, *RequestInfo, *ReplyInfo) error
	GetNode(context.Context, *RequestInfo, *ReplyNodeInfo) error
	GetLink(context.Context, *RequestInfo, *ReplyLinkInfo) error
	FindNodes(context.Context, *RequestInfo, *ReplyGraphInfo) error
	FindPath(context.Context, *ReqGraphPath, *ReplyGraphInfo) error
	FindGraph(context.Context, *RequestInfo, *ReplyGraphInfo) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
}

func RegisterGraphServiceHandler(s server.Server, hdlr GraphServiceHandler, opts ...server.HandlerOption) error {
	type graphService interface {
		AddNode(ctx context.Context, in *ReqNodeAdd, out *ReplyNodeInfo) error
		AddLink(ctx context.Context, in *ReqLinkAdd, out *ReplyLinkInfo) error
		RemoveNode(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		RemoveLink(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetNode(ctx context.Context, in *RequestInfo, out *ReplyNodeInfo) error
		GetLink(ctx context.Context, in *RequestInfo, out *ReplyLinkInfo) error
		FindNodes(ctx context.Context, in *RequestInfo, out *ReplyGraphInfo) error
		FindPath(ctx context.Context, in *ReqGraphPath, out *ReplyGraphInfo) error
		FindGraph(ctx context.Context, in *RequestInfo, out *ReplyGraphInfo) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
	}
	type GraphService struct {
		graphService
	}
	h := &graphServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GraphService{h}, opts...))
}

type graphServiceHandler struct {
	GraphServiceHandler
}

func (h *graphServiceHandler) AddNode(ctx context.Context, in *ReqNodeAdd, out *ReplyNodeInfo) error {
	return h.GraphServiceHandler.AddNode(ctx, in, out)
}

func (h *graphServiceHandler) AddLink(ctx context.Context, in *ReqLinkAdd, out *ReplyLinkInfo) error {
	return h.GraphServiceHandler.AddLink(ctx, in, out)
}

func (h *graphServiceHandler) RemoveNode(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.GraphServiceHandler.RemoveNode(ctx, in, out)
}

func (h *graphServiceHandler) RemoveLink(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.GraphServiceHandler.RemoveLink(ctx, in, out)
}

func (h *graphServiceHandler) GetNode(ctx context.Context, in *RequestInfo, out *ReplyNodeInfo) error {
	return h.GraphServiceHandler.GetNode(ctx, in, out)
}

func (h *graphServiceHandler) GetLink(ctx context.Context, in *RequestInfo, out *ReplyLinkInfo) error {
	return h.GraphServiceHandler.GetLink(ctx, in, out)
}

func (h *graphServiceHandler) FindNodes(ctx context.Context, in *RequestInfo, out *ReplyGraphInfo) error {
	return h.GraphServiceHandler.FindNodes(ctx, in, out)
}

func (h *graphServiceHandler) FindPath(ctx context.Context, in *ReqGraphPath, out *ReplyGraphInfo) error {
	return h.GraphServiceHandler.FindPath(ctx, in, out)
}

func (h *graphServiceHandler) FindGraph(ctx context.Context, in *RequestInfo, out *ReplyGraphInfo) error {
	return h.GraphServiceHandler.FindGraph(ctx, in, out)
}

func (h *graphServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.GraphServiceHandler.GetStatistic(ctx, in, out)
}
