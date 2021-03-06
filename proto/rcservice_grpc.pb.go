// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: rcservice.proto

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

// RcServiceClient is the client API for RcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RcServiceClient interface {
	Recommend(ctx context.Context, opts ...grpc.CallOption) (RcService_RecommendClient, error)
	RecommendByMany(ctx context.Context, in *RecommendManyRequest, opts ...grpc.CallOption) (*RecommendManyResponse, error)
}

type rcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRcServiceClient(cc grpc.ClientConnInterface) RcServiceClient {
	return &rcServiceClient{cc}
}

func (c *rcServiceClient) Recommend(ctx context.Context, opts ...grpc.CallOption) (RcService_RecommendClient, error) {
	stream, err := c.cc.NewStream(ctx, &RcService_ServiceDesc.Streams[0], "/rcserver.RcService/Recommend", opts...)
	if err != nil {
		return nil, err
	}
	x := &rcServiceRecommendClient{stream}
	return x, nil
}

type RcService_RecommendClient interface {
	Send(*RecommendRequest) error
	Recv() (*RecommendResponse, error)
	grpc.ClientStream
}

type rcServiceRecommendClient struct {
	grpc.ClientStream
}

func (x *rcServiceRecommendClient) Send(m *RecommendRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rcServiceRecommendClient) Recv() (*RecommendResponse, error) {
	m := new(RecommendResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rcServiceClient) RecommendByMany(ctx context.Context, in *RecommendManyRequest, opts ...grpc.CallOption) (*RecommendManyResponse, error) {
	out := new(RecommendManyResponse)
	err := c.cc.Invoke(ctx, "/rcserver.RcService/RecommendByMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RcServiceServer is the server API for RcService service.
// All implementations must embed UnimplementedRcServiceServer
// for forward compatibility
type RcServiceServer interface {
	Recommend(RcService_RecommendServer) error
	RecommendByMany(context.Context, *RecommendManyRequest) (*RecommendManyResponse, error)
	mustEmbedUnimplementedRcServiceServer()
}

// UnimplementedRcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRcServiceServer struct {
}

func (UnimplementedRcServiceServer) Recommend(RcService_RecommendServer) error {
	return status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}
func (UnimplementedRcServiceServer) RecommendByMany(context.Context, *RecommendManyRequest) (*RecommendManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendByMany not implemented")
}
func (UnimplementedRcServiceServer) mustEmbedUnimplementedRcServiceServer() {}

// UnsafeRcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RcServiceServer will
// result in compilation errors.
type UnsafeRcServiceServer interface {
	mustEmbedUnimplementedRcServiceServer()
}

func RegisterRcServiceServer(s grpc.ServiceRegistrar, srv RcServiceServer) {
	s.RegisterService(&RcService_ServiceDesc, srv)
}

func _RcService_Recommend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RcServiceServer).Recommend(&rcServiceRecommendServer{stream})
}

type RcService_RecommendServer interface {
	Send(*RecommendResponse) error
	Recv() (*RecommendRequest, error)
	grpc.ServerStream
}

type rcServiceRecommendServer struct {
	grpc.ServerStream
}

func (x *rcServiceRecommendServer) Send(m *RecommendResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rcServiceRecommendServer) Recv() (*RecommendRequest, error) {
	m := new(RecommendRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RcService_RecommendByMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RcServiceServer).RecommendByMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcserver.RcService/RecommendByMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RcServiceServer).RecommendByMany(ctx, req.(*RecommendManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RcService_ServiceDesc is the grpc.ServiceDesc for RcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rcserver.RcService",
	HandlerType: (*RcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecommendByMany",
			Handler:    _RcService_RecommendByMany_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Recommend",
			Handler:       _RcService_Recommend_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rcservice.proto",
}
