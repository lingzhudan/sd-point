package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/sd-point/interface/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "sd-point/api/sd-point/interface/v1"
)

type SdPointInterfaceService struct {
	pb.UnimplementedSdPointInterfaceServer

	uc *biz.UserUseCase
	pc *biz.PointUseCase

	log *log.Helper
}

func NewSdPointInterfaceService(uc *biz.UserUseCase, pc *biz.PointUseCase, logger log.Logger) *SdPointInterfaceService {
	return &SdPointInterfaceService{
		uc:  uc,
		pc:  pc,
		log: log.NewHelper(logger),
	}
}

func (s *SdPointInterfaceService) CreatePoint(ctx context.Context, req *pb.CreatePointRequest) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) UpdatePoint(ctx context.Context, req *pb.UpdatePointRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (*pb.GetPointReply, error) {
	return &pb.GetPointReply{}, nil
}
func (s *SdPointInterfaceService) ListPoint(ctx context.Context, req *pb.ListPointRequest) (*pb.ListPointReply, error) {
	return &pb.ListPointReply{}, nil
}
func (s *SdPointInterfaceService) CreateRecord(ctx context.Context, req *pb.CreateRecordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) DeleteRecord(ctx context.Context, req *pb.DeleteRecordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) UpdateRecord(ctx context.Context, req *pb.UpdateRecordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) ListRecord(ctx context.Context, req *pb.ListRecordRequest) (*pb.ListRecordReply, error) {
	return &pb.ListRecordReply{}, nil
}
func (s *SdPointInterfaceService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *SdPointInterfaceService) Logout(ctx context.Context, req *pb.LogoutRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) BindAccount(ctx context.Context, req *pb.BindAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *SdPointInterfaceService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *SdPointInterfaceService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
