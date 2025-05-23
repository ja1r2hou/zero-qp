// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: gameService.proto

package game

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GameRPC_Ping_FullMethodName = "/game.GameRPC/Ping"
)

// GameRPCClient is the client API for GameRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameRPCClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type gameRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGameRPCClient(cc grpc.ClientConnInterface) GameRPCClient {
	return &gameRPCClient{cc}
}

func (c *gameRPCClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, GameRPC_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameRPCServer is the server API for GameRPC service.
// All implementations must embed UnimplementedGameRPCServer
// for forward compatibility.
type GameRPCServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGameRPCServer()
}

// UnimplementedGameRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGameRPCServer struct{}

func (UnimplementedGameRPCServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGameRPCServer) mustEmbedUnimplementedGameRPCServer() {}
func (UnimplementedGameRPCServer) testEmbeddedByValue()                 {}

// UnsafeGameRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameRPCServer will
// result in compilation errors.
type UnsafeGameRPCServer interface {
	mustEmbedUnimplementedGameRPCServer()
}

func RegisterGameRPCServer(s grpc.ServiceRegistrar, srv GameRPCServer) {
	// If the following call pancis, it indicates UnimplementedGameRPCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GameRPC_ServiceDesc, srv)
}

func _GameRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameRPC_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameRPCServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GameRPC_ServiceDesc is the grpc.ServiceDesc for GameRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game.GameRPC",
	HandlerType: (*GameRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GameRPC_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gameService.proto",
}
