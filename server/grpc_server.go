// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: datamodel.proto

package server

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

// GuanyuClient is the client API for Guanyu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuanyuClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type guanyuClient struct {
	cc grpc.ClientConnInterface
}

func NewGuanyuClient(cc grpc.ClientConnInterface) GuanyuClient {
	return &guanyuClient{cc}
}

func (c *guanyuClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/server.Guanyu/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuanyuServer is the server API for Guanyu service.
// All implementations must embed UnimplementedGuanyuServer
// for forward compatibility
type GuanyuServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedGuanyuServer()
}

// UnimplementedGuanyuServer must be embedded to have forward compatible implementations.
type UnimplementedGuanyuServer struct {
}

func (UnimplementedGuanyuServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGuanyuServer) mustEmbedUnimplementedGuanyuServer() {}

// UnsafeGuanyuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuanyuServer will
// result in compilation errors.
type UnsafeGuanyuServer interface {
	mustEmbedUnimplementedGuanyuServer()
}

func RegisterGuanyuServer(s grpc.ServiceRegistrar, srv GuanyuServer) {
	s.RegisterService(&Guanyu_ServiceDesc, srv)
}

func _Guanyu_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuanyuServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Guanyu/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuanyuServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Guanyu_ServiceDesc is the grpc.ServiceDesc for Guanyu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Guanyu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Guanyu",
	HandlerType: (*GuanyuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Guanyu_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datamodel.proto",
}