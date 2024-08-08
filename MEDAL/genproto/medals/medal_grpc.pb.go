// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: medal-service/medal.proto

package medals

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

// MedalServiceClient is the client API for MedalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedalServiceClient interface {
	CreateMedal(ctx context.Context, in *CreateMedalRequest, opts ...grpc.CallOption) (*Medal, error)
	GetMedal(ctx context.Context, in *GetMedalRequest, opts ...grpc.CallOption) (*GetMedalResponse, error)
	GetMedals(ctx context.Context, in *GetMedalsRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error)
	UpdateMedal(ctx context.Context, in *UpdateMedalRequest, opts ...grpc.CallOption) (*UpdateMedalResponse, error)
	DeleteMedal(ctx context.Context, in *DeleteMedalRequest, opts ...grpc.CallOption) (*DeleteMedalResponse, error)
	GetMedalsByCountry(ctx context.Context, in *GetMedalsByCountryRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error)
	GetMedalsByAthlete(ctx context.Context, in *GetMedalsByAthleteRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error)
	GetMedalsByTimeRange(ctx context.Context, in *GetMedalsByTimeRangeRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error)
	RankingByCountry(ctx context.Context, in *GetRankingByCountryRequest, opts ...grpc.CallOption) (*GetRankingResponse, error)
}

type medalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedalServiceClient(cc grpc.ClientConnInterface) MedalServiceClient {
	return &medalServiceClient{cc}
}

func (c *medalServiceClient) CreateMedal(ctx context.Context, in *CreateMedalRequest, opts ...grpc.CallOption) (*Medal, error) {
	out := new(Medal)
	err := c.cc.Invoke(ctx, "/medal.MedalService/CreateMedal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) GetMedal(ctx context.Context, in *GetMedalRequest, opts ...grpc.CallOption) (*GetMedalResponse, error) {
	out := new(GetMedalResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/GetMedal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) GetMedals(ctx context.Context, in *GetMedalsRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error) {
	out := new(GetMedalsResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/GetMedals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) UpdateMedal(ctx context.Context, in *UpdateMedalRequest, opts ...grpc.CallOption) (*UpdateMedalResponse, error) {
	out := new(UpdateMedalResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/UpdateMedal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) DeleteMedal(ctx context.Context, in *DeleteMedalRequest, opts ...grpc.CallOption) (*DeleteMedalResponse, error) {
	out := new(DeleteMedalResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/DeleteMedal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) GetMedalsByCountry(ctx context.Context, in *GetMedalsByCountryRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error) {
	out := new(GetMedalsResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/GetMedalsByCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) GetMedalsByAthlete(ctx context.Context, in *GetMedalsByAthleteRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error) {
	out := new(GetMedalsResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/GetMedalsByAthlete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) GetMedalsByTimeRange(ctx context.Context, in *GetMedalsByTimeRangeRequest, opts ...grpc.CallOption) (*GetMedalsResponse, error) {
	out := new(GetMedalsResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/GetMedalsByTimeRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medalServiceClient) RankingByCountry(ctx context.Context, in *GetRankingByCountryRequest, opts ...grpc.CallOption) (*GetRankingResponse, error) {
	out := new(GetRankingResponse)
	err := c.cc.Invoke(ctx, "/medal.MedalService/RankingByCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedalServiceServer is the server API for MedalService service.
// All implementations must embed UnimplementedMedalServiceServer
// for forward compatibility
type MedalServiceServer interface {
	CreateMedal(context.Context, *CreateMedalRequest) (*Medal, error)
	GetMedal(context.Context, *GetMedalRequest) (*GetMedalResponse, error)
	GetMedals(context.Context, *GetMedalsRequest) (*GetMedalsResponse, error)
	UpdateMedal(context.Context, *UpdateMedalRequest) (*UpdateMedalResponse, error)
	DeleteMedal(context.Context, *DeleteMedalRequest) (*DeleteMedalResponse, error)
	GetMedalsByCountry(context.Context, *GetMedalsByCountryRequest) (*GetMedalsResponse, error)
	GetMedalsByAthlete(context.Context, *GetMedalsByAthleteRequest) (*GetMedalsResponse, error)
	GetMedalsByTimeRange(context.Context, *GetMedalsByTimeRangeRequest) (*GetMedalsResponse, error)
	RankingByCountry(context.Context, *GetRankingByCountryRequest) (*GetRankingResponse, error)
	mustEmbedUnimplementedMedalServiceServer()
}

// UnimplementedMedalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMedalServiceServer struct {
}

func (UnimplementedMedalServiceServer) CreateMedal(context.Context, *CreateMedalRequest) (*Medal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMedal not implemented")
}
func (UnimplementedMedalServiceServer) GetMedal(context.Context, *GetMedalRequest) (*GetMedalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedal not implemented")
}
func (UnimplementedMedalServiceServer) GetMedals(context.Context, *GetMedalsRequest) (*GetMedalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedals not implemented")
}
func (UnimplementedMedalServiceServer) UpdateMedal(context.Context, *UpdateMedalRequest) (*UpdateMedalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMedal not implemented")
}
func (UnimplementedMedalServiceServer) DeleteMedal(context.Context, *DeleteMedalRequest) (*DeleteMedalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedal not implemented")
}
func (UnimplementedMedalServiceServer) GetMedalsByCountry(context.Context, *GetMedalsByCountryRequest) (*GetMedalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedalsByCountry not implemented")
}
func (UnimplementedMedalServiceServer) GetMedalsByAthlete(context.Context, *GetMedalsByAthleteRequest) (*GetMedalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedalsByAthlete not implemented")
}
func (UnimplementedMedalServiceServer) GetMedalsByTimeRange(context.Context, *GetMedalsByTimeRangeRequest) (*GetMedalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedalsByTimeRange not implemented")
}
func (UnimplementedMedalServiceServer) RankingByCountry(context.Context, *GetRankingByCountryRequest) (*GetRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RankingByCountry not implemented")
}
func (UnimplementedMedalServiceServer) mustEmbedUnimplementedMedalServiceServer() {}

// UnsafeMedalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedalServiceServer will
// result in compilation errors.
type UnsafeMedalServiceServer interface {
	mustEmbedUnimplementedMedalServiceServer()
}

func RegisterMedalServiceServer(s grpc.ServiceRegistrar, srv MedalServiceServer) {
	s.RegisterService(&MedalService_ServiceDesc, srv)
}

func _MedalService_CreateMedal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMedalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).CreateMedal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/CreateMedal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).CreateMedal(ctx, req.(*CreateMedalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_GetMedal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).GetMedal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/GetMedal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).GetMedal(ctx, req.(*GetMedalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_GetMedals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).GetMedals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/GetMedals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).GetMedals(ctx, req.(*GetMedalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_UpdateMedal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMedalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).UpdateMedal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/UpdateMedal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).UpdateMedal(ctx, req.(*UpdateMedalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_DeleteMedal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMedalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).DeleteMedal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/DeleteMedal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).DeleteMedal(ctx, req.(*DeleteMedalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_GetMedalsByCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedalsByCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).GetMedalsByCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/GetMedalsByCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).GetMedalsByCountry(ctx, req.(*GetMedalsByCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_GetMedalsByAthlete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedalsByAthleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).GetMedalsByAthlete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/GetMedalsByAthlete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).GetMedalsByAthlete(ctx, req.(*GetMedalsByAthleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_GetMedalsByTimeRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedalsByTimeRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).GetMedalsByTimeRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/GetMedalsByTimeRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).GetMedalsByTimeRange(ctx, req.(*GetMedalsByTimeRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedalService_RankingByCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRankingByCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).RankingByCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medal.MedalService/RankingByCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).RankingByCountry(ctx, req.(*GetRankingByCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MedalService_ServiceDesc is the grpc.ServiceDesc for MedalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medal.MedalService",
	HandlerType: (*MedalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMedal",
			Handler:    _MedalService_CreateMedal_Handler,
		},
		{
			MethodName: "GetMedal",
			Handler:    _MedalService_GetMedal_Handler,
		},
		{
			MethodName: "GetMedals",
			Handler:    _MedalService_GetMedals_Handler,
		},
		{
			MethodName: "UpdateMedal",
			Handler:    _MedalService_UpdateMedal_Handler,
		},
		{
			MethodName: "DeleteMedal",
			Handler:    _MedalService_DeleteMedal_Handler,
		},
		{
			MethodName: "GetMedalsByCountry",
			Handler:    _MedalService_GetMedalsByCountry_Handler,
		},
		{
			MethodName: "GetMedalsByAthlete",
			Handler:    _MedalService_GetMedalsByAthlete_Handler,
		},
		{
			MethodName: "GetMedalsByTimeRange",
			Handler:    _MedalService_GetMedalsByTimeRange_Handler,
		},
		{
			MethodName: "RankingByCountry",
			Handler:    _MedalService_RankingByCountry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medal-service/medal.proto",
}
