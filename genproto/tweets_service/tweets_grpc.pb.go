// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: tweets.proto

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
	TweetsService_CreateTweet_FullMethodName  = "/tweets_service.TweetsService/CreateTweet"
	TweetsService_GetTweet_FullMethodName     = "/tweets_service.TweetsService/GetTweet"
	TweetsService_GetListTweet_FullMethodName = "/tweets_service.TweetsService/GetListTweet"
	TweetsService_UpdateTweet_FullMethodName  = "/tweets_service.TweetsService/UpdateTweet"
	TweetsService_DeleteTweet_FullMethodName  = "/tweets_service.TweetsService/DeleteTweet"
)

// TweetsServiceClient is the client API for TweetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition for Tweet
type TweetsServiceClient interface {
	CreateTweet(ctx context.Context, in *CreateTweetRequest, opts ...grpc.CallOption) (*CreateTweetResponse, error)
	GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*Tweet, error)
	GetListTweet(ctx context.Context, in *GetListTweetRequest, opts ...grpc.CallOption) (*TweetList, error)
	UpdateTweet(ctx context.Context, in *UpdateTweetRequest, opts ...grpc.CallOption) (*UpdateTweetResponse, error)
	DeleteTweet(ctx context.Context, in *DeleteTweetRequest, opts ...grpc.CallOption) (*DeleteTweetResponse, error)
}

type tweetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTweetsServiceClient(cc grpc.ClientConnInterface) TweetsServiceClient {
	return &tweetsServiceClient{cc}
}

func (c *tweetsServiceClient) CreateTweet(ctx context.Context, in *CreateTweetRequest, opts ...grpc.CallOption) (*CreateTweetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTweetResponse)
	err := c.cc.Invoke(ctx, TweetsService_CreateTweet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*Tweet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tweet)
	err := c.cc.Invoke(ctx, TweetsService_GetTweet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) GetListTweet(ctx context.Context, in *GetListTweetRequest, opts ...grpc.CallOption) (*TweetList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TweetList)
	err := c.cc.Invoke(ctx, TweetsService_GetListTweet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) UpdateTweet(ctx context.Context, in *UpdateTweetRequest, opts ...grpc.CallOption) (*UpdateTweetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTweetResponse)
	err := c.cc.Invoke(ctx, TweetsService_UpdateTweet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) DeleteTweet(ctx context.Context, in *DeleteTweetRequest, opts ...grpc.CallOption) (*DeleteTweetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTweetResponse)
	err := c.cc.Invoke(ctx, TweetsService_DeleteTweet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetsServiceServer is the server API for TweetsService service.
// All implementations must embed UnimplementedTweetsServiceServer
// for forward compatibility.
//
// Service definition for Tweet
type TweetsServiceServer interface {
	CreateTweet(context.Context, *CreateTweetRequest) (*CreateTweetResponse, error)
	GetTweet(context.Context, *GetTweetRequest) (*Tweet, error)
	GetListTweet(context.Context, *GetListTweetRequest) (*TweetList, error)
	UpdateTweet(context.Context, *UpdateTweetRequest) (*UpdateTweetResponse, error)
	DeleteTweet(context.Context, *DeleteTweetRequest) (*DeleteTweetResponse, error)
	mustEmbedUnimplementedTweetsServiceServer()
}

// UnimplementedTweetsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTweetsServiceServer struct{}

func (UnimplementedTweetsServiceServer) CreateTweet(context.Context, *CreateTweetRequest) (*CreateTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTweet not implemented")
}
func (UnimplementedTweetsServiceServer) GetTweet(context.Context, *GetTweetRequest) (*Tweet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweet not implemented")
}
func (UnimplementedTweetsServiceServer) GetListTweet(context.Context, *GetListTweetRequest) (*TweetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListTweet not implemented")
}
func (UnimplementedTweetsServiceServer) UpdateTweet(context.Context, *UpdateTweetRequest) (*UpdateTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTweet not implemented")
}
func (UnimplementedTweetsServiceServer) DeleteTweet(context.Context, *DeleteTweetRequest) (*DeleteTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTweet not implemented")
}
func (UnimplementedTweetsServiceServer) mustEmbedUnimplementedTweetsServiceServer() {}
func (UnimplementedTweetsServiceServer) testEmbeddedByValue()                       {}

// UnsafeTweetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TweetsServiceServer will
// result in compilation errors.
type UnsafeTweetsServiceServer interface {
	mustEmbedUnimplementedTweetsServiceServer()
}

func RegisterTweetsServiceServer(s grpc.ServiceRegistrar, srv TweetsServiceServer) {
	// If the following call pancis, it indicates UnimplementedTweetsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TweetsService_ServiceDesc, srv)
}

func _TweetsService_CreateTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).CreateTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetsService_CreateTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).CreateTweet(ctx, req.(*CreateTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_GetTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).GetTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetsService_GetTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).GetTweet(ctx, req.(*GetTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_GetListTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).GetListTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetsService_GetListTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).GetListTweet(ctx, req.(*GetListTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_UpdateTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).UpdateTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetsService_UpdateTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).UpdateTweet(ctx, req.(*UpdateTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_DeleteTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).DeleteTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetsService_DeleteTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).DeleteTweet(ctx, req.(*DeleteTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TweetsService_ServiceDesc is the grpc.ServiceDesc for TweetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TweetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tweets_service.TweetsService",
	HandlerType: (*TweetsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTweet",
			Handler:    _TweetsService_CreateTweet_Handler,
		},
		{
			MethodName: "GetTweet",
			Handler:    _TweetsService_GetTweet_Handler,
		},
		{
			MethodName: "GetListTweet",
			Handler:    _TweetsService_GetListTweet_Handler,
		},
		{
			MethodName: "UpdateTweet",
			Handler:    _TweetsService_UpdateTweet_Handler,
		},
		{
			MethodName: "DeleteTweet",
			Handler:    _TweetsService_DeleteTweet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tweets.proto",
}
