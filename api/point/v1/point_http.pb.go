// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.1
// source: point/v1/point.proto

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

const OperationPointCreatePoint = "/api.point.v1.Point/CreatePoint"
const OperationPointCreateRecords = "/api.point.v1.Point/CreateRecords"
const OperationPointDeletePoint = "/api.point.v1.Point/DeletePoint"
const OperationPointDeleteRecord = "/api.point.v1.Point/DeleteRecord"
const OperationPointGetPoint = "/api.point.v1.Point/GetPoint"
const OperationPointListPoint = "/api.point.v1.Point/ListPoint"
const OperationPointListRecord = "/api.point.v1.Point/ListRecord"
const OperationPointUpdatePoint = "/api.point.v1.Point/UpdatePoint"

type PointHTTPServer interface {
	CreatePoint(context.Context, *CreatePointRequest) (*emptypb.Empty, error)
	CreateRecords(context.Context, *CreateRecordsRequest) (*emptypb.Empty, error)
	DeletePoint(context.Context, *DeletePointRequest) (*emptypb.Empty, error)
	DeleteRecord(context.Context, *DeleteRecordRequest) (*emptypb.Empty, error)
	GetPoint(context.Context, *GetPointRequest) (*GetPointReply, error)
	ListPoint(context.Context, *ListPointRequest) (*ListPointReply, error)
	ListRecord(context.Context, *ListRecordRequest) (*ListRecordReply, error)
	UpdatePoint(context.Context, *UpdatePointRequest) (*emptypb.Empty, error)
}

func RegisterPointHTTPServer(s *http.Server, srv PointHTTPServer) {
	r := s.Route("/")
	r.PUT("/v1/point", _Point_CreatePoint0_HTTP_Handler(srv))
	r.POST("/v1/point", _Point_UpdatePoint0_HTTP_Handler(srv))
	r.DELETE("/v1/point", _Point_DeletePoint0_HTTP_Handler(srv))
	r.GET("/v1/point", _Point_GetPoint0_HTTP_Handler(srv))
	r.GET("/v1/point/list", _Point_ListPoint0_HTTP_Handler(srv))
	r.PUT("/v1/record", _Point_CreateRecords0_HTTP_Handler(srv))
	r.DELETE("/v1/record", _Point_DeleteRecord0_HTTP_Handler(srv))
	r.GET("/v1/record/list", _Point_ListRecord0_HTTP_Handler(srv))
}

func _Point_CreatePoint0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointCreatePoint)
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

func _Point_UpdatePoint0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointUpdatePoint)
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

func _Point_DeletePoint0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointDeletePoint)
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

func _Point_GetPoint0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointGetPoint)
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

func _Point_ListPoint0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointListPoint)
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

func _Point_CreateRecords0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRecordsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointCreateRecords)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRecords(ctx, req.(*CreateRecordsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Point_DeleteRecord0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRecordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointDeleteRecord)
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

func _Point_ListRecord0_HTTP_Handler(srv PointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRecordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPointListRecord)
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

type PointHTTPClient interface {
	CreatePoint(ctx context.Context, req *CreatePointRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	CreateRecords(ctx context.Context, req *CreateRecordsRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeletePoint(ctx context.Context, req *DeletePointRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteRecord(ctx context.Context, req *DeleteRecordRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetPoint(ctx context.Context, req *GetPointRequest, opts ...http.CallOption) (rsp *GetPointReply, err error)
	ListPoint(ctx context.Context, req *ListPointRequest, opts ...http.CallOption) (rsp *ListPointReply, err error)
	ListRecord(ctx context.Context, req *ListRecordRequest, opts ...http.CallOption) (rsp *ListRecordReply, err error)
	UpdatePoint(ctx context.Context, req *UpdatePointRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type PointHTTPClientImpl struct {
	cc *http.Client
}

func NewPointHTTPClient(client *http.Client) PointHTTPClient {
	return &PointHTTPClientImpl{client}
}

func (c *PointHTTPClientImpl) CreatePoint(ctx context.Context, in *CreatePointRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/point"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPointCreatePoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PointHTTPClientImpl) CreateRecords(ctx context.Context, in *CreateRecordsRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/record"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPointCreateRecords))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PointHTTPClientImpl) DeletePoint(ctx context.Context, in *DeletePointRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/point"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPointDeletePoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PointHTTPClientImpl) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/record"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPointDeleteRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PointHTTPClientImpl) GetPoint(ctx context.Context, in *GetPointRequest, opts ...http.CallOption) (*GetPointReply, error) {
	var out GetPointReply
	pattern := "/v1/point"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPointGetPoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PointHTTPClientImpl) ListPoint(ctx context.Context, in *ListPointRequest, opts ...http.CallOption) (*ListPointReply, error) {
	var out ListPointReply
	pattern := "/v1/point/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPointListPoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PointHTTPClientImpl) ListRecord(ctx context.Context, in *ListRecordRequest, opts ...http.CallOption) (*ListRecordReply, error) {
	var out ListRecordReply
	pattern := "/v1/record/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPointListRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PointHTTPClientImpl) UpdatePoint(ctx context.Context, in *UpdatePointRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/point"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPointUpdatePoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
