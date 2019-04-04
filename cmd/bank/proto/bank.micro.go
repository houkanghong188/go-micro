// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: bank.proto

/*
Package bank is a generated protocol buffer package.

It is generated from these files:
	bank.proto

It has these top-level messages:
	Request
	Response
	History
*/
package bank

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

// Client API for Bank service

type BankService interface {
	Show(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type bankService struct {
	c    client.Client
	name string
}

func NewBankService(name string, c client.Client) BankService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "bank"
	}
	return &bankService{
		c:    c,
		name: name,
	}
}

func (c *bankService) Show(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Bank.Show", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bank service

type BankHandler interface {
	Show(context.Context, *Request, *Response) error
}

func RegisterBankHandler(s server.Server, hdlr BankHandler, opts ...server.HandlerOption) error {
	type bank interface {
		Show(ctx context.Context, in *Request, out *Response) error
	}
	type Bank struct {
		bank
	}
	h := &bankHandler{hdlr}
	return s.Handle(s.NewHandler(&Bank{h}, opts...))
}

type bankHandler struct {
	BankHandler
}

func (h *bankHandler) Show(ctx context.Context, in *Request, out *Response) error {
	return h.BankHandler.Show(ctx, in, out)
}
