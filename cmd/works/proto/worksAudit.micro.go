// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: worksAudit.proto

/*
Package worksAudit is a generated protocol buffer package.

It is generated from these files:
	worksAudit.proto

It has these top-level messages:
	Request
	Response
	WorksResponse
	AuditBracket
	WorksBracket
	DailyPvUvBracket
	DailyRequest
	DailyResponse
*/
package worksAudit

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for WorksAudit service

type WorksAuditService interface {
	Show(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Index(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	WorksUpdate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type worksAuditService struct {
	c    client.Client
	name string
}

func NewWorksAuditService(name string, c client.Client) WorksAuditService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "worksAudit"
	}
	return &worksAuditService{
		c:    c,
		name: name,
	}
}

func (c *worksAuditService) Show(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "WorksAudit.Show", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worksAuditService) Index(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "WorksAudit.Index", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worksAuditService) WorksUpdate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "WorksAudit.WorksUpdate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorksAudit service

type WorksAuditHandler interface {
	Show(context.Context, *Request, *Response) error
	Index(context.Context, *Request, *Response) error
	WorksUpdate(context.Context, *Request, *Response) error
}

func RegisterWorksAuditHandler(s server.Server, hdlr WorksAuditHandler, opts ...server.HandlerOption) error {
	type worksAudit interface {
		Show(ctx context.Context, in *Request, out *Response) error
		Index(ctx context.Context, in *Request, out *Response) error
		WorksUpdate(ctx context.Context, in *Request, out *Response) error
	}
	type WorksAudit struct {
		worksAudit
	}
	h := &worksAuditHandler{hdlr}
	return s.Handle(s.NewHandler(&WorksAudit{h}, opts...))
}

type worksAuditHandler struct {
	WorksAuditHandler
}

func (h *worksAuditHandler) Show(ctx context.Context, in *Request, out *Response) error {
	return h.WorksAuditHandler.Show(ctx, in, out)
}

func (h *worksAuditHandler) Index(ctx context.Context, in *Request, out *Response) error {
	return h.WorksAuditHandler.Index(ctx, in, out)
}

func (h *worksAuditHandler) WorksUpdate(ctx context.Context, in *Request, out *Response) error {
	return h.WorksAuditHandler.WorksUpdate(ctx, in, out)
}

// Client API for Works service

type WorksService interface {
	WorksDetail(ctx context.Context, in *Request, opts ...client.CallOption) (*WorksResponse, error)
}

type worksService struct {
	c    client.Client
	name string
}

func NewWorksService(name string, c client.Client) WorksService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "worksAudit"
	}
	return &worksService{
		c:    c,
		name: name,
	}
}

func (c *worksService) WorksDetail(ctx context.Context, in *Request, opts ...client.CallOption) (*WorksResponse, error) {
	req := c.c.NewRequest(c.name, "Works.WorksDetail", in)
	out := new(WorksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Works service

type WorksHandler interface {
	WorksDetail(context.Context, *Request, *WorksResponse) error
}

func RegisterWorksHandler(s server.Server, hdlr WorksHandler, opts ...server.HandlerOption) error {
	type works interface {
		WorksDetail(ctx context.Context, in *Request, out *WorksResponse) error
	}
	type Works struct {
		works
	}
	h := &worksHandler{hdlr}
	return s.Handle(s.NewHandler(&Works{h}, opts...))
}

type worksHandler struct {
	WorksHandler
}

func (h *worksHandler) WorksDetail(ctx context.Context, in *Request, out *WorksResponse) error {
	return h.WorksHandler.WorksDetail(ctx, in, out)
}

// Client API for DailyPvUv service

type DailyPvUvService interface {
	Show(ctx context.Context, in *DailyRequest, opts ...client.CallOption) (*DailyResponse, error)
}

type dailyPvUvService struct {
	c    client.Client
	name string
}

func NewDailyPvUvService(name string, c client.Client) DailyPvUvService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "worksAudit"
	}
	return &dailyPvUvService{
		c:    c,
		name: name,
	}
}

func (c *dailyPvUvService) Show(ctx context.Context, in *DailyRequest, opts ...client.CallOption) (*DailyResponse, error) {
	req := c.c.NewRequest(c.name, "DailyPvUv.Show", in)
	out := new(DailyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DailyPvUv service

type DailyPvUvHandler interface {
	Show(context.Context, *DailyRequest, *DailyResponse) error
}

func RegisterDailyPvUvHandler(s server.Server, hdlr DailyPvUvHandler, opts ...server.HandlerOption) error {
	type dailyPvUv interface {
		Show(ctx context.Context, in *DailyRequest, out *DailyResponse) error
	}
	type DailyPvUv struct {
		dailyPvUv
	}
	h := &dailyPvUvHandler{hdlr}
	return s.Handle(s.NewHandler(&DailyPvUv{h}, opts...))
}

type dailyPvUvHandler struct {
	DailyPvUvHandler
}

func (h *dailyPvUvHandler) Show(ctx context.Context, in *DailyRequest, out *DailyResponse) error {
	return h.DailyPvUvHandler.Show(ctx, in, out)
}
