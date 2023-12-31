// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: monitor.proto

package protoconnect

import (
	context "context"
	errors "errors"
	proto "github.com/anukul/xmon/backend/proto"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// MonitorName is the fully-qualified name of the Monitor service.
	MonitorName = "xmon.Monitor"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MonitorListNodesProcedure is the fully-qualified name of the Monitor's ListNodes RPC.
	MonitorListNodesProcedure = "/xmon.Monitor/ListNodes"
	// MonitorListBlocksProcedure is the fully-qualified name of the Monitor's ListBlocks RPC.
	MonitorListBlocksProcedure = "/xmon.Monitor/ListBlocks"
	// MonitorListSlotsProcedure is the fully-qualified name of the Monitor's ListSlots RPC.
	MonitorListSlotsProcedure = "/xmon.Monitor/ListSlots"
)

// MonitorClient is a client for the xmon.Monitor service.
type MonitorClient interface {
	ListNodes(context.Context, *connect_go.Request[proto.ListNodesRequest]) (*connect_go.Response[proto.ListNodesResponse], error)
	ListBlocks(context.Context, *connect_go.Request[proto.ListBlocksRequest]) (*connect_go.Response[proto.ListBlocksResponse], error)
	ListSlots(context.Context, *connect_go.Request[proto.ListSlotsRequest]) (*connect_go.Response[proto.ListSlotsResponse], error)
}

// NewMonitorClient constructs a client for the xmon.Monitor service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMonitorClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MonitorClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &monitorClient{
		listNodes: connect_go.NewClient[proto.ListNodesRequest, proto.ListNodesResponse](
			httpClient,
			baseURL+MonitorListNodesProcedure,
			opts...,
		),
		listBlocks: connect_go.NewClient[proto.ListBlocksRequest, proto.ListBlocksResponse](
			httpClient,
			baseURL+MonitorListBlocksProcedure,
			opts...,
		),
		listSlots: connect_go.NewClient[proto.ListSlotsRequest, proto.ListSlotsResponse](
			httpClient,
			baseURL+MonitorListSlotsProcedure,
			opts...,
		),
	}
}

// monitorClient implements MonitorClient.
type monitorClient struct {
	listNodes  *connect_go.Client[proto.ListNodesRequest, proto.ListNodesResponse]
	listBlocks *connect_go.Client[proto.ListBlocksRequest, proto.ListBlocksResponse]
	listSlots  *connect_go.Client[proto.ListSlotsRequest, proto.ListSlotsResponse]
}

// ListNodes calls xmon.Monitor.ListNodes.
func (c *monitorClient) ListNodes(ctx context.Context, req *connect_go.Request[proto.ListNodesRequest]) (*connect_go.Response[proto.ListNodesResponse], error) {
	return c.listNodes.CallUnary(ctx, req)
}

// ListBlocks calls xmon.Monitor.ListBlocks.
func (c *monitorClient) ListBlocks(ctx context.Context, req *connect_go.Request[proto.ListBlocksRequest]) (*connect_go.Response[proto.ListBlocksResponse], error) {
	return c.listBlocks.CallUnary(ctx, req)
}

// ListSlots calls xmon.Monitor.ListSlots.
func (c *monitorClient) ListSlots(ctx context.Context, req *connect_go.Request[proto.ListSlotsRequest]) (*connect_go.Response[proto.ListSlotsResponse], error) {
	return c.listSlots.CallUnary(ctx, req)
}

// MonitorHandler is an implementation of the xmon.Monitor service.
type MonitorHandler interface {
	ListNodes(context.Context, *connect_go.Request[proto.ListNodesRequest]) (*connect_go.Response[proto.ListNodesResponse], error)
	ListBlocks(context.Context, *connect_go.Request[proto.ListBlocksRequest]) (*connect_go.Response[proto.ListBlocksResponse], error)
	ListSlots(context.Context, *connect_go.Request[proto.ListSlotsRequest]) (*connect_go.Response[proto.ListSlotsResponse], error)
}

// NewMonitorHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMonitorHandler(svc MonitorHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(MonitorListNodesProcedure, connect_go.NewUnaryHandler(
		MonitorListNodesProcedure,
		svc.ListNodes,
		opts...,
	))
	mux.Handle(MonitorListBlocksProcedure, connect_go.NewUnaryHandler(
		MonitorListBlocksProcedure,
		svc.ListBlocks,
		opts...,
	))
	mux.Handle(MonitorListSlotsProcedure, connect_go.NewUnaryHandler(
		MonitorListSlotsProcedure,
		svc.ListSlots,
		opts...,
	))
	return "/xmon.Monitor/", mux
}

// UnimplementedMonitorHandler returns CodeUnimplemented from all methods.
type UnimplementedMonitorHandler struct{}

func (UnimplementedMonitorHandler) ListNodes(context.Context, *connect_go.Request[proto.ListNodesRequest]) (*connect_go.Response[proto.ListNodesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("xmon.Monitor.ListNodes is not implemented"))
}

func (UnimplementedMonitorHandler) ListBlocks(context.Context, *connect_go.Request[proto.ListBlocksRequest]) (*connect_go.Response[proto.ListBlocksResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("xmon.Monitor.ListBlocks is not implemented"))
}

func (UnimplementedMonitorHandler) ListSlots(context.Context, *connect_go.Request[proto.ListSlotsRequest]) (*connect_go.Response[proto.ListSlotsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("xmon.Monitor.ListSlots is not implemented"))
}
