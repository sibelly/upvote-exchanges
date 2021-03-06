// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: exchanges.proto

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

// ExchangesServiceClient is the client API for ExchangesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangesServiceClient interface {
	Upvote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	ListExchanges(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ExchangesService_ListExchangesClient, error)
	ReadExchange(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadRes, error)
}

type exchangesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangesServiceClient(cc grpc.ClientConnInterface) ExchangesServiceClient {
	return &exchangesServiceClient{cc}
}

func (c *exchangesServiceClient) Upvote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, "/proto.ExchangesService/Upvote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangesServiceClient) ListExchanges(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ExchangesService_ListExchangesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExchangesService_ServiceDesc.Streams[0], "/proto.ExchangesService/ListExchanges", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangesServiceListExchangesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExchangesService_ListExchangesClient interface {
	Recv() (*Exchange, error)
	grpc.ClientStream
}

type exchangesServiceListExchangesClient struct {
	grpc.ClientStream
}

func (x *exchangesServiceListExchangesClient) Recv() (*Exchange, error) {
	m := new(Exchange)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exchangesServiceClient) ReadExchange(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (*ReadRes, error) {
	out := new(ReadRes)
	err := c.cc.Invoke(ctx, "/proto.ExchangesService/ReadExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangesServiceServer is the server API for ExchangesService service.
// All implementations must embed UnimplementedExchangesServiceServer
// for forward compatibility
type ExchangesServiceServer interface {
	Upvote(context.Context, *VoteRequest) (*VoteResponse, error)
	ListExchanges(*Empty, ExchangesService_ListExchangesServer) error
	ReadExchange(context.Context, *ReadReq) (*ReadRes, error)
	mustEmbedUnimplementedExchangesServiceServer()
}

// UnimplementedExchangesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExchangesServiceServer struct {
}

func (UnimplementedExchangesServiceServer) Upvote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upvote not implemented")
}
func (UnimplementedExchangesServiceServer) ListExchanges(*Empty, ExchangesService_ListExchangesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListExchanges not implemented")
}
func (UnimplementedExchangesServiceServer) ReadExchange(context.Context, *ReadReq) (*ReadRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadExchange not implemented")
}
func (UnimplementedExchangesServiceServer) mustEmbedUnimplementedExchangesServiceServer() {}

// UnsafeExchangesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangesServiceServer will
// result in compilation errors.
type UnsafeExchangesServiceServer interface {
	mustEmbedUnimplementedExchangesServiceServer()
}

func RegisterExchangesServiceServer(s grpc.ServiceRegistrar, srv ExchangesServiceServer) {
	s.RegisterService(&ExchangesService_ServiceDesc, srv)
}

func _ExchangesService_Upvote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangesServiceServer).Upvote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ExchangesService/Upvote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangesServiceServer).Upvote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangesService_ListExchanges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangesServiceServer).ListExchanges(m, &exchangesServiceListExchangesServer{stream})
}

type ExchangesService_ListExchangesServer interface {
	Send(*Exchange) error
	grpc.ServerStream
}

type exchangesServiceListExchangesServer struct {
	grpc.ServerStream
}

func (x *exchangesServiceListExchangesServer) Send(m *Exchange) error {
	return x.ServerStream.SendMsg(m)
}

func _ExchangesService_ReadExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangesServiceServer).ReadExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ExchangesService/ReadExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangesServiceServer).ReadExchange(ctx, req.(*ReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ExchangesService_ServiceDesc is the grpc.ServiceDesc for ExchangesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ExchangesService",
	HandlerType: (*ExchangesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upvote",
			Handler:    _ExchangesService_Upvote_Handler,
		},
		{
			MethodName: "ReadExchange",
			Handler:    _ExchangesService_ReadExchange_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListExchanges",
			Handler:       _ExchangesService_ListExchanges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "exchanges.proto",
}
