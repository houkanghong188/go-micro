// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auditConf.proto

/*
Package auditConf is a generated protocol buffer package.

It is generated from these files:
	auditConf.proto

It has these top-level messages:
	Request
	Response
*/
package auditConf

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

// Client API for AuditConf service

type AuditConfService interface {
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Show(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Index(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type auditConfService struct {
	c    client.Client
	name string
}

func NewAuditConfService(name string, c client.Client) AuditConfService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "auditconf"
	}
	return &auditConfService{
		c:    c,
		name: name,
	}
}

func (c *auditConfService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuditConf.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditConfService) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuditConf.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditConfService) Show(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuditConf.Show", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditConfService) Index(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuditConf.Index", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuditConf service

type AuditConfHandler interface {
	Create(context.Context, *Request, *Response) error
	Update(context.Context, *Request, *Response) error
	Show(context.Context, *Request, *Response) error
	Index(context.Context, *Request, *Response) error
}

func RegisterAuditConfHandler(s server.Server, hdlr AuditConfHandler, opts ...server.HandlerOption) error {
	type auditConf interface {
		Create(ctx context.Context, in *Request, out *Response) error
		Update(ctx context.Context, in *Request, out *Response) error
		Show(ctx context.Context, in *Request, out *Response) error
		Index(ctx context.Context, in *Request, out *Response) error
	}
	type AuditConf struct {
		auditConf
	}
	h := &auditConfHandler{hdlr}
	return s.Handle(s.NewHandler(&AuditConf{h}, opts...))
}

type auditConfHandler struct {
	AuditConfHandler
}

func (h *auditConfHandler) Create(ctx context.Context, in *Request, out *Response) error {
	return h.AuditConfHandler.Create(ctx, in, out)
}

func (h *auditConfHandler) Update(ctx context.Context, in *Request, out *Response) error {
	return h.AuditConfHandler.Update(ctx, in, out)
}

func (h *auditConfHandler) Show(ctx context.Context, in *Request, out *Response) error {
	return h.AuditConfHandler.Show(ctx, in, out)
}

func (h *auditConfHandler) Index(ctx context.Context, in *Request, out *Response) error {
	return h.AuditConfHandler.Index(ctx, in, out)
}
