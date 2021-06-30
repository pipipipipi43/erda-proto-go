// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: trace.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/trace/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// TraceService trace.proto
	TraceService() pb.TraceServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		traceService: pb.NewTraceServiceClient(cc),
	}
}

type serviceClients struct {
	traceService pb.TraceServiceClient
}

func (c *serviceClients) TraceService() pb.TraceServiceClient {
	return c.traceService
}

type traceServiceWrapper struct {
	client pb.TraceServiceClient
	opts   []grpc1.CallOption
}

func (s *traceServiceWrapper) GetSpans(ctx context.Context, req *pb.GetSpansRequest) (*pb.GetSpansResponse, error) {
	return s.client.GetSpans(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *traceServiceWrapper) GetTraces(ctx context.Context, req *pb.GetTracesRequest) (*pb.GetTracesResponse, error) {
	return s.client.GetTraces(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}