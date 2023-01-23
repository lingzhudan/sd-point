package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userv1 "sd-point/api/user/service/v1"
	"sd-point/app/sd-point/interface/internal/biz"
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
	// TODO 向微信后台服务器换取openID和手机号
	openID, phoneNumber := account.OpenIDCode, account.PhoneNumberCode
	if reply, err = r.data.uc.WechatLogin(ctx, &userv1.WechatLoginRequest{
		OpenId:      openID,
		PhoneNumber: phoneNumber,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	sessionId = reply.SessionId
	return
}

func (r *userRepo) WechatRegister(ctx context.Context, account *biz.WechatAccount) (uid uint32, err error) {
	var reply *userv1.RegisterReply
	// TODO 向微信后台服务器换取openID和手机号
	openID, phoneNumber := account.OpenIDCode, account.PhoneNumberCode
	if reply, err = r.data.uc.WechatRegister(ctx, &userv1.WechatRegisterRequest{
		OpenId:      openID,
		PhoneNumber: phoneNumber,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
		return
	}
	uid = reply.Uid
	return
}

func (r *userRepo) WechatBind(ctx context.Context, account *biz.WechatAccount) (err error) {
	// TODO 获取用户编号
	uid := uint32(0)
	// TODO 向微信后台服务器换取openID和手机号
	openID, phoneNumber := account.OpenIDCode, account.PhoneNumberCode
	if _, err = r.data.uc.WechatBind(ctx, &userv1.WechatBindRequest{
		Uid:         uid,
		OpenId:      openID,
		PhoneNumber: phoneNumber,
	}); err != nil {
		r.log.Errorf("grpc client error: %v", err)
	}
	return
}
