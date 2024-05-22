//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: wb/prices/v1/service.proto

package wbPrices

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

const (
	PricesService_ApiV2BufferGoodsTaskGet_FullMethodName  = "/wb.prices.v1.PricesService/ApiV2BufferGoodsTaskGet"
	PricesService_ApiV2BufferTasksGet_FullMethodName      = "/wb.prices.v1.PricesService/ApiV2BufferTasksGet"
	PricesService_ApiV2HistoryGoodsTaskGet_FullMethodName = "/wb.prices.v1.PricesService/ApiV2HistoryGoodsTaskGet"
	PricesService_ApiV2HistoryTasksGet_FullMethodName     = "/wb.prices.v1.PricesService/ApiV2HistoryTasksGet"
	PricesService_ApiV2ListGoodsFilterGet_FullMethodName  = "/wb.prices.v1.PricesService/ApiV2ListGoodsFilterGet"
	PricesService_ApiV2ListGoodsSizeNmGet_FullMethodName  = "/wb.prices.v1.PricesService/ApiV2ListGoodsSizeNmGet"
	PricesService_ApiV2UploadTaskPost_FullMethodName      = "/wb.prices.v1.PricesService/ApiV2UploadTaskPost"
	PricesService_ApiV2UploadTaskSizePost_FullMethodName  = "/wb.prices.v1.PricesService/ApiV2UploadTaskSizePost"
)

// PricesServiceClient is the client API for PricesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PricesServiceClient interface {
	ApiV2BufferGoodsTaskGet(ctx context.Context, in *ApiV2BufferGoodsTaskGetRequest, opts ...grpc.CallOption) (*ApiV2BufferGoodsTaskGet200Response, error)
	ApiV2BufferTasksGet(ctx context.Context, in *ApiV2BufferTasksGetRequest, opts ...grpc.CallOption) (*ApiV2BufferTasksGet200Response, error)
	ApiV2HistoryGoodsTaskGet(ctx context.Context, in *ApiV2HistoryGoodsTaskGetRequest, opts ...grpc.CallOption) (*ApiV2HistoryGoodsTaskGet200Response, error)
	ApiV2HistoryTasksGet(ctx context.Context, in *ApiV2HistoryTasksGetRequest, opts ...grpc.CallOption) (*ApiV2HistoryTasksGet200Response, error)
	ApiV2ListGoodsFilterGet(ctx context.Context, in *ApiV2ListGoodsFilterGetRequest, opts ...grpc.CallOption) (*ApiV2ListGoodsFilterGet200Response, error)
	ApiV2ListGoodsSizeNmGet(ctx context.Context, in *ApiV2ListGoodsSizeNmGetRequest, opts ...grpc.CallOption) (*ApiV2ListGoodsSizeNmGet200Response, error)
	ApiV2UploadTaskPost(ctx context.Context, in *ApiV2UploadTaskPostRequest, opts ...grpc.CallOption) (*TaskCreated, error)
	ApiV2UploadTaskSizePost(ctx context.Context, in *ApiV2UploadTaskSizePostRequest, opts ...grpc.CallOption) (*TaskCreated, error)
}

type pricesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPricesServiceClient(cc grpc.ClientConnInterface) PricesServiceClient {
	return &pricesServiceClient{cc}
}

func (c *pricesServiceClient) ApiV2BufferGoodsTaskGet(ctx context.Context, in *ApiV2BufferGoodsTaskGetRequest, opts ...grpc.CallOption) (*ApiV2BufferGoodsTaskGet200Response, error) {
	out := new(ApiV2BufferGoodsTaskGet200Response)
	err := c.cc.Invoke(ctx, PricesService_ApiV2BufferGoodsTaskGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricesServiceClient) ApiV2BufferTasksGet(ctx context.Context, in *ApiV2BufferTasksGetRequest, opts ...grpc.CallOption) (*ApiV2BufferTasksGet200Response, error) {
	out := new(ApiV2BufferTasksGet200Response)
	err := c.cc.Invoke(ctx, PricesService_ApiV2BufferTasksGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricesServiceClient) ApiV2HistoryGoodsTaskGet(ctx context.Context, in *ApiV2HistoryGoodsTaskGetRequest, opts ...grpc.CallOption) (*ApiV2HistoryGoodsTaskGet200Response, error) {
	out := new(ApiV2HistoryGoodsTaskGet200Response)
	err := c.cc.Invoke(ctx, PricesService_ApiV2HistoryGoodsTaskGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricesServiceClient) ApiV2HistoryTasksGet(ctx context.Context, in *ApiV2HistoryTasksGetRequest, opts ...grpc.CallOption) (*ApiV2HistoryTasksGet200Response, error) {
	out := new(ApiV2HistoryTasksGet200Response)
	err := c.cc.Invoke(ctx, PricesService_ApiV2HistoryTasksGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricesServiceClient) ApiV2ListGoodsFilterGet(ctx context.Context, in *ApiV2ListGoodsFilterGetRequest, opts ...grpc.CallOption) (*ApiV2ListGoodsFilterGet200Response, error) {
	out := new(ApiV2ListGoodsFilterGet200Response)
	err := c.cc.Invoke(ctx, PricesService_ApiV2ListGoodsFilterGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricesServiceClient) ApiV2ListGoodsSizeNmGet(ctx context.Context, in *ApiV2ListGoodsSizeNmGetRequest, opts ...grpc.CallOption) (*ApiV2ListGoodsSizeNmGet200Response, error) {
	out := new(ApiV2ListGoodsSizeNmGet200Response)
	err := c.cc.Invoke(ctx, PricesService_ApiV2ListGoodsSizeNmGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricesServiceClient) ApiV2UploadTaskPost(ctx context.Context, in *ApiV2UploadTaskPostRequest, opts ...grpc.CallOption) (*TaskCreated, error) {
	out := new(TaskCreated)
	err := c.cc.Invoke(ctx, PricesService_ApiV2UploadTaskPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricesServiceClient) ApiV2UploadTaskSizePost(ctx context.Context, in *ApiV2UploadTaskSizePostRequest, opts ...grpc.CallOption) (*TaskCreated, error) {
	out := new(TaskCreated)
	err := c.cc.Invoke(ctx, PricesService_ApiV2UploadTaskSizePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricesServiceServer is the server API for PricesService service.
// All implementations must embed UnimplementedPricesServiceServer
// for forward compatibility
type PricesServiceServer interface {
	ApiV2BufferGoodsTaskGet(context.Context, *ApiV2BufferGoodsTaskGetRequest) (*ApiV2BufferGoodsTaskGet200Response, error)
	ApiV2BufferTasksGet(context.Context, *ApiV2BufferTasksGetRequest) (*ApiV2BufferTasksGet200Response, error)
	ApiV2HistoryGoodsTaskGet(context.Context, *ApiV2HistoryGoodsTaskGetRequest) (*ApiV2HistoryGoodsTaskGet200Response, error)
	ApiV2HistoryTasksGet(context.Context, *ApiV2HistoryTasksGetRequest) (*ApiV2HistoryTasksGet200Response, error)
	ApiV2ListGoodsFilterGet(context.Context, *ApiV2ListGoodsFilterGetRequest) (*ApiV2ListGoodsFilterGet200Response, error)
	ApiV2ListGoodsSizeNmGet(context.Context, *ApiV2ListGoodsSizeNmGetRequest) (*ApiV2ListGoodsSizeNmGet200Response, error)
	ApiV2UploadTaskPost(context.Context, *ApiV2UploadTaskPostRequest) (*TaskCreated, error)
	ApiV2UploadTaskSizePost(context.Context, *ApiV2UploadTaskSizePostRequest) (*TaskCreated, error)
	mustEmbedUnimplementedPricesServiceServer()
}

// UnimplementedPricesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPricesServiceServer struct {
}

func (UnimplementedPricesServiceServer) ApiV2BufferGoodsTaskGet(context.Context, *ApiV2BufferGoodsTaskGetRequest) (*ApiV2BufferGoodsTaskGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2BufferGoodsTaskGet not implemented")
}
func (UnimplementedPricesServiceServer) ApiV2BufferTasksGet(context.Context, *ApiV2BufferTasksGetRequest) (*ApiV2BufferTasksGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2BufferTasksGet not implemented")
}
func (UnimplementedPricesServiceServer) ApiV2HistoryGoodsTaskGet(context.Context, *ApiV2HistoryGoodsTaskGetRequest) (*ApiV2HistoryGoodsTaskGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2HistoryGoodsTaskGet not implemented")
}
func (UnimplementedPricesServiceServer) ApiV2HistoryTasksGet(context.Context, *ApiV2HistoryTasksGetRequest) (*ApiV2HistoryTasksGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2HistoryTasksGet not implemented")
}
func (UnimplementedPricesServiceServer) ApiV2ListGoodsFilterGet(context.Context, *ApiV2ListGoodsFilterGetRequest) (*ApiV2ListGoodsFilterGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2ListGoodsFilterGet not implemented")
}
func (UnimplementedPricesServiceServer) ApiV2ListGoodsSizeNmGet(context.Context, *ApiV2ListGoodsSizeNmGetRequest) (*ApiV2ListGoodsSizeNmGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2ListGoodsSizeNmGet not implemented")
}
func (UnimplementedPricesServiceServer) ApiV2UploadTaskPost(context.Context, *ApiV2UploadTaskPostRequest) (*TaskCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2UploadTaskPost not implemented")
}
func (UnimplementedPricesServiceServer) ApiV2UploadTaskSizePost(context.Context, *ApiV2UploadTaskSizePostRequest) (*TaskCreated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2UploadTaskSizePost not implemented")
}
func (UnimplementedPricesServiceServer) mustEmbedUnimplementedPricesServiceServer() {}

// UnsafePricesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PricesServiceServer will
// result in compilation errors.
type UnsafePricesServiceServer interface {
	mustEmbedUnimplementedPricesServiceServer()
}

func RegisterPricesServiceServer(s grpc.ServiceRegistrar, srv PricesServiceServer) {
	s.RegisterService(&PricesService_ServiceDesc, srv)
}

func _PricesService_ApiV2BufferGoodsTaskGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2BufferGoodsTaskGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2BufferGoodsTaskGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2BufferGoodsTaskGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2BufferGoodsTaskGet(ctx, req.(*ApiV2BufferGoodsTaskGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricesService_ApiV2BufferTasksGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2BufferTasksGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2BufferTasksGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2BufferTasksGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2BufferTasksGet(ctx, req.(*ApiV2BufferTasksGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricesService_ApiV2HistoryGoodsTaskGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2HistoryGoodsTaskGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2HistoryGoodsTaskGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2HistoryGoodsTaskGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2HistoryGoodsTaskGet(ctx, req.(*ApiV2HistoryGoodsTaskGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricesService_ApiV2HistoryTasksGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2HistoryTasksGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2HistoryTasksGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2HistoryTasksGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2HistoryTasksGet(ctx, req.(*ApiV2HistoryTasksGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricesService_ApiV2ListGoodsFilterGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2ListGoodsFilterGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2ListGoodsFilterGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2ListGoodsFilterGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2ListGoodsFilterGet(ctx, req.(*ApiV2ListGoodsFilterGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricesService_ApiV2ListGoodsSizeNmGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2ListGoodsSizeNmGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2ListGoodsSizeNmGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2ListGoodsSizeNmGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2ListGoodsSizeNmGet(ctx, req.(*ApiV2ListGoodsSizeNmGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricesService_ApiV2UploadTaskPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2UploadTaskPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2UploadTaskPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2UploadTaskPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2UploadTaskPost(ctx, req.(*ApiV2UploadTaskPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricesService_ApiV2UploadTaskSizePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2UploadTaskSizePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricesServiceServer).ApiV2UploadTaskSizePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricesService_ApiV2UploadTaskSizePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricesServiceServer).ApiV2UploadTaskSizePost(ctx, req.(*ApiV2UploadTaskSizePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PricesService_ServiceDesc is the grpc.ServiceDesc for PricesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PricesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wb.prices.v1.PricesService",
	HandlerType: (*PricesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApiV2BufferGoodsTaskGet",
			Handler:    _PricesService_ApiV2BufferGoodsTaskGet_Handler,
		},
		{
			MethodName: "ApiV2BufferTasksGet",
			Handler:    _PricesService_ApiV2BufferTasksGet_Handler,
		},
		{
			MethodName: "ApiV2HistoryGoodsTaskGet",
			Handler:    _PricesService_ApiV2HistoryGoodsTaskGet_Handler,
		},
		{
			MethodName: "ApiV2HistoryTasksGet",
			Handler:    _PricesService_ApiV2HistoryTasksGet_Handler,
		},
		{
			MethodName: "ApiV2ListGoodsFilterGet",
			Handler:    _PricesService_ApiV2ListGoodsFilterGet_Handler,
		},
		{
			MethodName: "ApiV2ListGoodsSizeNmGet",
			Handler:    _PricesService_ApiV2ListGoodsSizeNmGet_Handler,
		},
		{
			MethodName: "ApiV2UploadTaskPost",
			Handler:    _PricesService_ApiV2UploadTaskPost_Handler,
		},
		{
			MethodName: "ApiV2UploadTaskSizePost",
			Handler:    _PricesService_ApiV2UploadTaskSizePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wb/prices/v1/service.proto",
}