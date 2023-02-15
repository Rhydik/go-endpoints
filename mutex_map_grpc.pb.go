// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: mutex_map.proto

package main

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

// MutexMapClient is the client API for MutexMap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutexMapClient interface {
	// Sends a greeting
	GetMutexMap(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ValueReply, error)
	SetMutexMap(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Empty, error)
}

type mutexMapClient struct {
	cc grpc.ClientConnInterface
}

func NewMutexMapClient(cc grpc.ClientConnInterface) MutexMapClient {
	return &mutexMapClient{cc}
}

func (c *mutexMapClient) GetMutexMap(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ValueReply, error) {
	out := new(ValueReply)
	err := c.cc.Invoke(ctx, "/internal.MutexMap/GetMutexMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutexMapClient) SetMutexMap(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/internal.MutexMap/SetMutexMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutexMapServer is the server API for MutexMap service.
// All implementations must embed UnimplementedMutexMapServer
// for forward compatibility
type MutexMapServer interface {
	// Sends a greeting
	GetMutexMap(context.Context, *KeyRequest) (*ValueReply, error)
	SetMutexMap(context.Context, *KeyRequest) (*Empty, error)
	mustEmbedUnimplementedMutexMapServer()
}

// UnimplementedMutexMapServer must be embedded to have forward compatible implementations.
type UnimplementedMutexMapServer struct {
}

func (UnimplementedMutexMapServer) GetMutexMap(context.Context, *KeyRequest) (*ValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMutexMap not implemented")
}
func (UnimplementedMutexMapServer) SetMutexMap(context.Context, *KeyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMutexMap not implemented")
}
func (UnimplementedMutexMapServer) mustEmbedUnimplementedMutexMapServer() {}

// UnsafeMutexMapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutexMapServer will
// result in compilation errors.
type UnsafeMutexMapServer interface {
	mustEmbedUnimplementedMutexMapServer()
}

func RegisterMutexMapServer(s grpc.ServiceRegistrar, srv MutexMapServer) {
	s.RegisterService(&MutexMap_ServiceDesc, srv)
}

func _MutexMap_GetMutexMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexMapServer).GetMutexMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.MutexMap/GetMutexMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexMapServer).GetMutexMap(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutexMap_SetMutexMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexMapServer).SetMutexMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.MutexMap/SetMutexMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexMapServer).SetMutexMap(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MutexMap_ServiceDesc is the grpc.ServiceDesc for MutexMap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MutexMap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.MutexMap",
	HandlerType: (*MutexMapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMutexMap",
			Handler:    _MutexMap_GetMutexMap_Handler,
		},
		{
			MethodName: "SetMutexMap",
			Handler:    _MutexMap_SetMutexMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mutex_map.proto",
}
