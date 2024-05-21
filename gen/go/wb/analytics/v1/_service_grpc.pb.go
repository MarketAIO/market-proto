//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam).
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: wb/analytics/v1/_service.proto

package wbANALYTICS

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
	DefaultService_ApiV1AnalyticsAcceptanceReportGet_FullMethodName       = "/wb.analytics.v1.DefaultService/ApiV1AnalyticsAcceptanceReportGet"
	DefaultService_ApiV1AnalyticsAntifraudDetailsGet_FullMethodName       = "/wb.analytics.v1.DefaultService/ApiV1AnalyticsAntifraudDetailsGet"
	DefaultService_ApiV1AnalyticsExciseReportPost_FullMethodName          = "/wb.analytics.v1.DefaultService/ApiV1AnalyticsExciseReportPost"
	DefaultService_ApiV1AnalyticsGoodsLabelingGet_FullMethodName          = "/wb.analytics.v1.DefaultService/ApiV1AnalyticsGoodsLabelingGet"
	DefaultService_ApiV1AnalyticsIncorrectAttachmentsGet_FullMethodName   = "/wb.analytics.v1.DefaultService/ApiV1AnalyticsIncorrectAttachmentsGet"
	DefaultService_ApiV1AnalyticsStorageCoefficientGet_FullMethodName     = "/wb.analytics.v1.DefaultService/ApiV1AnalyticsStorageCoefficientGet"
	DefaultService_ApiV1PaidStorageGet_FullMethodName                     = "/wb.analytics.v1.DefaultService/ApiV1PaidStorageGet"
	DefaultService_ApiV1PaidStorageTasksTaskIdDownloadGet_FullMethodName  = "/wb.analytics.v1.DefaultService/ApiV1PaidStorageTasksTaskIdDownloadGet"
	DefaultService_ApiV1PaidStorageTasksTaskIdStatusGet_FullMethodName    = "/wb.analytics.v1.DefaultService/ApiV1PaidStorageTasksTaskIdStatusGet"
	DefaultService_ApiV2NmReportDetailHistoryPost_FullMethodName          = "/wb.analytics.v1.DefaultService/ApiV2NmReportDetailHistoryPost"
	DefaultService_ApiV2NmReportDetailPost_FullMethodName                 = "/wb.analytics.v1.DefaultService/ApiV2NmReportDetailPost"
	DefaultService_ApiV2NmReportDownloadsFileDownloadIdGet_FullMethodName = "/wb.analytics.v1.DefaultService/ApiV2NmReportDownloadsFileDownloadIdGet"
	DefaultService_ApiV2NmReportDownloadsGet_FullMethodName               = "/wb.analytics.v1.DefaultService/ApiV2NmReportDownloadsGet"
	DefaultService_ApiV2NmReportDownloadsPost_FullMethodName              = "/wb.analytics.v1.DefaultService/ApiV2NmReportDownloadsPost"
	DefaultService_ApiV2NmReportDownloadsRetryPost_FullMethodName         = "/wb.analytics.v1.DefaultService/ApiV2NmReportDownloadsRetryPost"
	DefaultService_ApiV2NmReportGroupedHistoryPost_FullMethodName         = "/wb.analytics.v1.DefaultService/ApiV2NmReportGroupedHistoryPost"
	DefaultService_ApiV2NmReportGroupedPost_FullMethodName                = "/wb.analytics.v1.DefaultService/ApiV2NmReportGroupedPost"
)

// DefaultServiceClient is the client API for DefaultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DefaultServiceClient interface {
	ApiV1AnalyticsAcceptanceReportGet(ctx context.Context, in *ApiV1AnalyticsAcceptanceReportGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsAcceptanceReportGet200Response, error)
	ApiV1AnalyticsAntifraudDetailsGet(ctx context.Context, in *ApiV1AnalyticsAntifraudDetailsGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsAntifraudDetailsGet200Response, error)
	ApiV1AnalyticsExciseReportPost(ctx context.Context, in *ApiV1AnalyticsExciseReportPostRequest, opts ...grpc.CallOption) (*ExciseReportResponse, error)
	ApiV1AnalyticsGoodsLabelingGet(ctx context.Context, in *ApiV1AnalyticsGoodsLabelingGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsGoodsLabelingGet200Response, error)
	ApiV1AnalyticsIncorrectAttachmentsGet(ctx context.Context, in *ApiV1AnalyticsIncorrectAttachmentsGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsIncorrectAttachmentsGet200Response, error)
	ApiV1AnalyticsStorageCoefficientGet(ctx context.Context, in *ApiV1AnalyticsStorageCoefficientGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsStorageCoefficientGet200Response, error)
	ApiV1PaidStorageGet(ctx context.Context, in *ApiV1PaidStorageGetRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	ApiV1PaidStorageTasksTaskIdDownloadGet(ctx context.Context, in *ApiV1PaidStorageTasksTaskIdDownloadGetRequest, opts ...grpc.CallOption) (*ApiV1PaidStorageTasksTaskIdDownloadGetResponse, error)
	ApiV1PaidStorageTasksTaskIdStatusGet(ctx context.Context, in *ApiV1PaidStorageTasksTaskIdStatusGetRequest, opts ...grpc.CallOption) (*GetTasksResponse, error)
	ApiV2NmReportDetailHistoryPost(ctx context.Context, in *ApiV2NmReportDetailHistoryPostRequest, opts ...grpc.CallOption) (*NmReportDetailHistoryResponse, error)
	ApiV2NmReportDetailPost(ctx context.Context, in *ApiV2NmReportDetailPostRequest, opts ...grpc.CallOption) (*NmReportDetailResponse, error)
	ApiV2NmReportDownloadsFileDownloadIdGet(ctx context.Context, in *ApiV2NmReportDownloadsFileDownloadIdGetRequest, opts ...grpc.CallOption) (*ApiV2NmReportDownloadsFileDownloadIdGetResponse, error)
	ApiV2NmReportDownloadsGet(ctx context.Context, in *ApiV2NmReportDownloadsGetRequest, opts ...grpc.CallOption) (*NmReportGetReportsResponse, error)
	ApiV2NmReportDownloadsPost(ctx context.Context, in *ApiV2NmReportDownloadsPostRequest, opts ...grpc.CallOption) (*NmReportCreateReportResponse, error)
	ApiV2NmReportDownloadsRetryPost(ctx context.Context, in *ApiV2NmReportDownloadsRetryPostRequest, opts ...grpc.CallOption) (*NmReportRetryReportResponse, error)
	ApiV2NmReportGroupedHistoryPost(ctx context.Context, in *ApiV2NmReportGroupedHistoryPostRequest, opts ...grpc.CallOption) (*NmReportGroupedHistoryResponse, error)
	ApiV2NmReportGroupedPost(ctx context.Context, in *ApiV2NmReportGroupedPostRequest, opts ...grpc.CallOption) (*NmReportGroupedResponse, error)
}

type defaultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDefaultServiceClient(cc grpc.ClientConnInterface) DefaultServiceClient {
	return &defaultServiceClient{cc}
}

func (c *defaultServiceClient) ApiV1AnalyticsAcceptanceReportGet(ctx context.Context, in *ApiV1AnalyticsAcceptanceReportGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsAcceptanceReportGet200Response, error) {
	out := new(ApiV1AnalyticsAcceptanceReportGet200Response)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1AnalyticsAcceptanceReportGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1AnalyticsAntifraudDetailsGet(ctx context.Context, in *ApiV1AnalyticsAntifraudDetailsGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsAntifraudDetailsGet200Response, error) {
	out := new(ApiV1AnalyticsAntifraudDetailsGet200Response)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1AnalyticsAntifraudDetailsGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1AnalyticsExciseReportPost(ctx context.Context, in *ApiV1AnalyticsExciseReportPostRequest, opts ...grpc.CallOption) (*ExciseReportResponse, error) {
	out := new(ExciseReportResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1AnalyticsExciseReportPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1AnalyticsGoodsLabelingGet(ctx context.Context, in *ApiV1AnalyticsGoodsLabelingGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsGoodsLabelingGet200Response, error) {
	out := new(ApiV1AnalyticsGoodsLabelingGet200Response)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1AnalyticsGoodsLabelingGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1AnalyticsIncorrectAttachmentsGet(ctx context.Context, in *ApiV1AnalyticsIncorrectAttachmentsGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsIncorrectAttachmentsGet200Response, error) {
	out := new(ApiV1AnalyticsIncorrectAttachmentsGet200Response)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1AnalyticsIncorrectAttachmentsGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1AnalyticsStorageCoefficientGet(ctx context.Context, in *ApiV1AnalyticsStorageCoefficientGetRequest, opts ...grpc.CallOption) (*ApiV1AnalyticsStorageCoefficientGet200Response, error) {
	out := new(ApiV1AnalyticsStorageCoefficientGet200Response)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1AnalyticsStorageCoefficientGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1PaidStorageGet(ctx context.Context, in *ApiV1PaidStorageGetRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1PaidStorageGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1PaidStorageTasksTaskIdDownloadGet(ctx context.Context, in *ApiV1PaidStorageTasksTaskIdDownloadGetRequest, opts ...grpc.CallOption) (*ApiV1PaidStorageTasksTaskIdDownloadGetResponse, error) {
	out := new(ApiV1PaidStorageTasksTaskIdDownloadGetResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1PaidStorageTasksTaskIdDownloadGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV1PaidStorageTasksTaskIdStatusGet(ctx context.Context, in *ApiV1PaidStorageTasksTaskIdStatusGetRequest, opts ...grpc.CallOption) (*GetTasksResponse, error) {
	out := new(GetTasksResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV1PaidStorageTasksTaskIdStatusGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportDetailHistoryPost(ctx context.Context, in *ApiV2NmReportDetailHistoryPostRequest, opts ...grpc.CallOption) (*NmReportDetailHistoryResponse, error) {
	out := new(NmReportDetailHistoryResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportDetailHistoryPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportDetailPost(ctx context.Context, in *ApiV2NmReportDetailPostRequest, opts ...grpc.CallOption) (*NmReportDetailResponse, error) {
	out := new(NmReportDetailResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportDetailPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportDownloadsFileDownloadIdGet(ctx context.Context, in *ApiV2NmReportDownloadsFileDownloadIdGetRequest, opts ...grpc.CallOption) (*ApiV2NmReportDownloadsFileDownloadIdGetResponse, error) {
	out := new(ApiV2NmReportDownloadsFileDownloadIdGetResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportDownloadsFileDownloadIdGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportDownloadsGet(ctx context.Context, in *ApiV2NmReportDownloadsGetRequest, opts ...grpc.CallOption) (*NmReportGetReportsResponse, error) {
	out := new(NmReportGetReportsResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportDownloadsGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportDownloadsPost(ctx context.Context, in *ApiV2NmReportDownloadsPostRequest, opts ...grpc.CallOption) (*NmReportCreateReportResponse, error) {
	out := new(NmReportCreateReportResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportDownloadsPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportDownloadsRetryPost(ctx context.Context, in *ApiV2NmReportDownloadsRetryPostRequest, opts ...grpc.CallOption) (*NmReportRetryReportResponse, error) {
	out := new(NmReportRetryReportResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportDownloadsRetryPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportGroupedHistoryPost(ctx context.Context, in *ApiV2NmReportGroupedHistoryPostRequest, opts ...grpc.CallOption) (*NmReportGroupedHistoryResponse, error) {
	out := new(NmReportGroupedHistoryResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportGroupedHistoryPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) ApiV2NmReportGroupedPost(ctx context.Context, in *ApiV2NmReportGroupedPostRequest, opts ...grpc.CallOption) (*NmReportGroupedResponse, error) {
	out := new(NmReportGroupedResponse)
	err := c.cc.Invoke(ctx, DefaultService_ApiV2NmReportGroupedPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultServiceServer is the server API for DefaultService service.
// All implementations must embed UnimplementedDefaultServiceServer
// for forward compatibility
type DefaultServiceServer interface {
	ApiV1AnalyticsAcceptanceReportGet(context.Context, *ApiV1AnalyticsAcceptanceReportGetRequest) (*ApiV1AnalyticsAcceptanceReportGet200Response, error)
	ApiV1AnalyticsAntifraudDetailsGet(context.Context, *ApiV1AnalyticsAntifraudDetailsGetRequest) (*ApiV1AnalyticsAntifraudDetailsGet200Response, error)
	ApiV1AnalyticsExciseReportPost(context.Context, *ApiV1AnalyticsExciseReportPostRequest) (*ExciseReportResponse, error)
	ApiV1AnalyticsGoodsLabelingGet(context.Context, *ApiV1AnalyticsGoodsLabelingGetRequest) (*ApiV1AnalyticsGoodsLabelingGet200Response, error)
	ApiV1AnalyticsIncorrectAttachmentsGet(context.Context, *ApiV1AnalyticsIncorrectAttachmentsGetRequest) (*ApiV1AnalyticsIncorrectAttachmentsGet200Response, error)
	ApiV1AnalyticsStorageCoefficientGet(context.Context, *ApiV1AnalyticsStorageCoefficientGetRequest) (*ApiV1AnalyticsStorageCoefficientGet200Response, error)
	ApiV1PaidStorageGet(context.Context, *ApiV1PaidStorageGetRequest) (*CreateTaskResponse, error)
	ApiV1PaidStorageTasksTaskIdDownloadGet(context.Context, *ApiV1PaidStorageTasksTaskIdDownloadGetRequest) (*ApiV1PaidStorageTasksTaskIdDownloadGetResponse, error)
	ApiV1PaidStorageTasksTaskIdStatusGet(context.Context, *ApiV1PaidStorageTasksTaskIdStatusGetRequest) (*GetTasksResponse, error)
	ApiV2NmReportDetailHistoryPost(context.Context, *ApiV2NmReportDetailHistoryPostRequest) (*NmReportDetailHistoryResponse, error)
	ApiV2NmReportDetailPost(context.Context, *ApiV2NmReportDetailPostRequest) (*NmReportDetailResponse, error)
	ApiV2NmReportDownloadsFileDownloadIdGet(context.Context, *ApiV2NmReportDownloadsFileDownloadIdGetRequest) (*ApiV2NmReportDownloadsFileDownloadIdGetResponse, error)
	ApiV2NmReportDownloadsGet(context.Context, *ApiV2NmReportDownloadsGetRequest) (*NmReportGetReportsResponse, error)
	ApiV2NmReportDownloadsPost(context.Context, *ApiV2NmReportDownloadsPostRequest) (*NmReportCreateReportResponse, error)
	ApiV2NmReportDownloadsRetryPost(context.Context, *ApiV2NmReportDownloadsRetryPostRequest) (*NmReportRetryReportResponse, error)
	ApiV2NmReportGroupedHistoryPost(context.Context, *ApiV2NmReportGroupedHistoryPostRequest) (*NmReportGroupedHistoryResponse, error)
	ApiV2NmReportGroupedPost(context.Context, *ApiV2NmReportGroupedPostRequest) (*NmReportGroupedResponse, error)
	mustEmbedUnimplementedDefaultServiceServer()
}

// UnimplementedDefaultServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDefaultServiceServer struct {
}

func (UnimplementedDefaultServiceServer) ApiV1AnalyticsAcceptanceReportGet(context.Context, *ApiV1AnalyticsAcceptanceReportGetRequest) (*ApiV1AnalyticsAcceptanceReportGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1AnalyticsAcceptanceReportGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1AnalyticsAntifraudDetailsGet(context.Context, *ApiV1AnalyticsAntifraudDetailsGetRequest) (*ApiV1AnalyticsAntifraudDetailsGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1AnalyticsAntifraudDetailsGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1AnalyticsExciseReportPost(context.Context, *ApiV1AnalyticsExciseReportPostRequest) (*ExciseReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1AnalyticsExciseReportPost not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1AnalyticsGoodsLabelingGet(context.Context, *ApiV1AnalyticsGoodsLabelingGetRequest) (*ApiV1AnalyticsGoodsLabelingGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1AnalyticsGoodsLabelingGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1AnalyticsIncorrectAttachmentsGet(context.Context, *ApiV1AnalyticsIncorrectAttachmentsGetRequest) (*ApiV1AnalyticsIncorrectAttachmentsGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1AnalyticsIncorrectAttachmentsGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1AnalyticsStorageCoefficientGet(context.Context, *ApiV1AnalyticsStorageCoefficientGetRequest) (*ApiV1AnalyticsStorageCoefficientGet200Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1AnalyticsStorageCoefficientGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1PaidStorageGet(context.Context, *ApiV1PaidStorageGetRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1PaidStorageGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1PaidStorageTasksTaskIdDownloadGet(context.Context, *ApiV1PaidStorageTasksTaskIdDownloadGetRequest) (*ApiV1PaidStorageTasksTaskIdDownloadGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1PaidStorageTasksTaskIdDownloadGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV1PaidStorageTasksTaskIdStatusGet(context.Context, *ApiV1PaidStorageTasksTaskIdStatusGetRequest) (*GetTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV1PaidStorageTasksTaskIdStatusGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportDetailHistoryPost(context.Context, *ApiV2NmReportDetailHistoryPostRequest) (*NmReportDetailHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportDetailHistoryPost not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportDetailPost(context.Context, *ApiV2NmReportDetailPostRequest) (*NmReportDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportDetailPost not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportDownloadsFileDownloadIdGet(context.Context, *ApiV2NmReportDownloadsFileDownloadIdGetRequest) (*ApiV2NmReportDownloadsFileDownloadIdGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportDownloadsFileDownloadIdGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportDownloadsGet(context.Context, *ApiV2NmReportDownloadsGetRequest) (*NmReportGetReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportDownloadsGet not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportDownloadsPost(context.Context, *ApiV2NmReportDownloadsPostRequest) (*NmReportCreateReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportDownloadsPost not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportDownloadsRetryPost(context.Context, *ApiV2NmReportDownloadsRetryPostRequest) (*NmReportRetryReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportDownloadsRetryPost not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportGroupedHistoryPost(context.Context, *ApiV2NmReportGroupedHistoryPostRequest) (*NmReportGroupedHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportGroupedHistoryPost not implemented")
}
func (UnimplementedDefaultServiceServer) ApiV2NmReportGroupedPost(context.Context, *ApiV2NmReportGroupedPostRequest) (*NmReportGroupedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiV2NmReportGroupedPost not implemented")
}
func (UnimplementedDefaultServiceServer) mustEmbedUnimplementedDefaultServiceServer() {}

// UnsafeDefaultServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DefaultServiceServer will
// result in compilation errors.
type UnsafeDefaultServiceServer interface {
	mustEmbedUnimplementedDefaultServiceServer()
}

func RegisterDefaultServiceServer(s grpc.ServiceRegistrar, srv DefaultServiceServer) {
	s.RegisterService(&DefaultService_ServiceDesc, srv)
}

func _DefaultService_ApiV1AnalyticsAcceptanceReportGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1AnalyticsAcceptanceReportGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1AnalyticsAcceptanceReportGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1AnalyticsAcceptanceReportGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1AnalyticsAcceptanceReportGet(ctx, req.(*ApiV1AnalyticsAcceptanceReportGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1AnalyticsAntifraudDetailsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1AnalyticsAntifraudDetailsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1AnalyticsAntifraudDetailsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1AnalyticsAntifraudDetailsGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1AnalyticsAntifraudDetailsGet(ctx, req.(*ApiV1AnalyticsAntifraudDetailsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1AnalyticsExciseReportPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1AnalyticsExciseReportPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1AnalyticsExciseReportPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1AnalyticsExciseReportPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1AnalyticsExciseReportPost(ctx, req.(*ApiV1AnalyticsExciseReportPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1AnalyticsGoodsLabelingGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1AnalyticsGoodsLabelingGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1AnalyticsGoodsLabelingGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1AnalyticsGoodsLabelingGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1AnalyticsGoodsLabelingGet(ctx, req.(*ApiV1AnalyticsGoodsLabelingGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1AnalyticsIncorrectAttachmentsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1AnalyticsIncorrectAttachmentsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1AnalyticsIncorrectAttachmentsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1AnalyticsIncorrectAttachmentsGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1AnalyticsIncorrectAttachmentsGet(ctx, req.(*ApiV1AnalyticsIncorrectAttachmentsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1AnalyticsStorageCoefficientGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1AnalyticsStorageCoefficientGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1AnalyticsStorageCoefficientGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1AnalyticsStorageCoefficientGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1AnalyticsStorageCoefficientGet(ctx, req.(*ApiV1AnalyticsStorageCoefficientGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1PaidStorageGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1PaidStorageGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1PaidStorageGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1PaidStorageGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1PaidStorageGet(ctx, req.(*ApiV1PaidStorageGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1PaidStorageTasksTaskIdDownloadGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1PaidStorageTasksTaskIdDownloadGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1PaidStorageTasksTaskIdDownloadGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1PaidStorageTasksTaskIdDownloadGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1PaidStorageTasksTaskIdDownloadGet(ctx, req.(*ApiV1PaidStorageTasksTaskIdDownloadGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV1PaidStorageTasksTaskIdStatusGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV1PaidStorageTasksTaskIdStatusGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV1PaidStorageTasksTaskIdStatusGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV1PaidStorageTasksTaskIdStatusGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV1PaidStorageTasksTaskIdStatusGet(ctx, req.(*ApiV1PaidStorageTasksTaskIdStatusGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportDetailHistoryPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportDetailHistoryPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportDetailHistoryPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportDetailHistoryPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportDetailHistoryPost(ctx, req.(*ApiV2NmReportDetailHistoryPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportDetailPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportDetailPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportDetailPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportDetailPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportDetailPost(ctx, req.(*ApiV2NmReportDetailPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportDownloadsFileDownloadIdGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportDownloadsFileDownloadIdGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsFileDownloadIdGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportDownloadsFileDownloadIdGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsFileDownloadIdGet(ctx, req.(*ApiV2NmReportDownloadsFileDownloadIdGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportDownloadsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportDownloadsGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportDownloadsGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsGet(ctx, req.(*ApiV2NmReportDownloadsGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportDownloadsPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportDownloadsPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportDownloadsPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsPost(ctx, req.(*ApiV2NmReportDownloadsPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportDownloadsRetryPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportDownloadsRetryPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsRetryPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportDownloadsRetryPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportDownloadsRetryPost(ctx, req.(*ApiV2NmReportDownloadsRetryPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportGroupedHistoryPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportGroupedHistoryPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportGroupedHistoryPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportGroupedHistoryPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportGroupedHistoryPost(ctx, req.(*ApiV2NmReportGroupedHistoryPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_ApiV2NmReportGroupedPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiV2NmReportGroupedPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).ApiV2NmReportGroupedPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DefaultService_ApiV2NmReportGroupedPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).ApiV2NmReportGroupedPost(ctx, req.(*ApiV2NmReportGroupedPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DefaultService_ServiceDesc is the grpc.ServiceDesc for DefaultService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DefaultService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wb.analytics.v1.DefaultService",
	HandlerType: (*DefaultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApiV1AnalyticsAcceptanceReportGet",
			Handler:    _DefaultService_ApiV1AnalyticsAcceptanceReportGet_Handler,
		},
		{
			MethodName: "ApiV1AnalyticsAntifraudDetailsGet",
			Handler:    _DefaultService_ApiV1AnalyticsAntifraudDetailsGet_Handler,
		},
		{
			MethodName: "ApiV1AnalyticsExciseReportPost",
			Handler:    _DefaultService_ApiV1AnalyticsExciseReportPost_Handler,
		},
		{
			MethodName: "ApiV1AnalyticsGoodsLabelingGet",
			Handler:    _DefaultService_ApiV1AnalyticsGoodsLabelingGet_Handler,
		},
		{
			MethodName: "ApiV1AnalyticsIncorrectAttachmentsGet",
			Handler:    _DefaultService_ApiV1AnalyticsIncorrectAttachmentsGet_Handler,
		},
		{
			MethodName: "ApiV1AnalyticsStorageCoefficientGet",
			Handler:    _DefaultService_ApiV1AnalyticsStorageCoefficientGet_Handler,
		},
		{
			MethodName: "ApiV1PaidStorageGet",
			Handler:    _DefaultService_ApiV1PaidStorageGet_Handler,
		},
		{
			MethodName: "ApiV1PaidStorageTasksTaskIdDownloadGet",
			Handler:    _DefaultService_ApiV1PaidStorageTasksTaskIdDownloadGet_Handler,
		},
		{
			MethodName: "ApiV1PaidStorageTasksTaskIdStatusGet",
			Handler:    _DefaultService_ApiV1PaidStorageTasksTaskIdStatusGet_Handler,
		},
		{
			MethodName: "ApiV2NmReportDetailHistoryPost",
			Handler:    _DefaultService_ApiV2NmReportDetailHistoryPost_Handler,
		},
		{
			MethodName: "ApiV2NmReportDetailPost",
			Handler:    _DefaultService_ApiV2NmReportDetailPost_Handler,
		},
		{
			MethodName: "ApiV2NmReportDownloadsFileDownloadIdGet",
			Handler:    _DefaultService_ApiV2NmReportDownloadsFileDownloadIdGet_Handler,
		},
		{
			MethodName: "ApiV2NmReportDownloadsGet",
			Handler:    _DefaultService_ApiV2NmReportDownloadsGet_Handler,
		},
		{
			MethodName: "ApiV2NmReportDownloadsPost",
			Handler:    _DefaultService_ApiV2NmReportDownloadsPost_Handler,
		},
		{
			MethodName: "ApiV2NmReportDownloadsRetryPost",
			Handler:    _DefaultService_ApiV2NmReportDownloadsRetryPost_Handler,
		},
		{
			MethodName: "ApiV2NmReportGroupedHistoryPost",
			Handler:    _DefaultService_ApiV2NmReportGroupedHistoryPost_Handler,
		},
		{
			MethodName: "ApiV2NmReportGroupedPost",
			Handler:    _DefaultService_ApiV2NmReportGroupedPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wb/analytics/v1/_service.proto",
}