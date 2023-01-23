package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/user/service/internal/biz"

	pb "sd-point/api/user/service/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.UserUseCase

	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *UserService) WechatLogin(ctx context.Context, req *pb.WechatLoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *UserService) WechatRegister(ctx context.Context, req *pb.WechatRegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *UserService) WechatBind(ctx context.Context, req *pb.WechatBindRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
