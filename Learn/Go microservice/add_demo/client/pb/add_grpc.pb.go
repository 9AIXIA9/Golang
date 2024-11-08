// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.0--rc2
// source: pb/add.proto

package pb

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

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterClient interface {
	// 定义方法
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type counterClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterClient(cc grpc.ClientConnInterface) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/pb.Counter/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServer is the server API for Counter service.
// All implementations must embed UnimplementedCounterServer
// for forward compatibility
type CounterServer interface {
	// 定义方法
	Add(context.Context, *AddRequest) (*AddResponse, error)
	mustEmbedUnimplementedCounterServer()
}

// UnimplementedCounterServer must be embedded to have forward compatible implementations.
type UnimplementedCounterServer struct {
}

func (UnimplementedCounterServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCounterServer) mustEmbedUnimplementedCounterServer() {}

// UnsafeCounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServer will
// result in compilation errors.
type UnsafeCounterServer interface {
	mustEmbedUnimplementedCounterServer()
}

func RegisterCounterServer(s grpc.ServiceRegistrar, srv CounterServer) {
	s.RegisterService(&Counter_ServiceDesc, srv)
}

func _Counter_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Counter/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Counter_ServiceDesc is the grpc.ServiceDesc for Counter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Counter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Counter_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/add.proto",
}
