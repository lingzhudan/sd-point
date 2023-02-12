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

	pc  *biz.PointUseCase
	rc  *biz.RecordUseCase
	tc  *biz.TotalUseCase
	log *log.Helper
}

func NewPointService(pc *biz.PointUseCase, rc *biz.RecordUseCase,
	tc *biz.TotalUseCase, logger log.Logger) *PointService {
	return &PointService{
		pc:  pc,
		rc:  rc,
		tc:  tc,
		log: log.NewHelper(logger)}
}

// 点数方法

func (s *PointService) CreatePoint(ctx context.Context, req *v1.CreatePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	point := &biz.Point{
		PID:  req.Point.Pid,
		UID:  req.Point.Uid,
		Name: req.Point.Name,
		Desc: req.Point.Desc,
	}
	if err = s.pc.Create(ctx, point); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) UpdatePoint(ctx context.Context, req *v1.UpdatePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	point := &biz.Point{
		PID:  req.Point.Pid,
		Name: req.Point.Name,
		Desc: req.Point.Desc,
	}
	if err = s.pc.Update(ctx, point); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeletePoint(ctx context.Context, req *v1.DeletePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.pc.Delete(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if err = s.tc.Del(ctx, []uint32{req.Pid}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) GetPoint(ctx context.Context, req *v1.GetPointRequest) (rep *v1.GetPointReply, err error) {
	rep = new(v1.GetPointReply)
	var point *biz.Point
	if point, err = s.pc.Get(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	var ptMap map[uint32]int32
	if ptMap, err = s.tc.MGet(ctx, []uint32{req.Pid}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	s.log.Debugf("pt: %v", ptMap[req.Pid])
	rep.Point = &v1.GetPointReply_Point{
		Pid:       point.PID,
		Name:      point.Name,
		Total:     ptMap[point.PID],
		Desc:      point.Desc,
		CreatedAt: uint64(point.CreatedAt.Unix()),
		UpdatedAt: uint64(point.UpdatedAt.Unix()),
		DeletedAt: uint64(point.DeletedAt.Time.Unix()),
	}
	return
}

func (s *PointService) ListPoint(ctx context.Context, req *v1.ListPointRequest) (rep *v1.ListPointReply, err error) {
	rep = &v1.ListPointReply{Finished: true}
	var points []*biz.Point
	if points, err = s.pc.List(ctx, &biz.PointCond{
		Begin: req.Begin,
		Count: req.Count + 1,
		PIDs:  req.Pids,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	var pids []uint32
	for _, p := range points {
		pids = append(pids, p.PID)
	}
	var ptMap map[uint32]int32
	ptMap, err = s.tc.MGet(ctx, pids)
	if err != nil {
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
			Total:     ptMap[p.PID],
			Desc:      p.Desc,
			CreatedAt: uint64(p.CreatedAt.Unix()),
			UpdatedAt: uint64(p.UpdatedAt.Unix()),
			DeletedAt: uint64(p.DeletedAt.Time.Unix()),
		})
	}
	rep.Count = uint32(len(rep.Points))
	return
}
