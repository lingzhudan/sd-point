package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	userv1 "sd-point/api/user/service/v1"
	"sd-point/app/sd-point/interface/internal/biz"
	"sd-point/app/sd-point/interface/internal/define"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetPublicKey(ctx context.Context) (key []byte, err error) {
	var reply *userv1.GetPublicKeyReply
	if reply, err = r.data.uc.GetPublicKey(ctx, &emptypb.Empty{}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
	}
	key = reply.PublicKey
	return
}

func (r *userRepo) GetSession(ctx context.Context, sessionId string) (s *biz.Session, err error) {
	s = new(biz.Session)
	var reply *userv1.GetSessionReply
	if reply, err = r.data.uc.GetSession(ctx, &userv1.GetSessionRequest{SessionId: sessionId}); err != nil {
		if define.IsErrRecordNotFound(err) {
			r.log.Error("session not found")
			return
		}
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	s.UID = reply.Session.Uid
	return
}

func (r *userRepo) GetUser(ctx context.Context, uid uint32) (user *biz.User, err error) {
	var reply *userv1.GetUserReply
	if reply, err = r.data.uc.GetUser(ctx, &userv1.GetUserRequest{Uid: uid}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	u := reply.User
	user = &biz.User{
		UID:      u.Uid,
		Username: u.Username,
	}
	return
}

func (r *userRepo) ListUser(ctx context.Context, cond *biz.UserCond) (users []*biz.User, err error) {
	var reply *userv1.ListUserReply
	if reply, err = r.data.uc.ListUser(ctx, &userv1.ListUserRequest{Uids: cond.UIDs}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	for _, u := range reply.Users {
		users = append(users, &biz.User{
			UID:      u.Uid,
			Username: u.Username,
		})
	}
	return
}

func (r *userRepo) Logout(ctx context.Context, sessionId string) (err error) {
	if _, err = r.data.uc.Logout(ctx, &userv1.LogoutRequest{SessionId: sessionId}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (r *userRepo) Login(ctx context.Context, account *biz.OriginAccount) (sessionId string, err error) {
	var reply *userv1.LoginReply
	if reply, err = r.data.uc.Login(ctx, &userv1.LoginRequest{
		Account:  account.Account,
		Password: account.Password,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	sessionId = reply.SessionId
	return
}

func (r *userRepo) Register(ctx context.Context, account *biz.OriginAccount) (uid uint32, err error) {
	var reply *userv1.RegisterReply
	if reply, err = r.data.uc.Register(ctx, &userv1.RegisterRequest{
		Account:  account.Account,
		Password: account.Password,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	uid = reply.Uid
	return
}

func (r *userRepo) WechatLogin(ctx context.Context, account *biz.WechatAccount) (sessionId string, err error) {
	var reply *userv1.LoginReply
	if reply, err = r.data.uc.WechatLogin(ctx, &userv1.WechatLoginRequest{
		OpenId: account.Openid,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	sessionId = reply.SessionId
	return
}

func (r *userRepo) WechatRegister(ctx context.Context, account *biz.WechatAccount) (uid uint32, err error) {
	var reply *userv1.RegisterReply
	if reply, err = r.data.uc.WechatRegister(ctx, &userv1.WechatRegisterRequest{
		Openid: account.Openid,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	uid = reply.Uid
	return
}

func (r *userRepo) WechatBind(ctx context.Context, uid uint32, account *biz.WechatAccount) (err error) {
	if _, err = r.data.uc.WechatBind(ctx, &userv1.WechatBindRequest{
		Uid:    uid,
		Openid: account.Openid,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
	}
	return
}
