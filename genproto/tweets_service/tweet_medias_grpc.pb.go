// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: tweet_medias.proto

package tweets_service

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
	TweetMediasService_CreateTweetMedia_FullMethodName            = "/tweets_service.TweetMediasService/CreateTweetMedia"
	TweetMediasService_GetTweetMedia_FullMethodName               = "/tweets_service.TweetMediasService/GetTweetMedia"
	TweetMediasService_GetListTweetMedia_FullMethodName           = "/tweets_service.TweetMediasService/GetListTweetMedia"
	TweetMediasService_UpdateTweetMedia_FullMethodName            = "/tweets_service.TweetMediasService/UpdateTweetMedia"
	TweetMediasService_DeleteTweetMedia_FullMethodName            = "/tweets_service.TweetMediasService/DeleteTweetMedia"
	TweetMediasService_DeleteTweetMediaWithTweetID_FullMethodName = "/tweets_service.TweetMediasService/DeleteTweetMediaWithTweetID"
)

// TweetMediasServiceClient is the client API for TweetMediasService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition for TweetMedia
type TweetMediasServiceClient interface {
	CreateTweetMedia(ctx context.Context, in *CreateTweetMediaRequest, opts ...grpc.CallOption) (*CreateTweetMediaResponse, error)
	GetTweetMedia(ctx context.Context, in *GetTweetMediaRequest, opts ...grpc.CallOption) (*TweetMedia, error)
	GetListTweetMedia(ctx context.Context, in *GetListTweetMediaRequest, opts ...grpc.CallOption) (*TweetMediaList, error)
	UpdateTweetMedia(ctx context.Context, in *UpdateTweetMediaRequest, opts ...grpc.CallOption) (*UpdateTweetMediaResponse, error)
	DeleteTweetMedia(ctx context.Context, in *DeleteTweetMediaRequest, opts ...grpc.CallOption) (*DeleteTweetMediaResponse, error)
	DeleteTweetMediaWithTweetID(ctx context.Context, in *DeleteTweetMediaWithTweetIDRequest, opts ...grpc.CallOption) (*DeleteTweetMediaWithTweetIDResponse, error)
}

type tweetMediasServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTweetMediasServiceClient(cc grpc.ClientConnInterface) TweetMediasServiceClient {
	return &tweetMediasServiceClient{cc}
}

func (c *tweetMediasServiceClient) CreateTweetMedia(ctx context.Context, in *CreateTweetMediaRequest, opts ...grpc.CallOption) (*CreateTweetMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTweetMediaResponse)
	err := c.cc.Invoke(ctx, TweetMediasService_CreateTweetMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetMediasServiceClient) GetTweetMedia(ctx context.Context, in *GetTweetMediaRequest, opts ...grpc.CallOption) (*TweetMedia, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TweetMedia)
	err := c.cc.Invoke(ctx, TweetMediasService_GetTweetMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetMediasServiceClient) GetListTweetMedia(ctx context.Context, in *GetListTweetMediaRequest, opts ...grpc.CallOption) (*TweetMediaList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TweetMediaList)
	err := c.cc.Invoke(ctx, TweetMediasService_GetListTweetMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetMediasServiceClient) UpdateTweetMedia(ctx context.Context, in *UpdateTweetMediaRequest, opts ...grpc.CallOption) (*UpdateTweetMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTweetMediaResponse)
	err := c.cc.Invoke(ctx, TweetMediasService_UpdateTweetMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetMediasServiceClient) DeleteTweetMedia(ctx context.Context, in *DeleteTweetMediaRequest, opts ...grpc.CallOption) (*DeleteTweetMediaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTweetMediaResponse)
	err := c.cc.Invoke(ctx, TweetMediasService_DeleteTweetMedia_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetMediasServiceClient) DeleteTweetMediaWithTweetID(ctx context.Context, in *DeleteTweetMediaWithTweetIDRequest, opts ...grpc.CallOption) (*DeleteTweetMediaWithTweetIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTweetMediaWithTweetIDResponse)
	err := c.cc.Invoke(ctx, TweetMediasService_DeleteTweetMediaWithTweetID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetMediasServiceServer is the server API for TweetMediasService service.
// All implementations must embed UnimplementedTweetMediasServiceServer
// for forward compatibility.
//
// Service definition for TweetMedia
type TweetMediasServiceServer interface {
	CreateTweetMedia(context.Context, *CreateTweetMediaRequest) (*CreateTweetMediaResponse, error)
	GetTweetMedia(context.Context, *GetTweetMediaRequest) (*TweetMedia, error)
	GetListTweetMedia(context.Context, *GetListTweetMediaRequest) (*TweetMediaList, error)
	UpdateTweetMedia(context.Context, *UpdateTweetMediaRequest) (*UpdateTweetMediaResponse, error)
	DeleteTweetMedia(context.Context, *DeleteTweetMediaRequest) (*DeleteTweetMediaResponse, error)
	DeleteTweetMediaWithTweetID(context.Context, *DeleteTweetMediaWithTweetIDRequest) (*DeleteTweetMediaWithTweetIDResponse, error)
	mustEmbedUnimplementedTweetMediasServiceServer()
}

// UnimplementedTweetMediasServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTweetMediasServiceServer struct{}

func (UnimplementedTweetMediasServiceServer) CreateTweetMedia(context.Context, *CreateTweetMediaRequest) (*CreateTweetMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTweetMedia not implemented")
}
func (UnimplementedTweetMediasServiceServer) GetTweetMedia(context.Context, *GetTweetMediaRequest) (*TweetMedia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweetMedia not implemented")
}
func (UnimplementedTweetMediasServiceServer) GetListTweetMedia(context.Context, *GetListTweetMediaRequest) (*TweetMediaList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTweetMedia not implemented")
}
func (UnimplementedTweetMediasServiceServer) UpdateTweetMedia(context.Context, *UpdateTweetMediaRequest) (*UpdateTweetMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTweetMedia not implemented")
}
func (UnimplementedTweetMediasServiceServer) DeleteTweetMedia(context.Context, *DeleteTweetMediaRequest) (*DeleteTweetMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTweetMedia not implemented")
}
func (UnimplementedTweetMediasServiceServer) DeleteTweetMediaWithTweetID(context.Context, *DeleteTweetMediaWithTweetIDRequest) (*DeleteTweetMediaWithTweetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTweetMediaWithTweetID not implemented")
}
func (UnimplementedTweetMediasServiceServer) mustEmbedUnimplementedTweetMediasServiceServer() {}
func (UnimplementedTweetMediasServiceServer) testEmbeddedByValue()                            {}

// UnsafeTweetMediasServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TweetMediasServiceServer will
// result in compilation errors.
type UnsafeTweetMediasServiceServer interface {
	mustEmbedUnimplementedTweetMediasServiceServer()
}

func RegisterTweetMediasServiceServer(s grpc.ServiceRegistrar, srv TweetMediasServiceServer) {
	// If the following call pancis, it indicates UnimplementedTweetMediasServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TweetMediasService_ServiceDesc, srv)
}

func _TweetMediasService_CreateTweetMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTweetMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetMediasServiceServer).CreateTweetMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetMediasService_CreateTweetMedia_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetMediasServiceServer).CreateTweetMedia(ctx, req.(*CreateTweetMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetMediasService_GetTweetMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetMediasServiceServer).GetTweetMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetMediasService_GetTweetMedia_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetMediasServiceServer).GetTweetMedia(ctx, req.(*GetTweetMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetMediasService_GetListTweetMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTweetMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetMediasServiceServer).GetListTweetMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetMediasService_GetListTweetMedia_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetMediasServiceServer).GetListTweetMedia(ctx, req.(*GetListTweetMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetMediasService_UpdateTweetMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTweetMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetMediasServiceServer).UpdateTweetMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetMediasService_UpdateTweetMedia_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetMediasServiceServer).UpdateTweetMedia(ctx, req.(*UpdateTweetMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetMediasService_DeleteTweetMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTweetMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetMediasServiceServer).DeleteTweetMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetMediasService_DeleteTweetMedia_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetMediasServiceServer).DeleteTweetMedia(ctx, req.(*DeleteTweetMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetMediasService_DeleteTweetMediaWithTweetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTweetMediaWithTweetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetMediasServiceServer).DeleteTweetMediaWithTweetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetMediasService_DeleteTweetMediaWithTweetID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetMediasServiceServer).DeleteTweetMediaWithTweetID(ctx, req.(*DeleteTweetMediaWithTweetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TweetMediasService_ServiceDesc is the grpc.ServiceDesc for TweetMediasService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TweetMediasService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tweets_service.TweetMediasService",
	HandlerType: (*TweetMediasServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTweetMedia",
			Handler:    _TweetMediasService_CreateTweetMedia_Handler,
		},
		{
			MethodName: "GetTweetMedia",
			Handler:    _TweetMediasService_GetTweetMedia_Handler,
		},
		{
			MethodName: "GetListTweetMedia",
			Handler:    _TweetMediasService_GetListTweetMedia_Handler,
		},
		{
			MethodName: "UpdateTweetMedia",
			Handler:    _TweetMediasService_UpdateTweetMedia_Handler,
		},
		{
			MethodName: "DeleteTweetMedia",
			Handler:    _TweetMediasService_DeleteTweetMedia_Handler,
		},
		{
			MethodName: "DeleteTweetMediaWithTweetID",
			Handler:    _TweetMediasService_DeleteTweetMediaWithTweetID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tweet_medias.proto",
}
