// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: partner/service/v1/partner.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PartnerService_AddTeam_FullMethodName     = "/partner.v1.PartnerService/AddTeam"
	PartnerService_DeleteTeam_FullMethodName  = "/partner.v1.PartnerService/DeleteTeam"
	PartnerService_UpdateTeam_FullMethodName  = "/partner.v1.PartnerService/UpdateTeam"
	PartnerService_GetTeam_FullMethodName     = "/partner.v1.PartnerService/GetTeam"
	PartnerService_GetTeamList_FullMethodName = "/partner.v1.PartnerService/GetTeamList"
	PartnerService_JoinTeam_FullMethodName    = "/partner.v1.PartnerService/JoinTeam"
)

// PartnerServiceClient is the client API for PartnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartnerServiceClient interface {
	AddTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTeam(ctx context.Context, in *DeleteTeamReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTeam(ctx context.Context, in *GetTeamReq, opts ...grpc.CallOption) (*GetTeamResponse, error)
	GetTeamList(ctx context.Context, in *GetTeamListReq, opts ...grpc.CallOption) (*GetTeamListResponse, error)
	JoinTeam(ctx context.Context, in *JoinTeamReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type partnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerServiceClient(cc grpc.ClientConnInterface) PartnerServiceClient {
	return &partnerServiceClient{cc}
}

func (c *partnerServiceClient) AddTeam(ctx context.Context, in *Team, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PartnerService_AddTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) DeleteTeam(ctx context.Context, in *DeleteTeamReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PartnerService_DeleteTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) UpdateTeam(ctx context.Context, in *UpdateTeamReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PartnerService_UpdateTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) GetTeam(ctx context.Context, in *GetTeamReq, opts ...grpc.CallOption) (*GetTeamResponse, error) {
	out := new(GetTeamResponse)
	err := c.cc.Invoke(ctx, PartnerService_GetTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) GetTeamList(ctx context.Context, in *GetTeamListReq, opts ...grpc.CallOption) (*GetTeamListResponse, error) {
	out := new(GetTeamListResponse)
	err := c.cc.Invoke(ctx, PartnerService_GetTeamList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) JoinTeam(ctx context.Context, in *JoinTeamReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PartnerService_JoinTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerServiceServer is the server API for PartnerService service.
// All implementations must embed UnimplementedPartnerServiceServer
// for forward compatibility
type PartnerServiceServer interface {
	AddTeam(context.Context, *Team) (*emptypb.Empty, error)
	DeleteTeam(context.Context, *DeleteTeamReq) (*emptypb.Empty, error)
	UpdateTeam(context.Context, *UpdateTeamReq) (*emptypb.Empty, error)
	GetTeam(context.Context, *GetTeamReq) (*GetTeamResponse, error)
	GetTeamList(context.Context, *GetTeamListReq) (*GetTeamListResponse, error)
	JoinTeam(context.Context, *JoinTeamReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedPartnerServiceServer()
}

// UnimplementedPartnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartnerServiceServer struct {
}

func (UnimplementedPartnerServiceServer) AddTeam(context.Context, *Team) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeam not implemented")
}
func (UnimplementedPartnerServiceServer) DeleteTeam(context.Context, *DeleteTeamReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedPartnerServiceServer) UpdateTeam(context.Context, *UpdateTeamReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedPartnerServiceServer) GetTeam(context.Context, *GetTeamReq) (*GetTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (UnimplementedPartnerServiceServer) GetTeamList(context.Context, *GetTeamListReq) (*GetTeamListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamList not implemented")
}
func (UnimplementedPartnerServiceServer) JoinTeam(context.Context, *JoinTeamReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinTeam not implemented")
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

func _PartnerService_AddTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Team)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).AddTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_AddTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).AddTeam(ctx, req.(*Team))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_DeleteTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).DeleteTeam(ctx, req.(*DeleteTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_UpdateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).UpdateTeam(ctx, req.(*UpdateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_GetTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).GetTeam(ctx, req.(*GetTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_GetTeamList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).GetTeamList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_GetTeamList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).GetTeamList(ctx, req.(*GetTeamListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_JoinTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).JoinTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PartnerService_JoinTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).JoinTeam(ctx, req.(*JoinTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PartnerService_ServiceDesc is the grpc.ServiceDesc for PartnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "partner.v1.PartnerService",
	HandlerType: (*PartnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTeam",
			Handler:    _PartnerService_AddTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _PartnerService_DeleteTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _PartnerService_UpdateTeam_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _PartnerService_GetTeam_Handler,
		},
		{
			MethodName: "GetTeamList",
			Handler:    _PartnerService_GetTeamList_Handler,
		},
		{
			MethodName: "JoinTeam",
			Handler:    _PartnerService_JoinTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partner/service/v1/partner.proto",
}
