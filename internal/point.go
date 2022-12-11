package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "sd-point/api/point/v1"
)

type PointService struct {
	pb.UnimplementedPointServer
}

func NewPointService() *PointService {
	return &PointService{}
}

func (s *PointService) CreatePoint(ctx context.Context, req *pb.CreatePointRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *PointService) UpdatePoint(ctx context.Context, req *pb.UpdatePointRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *PointService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *PointService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (*pb.GetPointReply, error) {
	return &pb.GetPointReply{}, nil
}
func (s *PointService) ListPoint(ctx context.Context, req *pb.ListPointRequest) (*pb.ListPointReply, error) {
	return &pb.ListPointReply{}, nil
}
