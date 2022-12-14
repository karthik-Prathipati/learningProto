// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: proto/user.proto

package proto

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

// NumberManipulationClient is the client API for NumberManipulation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NumberManipulationClient interface {
	Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Number, error)
	PrimeNumbers(ctx context.Context, in *Number, opts ...grpc.CallOption) (NumberManipulation_PrimeNumbersClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (NumberManipulation_ComputeAverageClient, error)
	FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (NumberManipulation_FindMaxNumberClient, error)
}

type numberManipulationClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberManipulationClient(cc grpc.ClientConnInterface) NumberManipulationClient {
	return &numberManipulationClient{cc}
}

func (c *numberManipulationClient) Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Number, error) {
	out := new(Number)
	err := c.cc.Invoke(ctx, "/proto.NumberManipulation/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *numberManipulationClient) PrimeNumbers(ctx context.Context, in *Number, opts ...grpc.CallOption) (NumberManipulation_PrimeNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &NumberManipulation_ServiceDesc.Streams[0], "/proto.NumberManipulation/PrimeNumbers", opts...)
	if err != nil {
		return nil, err
	}
	x := &numberManipulationPrimeNumbersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NumberManipulation_PrimeNumbersClient interface {
	Recv() (*Number, error)
	grpc.ClientStream
}

type numberManipulationPrimeNumbersClient struct {
	grpc.ClientStream
}

func (x *numberManipulationPrimeNumbersClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *numberManipulationClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (NumberManipulation_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &NumberManipulation_ServiceDesc.Streams[1], "/proto.NumberManipulation/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &numberManipulationComputeAverageClient{stream}
	return x, nil
}

type NumberManipulation_ComputeAverageClient interface {
	Send(*Number) error
	CloseAndRecv() (*Number, error)
	grpc.ClientStream
}

type numberManipulationComputeAverageClient struct {
	grpc.ClientStream
}

func (x *numberManipulationComputeAverageClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *numberManipulationComputeAverageClient) CloseAndRecv() (*Number, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *numberManipulationClient) FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (NumberManipulation_FindMaxNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &NumberManipulation_ServiceDesc.Streams[2], "/proto.NumberManipulation/FindMaxNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &numberManipulationFindMaxNumberClient{stream}
	return x, nil
}

type NumberManipulation_FindMaxNumberClient interface {
	Send(*Number) error
	Recv() (*Number, error)
	grpc.ClientStream
}

type numberManipulationFindMaxNumberClient struct {
	grpc.ClientStream
}

func (x *numberManipulationFindMaxNumberClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *numberManipulationFindMaxNumberClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NumberManipulationServer is the server API for NumberManipulation service.
// All implementations must embed UnimplementedNumberManipulationServer
// for forward compatibility
type NumberManipulationServer interface {
	Add(context.Context, *Numbers) (*Number, error)
	PrimeNumbers(*Number, NumberManipulation_PrimeNumbersServer) error
	ComputeAverage(NumberManipulation_ComputeAverageServer) error
	FindMaxNumber(NumberManipulation_FindMaxNumberServer) error
	mustEmbedUnimplementedNumberManipulationServer()
}

// UnimplementedNumberManipulationServer must be embedded to have forward compatible implementations.
type UnimplementedNumberManipulationServer struct {
}

func (UnimplementedNumberManipulationServer) Add(context.Context, *Numbers) (*Number, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedNumberManipulationServer) PrimeNumbers(*Number, NumberManipulation_PrimeNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumbers not implemented")
}
func (UnimplementedNumberManipulationServer) ComputeAverage(NumberManipulation_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (UnimplementedNumberManipulationServer) FindMaxNumber(NumberManipulation_FindMaxNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaxNumber not implemented")
}
func (UnimplementedNumberManipulationServer) mustEmbedUnimplementedNumberManipulationServer() {}

// UnsafeNumberManipulationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumberManipulationServer will
// result in compilation errors.
type UnsafeNumberManipulationServer interface {
	mustEmbedUnimplementedNumberManipulationServer()
}

func RegisterNumberManipulationServer(s grpc.ServiceRegistrar, srv NumberManipulationServer) {
	s.RegisterService(&NumberManipulation_ServiceDesc, srv)
}

func _NumberManipulation_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberManipulationServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NumberManipulation/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberManipulationServer).Add(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _NumberManipulation_PrimeNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Number)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NumberManipulationServer).PrimeNumbers(m, &numberManipulationPrimeNumbersServer{stream})
}

type NumberManipulation_PrimeNumbersServer interface {
	Send(*Number) error
	grpc.ServerStream
}

type numberManipulationPrimeNumbersServer struct {
	grpc.ServerStream
}

func (x *numberManipulationPrimeNumbersServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func _NumberManipulation_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NumberManipulationServer).ComputeAverage(&numberManipulationComputeAverageServer{stream})
}

type NumberManipulation_ComputeAverageServer interface {
	SendAndClose(*Number) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type numberManipulationComputeAverageServer struct {
	grpc.ServerStream
}

func (x *numberManipulationComputeAverageServer) SendAndClose(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func (x *numberManipulationComputeAverageServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NumberManipulation_FindMaxNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NumberManipulationServer).FindMaxNumber(&numberManipulationFindMaxNumberServer{stream})
}

type NumberManipulation_FindMaxNumberServer interface {
	Send(*Number) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type numberManipulationFindMaxNumberServer struct {
	grpc.ServerStream
}

func (x *numberManipulationFindMaxNumberServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func (x *numberManipulationFindMaxNumberServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NumberManipulation_ServiceDesc is the grpc.ServiceDesc for NumberManipulation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NumberManipulation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NumberManipulation",
	HandlerType: (*NumberManipulationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _NumberManipulation_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumbers",
			Handler:       _NumberManipulation_PrimeNumbers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _NumberManipulation_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaxNumber",
			Handler:       _NumberManipulation_FindMaxNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/user.proto",
}
