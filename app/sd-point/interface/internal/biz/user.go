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
	GetPublicKey(ctx context.Context) (key []byte, err error)
	GetSession(ctx context.Context, sessionId string) (s *Session, err error)
	GetUser(ctx context.Context, uid uint32) (user *User, err error)
	ListUser(ctx context.Context, cond *UserCond) (users []*User, err error)
	Login(ctx context.Context, account *OriginAccount) (sessionId string, err error)
	Logout(ctx context.Context, sessionId string) (err error)
	Register(ctx context.Context, account *OriginAccount) (uid uint32, err error)
	WechatLogin(ctx context.Context, account *WechatAccount) (sessionId string, err error)
	WechatRegister(ctx context.Context, account *WechatAccount) (uid uint32, err error)
	WechatBind(ctx context.Context, uid uint32, account *WechatAccount) (err error)
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

// WechatAccountCode 使用微信进行账号操作所需要的字段
type WechatAccountCode struct {
	// 向微信后台服务器换取openID的code
	OpenidCode string
	// 向微信后台服务器换取手机号的code
	PhoneNumberCode string
}

// WechatAccount 微信账号
type WechatAccount struct {
	// 微信openID
	Openid string
	// 微信手机号
	PhoneNumber string
}

type Session struct {
	// 用户编号
	UID uint32
}

func (uc *UserUseCase) GetPublicKey(ctx context.Context) (key []byte, err error) {
	if key, err = uc.repo.GetPublicKey(ctx); err != nil {
		uc.log.Errorf("failed to get user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) GetSession(ctx context.Context, sessionId string) (s *Session, err error) {
	if s, err = uc.repo.GetSession(ctx, sessionId); err != nil {
		uc.log.Errorf("failed to get user, error: %v", err)
	}
	return
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

func (uc *UserUseCase) WechatBind(ctx context.Context, uid uint32, account *WechatAccount) (err error) {
	if err = uc.repo.WechatBind(ctx, uid, account); err != nil {
		uc.log.Errorf("failed to login by wechat, error: %v", err)
	}
	return
}

func (uc *UserUseCase) WechatPhoneNumberBind(ctx context.Context, account *WechatAccount) (err error) {
	// TODO
	return
}
