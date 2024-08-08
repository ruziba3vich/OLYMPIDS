// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: stats/stat.proto

package stats

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	StatsService_CreateTeamStats_FullMethodName       = "/StatsService/CreateTeamStats"
	StatsService_CreatePlayerOnlyStats_FullMethodName = "/StatsService/CreatePlayerOnlyStats"
	StatsService_CreateRaceStats_FullMethodName       = "/StatsService/CreateRaceStats"
	StatsService_UpdateTeamStats_FullMethodName       = "/StatsService/UpdateTeamStats"
	StatsService_GetStatsByEventId_FullMethodName     = "/StatsService/GetStatsByEventId"
)

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsServiceClient interface {
	CreateTeamStats(ctx context.Context, in *TeamEvent, opts ...grpc.CallOption) (*TeamEvent, error)
	CreatePlayerOnlyStats(ctx context.Context, in *PlayerOnly, opts ...grpc.CallOption) (*PlayerOnly, error)
	CreateRaceStats(ctx context.Context, in *Race, opts ...grpc.CallOption) (*Race, error)
	UpdateTeamStats(ctx context.Context, in *UpdateTeamStatsRequest, opts ...grpc.CallOption) (*Team, error)
	GetStatsByEventId(ctx context.Context, in *GetStatsByEventIdRequest, opts ...grpc.CallOption) (StatsService_GetStatsByEventIdClient, error)
}

type statsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsServiceClient(cc grpc.ClientConnInterface) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) CreateTeamStats(ctx context.Context, in *TeamEvent, opts ...grpc.CallOption) (*TeamEvent, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeamEvent)
	err := c.cc.Invoke(ctx, StatsService_CreateTeamStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) CreatePlayerOnlyStats(ctx context.Context, in *PlayerOnly, opts ...grpc.CallOption) (*PlayerOnly, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerOnly)
	err := c.cc.Invoke(ctx, StatsService_CreatePlayerOnlyStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) CreateRaceStats(ctx context.Context, in *Race, opts ...grpc.CallOption) (*Race, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Race)
	err := c.cc.Invoke(ctx, StatsService_CreateRaceStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) UpdateTeamStats(ctx context.Context, in *UpdateTeamStatsRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, StatsService_UpdateTeamStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetStatsByEventId(ctx context.Context, in *GetStatsByEventIdRequest, opts ...grpc.CallOption) (StatsService_GetStatsByEventIdClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StatsService_ServiceDesc.Streams[0], StatsService_GetStatsByEventId_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &statsServiceGetStatsByEventIdClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StatsService_GetStatsByEventIdClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type statsServiceGetStatsByEventIdClient struct {
	grpc.ClientStream
}

func (x *statsServiceGetStatsByEventIdClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatsServiceServer is the server API for StatsService service.
// All implementations must embed UnimplementedStatsServiceServer
// for forward compatibility
type StatsServiceServer interface {
	CreateTeamStats(context.Context, *TeamEvent) (*TeamEvent, error)
	CreatePlayerOnlyStats(context.Context, *PlayerOnly) (*PlayerOnly, error)
	CreateRaceStats(context.Context, *Race) (*Race, error)
	UpdateTeamStats(context.Context, *UpdateTeamStatsRequest) (*Team, error)
	GetStatsByEventId(*GetStatsByEventIdRequest, StatsService_GetStatsByEventIdServer) error
	mustEmbedUnimplementedStatsServiceServer()
}

// UnimplementedStatsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatsServiceServer struct {
}

func (UnimplementedStatsServiceServer) CreateTeamStats(context.Context, *TeamEvent) (*TeamEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeamStats not implemented")
}
func (UnimplementedStatsServiceServer) CreatePlayerOnlyStats(context.Context, *PlayerOnly) (*PlayerOnly, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayerOnlyStats not implemented")
}
func (UnimplementedStatsServiceServer) CreateRaceStats(context.Context, *Race) (*Race, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRaceStats not implemented")
}
func (UnimplementedStatsServiceServer) UpdateTeamStats(context.Context, *UpdateTeamStatsRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeamStats not implemented")
}
func (UnimplementedStatsServiceServer) GetStatsByEventId(*GetStatsByEventIdRequest, StatsService_GetStatsByEventIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatsByEventId not implemented")
}
func (UnimplementedStatsServiceServer) mustEmbedUnimplementedStatsServiceServer() {}

// UnsafeStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServiceServer will
// result in compilation errors.
type UnsafeStatsServiceServer interface {
	mustEmbedUnimplementedStatsServiceServer()
}

func RegisterStatsServiceServer(s grpc.ServiceRegistrar, srv StatsServiceServer) {
	s.RegisterService(&StatsService_ServiceDesc, srv)
}

func _StatsService_CreateTeamStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).CreateTeamStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_CreateTeamStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).CreateTeamStats(ctx, req.(*TeamEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_CreatePlayerOnlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerOnly)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).CreatePlayerOnlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_CreatePlayerOnlyStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).CreatePlayerOnlyStats(ctx, req.(*PlayerOnly))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_CreateRaceStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Race)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).CreateRaceStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_CreateRaceStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).CreateRaceStats(ctx, req.(*Race))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_UpdateTeamStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).UpdateTeamStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StatsService_UpdateTeamStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).UpdateTeamStats(ctx, req.(*UpdateTeamStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetStatsByEventId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStatsByEventIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatsServiceServer).GetStatsByEventId(m, &statsServiceGetStatsByEventIdServer{ServerStream: stream})
}

type StatsService_GetStatsByEventIdServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type statsServiceGetStatsByEventIdServer struct {
	grpc.ServerStream
}

func (x *statsServiceGetStatsByEventIdServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// StatsService_ServiceDesc is the grpc.ServiceDesc for StatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeamStats",
			Handler:    _StatsService_CreateTeamStats_Handler,
		},
		{
			MethodName: "CreatePlayerOnlyStats",
			Handler:    _StatsService_CreatePlayerOnlyStats_Handler,
		},
		{
			MethodName: "CreateRaceStats",
			Handler:    _StatsService_CreateRaceStats_Handler,
		},
		{
			MethodName: "UpdateTeamStats",
			Handler:    _StatsService_UpdateTeamStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatsByEventId",
			Handler:       _StatsService_GetStatsByEventId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stats/stat.proto",
}
