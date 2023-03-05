// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.1
// source: v1/sd-point_interface.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSdPointInterfaceBindAccount = "/api.sd_point.interface.v1.SdPointInterface/BindAccount"
const OperationSdPointInterfaceCreatePoint = "/api.sd_point.interface.v1.SdPointInterface/CreatePoint"
const OperationSdPointInterfaceCreateRecord = "/api.sd_point.interface.v1.SdPointInterface/CreateRecord"
const OperationSdPointInterfaceDeletePoint = "/api.sd_point.interface.v1.SdPointInterface/DeletePoint"
const OperationSdPointInterfaceDeleteRecord = "/api.sd_point.interface.v1.SdPointInterface/DeleteRecord"
const OperationSdPointInterfaceGetPoint = "/api.sd_point.interface.v1.SdPointInterface/GetPoint"
const OperationSdPointInterfaceGetPublicKey = "/api.sd_point.interface.v1.SdPointInterface/GetPublicKey"
const OperationSdPointInterfaceGetUser = "/api.sd_point.interface.v1.SdPointInterface/GetUser"
const OperationSdPointInterfaceListPoint = "/api.sd_point.interface.v1.SdPointInterface/ListPoint"
const OperationSdPointInterfaceListRecord = "/api.sd_point.interface.v1.SdPointInterface/ListRecord"
const OperationSdPointInterfaceListUser = "/api.sd_point.interface.v1.SdPointInterface/ListUser"
const OperationSdPointInterfaceLogin = "/api.sd_point.interface.v1.SdPointInterface/Login"
const OperationSdPointInterfaceLogout = "/api.sd_point.interface.v1.SdPointInterface/Logout"
const OperationSdPointInterfaceRegister = "/api.sd_point.interface.v1.SdPointInterface/Register"
const OperationSdPointInterfaceUpdatePoint = "/api.sd_point.interface.v1.SdPointInterface/UpdatePoint"
const OperationSdPointInterfaceUpdateRecord = "/api.sd_point.interface.v1.SdPointInterface/UpdateRecord"

type SdPointInterfaceHTTPServer interface {
	BindAccount(context.Context, *BindAccountRequest) (*emptypb.Empty, error)
	CreatePoint(context.Context, *CreatePointRequest) (*emptypb.Empty, error)
	CreateRecord(context.Context, *CreateRecordRequest) (*emptypb.Empty, error)
	DeletePoint(context.Context, *DeletePointRequest) (*emptypb.Empty, error)
	DeleteRecord(context.Context, *DeleteRecordRequest) (*emptypb.Empty, error)
	GetPoint(context.Context, *GetPointRequest) (*GetPointReply, error)
	GetPublicKey(context.Context, *emptypb.Empty) (*GetPublicKeyReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	ListPoint(context.Context, *ListPointRequest) (*ListPointReply, error)
	ListRecord(context.Context, *ListRecordRequest) (*ListRecordReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	UpdatePoint(context.Context, *UpdatePointRequest) (*emptypb.Empty, error)
	UpdateRecord(context.Context, *UpdateRecordRequest) (*emptypb.Empty, error)
}

func RegisterSdPointInterfaceHTTPServer(s *http.Server, srv SdPointInterfaceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/point", _SdPointInterface_CreatePoint0_HTTP_Handler(srv))
	r.PUT("/v1/point/{pid}", _SdPointInterface_UpdatePoint0_HTTP_Handler(srv))
	r.DELETE("/v1/point/{pid}", _SdPointInterface_DeletePoint0_HTTP_Handler(srv))
	r.GET("/v1/point/{pid}", _SdPointInterface_GetPoint0_HTTP_Handler(srv))
	r.GET("/v1/point", _SdPointInterface_ListPoint0_HTTP_Handler(srv))
	r.POST("/v1/record", _SdPointInterface_CreateRecord0_HTTP_Handler(srv))
	r.DELETE("/v1/record/{rid}", _SdPointInterface_DeleteRecord0_HTTP_Handler(srv))
	r.PUT("/v1/record/{rid}", _SdPointInterface_UpdateRecord0_HTTP_Handler(srv))
	r.GET("/v1/record", _SdPointInterface_ListRecord0_HTTP_Handler(srv))
	r.GET("/v1/user/public_key", _SdPointInterface_GetPublicKey0_HTTP_Handler(srv))
	r.POST("/v1/user/login/{login_type}", _SdPointInterface_Login0_HTTP_Handler(srv))
	r.POST("/v1/user/logout", _SdPointInterface_Logout0_HTTP_Handler(srv))
	r.POST("/v1/user/register/{register_type}", _SdPointInterface_Register0_HTTP_Handler(srv))
	r.POST("/v1/user/bind/{bind_type}", _SdPointInterface_BindAccount0_HTTP_Handler(srv))
	r.GET("/v1/user/{uid}", _SdPointInterface_GetUser0_HTTP_Handler(srv))
	r.GET("/v1/user/list", _SdPointInterface_ListUser0_HTTP_Handler(srv))
}

func _SdPointInterface_CreatePoint0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceCreatePoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePoint(ctx, req.(*CreatePointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_UpdatePoint0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceUpdatePoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePoint(ctx, req.(*UpdatePointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_DeletePoint0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceDeletePoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePoint(ctx, req.(*DeletePointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_GetPoint0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceGetPoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPoint(ctx, req.(*GetPointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPointReply)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_ListPoint0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceListPoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPoint(ctx, req.(*ListPointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPointReply)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_CreateRecord0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRecordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceCreateRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRecord(ctx, req.(*CreateRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_DeleteRecord0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRecordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceDeleteRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRecord(ctx, req.(*DeleteRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_UpdateRecord0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRecordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceUpdateRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateRecord(ctx, req.(*UpdateRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_ListRecord0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRecordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceListRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRecord(ctx, req.(*ListRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRecordReply)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_GetPublicKey0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceGetPublicKey)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPublicKey(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPublicKeyReply)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_Login0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_Logout0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_Register0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_BindAccount0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BindAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceBindAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BindAccount(ctx, req.(*BindAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_GetUser0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _SdPointInterface_ListUser0_HTTP_Handler(srv SdPointInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSdPointInterfaceListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

type SdPointInterfaceHTTPClient interface {
	BindAccount(ctx context.Context, req *BindAccountRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CreatePoint(ctx context.Context, req *CreatePointRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CreateRecord(ctx context.Context, req *CreateRecordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeletePoint(ctx context.Context, req *DeletePointRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteRecord(ctx context.Context, req *DeleteRecordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetPoint(ctx context.Context, req *GetPointRequest, opts ...http.CallOption) (rsp *GetPointReply, err error)
	GetPublicKey(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetPublicKeyReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	ListPoint(ctx context.Context, req *ListPointRequest, opts ...http.CallOption) (rsp *ListPointReply, err error)
	ListRecord(ctx context.Context, req *ListRecordRequest, opts ...http.CallOption) (rsp *ListRecordReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	UpdatePoint(ctx context.Context, req *UpdatePointRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdateRecord(ctx context.Context, req *UpdateRecordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type SdPointInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewSdPointInterfaceHTTPClient(client *http.Client) SdPointInterfaceHTTPClient {
	return &SdPointInterfaceHTTPClientImpl{client}
}

func (c *SdPointInterfaceHTTPClientImpl) BindAccount(ctx context.Context, in *BindAccountRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/user/bind/{bind_type}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceBindAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) CreatePoint(ctx context.Context, in *CreatePointRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/point"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceCreatePoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/record"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceCreateRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) DeletePoint(ctx context.Context, in *DeletePointRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/point/{pid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceDeletePoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/record/{rid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceDeleteRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) GetPoint(ctx context.Context, in *GetPointRequest, opts ...http.CallOption) (*GetPointReply, error) {
	var out GetPointReply
	pattern := "/v1/point/{pid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceGetPoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) GetPublicKey(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetPublicKeyReply, error) {
	var out GetPublicKeyReply
	pattern := "/v1/user/public_key"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceGetPublicKey))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/v1/user/{uid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) ListPoint(ctx context.Context, in *ListPointRequest, opts ...http.CallOption) (*ListPointReply, error) {
	var out ListPointReply
	pattern := "/v1/point"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceListPoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) ListRecord(ctx context.Context, in *ListRecordRequest, opts ...http.CallOption) (*ListRecordReply, error) {
	var out ListRecordReply
	pattern := "/v1/record"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceListRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/v1/user/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSdPointInterfaceListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/user/login/{login_type}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) Logout(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/user/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/user/register/{register_type}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) UpdatePoint(ctx context.Context, in *UpdatePointRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/point/{pid}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceUpdatePoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SdPointInterfaceHTTPClientImpl) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/record/{rid}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSdPointInterfaceUpdateRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
