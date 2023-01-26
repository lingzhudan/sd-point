package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

type UserRepo interface {
	ListUser(ctx context.Context, cond *UserCond) (users []*User, err error)
	GetUser(ctx context.Context, uid uint32) (user *User, err error)
	Login(ctx context.Context, account *OriginAccount) (sessionId string, err error)
	Logout(ctx context.Context, sessionId string) (err error)
	Register(ctx context.Context, account *OriginAccount) (uid uint32, err error)
	WechatLogin(ctx context.Context, account *WechatAccount) (sessionId string, err error)
	WechatRegister(ctx context.Context, account *WechatAccount) (uid uint32, err error)
	WechatBind(ctx context.Context, account *WechatAccount) (err error)
}

type User struct {
	// 用户编号
	UID uint32
	// 用户名称
	Username string
}

// OriginAccount 自律点项目原生账号系统认证
type OriginAccount struct {
	// 用户账号
	Account string
	// 用户密码 经过公钥加密
	Password string
}

// WechatAccount 使用微信进行账号操作所需要的字段
type WechatAccount struct {
	// 向微信后台服务器换取openID的code
	OpenIDCode string
	// 向微信后台服务器换取手机号的code
	PhoneNumberCode string
}

func (uc *UserUseCase) Get(ctx context.Context, uid uint32) (user *User, err error) {
	if user, err = uc.repo.GetUser(ctx, uid); err != nil {
		uc.log.Errorf("failed to get user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) List(ctx context.Context, cond *UserCond) (users []*User, err error) {
	if users, err = uc.repo.ListUser(ctx, cond); err != nil {
		uc.log.Errorf("failed to get user list, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Login(ctx context.Context, account *OriginAccount) (sessionId string, err error) {
	if sessionId, err = uc.repo.Login(ctx, account); err != nil {
		uc.log.Errorf("failed to login, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Logout(ctx context.Context, sessionId string) (err error) {
	if err = uc.repo.Logout(ctx, sessionId); err != nil {
		uc.log.Errorf("failed to logout, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Register(ctx context.Context, account *OriginAccount) (uid uint32, err error) {
	if uid, err = uc.repo.Register(ctx, account); err != nil {
		uc.log.Errorf("failed to register, error: %v", err)
	}
	return
}

func (uc *UserUseCase) WechatLogin(ctx context.Context, account *WechatAccount) (sessionId string, err error) {
	if sessionId, err = uc.repo.WechatLogin(ctx, account); err != nil {
		uc.log.Errorf("failed to login by wechat, error: %v", err)
	}
	return
}

func (uc *UserUseCase) WechatPhoneNumberLogin(ctx context.Context, account *WechatAccount) (sessionId string, err error) {
	// TODO
	return
}

func (uc *UserUseCase) WechatRegister(ctx context.Context, account *WechatAccount) (uid uint32, err error) {
	if uid, err = uc.repo.WechatRegister(ctx, account); err != nil {
		uc.log.Errorf("failed to login by wechat, error: %v", err)
	}
	return
}

func (uc *UserUseCase) WechatPhoneNumberRegister(ctx context.Context, account *WechatAccount) (uid uint32, err error) {
	// TODO
	return
}

func (uc *UserUseCase) WechatBind(ctx context.Context, account *WechatAccount) (err error) {
	if err = uc.repo.WechatBind(ctx, account); err != nil {
		uc.log.Errorf("failed to login by wechat, error: %v", err)
	}
	return
}

func (uc *UserUseCase) WechatPhoneNumberBind(ctx context.Context, account *WechatAccount) (err error) {
	// TODO
	return
}
