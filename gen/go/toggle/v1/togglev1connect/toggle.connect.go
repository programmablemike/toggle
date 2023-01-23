// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: toggle/v1/toggle.proto

package togglev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
	v1 "toggle/gen/go/toggle/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ToggleServiceName is the fully-qualified name of the ToggleService service.
	ToggleServiceName = "toggle.v1.ToggleService"
)

// ToggleServiceClient is a client for the toggle.v1.ToggleService service.
type ToggleServiceClient interface {
	CreateScope(context.Context, *connect_go.Request[v1.CreateScopeRequest]) (*connect_go.Response[v1.CreateScopeResponse], error)
}

// NewToggleServiceClient constructs a client for the toggle.v1.ToggleService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewToggleServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ToggleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &toggleServiceClient{
		createScope: connect_go.NewClient[v1.CreateScopeRequest, v1.CreateScopeResponse](
			httpClient,
			baseURL+"/toggle.v1.ToggleService/CreateScope",
			opts...,
		),
	}
}

// toggleServiceClient implements ToggleServiceClient.
type toggleServiceClient struct {
	createScope *connect_go.Client[v1.CreateScopeRequest, v1.CreateScopeResponse]
}

// CreateScope calls toggle.v1.ToggleService.CreateScope.
func (c *toggleServiceClient) CreateScope(ctx context.Context, req *connect_go.Request[v1.CreateScopeRequest]) (*connect_go.Response[v1.CreateScopeResponse], error) {
	return c.createScope.CallUnary(ctx, req)
}

// ToggleServiceHandler is an implementation of the toggle.v1.ToggleService service.
type ToggleServiceHandler interface {
	CreateScope(context.Context, *connect_go.Request[v1.CreateScopeRequest]) (*connect_go.Response[v1.CreateScopeResponse], error)
}

// NewToggleServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewToggleServiceHandler(svc ToggleServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/toggle.v1.ToggleService/CreateScope", connect_go.NewUnaryHandler(
		"/toggle.v1.ToggleService/CreateScope",
		svc.CreateScope,
		opts...,
	))
	return "/toggle.v1.ToggleService/", mux
}

// UnimplementedToggleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedToggleServiceHandler struct{}

func (UnimplementedToggleServiceHandler) CreateScope(context.Context, *connect_go.Request[v1.CreateScopeRequest]) (*connect_go.Response[v1.CreateScopeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("toggle.v1.ToggleService.CreateScope is not implemented"))
}
