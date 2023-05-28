// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: partner/service/v1/partner.proto

package v1

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const ()

// PartnerServiceClient is the client API for PartnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartnerServiceClient interface {
}

type partnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerServiceClient(cc grpc.ClientConnInterface) PartnerServiceClient {
	return &partnerServiceClient{cc}
}

// PartnerServiceServer is the server API for PartnerService service.
// All implementations must embed UnimplementedPartnerServiceServer
// for forward compatibility
type PartnerServiceServer interface {
	mustEmbedUnimplementedPartnerServiceServer()
}

// UnimplementedPartnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartnerServiceServer struct {
}

func (UnimplementedPartnerServiceServer) mustEmbedUnimplementedPartnerServiceServer() {}

// UnsafePartnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartnerServiceServer will
// result in compilation errors.
type UnsafePartnerServiceServer interface {
	mustEmbedUnimplementedPartnerServiceServer()
}

func RegisterPartnerServiceServer(s grpc.ServiceRegistrar, srv PartnerServiceServer) {
	s.RegisterService(&PartnerService_ServiceDesc, srv)
}

// PartnerService_ServiceDesc is the grpc.ServiceDesc for PartnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "partner.v1.PartnerService",
	HandlerType: (*PartnerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "partner/service/v1/partner.proto",
}
