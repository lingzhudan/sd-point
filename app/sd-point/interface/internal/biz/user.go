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
	return uc.repo.GetPublicKey(ctx)
}

func (uc *UserUseCase) GetSession(ctx context.Context, sessionId string) (s *Session, err error) {
	return uc.repo.GetSession(ctx, sessionId)
}

func (uc *UserUseCase) Get(ctx context.Context, uid uint32) (user *User, err error) {
	return uc.repo.GetUser(ctx, uid)
}

func (uc *UserUseCase) List(ctx context.Context, cond *UserCond) (users []*User, err error) {
	return uc.repo.ListUser(ctx, cond)
}

func (uc *UserUseCase) Login(ctx context.Context, account *OriginAccount) (sessionId string, err error) {
	return uc.repo.Login(ctx, account)
}

func (uc *UserUseCase) Logout(ctx context.Context, sessionId string) (err error) {
	return uc.repo.Logout(ctx, sessionId)
}

func (uc *UserUseCase) Register(ctx context.Context, account *OriginAccount) (uid uint32, err error) {
	return uc.repo.Register(ctx, account)
}

func (uc *UserUseCase) WechatLogin(ctx context.Context, account *WechatAccount) (sessionId string, err error) {
	return uc.repo.WechatLogin(ctx, account)
}

func (uc *UserUseCase) WechatPhoneNumberLogin(ctx context.Context, account *WechatAccount) (sessionId string, err error) {
	// TODO
	return
}

func (uc *UserUseCase) WechatRegister(ctx context.Context, account *WechatAccount) (uid uint32, err error) {
	return uc.repo.WechatRegister(ctx, account)
}

func (uc *UserUseCase) WechatPhoneNumberRegister(ctx context.Context, account *WechatAccount) (uid uint32, err error) {
	// TODO
	return
}

func (uc *UserUseCase) WechatBind(ctx context.Context, uid uint32, account *WechatAccount) (err error) {
	return uc.repo.WechatBind(ctx, uid, account)
}

func (uc *UserUseCase) WechatPhoneNumberBind(ctx context.Context, account *WechatAccount) (err error) {
	// TODO
	return
}
