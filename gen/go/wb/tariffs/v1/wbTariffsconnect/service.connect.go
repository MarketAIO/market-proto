//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: wb/tariffs/v1/service.proto

package wbTariffsconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/MarketAIO/market-proto/gen/go/wb/tariffs/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TariffsServiceName is the fully-qualified name of the TariffsService service.
	TariffsServiceName = "wb.tariffs.v1.TariffsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TariffsServiceApiV1TariffsBoxGetProcedure is the fully-qualified name of the TariffsService's
	// ApiV1TariffsBoxGet RPC.
	TariffsServiceApiV1TariffsBoxGetProcedure = "/wb.tariffs.v1.TariffsService/ApiV1TariffsBoxGet"
	// TariffsServiceApiV1TariffsCommissionGetProcedure is the fully-qualified name of the
	// TariffsService's ApiV1TariffsCommissionGet RPC.
	TariffsServiceApiV1TariffsCommissionGetProcedure = "/wb.tariffs.v1.TariffsService/ApiV1TariffsCommissionGet"
	// TariffsServiceApiV1TariffsPalletGetProcedure is the fully-qualified name of the TariffsService's
	// ApiV1TariffsPalletGet RPC.
	TariffsServiceApiV1TariffsPalletGetProcedure = "/wb.tariffs.v1.TariffsService/ApiV1TariffsPalletGet"
	// TariffsServiceApiV1TariffsReturnGetProcedure is the fully-qualified name of the TariffsService's
	// ApiV1TariffsReturnGet RPC.
	TariffsServiceApiV1TariffsReturnGetProcedure = "/wb.tariffs.v1.TariffsService/ApiV1TariffsReturnGet"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	tariffsServiceServiceDescriptor                         = v1.File_wb_tariffs_v1_service_proto.Services().ByName("TariffsService")
	tariffsServiceApiV1TariffsBoxGetMethodDescriptor        = tariffsServiceServiceDescriptor.Methods().ByName("ApiV1TariffsBoxGet")
	tariffsServiceApiV1TariffsCommissionGetMethodDescriptor = tariffsServiceServiceDescriptor.Methods().ByName("ApiV1TariffsCommissionGet")
	tariffsServiceApiV1TariffsPalletGetMethodDescriptor     = tariffsServiceServiceDescriptor.Methods().ByName("ApiV1TariffsPalletGet")
	tariffsServiceApiV1TariffsReturnGetMethodDescriptor     = tariffsServiceServiceDescriptor.Methods().ByName("ApiV1TariffsReturnGet")
)

// TariffsServiceClient is a client for the wb.tariffs.v1.TariffsService service.
type TariffsServiceClient interface {
	ApiV1TariffsBoxGet(context.Context, *connect.Request[v1.ApiV1TariffsBoxGetRequest]) (*connect.Response[v1.TariffsBoxResponse], error)
	ApiV1TariffsCommissionGet(context.Context, *connect.Request[v1.ApiV1TariffsCommissionGetRequest]) (*connect.Response[v1.ApiV1TariffsCommissionGet200Response], error)
	ApiV1TariffsPalletGet(context.Context, *connect.Request[v1.ApiV1TariffsPalletGetRequest]) (*connect.Response[v1.TariffsPalletResponse], error)
	ApiV1TariffsReturnGet(context.Context, *connect.Request[v1.ApiV1TariffsReturnGetRequest]) (*connect.Response[v1.ReturnTariffsResponse], error)
}

// NewTariffsServiceClient constructs a client for the wb.tariffs.v1.TariffsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTariffsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TariffsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tariffsServiceClient{
		apiV1TariffsBoxGet: connect.NewClient[v1.ApiV1TariffsBoxGetRequest, v1.TariffsBoxResponse](
			httpClient,
			baseURL+TariffsServiceApiV1TariffsBoxGetProcedure,
			connect.WithSchema(tariffsServiceApiV1TariffsBoxGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		apiV1TariffsCommissionGet: connect.NewClient[v1.ApiV1TariffsCommissionGetRequest, v1.ApiV1TariffsCommissionGet200Response](
			httpClient,
			baseURL+TariffsServiceApiV1TariffsCommissionGetProcedure,
			connect.WithSchema(tariffsServiceApiV1TariffsCommissionGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		apiV1TariffsPalletGet: connect.NewClient[v1.ApiV1TariffsPalletGetRequest, v1.TariffsPalletResponse](
			httpClient,
			baseURL+TariffsServiceApiV1TariffsPalletGetProcedure,
			connect.WithSchema(tariffsServiceApiV1TariffsPalletGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		apiV1TariffsReturnGet: connect.NewClient[v1.ApiV1TariffsReturnGetRequest, v1.ReturnTariffsResponse](
			httpClient,
			baseURL+TariffsServiceApiV1TariffsReturnGetProcedure,
			connect.WithSchema(tariffsServiceApiV1TariffsReturnGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// tariffsServiceClient implements TariffsServiceClient.
type tariffsServiceClient struct {
	apiV1TariffsBoxGet        *connect.Client[v1.ApiV1TariffsBoxGetRequest, v1.TariffsBoxResponse]
	apiV1TariffsCommissionGet *connect.Client[v1.ApiV1TariffsCommissionGetRequest, v1.ApiV1TariffsCommissionGet200Response]
	apiV1TariffsPalletGet     *connect.Client[v1.ApiV1TariffsPalletGetRequest, v1.TariffsPalletResponse]
	apiV1TariffsReturnGet     *connect.Client[v1.ApiV1TariffsReturnGetRequest, v1.ReturnTariffsResponse]
}

// ApiV1TariffsBoxGet calls wb.tariffs.v1.TariffsService.ApiV1TariffsBoxGet.
func (c *tariffsServiceClient) ApiV1TariffsBoxGet(ctx context.Context, req *connect.Request[v1.ApiV1TariffsBoxGetRequest]) (*connect.Response[v1.TariffsBoxResponse], error) {
	return c.apiV1TariffsBoxGet.CallUnary(ctx, req)
}

// ApiV1TariffsCommissionGet calls wb.tariffs.v1.TariffsService.ApiV1TariffsCommissionGet.
func (c *tariffsServiceClient) ApiV1TariffsCommissionGet(ctx context.Context, req *connect.Request[v1.ApiV1TariffsCommissionGetRequest]) (*connect.Response[v1.ApiV1TariffsCommissionGet200Response], error) {
	return c.apiV1TariffsCommissionGet.CallUnary(ctx, req)
}

// ApiV1TariffsPalletGet calls wb.tariffs.v1.TariffsService.ApiV1TariffsPalletGet.
func (c *tariffsServiceClient) ApiV1TariffsPalletGet(ctx context.Context, req *connect.Request[v1.ApiV1TariffsPalletGetRequest]) (*connect.Response[v1.TariffsPalletResponse], error) {
	return c.apiV1TariffsPalletGet.CallUnary(ctx, req)
}

// ApiV1TariffsReturnGet calls wb.tariffs.v1.TariffsService.ApiV1TariffsReturnGet.
func (c *tariffsServiceClient) ApiV1TariffsReturnGet(ctx context.Context, req *connect.Request[v1.ApiV1TariffsReturnGetRequest]) (*connect.Response[v1.ReturnTariffsResponse], error) {
	return c.apiV1TariffsReturnGet.CallUnary(ctx, req)
}

// TariffsServiceHandler is an implementation of the wb.tariffs.v1.TariffsService service.
type TariffsServiceHandler interface {
	ApiV1TariffsBoxGet(context.Context, *connect.Request[v1.ApiV1TariffsBoxGetRequest]) (*connect.Response[v1.TariffsBoxResponse], error)
	ApiV1TariffsCommissionGet(context.Context, *connect.Request[v1.ApiV1TariffsCommissionGetRequest]) (*connect.Response[v1.ApiV1TariffsCommissionGet200Response], error)
	ApiV1TariffsPalletGet(context.Context, *connect.Request[v1.ApiV1TariffsPalletGetRequest]) (*connect.Response[v1.TariffsPalletResponse], error)
	ApiV1TariffsReturnGet(context.Context, *connect.Request[v1.ApiV1TariffsReturnGetRequest]) (*connect.Response[v1.ReturnTariffsResponse], error)
}

// NewTariffsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTariffsServiceHandler(svc TariffsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tariffsServiceApiV1TariffsBoxGetHandler := connect.NewUnaryHandler(
		TariffsServiceApiV1TariffsBoxGetProcedure,
		svc.ApiV1TariffsBoxGet,
		connect.WithSchema(tariffsServiceApiV1TariffsBoxGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tariffsServiceApiV1TariffsCommissionGetHandler := connect.NewUnaryHandler(
		TariffsServiceApiV1TariffsCommissionGetProcedure,
		svc.ApiV1TariffsCommissionGet,
		connect.WithSchema(tariffsServiceApiV1TariffsCommissionGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tariffsServiceApiV1TariffsPalletGetHandler := connect.NewUnaryHandler(
		TariffsServiceApiV1TariffsPalletGetProcedure,
		svc.ApiV1TariffsPalletGet,
		connect.WithSchema(tariffsServiceApiV1TariffsPalletGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	tariffsServiceApiV1TariffsReturnGetHandler := connect.NewUnaryHandler(
		TariffsServiceApiV1TariffsReturnGetProcedure,
		svc.ApiV1TariffsReturnGet,
		connect.WithSchema(tariffsServiceApiV1TariffsReturnGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/wb.tariffs.v1.TariffsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TariffsServiceApiV1TariffsBoxGetProcedure:
			tariffsServiceApiV1TariffsBoxGetHandler.ServeHTTP(w, r)
		case TariffsServiceApiV1TariffsCommissionGetProcedure:
			tariffsServiceApiV1TariffsCommissionGetHandler.ServeHTTP(w, r)
		case TariffsServiceApiV1TariffsPalletGetProcedure:
			tariffsServiceApiV1TariffsPalletGetHandler.ServeHTTP(w, r)
		case TariffsServiceApiV1TariffsReturnGetProcedure:
			tariffsServiceApiV1TariffsReturnGetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTariffsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTariffsServiceHandler struct{}

func (UnimplementedTariffsServiceHandler) ApiV1TariffsBoxGet(context.Context, *connect.Request[v1.ApiV1TariffsBoxGetRequest]) (*connect.Response[v1.TariffsBoxResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.tariffs.v1.TariffsService.ApiV1TariffsBoxGet is not implemented"))
}

func (UnimplementedTariffsServiceHandler) ApiV1TariffsCommissionGet(context.Context, *connect.Request[v1.ApiV1TariffsCommissionGetRequest]) (*connect.Response[v1.ApiV1TariffsCommissionGet200Response], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.tariffs.v1.TariffsService.ApiV1TariffsCommissionGet is not implemented"))
}

func (UnimplementedTariffsServiceHandler) ApiV1TariffsPalletGet(context.Context, *connect.Request[v1.ApiV1TariffsPalletGetRequest]) (*connect.Response[v1.TariffsPalletResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.tariffs.v1.TariffsService.ApiV1TariffsPalletGet is not implemented"))
}

func (UnimplementedTariffsServiceHandler) ApiV1TariffsReturnGet(context.Context, *connect.Request[v1.ApiV1TariffsReturnGetRequest]) (*connect.Response[v1.ReturnTariffsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("wb.tariffs.v1.TariffsService.ApiV1TariffsReturnGet is not implemented"))
}