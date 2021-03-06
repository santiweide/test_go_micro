// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/greeter.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterService interface {
	// Sends a greeting
	TestStruct(ctx context.Context, in *StructRequest, opts ...client.CallOption) (*StructResponse, error)
	// Sends another greeting
	TestString(ctx context.Context, in *StringRequest, opts ...client.CallOption) (*StringResponse, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "com.dut"
	}
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) TestStruct(ctx context.Context, in *StructRequest, opts ...client.CallOption) (*StructResponse, error) {
	req := c.c.NewRequest(c.name, "Greeter.TestStruct", in)
	out := new(StructResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterService) TestString(ctx context.Context, in *StringRequest, opts ...client.CallOption) (*StringResponse, error) {
	req := c.c.NewRequest(c.name, "Greeter.TestString", in)
	out := new(StringResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	// Sends a greeting
	TestStruct(context.Context, *StructRequest, *StructResponse) error
	// Sends another greeting
	TestString(context.Context, *StringRequest, *StringResponse) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		TestStruct(ctx context.Context, in *StructRequest, out *StructResponse) error
		TestString(ctx context.Context, in *StringRequest, out *StringResponse) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) TestStruct(ctx context.Context, in *StructRequest, out *StructResponse) error {
	return h.GreeterHandler.TestStruct(ctx, in, out)
}

func (h *greeterHandler) TestString(ctx context.Context, in *StringRequest, out *StringResponse) error {
	return h.GreeterHandler.TestString(ctx, in, out)
}
