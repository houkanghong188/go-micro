// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: witnesses.proto

/*
Package works_witnesses is a generated protocol buffer package.

It is generated from these files:
	witnesses.proto

It has these top-level messages:
	WitnessesRequest
	WitnessesResponse
	Witness
*/
package works_witnesses

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

// Client API for Witnesses service

type WitnessesService interface {
	Index(ctx context.Context, in *WitnessesRequest, opts ...client.CallOption) (*WitnessesResponse, error)
}

type witnessesService struct {
	c    client.Client
	name string
}

func NewWitnessesService(name string, c client.Client) WitnessesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "works.witnesses"
	}
	return &witnessesService{
		c:    c,
		name: name,
	}
}

func (c *witnessesService) Index(ctx context.Context, in *WitnessesRequest, opts ...client.CallOption) (*WitnessesResponse, error) {
	req := c.c.NewRequest(c.name, "Witnesses.Index", in)
	out := new(WitnessesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Witnesses service

type WitnessesHandler interface {
	Index(context.Context, *WitnessesRequest, *WitnessesResponse) error
}

func RegisterWitnessesHandler(s server.Server, hdlr WitnessesHandler, opts ...server.HandlerOption) error {
	type witnesses interface {
		Index(ctx context.Context, in *WitnessesRequest, out *WitnessesResponse) error
	}
	type Witnesses struct {
		witnesses
	}
	h := &witnessesHandler{hdlr}
	return s.Handle(s.NewHandler(&Witnesses{h}, opts...))
}

type witnessesHandler struct {
	WitnessesHandler
}

func (h *witnessesHandler) Index(ctx context.Context, in *WitnessesRequest, out *WitnessesResponse) error {
	return h.WitnessesHandler.Index(ctx, in, out)
}