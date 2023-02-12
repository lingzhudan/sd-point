package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "sd-point/api/user/service/v1"
	"sd-point/app/user/service/internal/biz"
	"sd-point/app/user/service/internal/define"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.UserUseCase
	sc *biz.SessionUseCase

	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, sc *biz.SessionUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		sc:  sc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) GetPublicKey(ctx context.Context, _ *emptypb.Empty) (rep *pb.GetPublicKeyReply, _ error) {
	rep = new(pb.GetPublicKeyReply)
	rep.PublicKey = s.uc.GetPublicKeyBytes(ctx)
	return
}

func (s *UserService) GetSession(ctx context.Context, req *pb.GetSessionRequest) (rep *pb.GetSessionReply, err error) {
	rep = new(pb.GetSessionReply)
	var session *biz.Session
	if session, err = s.sc.GetSession(ctx, req.SessionId); err != nil {
		if define.IsErrRecordNotFound(err) {
			s.log.Errorf("session not found")
			return
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	rep.Session = &pb.GetSessionReply_Session{Uid: session.UID}
	return
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (rep *pb.LoginReply, err error) {
	rep = new(pb.LoginReply)
	var u *biz.User
	var sessionId string
	u = &biz.User{
		Account:  req.Account,
		Password: req.Password,
	}
	if err = s.uc.Login(ctx, u); err != nil {
		if define.IsErrRecordNotFound(err) {
			s.log.Errorf(pb.ErrorReason_PASSWORD_ERROR.String())
			err = pb.ErrorPasswordError("", "")
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	if sessionId, err = s.sc.NewSessionID(ctx, u); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	rep.SessionId = sessionId
	return
}
func (s *UserService) WechatLogin(ctx context.Context, req *pb.WechatLoginRequest) (rep *pb.LoginReply, err error) {
	rep = new(pb.LoginReply)
	var user *biz.User
	var sessionId string
	if user, err = s.uc.Get(ctx, &biz.UserCond{
		OpenIDs: []string{req.OpenId},
	}); err != nil {
		if define.IsErrRecordNotFound(err) {
			s.log.Errorf(pb.ErrorReason_WECHAT_CODE_ERROR.String())
			err = pb.ErrorWechatCodeError("", "")
			return
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	if sessionId, err = s.sc.NewSessionID(ctx, user); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	rep.SessionId = sessionId
	return
}
func (s *UserService) PhoneNumberLogin(ctx context.Context, req *pb.PhoneNumberLoginRequest) (rep *pb.LoginReply, err error) {
	rep = new(pb.LoginReply)
	var user *biz.User
	var sessionId string
	if user, err = s.uc.Get(ctx, &biz.UserCond{
		PhoneNumbers: []string{req.PhoneNumber},
	}); err != nil {
		if define.IsErrRecordNotFound(err) {
			s.log.Errorf(pb.ErrorReason_WECHAT_CODE_ERROR.String())
			err = pb.ErrorWechatCodeError("", "")
			return
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	if sessionId, err = s.sc.NewSessionID(ctx, user); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	rep.SessionId = sessionId
	return
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (_ *emptypb.Empty, err error) {
	if err = s.sc.DelSessionID(ctx, req.SessionId); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (rep *pb.RegisterReply, err error) {
	rep = new(pb.RegisterReply)
	user := &biz.User{
		Account:  req.Account,
		Password: req.Password,
	}
	if err = s.uc.Register(ctx, user); err != nil {
		if define.IsErrDuplicateKey(err) {
			s.log.Errorf("duplicate account: %s", user.Account)
			err = pb.ErrorAccountRegistered("", "")
			return
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	rep.Uid = user.UID
	return
}
func (s *UserService) WechatRegister(ctx context.Context, req *pb.WechatRegisterRequest) (rep *pb.RegisterReply, err error) {
	rep = new(pb.RegisterReply)
	user := &biz.User{
		Openid: req.Openid,
	}
	if err = s.uc.WechatRegister(ctx, user); err != nil {
		if define.IsErrDuplicateKey(err) {
			s.log.Errorf("duplicate account: %s", user.Account)
			err = pb.ErrorWechatRegistered("", "")
			return
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	rep.Uid = user.UID
	return
}
func (s *UserService) PhoneNumberRegister(ctx context.Context, req *pb.PhoneNumberRegisterRequest) (rep *pb.RegisterReply, err error) {
	rep = new(pb.RegisterReply)
	user := &biz.User{
		PhoneNumber: req.PhoneNumber,
	}
	if err = s.uc.PhoneNumberRegister(ctx, user); err != nil {
		if define.IsErrDuplicateKey(err) {
			s.log.Errorf("duplicate account: %s", user.Account)
			err = pb.ErrorPhoneNumberRegistered("", "")
			return
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	rep.Uid = user.UID
	return
}
func (s *UserService) WechatBind(ctx context.Context, req *pb.WechatBindRequest) (_ *emptypb.Empty, err error) {
	if err = s.uc.WechatBind(ctx, &biz.User{UID: req.Uid, Openid: req.Openid}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *UserService) WechatSwitchBind(ctx context.Context, req *pb.WechatSwitchBindRequest) (_ *emptypb.Empty, err error) {
	// TODO
	return
}
func (s *UserService) PhoneNumberBind(ctx context.Context, req *pb.PhoneNumberBindRequest) (_ *emptypb.Empty, err error) {
	if err = s.uc.PhoneNumberBind(ctx, &biz.User{UID: req.Uid, PhoneNumber: req.PhoneNumber}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *UserService) PhoneNumberSwitchBind(ctx context.Context, req *pb.PhoneNumberSwitchBindRequest) (_ *emptypb.Empty, err error) {
	// TODO
	return
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (rep *pb.GetUserReply, err error) {
	rep = new(pb.GetUserReply)
	var user *biz.User
	if user, err = s.uc.Get(ctx, &biz.UserCond{UIDs: []uint32{req.Uid}}); err != nil {
		if define.IsErrRecordNotFound(err) {
			s.log.Error(pb.ErrorReason_USER_NOT_FOUND.String())
			err = pb.ErrorUserNotFound("", "")
			return
		}
		s.log.Errorf("internal error: %v", err)
		return
	}
	rep.User = &pb.GetUserReply_User{
		Uid:      user.UID,
		Username: user.Name,
	}
	return
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (rep *pb.ListUserReply, err error) {
	rep = new(pb.ListUserReply)
	var users []*biz.User
	if users, err = s.uc.List(ctx, &biz.UserCond{UIDs: req.Uids}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	for _, u := range users {
		rep.Users = append(rep.Users, &pb.GetUserReply_User{
			Uid:      u.UID,
			Username: u.Name,
		})
	}
	return
}
