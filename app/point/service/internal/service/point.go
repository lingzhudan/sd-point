package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"sd-point/api/point/service/v1"
	"sd-point/app/point/service/internal/biz"
)

type PointService struct {
	v1.UnimplementedPointServer

	uc  *biz.PointUseCase
	log *log.Helper
}

func NewPointService(pc *biz.PointUseCase, logger log.Logger) *PointService {
	return &PointService{
		uc:  pc,
		log: log.NewHelper(logger)}
}

// 点数方法

func (s *PointService) CreatePoint(ctx context.Context, req *v1.CreatePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.uc.CreatePoint(ctx, req.Uid, req.Name, req.Desc); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) UpdatePoint(ctx context.Context, req *v1.UpdatePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.uc.UpdatePoint(ctx, req.Pid, req.Name, req.Desc); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeletePoint(ctx context.Context, req *v1.DeletePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.uc.DeletePoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) GetPoint(ctx context.Context, req *v1.GetPointRequest) (rep *v1.GetPointReply, err error) {
	rep = new(v1.GetPointReply)
	var point *biz.Point
	if point, err = s.uc.GetPoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	rep.Point = &v1.GetPointReply_Point{
		Pid:       point.PID,
		Name:      point.Name,
		Total:     point.Total,
		Desc:      point.Desc,
		CreatedAt: uint64(point.CreatedAt.Unix()),
		UpdatedAt: uint64(point.UpdatedAt.Unix()),
	}
	return
}

func (s *PointService) ListPoint(ctx context.Context, req *v1.ListPointRequest) (rep *v1.ListPointReply, err error) {
	rep = &v1.ListPointReply{Finished: true}
	var points []*biz.Point
	if points, err = s.uc.ListPoint(ctx, int(req.Begin), int(req.Count+1), req.Uid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	for i, p := range points {
		if i >= int(req.Count) {
			rep.Finished = false
			break
		}
		rep.Points = append(rep.Points, &v1.GetPointReply_Point{
			Pid:       p.PID,
			Name:      p.Name,
			Total:     p.Total,
			Desc:      p.Desc,
			CreatedAt: uint64(p.CreatedAt.Unix()),
			UpdatedAt: uint64(p.UpdatedAt.Unix()),
		})
	}
	rep.Count = uint32(len(rep.Points))
	return
}
