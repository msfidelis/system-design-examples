// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: imc.proto

package imc

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
	IMCService_Calcular_FullMethodName = "/imc.IMCService/Calcular"
)

// IMCServiceClient is the client API for IMCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IMCServiceClient interface {
	// IMC calcula o quadrado de um número.
	Calcular(ctx context.Context, in *IMCRequest, opts ...grpc.CallOption) (*IMCResponse, error)
}

type iMCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIMCServiceClient(cc grpc.ClientConnInterface) IMCServiceClient {
	return &iMCServiceClient{cc}
}

func (c *iMCServiceClient) Calcular(ctx context.Context, in *IMCRequest, opts ...grpc.CallOption) (*IMCResponse, error) {
	out := new(IMCResponse)
	err := c.cc.Invoke(ctx, IMCService_Calcular_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IMCServiceServer is the server API for IMCService service.
// All implementations must embed UnimplementedIMCServiceServer
// for forward compatibility
type IMCServiceServer interface {
	// IMC calcula o quadrado de um número.
	Calcular(context.Context, *IMCRequest) (*IMCResponse, error)
	mustEmbedUnimplementedIMCServiceServer()
}

// UnimplementedIMCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIMCServiceServer struct {
}

func (UnimplementedIMCServiceServer) Calcular(context.Context, *IMCRequest) (*IMCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calcular not implemented")
}
func (UnimplementedIMCServiceServer) mustEmbedUnimplementedIMCServiceServer() {}

// UnsafeIMCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IMCServiceServer will
// result in compilation errors.
type UnsafeIMCServiceServer interface {
	mustEmbedUnimplementedIMCServiceServer()
}

func RegisterIMCServiceServer(s grpc.ServiceRegistrar, srv IMCServiceServer) {
	s.RegisterService(&IMCService_ServiceDesc, srv)
}

func _IMCService_Calcular_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IMCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMCServiceServer).Calcular(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IMCService_Calcular_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMCServiceServer).Calcular(ctx, req.(*IMCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IMCService_ServiceDesc is the grpc.ServiceDesc for IMCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IMCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "imc.IMCService",
	HandlerType: (*IMCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calcular",
			Handler:    _IMCService_Calcular_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imc.proto",
}
