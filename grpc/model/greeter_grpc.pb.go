// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	TestStruct(ctx context.Context, in *StructRequest, opts ...grpc.CallOption) (*StructResponse, error)
	// Sends another greeting
	TestString(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) TestStruct(ctx context.Context, in *StructRequest, opts ...grpc.CallOption) (*StructResponse, error) {
	out := new(StructResponse)
	err := c.cc.Invoke(ctx, "/com.dut.Greeter/TestStruct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) TestString(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/com.dut.Greeter/TestString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	TestStruct(context.Context, *StructRequest) (*StructResponse, error)
	// Sends another greeting
	TestString(context.Context, *StringRequest) (*StringResponse, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) TestStruct(context.Context, *StructRequest) (*StructResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestStruct not implemented")
}
func (UnimplementedGreeterServer) TestString(context.Context, *StringRequest) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestString not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_TestStruct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StructRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).TestStruct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.dut.Greeter/TestStruct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).TestStruct(ctx, req.(*StructRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_TestString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).TestString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.dut.Greeter/TestString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).TestString(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.dut.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestStruct",
			Handler:    _Greeter_TestStruct_Handler,
		},
		{
			MethodName: "TestString",
			Handler:    _Greeter_TestString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/greeter.proto",
}
