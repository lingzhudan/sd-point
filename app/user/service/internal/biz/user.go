package biz

import (
	"context"
	_ "embed"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/user/service/internal/define"
	"time"
)

type UserUseCase struct {
	ur  UserRepo
	sc  SessionUseCase
	log *log.Helper
}

func NewUserUseCase(ur UserRepo, sc SessionUseCase, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		ur:  ur,
		sc:  sc,
		log: log.NewHelper(logger),
	}
}

type UserRepo interface {
	ListUser(ctx context.Context, uids []uint32) (users []*User, err error)
	GetUser(ctx context.Context, uid uint32) (user *User, err error)
	GetUserByAccount(ctx context.Context, account string) (user *User, err error)
	GetUserByWechat(ctx context.Context, openid string) (user *User, err error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (user *User, err error)
	Register(ctx context.Context, account string, password string) (uid uint32, err error)
	RegisterByWechat(ctx context.Context, openid string) (uid uint32, err error)
	RegisterByPhoneNumber(ctx context.Context, phoneNumber string) (uid uint32, err error)
	BindWechat(ctx context.Context, uid uint32, openid string) (err error)
	BindPhoneNumber(ctx context.Context, uid uint32, phoneNumber string) (err error)
	UnbindWechat(ctx context.Context, openid string) (err error)
	UnbindPhoneNumber(ctx context.Context, phoneNumber string) (err error)
	DeleteUser(ctx context.Context, uid uint32) (err error)
}

type User struct {
	UID         uint32
	Name        string
	Account     string
	Password    string
	Openid      string
	PhoneNumber string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (uc *UserUseCase) GetUser(ctx context.Context, uid uint32) (user *User, err error) {
	return uc.ur.GetUser(ctx, uid)
}

func (uc *UserUseCase) ListUser(ctx context.Context, uids []uint32) (users []*User, err error) {
	return uc.ur.ListUser(ctx, uids)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, uid uint32) (err error) {
	return uc.ur.DeleteUser(ctx, uid)
}

func (uc *UserUseCase) Login(ctx context.Context, account string, password string) (sessionId string, err error) {
	u, err := uc.ur.GetUserByAccount(ctx, account)
	if err != nil {
		return
	} else if u.Password != password {
		return "", define.ErrPasswordIncorrect
	}
	return uc.newSessionID(ctx, u.UID)
}
func (uc *UserUseCase) WechatLogin(ctx context.Context, openid string) (sessionId string, err error) {
	u, err := uc.ur.GetUserByWechat(ctx, openid)
	if err != nil {
		return
	}
	return uc.newSessionID(ctx, u.UID)
}
func (uc *UserUseCase) PhoneNumberLogin(ctx context.Context, phoneNumber string) (sessionId string, err error) {
	u, err := uc.ur.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return
	}
	return uc.newSessionID(ctx, u.UID)
}
func (uc *UserUseCase) Register(ctx context.Context, account string, password string) (uid uint32, err error) {
	return uc.ur.Register(ctx, account, password)
}
func (uc *UserUseCase) WechatRegister(ctx context.Context, openid string) (uid uint32, err error) {
	return uc.ur.RegisterByWechat(ctx, openid)
}
func (uc *UserUseCase) PhoneNumberRegister(ctx context.Context, phoneNumber string) (uid uint32, err error) {
	return uc.ur.RegisterByPhoneNumber(ctx, phoneNumber)
}
func (uc *UserUseCase) BindWechat(ctx context.Context, uid uint32, openid string) (err error) {
	if err = uc.ur.UnbindWechat(ctx, openid); err != nil {
		return
	}
	return uc.ur.BindWechat(ctx, uid, openid)
}
func (uc *UserUseCase) BindPhoneNumber(ctx context.Context, uid uint32, phoneNumber string) (err error) {
	if err = uc.ur.UnbindPhoneNumber(ctx, phoneNumber); err != nil {
		return
	}
	return uc.ur.BindPhoneNumber(ctx, uid, phoneNumber)
}

func (uc *UserUseCase) GetSession(ctx context.Context, sessionId string) (*Session, error) {
	return uc.sc.GetSession(ctx, sessionId)
}

func (uc *UserUseCase) newSessionID(ctx context.Context, uid uint32) (string, error) {
	return uc.sc.NewSessionID(ctx, uid)
}

func (uc *UserUseCase) Logout(ctx context.Context, sessionId string) error {
	return uc.sc.DelSessionID(ctx, sessionId)
}
