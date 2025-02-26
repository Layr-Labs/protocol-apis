// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: eigenlayer/sidecar/v1/slashing/rpc.proto

/*
Package slashing is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package slashing

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_Slashing_ListStakerSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, client SlashingClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListStakerSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["stakerAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "stakerAddress")
	}

	protoReq.StakerAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "stakerAddress", err)
	}

	msg, err := client.ListStakerSlashingHistory(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Slashing_ListStakerSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, server SlashingServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListStakerSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["stakerAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "stakerAddress")
	}

	protoReq.StakerAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "stakerAddress", err)
	}

	msg, err := server.ListStakerSlashingHistory(ctx, &protoReq)
	return msg, metadata, err

}

func request_Slashing_ListOperatorSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, client SlashingClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListOperatorSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["operatorAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "operatorAddress")
	}

	protoReq.OperatorAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "operatorAddress", err)
	}

	msg, err := client.ListOperatorSlashingHistory(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Slashing_ListOperatorSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, server SlashingServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListOperatorSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["operatorAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "operatorAddress")
	}

	protoReq.OperatorAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "operatorAddress", err)
	}

	msg, err := server.ListOperatorSlashingHistory(ctx, &protoReq)
	return msg, metadata, err

}

func request_Slashing_ListAvsSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, client SlashingClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListAvsSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["avsAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "avsAddress")
	}

	protoReq.AvsAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "avsAddress", err)
	}

	msg, err := client.ListAvsSlashingHistory(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Slashing_ListAvsSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, server SlashingServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListAvsSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["avsAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "avsAddress")
	}

	protoReq.AvsAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "avsAddress", err)
	}

	msg, err := server.ListAvsSlashingHistory(ctx, &protoReq)
	return msg, metadata, err

}

func request_Slashing_ListAvsOperatorSetSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, client SlashingClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListAvsOperatorSetSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["avsAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "avsAddress")
	}

	protoReq.AvsAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "avsAddress", err)
	}

	val, ok = pathParams["operatorSetId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "operatorSetId")
	}

	protoReq.OperatorSetId, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "operatorSetId", err)
	}

	msg, err := client.ListAvsOperatorSetSlashingHistory(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Slashing_ListAvsOperatorSetSlashingHistory_0(ctx context.Context, marshaler runtime.Marshaler, server SlashingServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListAvsOperatorSetSlashingHistoryRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["avsAddress"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "avsAddress")
	}

	protoReq.AvsAddress, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "avsAddress", err)
	}

	val, ok = pathParams["operatorSetId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "operatorSetId")
	}

	protoReq.OperatorSetId, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "operatorSetId", err)
	}

	msg, err := server.ListAvsOperatorSetSlashingHistory(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterSlashingHandlerServer registers the http handlers for service Slashing to "mux".
// UnaryRPC     :call SlashingServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterSlashingHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterSlashingHandlerServer(ctx context.Context, mux *runtime.ServeMux, server SlashingServer) error {

	mux.Handle("GET", pattern_Slashing_ListStakerSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListStakerSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/stakers/{stakerAddress}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Slashing_ListStakerSlashingHistory_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListStakerSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Slashing_ListOperatorSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListOperatorSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/operators/{operatorAddress}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Slashing_ListOperatorSlashingHistory_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListOperatorSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Slashing_ListAvsSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListAvsSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/avss/{avsAddress}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Slashing_ListAvsSlashingHistory_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListAvsSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Slashing_ListAvsOperatorSetSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListAvsOperatorSetSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/avss/{avsAddress}/operator-sets/{operatorSetId}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Slashing_ListAvsOperatorSetSlashingHistory_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListAvsOperatorSetSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterSlashingHandlerFromEndpoint is same as RegisterSlashingHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSlashingHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSlashingHandler(ctx, mux, conn)
}

// RegisterSlashingHandler registers the http handlers for service Slashing to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSlashingHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSlashingHandlerClient(ctx, mux, NewSlashingClient(conn))
}

// RegisterSlashingHandlerClient registers the http handlers for service Slashing
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SlashingClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SlashingClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SlashingClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterSlashingHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SlashingClient) error {

	mux.Handle("GET", pattern_Slashing_ListStakerSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListStakerSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/stakers/{stakerAddress}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Slashing_ListStakerSlashingHistory_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListStakerSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Slashing_ListOperatorSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListOperatorSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/operators/{operatorAddress}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Slashing_ListOperatorSlashingHistory_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListOperatorSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Slashing_ListAvsSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListAvsSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/avss/{avsAddress}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Slashing_ListAvsSlashingHistory_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListAvsSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Slashing_ListAvsOperatorSetSlashingHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/eigenlayer.sidecar.v1.slashing.Slashing/ListAvsOperatorSetSlashingHistory", runtime.WithHTTPPathPattern("/v1/slashing/avss/{avsAddress}/operator-sets/{operatorSetId}/slashing-history"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Slashing_ListAvsOperatorSetSlashingHistory_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Slashing_ListAvsOperatorSetSlashingHistory_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Slashing_ListStakerSlashingHistory_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"v1", "slashing", "stakers", "stakerAddress", "slashing-history"}, ""))

	pattern_Slashing_ListOperatorSlashingHistory_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"v1", "slashing", "operators", "operatorAddress", "slashing-history"}, ""))

	pattern_Slashing_ListAvsSlashingHistory_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"v1", "slashing", "avss", "avsAddress", "slashing-history"}, ""))

	pattern_Slashing_ListAvsOperatorSetSlashingHistory_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 5, 2, 6}, []string{"v1", "slashing", "avss", "avsAddress", "operator-sets", "operatorSetId", "slashing-history"}, ""))
)

var (
	forward_Slashing_ListStakerSlashingHistory_0 = runtime.ForwardResponseMessage

	forward_Slashing_ListOperatorSlashingHistory_0 = runtime.ForwardResponseMessage

	forward_Slashing_ListAvsSlashingHistory_0 = runtime.ForwardResponseMessage

	forward_Slashing_ListAvsOperatorSetSlashingHistory_0 = runtime.ForwardResponseMessage
)
