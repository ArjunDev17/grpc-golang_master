// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: calculator.proto

package proto

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
	SumServices_Sum_FullMethodName    = "/calculator.SumServices/Sum"
	SumServices_Primes_FullMethodName = "/calculator.SumServices/Primes"
)

// SumServicesClient is the client API for SumServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SumServicesClient interface {
	// rpc MethodName (Request) returns (Response);
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrimeResponse], error)
}

type sumServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServicesClient(cc grpc.ClientConnInterface) SumServicesClient {
	return &sumServicesClient{cc}
}

func (c *sumServicesClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, SumServices_Sum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServicesClient) Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PrimeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SumServices_ServiceDesc.Streams[0], SumServices_Primes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PrimeRequest, PrimeResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SumServices_PrimesClient = grpc.ServerStreamingClient[PrimeResponse]

// SumServicesServer is the server API for SumServices service.
// All implementations must embed UnimplementedSumServicesServer
// for forward compatibility.
type SumServicesServer interface {
	// rpc MethodName (Request) returns (Response);
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Primes(*PrimeRequest, grpc.ServerStreamingServer[PrimeResponse]) error
	mustEmbedUnimplementedSumServicesServer()
}

// UnimplementedSumServicesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSumServicesServer struct{}

func (UnimplementedSumServicesServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedSumServicesServer) Primes(*PrimeRequest, grpc.ServerStreamingServer[PrimeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedSumServicesServer) mustEmbedUnimplementedSumServicesServer() {}
func (UnimplementedSumServicesServer) testEmbeddedByValue()                     {}

// UnsafeSumServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SumServicesServer will
// result in compilation errors.
type UnsafeSumServicesServer interface {
	mustEmbedUnimplementedSumServicesServer()
}

func RegisterSumServicesServer(s grpc.ServiceRegistrar, srv SumServicesServer) {
	// If the following call pancis, it indicates UnimplementedSumServicesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SumServices_ServiceDesc, srv)
}

func _SumServices_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServicesServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SumServices_Sum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServicesServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumServices_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServicesServer).Primes(m, &grpc.GenericServerStream[PrimeRequest, PrimeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SumServices_PrimesServer = grpc.ServerStreamingServer[PrimeResponse]

// SumServices_ServiceDesc is the grpc.ServiceDesc for SumServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SumServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.SumServices",
	HandlerType: (*SumServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumServices_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _SumServices_Primes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
