// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ExchangerClient is the client API for Exchanger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangerClient interface {
	// Exchange exchanges a amount of money from one currency to another.
	Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error)
}

type exchangerClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangerClient(cc grpc.ClientConnInterface) ExchangerClient {
	return &exchangerClient{cc}
}

func (c *exchangerClient) Exchange(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error) {
	out := new(ExchangeResponse)
	err := c.cc.Invoke(ctx, "/proto.Exchanger/Exchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangerServer is the server API for Exchanger service.
// All implementations must embed UnimplementedExchangerServer
// for forward compatibility
type ExchangerServer interface {
	// Exchange exchanges a amount of money from one currency to another.
	Exchange(context.Context, *ExchangeRequest) (*ExchangeResponse, error)
	mustEmbedUnimplementedExchangerServer()
}

// UnimplementedExchangerServer must be embedded to have forward compatible implementations.
type UnimplementedExchangerServer struct {
}

func (UnimplementedExchangerServer) Exchange(context.Context, *ExchangeRequest) (*ExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}
func (UnimplementedExchangerServer) mustEmbedUnimplementedExchangerServer() {}

// UnsafeExchangerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangerServer will
// result in compilation errors.
type UnsafeExchangerServer interface {
	mustEmbedUnimplementedExchangerServer()
}

func RegisterExchangerServer(s grpc.ServiceRegistrar, srv ExchangerServer) {
	s.RegisterService(&Exchanger_ServiceDesc, srv)
}

func _Exchanger_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangerServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Exchanger/Exchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangerServer).Exchange(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Exchanger_ServiceDesc is the grpc.ServiceDesc for Exchanger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exchanger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Exchanger",
	HandlerType: (*ExchangerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exchange",
			Handler:    _Exchanger_Exchange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}
