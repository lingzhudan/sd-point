package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "sd-point/api/point/v1"
)

type PointService struct {
	pb.UnimplementedPointServer
}

func NewPointService() *PointService {
	return &PointService{}
}

func (s *PointService) CreatePoints(ctx context.Context, req *pb.CreatePointsRequest) (*emptypb.Empty, error) {

	log.Debugf("point: %+v", req.Point)
	log.Debugf("point validateAll: %+v", req.ValidateAll())

	return &emptypb.Empty{}, nil
}
func (s *PointService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Debugf("name: %+v", req.Name)
	return &pb.HelloReply{Message: req.Name}, nil
}
func (s *PointService) UpdatePoint(ctx context.Context, req *pb.UpdatePointRequest) (*pb.UpdatePointReply, error) {
	return &pb.UpdatePointReply{}, nil
}
func (s *PointService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (*pb.DeletePointReply, error) {
	return &pb.DeletePointReply{}, nil
}
func (s *PointService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (*pb.GetPointReply, error) {
	return &pb.GetPointReply{}, nil
}
func (s *PointService) ListPoint(ctx context.Context, req *pb.ListPointRequest) (*pb.ListPointReply, error) {
	return &pb.ListPointReply{}, nil
}
