// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: v1/sd-point_interface.proto

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

// SdPointInterfaceClient is the client API for SdPointInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdPointInterfaceClient interface {
	// 创建点数
	CreatePoint(ctx context.Context, in *CreatePointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新点数
	UpdatePoint(ctx context.Context, in *UpdatePointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除点数
	DeletePoint(ctx context.Context, in *DeletePointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取点数
	GetPoint(ctx context.Context, in *GetPointRequest, opts ...grpc.CallOption) (*GetPointReply, error)
	// 获取点数列表
	ListPoint(ctx context.Context, in *ListPointRequest, opts ...grpc.CallOption) (*ListPointReply, error)
	// 创建记录
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除记录
	DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新记录
	UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取记录列表
	ListRecord(ctx context.Context, in *ListRecordRequest, opts ...grpc.CallOption) (*ListRecordReply, error)
	// 获取公钥
	GetPublicKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPublicKeyReply, error)
	// 用户登录
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 用户登出
	Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 用户注册
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	// 用户绑定其他登录方式
	BindAccount(ctx context.Context, in *BindAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 用户信息
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	// 用户列表
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
}

type sdPointInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewSdPointInterfaceClient(cc grpc.ClientConnInterface) SdPointInterfaceClient {
	return &sdPointInterfaceClient{cc}
}

func (c *sdPointInterfaceClient) CreatePoint(ctx context.Context, in *CreatePointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/CreatePoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) UpdatePoint(ctx context.Context, in *UpdatePointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/UpdatePoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) DeletePoint(ctx context.Context, in *DeletePointRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/DeletePoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) GetPoint(ctx context.Context, in *GetPointRequest, opts ...grpc.CallOption) (*GetPointReply, error) {
	out := new(GetPointReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/GetPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) ListPoint(ctx context.Context, in *ListPointRequest, opts ...grpc.CallOption) (*ListPointReply, error) {
	out := new(ListPointReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/ListPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/DeleteRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/UpdateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) ListRecord(ctx context.Context, in *ListRecordRequest, opts ...grpc.CallOption) (*ListRecordReply, error) {
	out := new(ListRecordReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/ListRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) GetPublicKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetPublicKeyReply, error) {
	out := new(GetPublicKeyReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) Logout(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) BindAccount(ctx context.Context, in *BindAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/BindAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sdPointInterfaceClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, "/api.sd_point.interface.v1.SdPointInterface/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdPointInterfaceServer is the server API for SdPointInterface service.
// All implementations must embed UnimplementedSdPointInterfaceServer
// for forward compatibility
type SdPointInterfaceServer interface {
	// 创建点数
	CreatePoint(context.Context, *CreatePointRequest) (*emptypb.Empty, error)
	// 更新点数
	UpdatePoint(context.Context, *UpdatePointRequest) (*emptypb.Empty, error)
	// 删除点数
	DeletePoint(context.Context, *DeletePointRequest) (*emptypb.Empty, error)
	// 获取点数
	GetPoint(context.Context, *GetPointRequest) (*GetPointReply, error)
	// 获取点数列表
	ListPoint(context.Context, *ListPointRequest) (*ListPointReply, error)
	// 创建记录
	CreateRecord(context.Context, *CreateRecordRequest) (*emptypb.Empty, error)
	// 删除记录
	DeleteRecord(context.Context, *DeleteRecordRequest) (*emptypb.Empty, error)
	// 更新记录
	UpdateRecord(context.Context, *UpdateRecordRequest) (*emptypb.Empty, error)
	// 获取记录列表
	ListRecord(context.Context, *ListRecordRequest) (*ListRecordReply, error)
	// 获取公钥
	GetPublicKey(context.Context, *emptypb.Empty) (*GetPublicKeyReply, error)
	// 用户登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// 用户登出
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// 用户注册
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	// 用户绑定其他登录方式
	BindAccount(context.Context, *BindAccountRequest) (*emptypb.Empty, error)
	// 用户信息
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	// 用户列表
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	mustEmbedUnimplementedSdPointInterfaceServer()
}

// UnimplementedSdPointInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedSdPointInterfaceServer struct {
}

func (UnimplementedSdPointInterfaceServer) CreatePoint(context.Context, *CreatePointRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePoint not implemented")
}
func (UnimplementedSdPointInterfaceServer) UpdatePoint(context.Context, *UpdatePointRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePoint not implemented")
}
func (UnimplementedSdPointInterfaceServer) DeletePoint(context.Context, *DeletePointRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePoint not implemented")
}
func (UnimplementedSdPointInterfaceServer) GetPoint(context.Context, *GetPointRequest) (*GetPointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoint not implemented")
}
func (UnimplementedSdPointInterfaceServer) ListPoint(context.Context, *ListPointRequest) (*ListPointReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPoint not implemented")
}
func (UnimplementedSdPointInterfaceServer) CreateRecord(context.Context, *CreateRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedSdPointInterfaceServer) DeleteRecord(context.Context, *DeleteRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (UnimplementedSdPointInterfaceServer) UpdateRecord(context.Context, *UpdateRecordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}
func (UnimplementedSdPointInterfaceServer) ListRecord(context.Context, *ListRecordRequest) (*ListRecordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecord not implemented")
}
func (UnimplementedSdPointInterfaceServer) GetPublicKey(context.Context, *emptypb.Empty) (*GetPublicKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedSdPointInterfaceServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSdPointInterfaceServer) Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSdPointInterfaceServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedSdPointInterfaceServer) BindAccount(context.Context, *BindAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindAccount not implemented")
}
func (UnimplementedSdPointInterfaceServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedSdPointInterfaceServer) ListUser(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedSdPointInterfaceServer) mustEmbedUnimplementedSdPointInterfaceServer() {}

// UnsafeSdPointInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdPointInterfaceServer will
// result in compilation errors.
type UnsafeSdPointInterfaceServer interface {
	mustEmbedUnimplementedSdPointInterfaceServer()
}

func RegisterSdPointInterfaceServer(s grpc.ServiceRegistrar, srv SdPointInterfaceServer) {
	s.RegisterService(&SdPointInterface_ServiceDesc, srv)
}

func _SdPointInterface_CreatePoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).CreatePoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/CreatePoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).CreatePoint(ctx, req.(*CreatePointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_UpdatePoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).UpdatePoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/UpdatePoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).UpdatePoint(ctx, req.(*UpdatePointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_DeletePoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).DeletePoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/DeletePoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).DeletePoint(ctx, req.(*DeletePointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_GetPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).GetPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/GetPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).GetPoint(ctx, req.(*GetPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_ListPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).ListPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/ListPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).ListPoint(ctx, req.(*ListPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).DeleteRecord(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/UpdateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).UpdateRecord(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_ListRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).ListRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/ListRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).ListRecord(ctx, req.(*ListRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).GetPublicKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).Logout(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_BindAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).BindAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/BindAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).BindAccount(ctx, req.(*BindAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SdPointInterface_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdPointInterfaceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sd_point.interface.v1.SdPointInterface/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdPointInterfaceServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SdPointInterface_ServiceDesc is the grpc.ServiceDesc for SdPointInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SdPointInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sd_point.interface.v1.SdPointInterface",
	HandlerType: (*SdPointInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePoint",
			Handler:    _SdPointInterface_CreatePoint_Handler,
		},
		{
			MethodName: "UpdatePoint",
			Handler:    _SdPointInterface_UpdatePoint_Handler,
		},
		{
			MethodName: "DeletePoint",
			Handler:    _SdPointInterface_DeletePoint_Handler,
		},
		{
			MethodName: "GetPoint",
			Handler:    _SdPointInterface_GetPoint_Handler,
		},
		{
			MethodName: "ListPoint",
			Handler:    _SdPointInterface_ListPoint_Handler,
		},
		{
			MethodName: "CreateRecord",
			Handler:    _SdPointInterface_CreateRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _SdPointInterface_DeleteRecord_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _SdPointInterface_UpdateRecord_Handler,
		},
		{
			MethodName: "ListRecord",
			Handler:    _SdPointInterface_ListRecord_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _SdPointInterface_GetPublicKey_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SdPointInterface_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SdPointInterface_Logout_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _SdPointInterface_Register_Handler,
		},
		{
			MethodName: "BindAccount",
			Handler:    _SdPointInterface_BindAccount_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _SdPointInterface_GetUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _SdPointInterface_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/sd-point_interface.proto",
}
