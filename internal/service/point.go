package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "sd-point/api/point/v1"
	"sd-point/internal/biz"
	"time"
)

type PointService struct {
	pb.UnimplementedPointServer

	uc  *biz.PointUsecase
	log *log.Helper
}

func NewPointService(uc *biz.PointUsecase, logger log.Logger) *PointService {
	return &PointService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PointService) CreatePoints(ctx context.Context, req *pb.CreatePointsRequest) (reply *emptypb.Empty, err error) {
	var points []*biz.Point
	for _, p := range req.Point {
		newPoint := &biz.Point{
			PID:       p.Pid,
			UID:       0,
			ClickedAt: time.Unix(p.ClickedAt, 0),
			Num:       int16(p.Num),
			Desc:      p.Desc,
		}
		points = append(points, newPoint)
	}
	if err = s.uc.Create(ctx, points); err != nil {
		s.log.Error("internal error: %v", err)
		return nil, pb.ErrorContentMissing("", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *PointService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Debugf("name: %+v", req.Name)
	return &pb.HelloReply{Message: req.Name}, nil
}

func (s *PointService) UpdatePoint(ctx context.Context, req *pb.UpdatePointRequest) (reply *emptypb.Empty, err error) {
	point := &biz.Point{
		PID:       req.Point.Pid,
		UID:       0,
		ClickedAt: time.Unix(req.Point.ClickedAt, 0),
		Num:       int16(req.Point.Num),
		Desc:      req.Point.Desc,
	}
	if err = s.uc.Update(ctx, req.Point.Pid, point); err != nil {
		s.log.Debugf("internal error: %v", err)
		return nil, pb.ErrorUserNotFound("", "")
	}
	return &emptypb.Empty{}, nil
}

func (s *PointService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (reply *emptypb.Empty, err error) {
	if err = s.uc.Delete(ctx, req.Pid); err != nil {
		s.log.Debugf("internal error: %v", err)
		return nil, pb.ErrorUserNotFound("", "")
	}
	return &emptypb.Empty{}, nil
}

func (s *PointService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (reply *pb.GetPointReply, err error) {
	var point *biz.Point
	if point, err = s.uc.Get(ctx, req.Pid); err != nil {
		s.log.Debugf("internal error: %v", err)
		return nil, pb.ErrorUserNotFound("", "")
	}
	newPoint := &pb.PointInfo{
		Pid:       point.PID,
		Num:       int32(point.Num),
		Desc:      point.Desc,
		ClickedAt: point.ClickedAt.Unix(),
		CreatedAt: point.CreatedAt.Unix(),
		UpdatedAt: point.UpdatedAt.Unix(),
		DeletedAt: point.DeletedAt.Time.Unix(),
	}

	return &pb.GetPointReply{Point: newPoint}, nil
}

func (s *PointService) ListPoint(ctx context.Context, req *pb.ListPointRequest) (reply *pb.ListPointReply, err error) {
	var data []*pb.PointInfo
	var points []*biz.Point
	if points, err = s.uc.List(ctx); err != nil {
		s.log.Debugf("internal error: %v", err)
		return nil, pb.ErrorUserNotFound("", "")
	}
	for _, point := range points {
		newPoint := &pb.PointInfo{
			Pid:       point.PID,
			Num:       int32(point.Num),
			Desc:      point.Desc,
			ClickedAt: point.ClickedAt.Unix(),
			CreatedAt: point.CreatedAt.Unix(),
			UpdatedAt: point.UpdatedAt.Unix(),
			DeletedAt: point.DeletedAt.Time.Unix(),
		}
		data = append(data, newPoint)
	}
	return &pb.ListPointReply{Point: data}, nil
}
