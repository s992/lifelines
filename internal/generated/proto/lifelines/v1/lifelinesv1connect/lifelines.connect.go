// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: lifelines/v1/lifelines.proto

package lifelinesv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/s992/lifelines/internal/generated/proto/lifelines/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// TagServiceName is the fully-qualified name of the TagService service.
	TagServiceName = "lifelines.v1.TagService"
	// LogLineServiceName is the fully-qualified name of the LogLineService service.
	LogLineServiceName = "lifelines.v1.LogLineService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TagServiceListTagsProcedure is the fully-qualified name of the TagService's ListTags RPC.
	TagServiceListTagsProcedure = "/lifelines.v1.TagService/ListTags"
	// TagServiceSearchTagsProcedure is the fully-qualified name of the TagService's SearchTags RPC.
	TagServiceSearchTagsProcedure = "/lifelines.v1.TagService/SearchTags"
	// TagServiceCreateTagProcedure is the fully-qualified name of the TagService's CreateTag RPC.
	TagServiceCreateTagProcedure = "/lifelines.v1.TagService/CreateTag"
	// LogLineServiceListLogLinesProcedure is the fully-qualified name of the LogLineService's
	// ListLogLines RPC.
	LogLineServiceListLogLinesProcedure = "/lifelines.v1.LogLineService/ListLogLines"
	// LogLineServiceCreateLogLineProcedure is the fully-qualified name of the LogLineService's
	// CreateLogLine RPC.
	LogLineServiceCreateLogLineProcedure = "/lifelines.v1.LogLineService/CreateLogLine"
)

// TagServiceClient is a client for the lifelines.v1.TagService service.
type TagServiceClient interface {
	ListTags(context.Context, *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.ListTagsResponse], error)
	SearchTags(context.Context, *connect.Request[v1.SearchTagsRequest]) (*connect.Response[v1.SearchTagsResponse], error)
	CreateTag(context.Context, *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.CreateTagResponse], error)
}

// NewTagServiceClient constructs a client for the lifelines.v1.TagService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTagServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TagServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tagServiceClient{
		listTags: connect.NewClient[v1.ListTagsRequest, v1.ListTagsResponse](
			httpClient,
			baseURL+TagServiceListTagsProcedure,
			opts...,
		),
		searchTags: connect.NewClient[v1.SearchTagsRequest, v1.SearchTagsResponse](
			httpClient,
			baseURL+TagServiceSearchTagsProcedure,
			opts...,
		),
		createTag: connect.NewClient[v1.CreateTagRequest, v1.CreateTagResponse](
			httpClient,
			baseURL+TagServiceCreateTagProcedure,
			opts...,
		),
	}
}

// tagServiceClient implements TagServiceClient.
type tagServiceClient struct {
	listTags   *connect.Client[v1.ListTagsRequest, v1.ListTagsResponse]
	searchTags *connect.Client[v1.SearchTagsRequest, v1.SearchTagsResponse]
	createTag  *connect.Client[v1.CreateTagRequest, v1.CreateTagResponse]
}

// ListTags calls lifelines.v1.TagService.ListTags.
func (c *tagServiceClient) ListTags(ctx context.Context, req *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.ListTagsResponse], error) {
	return c.listTags.CallUnary(ctx, req)
}

// SearchTags calls lifelines.v1.TagService.SearchTags.
func (c *tagServiceClient) SearchTags(ctx context.Context, req *connect.Request[v1.SearchTagsRequest]) (*connect.Response[v1.SearchTagsResponse], error) {
	return c.searchTags.CallUnary(ctx, req)
}

// CreateTag calls lifelines.v1.TagService.CreateTag.
func (c *tagServiceClient) CreateTag(ctx context.Context, req *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.CreateTagResponse], error) {
	return c.createTag.CallUnary(ctx, req)
}

// TagServiceHandler is an implementation of the lifelines.v1.TagService service.
type TagServiceHandler interface {
	ListTags(context.Context, *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.ListTagsResponse], error)
	SearchTags(context.Context, *connect.Request[v1.SearchTagsRequest]) (*connect.Response[v1.SearchTagsResponse], error)
	CreateTag(context.Context, *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.CreateTagResponse], error)
}

// NewTagServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTagServiceHandler(svc TagServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tagServiceListTagsHandler := connect.NewUnaryHandler(
		TagServiceListTagsProcedure,
		svc.ListTags,
		opts...,
	)
	tagServiceSearchTagsHandler := connect.NewUnaryHandler(
		TagServiceSearchTagsProcedure,
		svc.SearchTags,
		opts...,
	)
	tagServiceCreateTagHandler := connect.NewUnaryHandler(
		TagServiceCreateTagProcedure,
		svc.CreateTag,
		opts...,
	)
	return "/lifelines.v1.TagService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TagServiceListTagsProcedure:
			tagServiceListTagsHandler.ServeHTTP(w, r)
		case TagServiceSearchTagsProcedure:
			tagServiceSearchTagsHandler.ServeHTTP(w, r)
		case TagServiceCreateTagProcedure:
			tagServiceCreateTagHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTagServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTagServiceHandler struct{}

func (UnimplementedTagServiceHandler) ListTags(context.Context, *connect.Request[v1.ListTagsRequest]) (*connect.Response[v1.ListTagsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("lifelines.v1.TagService.ListTags is not implemented"))
}

func (UnimplementedTagServiceHandler) SearchTags(context.Context, *connect.Request[v1.SearchTagsRequest]) (*connect.Response[v1.SearchTagsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("lifelines.v1.TagService.SearchTags is not implemented"))
}

func (UnimplementedTagServiceHandler) CreateTag(context.Context, *connect.Request[v1.CreateTagRequest]) (*connect.Response[v1.CreateTagResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("lifelines.v1.TagService.CreateTag is not implemented"))
}

// LogLineServiceClient is a client for the lifelines.v1.LogLineService service.
type LogLineServiceClient interface {
	ListLogLines(context.Context, *connect.Request[v1.ListLogLinesRequest]) (*connect.Response[v1.ListLogLinesResponse], error)
	CreateLogLine(context.Context, *connect.Request[v1.CreateLogLineRequest]) (*connect.Response[v1.CreateLogLineResponse], error)
}

// NewLogLineServiceClient constructs a client for the lifelines.v1.LogLineService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLogLineServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LogLineServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &logLineServiceClient{
		listLogLines: connect.NewClient[v1.ListLogLinesRequest, v1.ListLogLinesResponse](
			httpClient,
			baseURL+LogLineServiceListLogLinesProcedure,
			opts...,
		),
		createLogLine: connect.NewClient[v1.CreateLogLineRequest, v1.CreateLogLineResponse](
			httpClient,
			baseURL+LogLineServiceCreateLogLineProcedure,
			opts...,
		),
	}
}

// logLineServiceClient implements LogLineServiceClient.
type logLineServiceClient struct {
	listLogLines  *connect.Client[v1.ListLogLinesRequest, v1.ListLogLinesResponse]
	createLogLine *connect.Client[v1.CreateLogLineRequest, v1.CreateLogLineResponse]
}

// ListLogLines calls lifelines.v1.LogLineService.ListLogLines.
func (c *logLineServiceClient) ListLogLines(ctx context.Context, req *connect.Request[v1.ListLogLinesRequest]) (*connect.Response[v1.ListLogLinesResponse], error) {
	return c.listLogLines.CallUnary(ctx, req)
}

// CreateLogLine calls lifelines.v1.LogLineService.CreateLogLine.
func (c *logLineServiceClient) CreateLogLine(ctx context.Context, req *connect.Request[v1.CreateLogLineRequest]) (*connect.Response[v1.CreateLogLineResponse], error) {
	return c.createLogLine.CallUnary(ctx, req)
}

// LogLineServiceHandler is an implementation of the lifelines.v1.LogLineService service.
type LogLineServiceHandler interface {
	ListLogLines(context.Context, *connect.Request[v1.ListLogLinesRequest]) (*connect.Response[v1.ListLogLinesResponse], error)
	CreateLogLine(context.Context, *connect.Request[v1.CreateLogLineRequest]) (*connect.Response[v1.CreateLogLineResponse], error)
}

// NewLogLineServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLogLineServiceHandler(svc LogLineServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	logLineServiceListLogLinesHandler := connect.NewUnaryHandler(
		LogLineServiceListLogLinesProcedure,
		svc.ListLogLines,
		opts...,
	)
	logLineServiceCreateLogLineHandler := connect.NewUnaryHandler(
		LogLineServiceCreateLogLineProcedure,
		svc.CreateLogLine,
		opts...,
	)
	return "/lifelines.v1.LogLineService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LogLineServiceListLogLinesProcedure:
			logLineServiceListLogLinesHandler.ServeHTTP(w, r)
		case LogLineServiceCreateLogLineProcedure:
			logLineServiceCreateLogLineHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLogLineServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLogLineServiceHandler struct{}

func (UnimplementedLogLineServiceHandler) ListLogLines(context.Context, *connect.Request[v1.ListLogLinesRequest]) (*connect.Response[v1.ListLogLinesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("lifelines.v1.LogLineService.ListLogLines is not implemented"))
}

func (UnimplementedLogLineServiceHandler) CreateLogLine(context.Context, *connect.Request[v1.CreateLogLineRequest]) (*connect.Response[v1.CreateLogLineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("lifelines.v1.LogLineService.CreateLogLine is not implemented"))
}