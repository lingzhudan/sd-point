package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/base64"
	"encoding/pem"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "sd-point/api/user/service/v1"
	"sd-point/app/user/service/internal/biz"
	"sd-point/app/user/service/internal/define"

	"google.golang.org/protobuf/types/known/emptypb"
)

//go:embed private.pem
var privKeyBytes []byte

//go:embed public.pem
var publicKeyBytes []byte

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.UserUseCase

	privKey        *rsa.PrivateKey
	publicKeyBytes []byte
	log            *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	//获取私钥
	block, _ := pem.Decode(privKeyBytes)
	if block == nil {
		panic(errors.New(500, "private key error", ""))
	}
	//解析PKCS1格式的私钥
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return &UserService{
		uc:             uc,
		privKey:        privKey,
		publicKeyBytes: publicKeyBytes,
		log:            log.NewHelper(logger),
	}
}

func (s *UserService) GetPublicKey(_ context.Context, _ *emptypb.Empty) (rep *pb.GetPublicKeyReply, _ error) {
	rep = new(pb.GetPublicKeyReply)
	rep.PublicKey = s.publicKeyBytes
	return
}

func (s *UserService) GetSession(ctx context.Context, req *pb.GetSessionRequest) (rep *pb.GetSessionReply, err error) {
	rep = new(pb.GetSessionReply)
	rep.Session = new(pb.GetSessionReply_Session)
	session, err := s.uc.GetSession(ctx, req.SessionId)
	if err != nil {
		if errors.Is(err, define.ErrRecordNotFound) {
			err = pb.ErrorNotLoggedIn("用户未登录")
		}
		return
	}
	rep.Session.Uid = session.UID
	return
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (rep *pb.LoginReply, err error) {
	rep = new(pb.LoginReply)
	rep.SessionId, err = s.uc.Login(ctx, req.Account, s.rsaDecrypt(req.Password))
	if err != nil {
		if errors.Is(err, define.ErrRecordNotFound) {
			err = pb.ErrorUserNotFound("账号: %s 错误", req.Account)
		}
		if errors.Is(err, define.ErrPasswordIncorrect) {
			err = pb.ErrorPasswordError("密码错误")
		}
	}
	return
}
func (s *UserService) WechatLogin(ctx context.Context, req *pb.WechatLoginRequest) (rep *pb.LoginReply, err error) {
	rep = new(pb.LoginReply)
	rep.SessionId, err = s.uc.WechatLogin(ctx, req.OpenId)
	if err != nil {
		if errors.Is(err, define.ErrRecordNotFound) {
			err = pb.ErrorUserNotFound("微信用户openid: %s 尚未注册", req.OpenId)
		}
	}
	return
}
func (s *UserService) PhoneNumberLogin(ctx context.Context, req *pb.PhoneNumberLoginRequest) (rep *pb.LoginReply, err error) {
	rep = new(pb.LoginReply)
	rep.SessionId, err = s.uc.PhoneNumberLogin(ctx, req.PhoneNumber)
	if err != nil {
		if errors.Is(err, define.ErrRecordNotFound) {
			err = pb.ErrorUserNotFound("手机号: %s 尚未注册", req.PhoneNumber)
		}
	}
	return
}
func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (rep *emptypb.Empty, err error) {
	return new(emptypb.Empty), s.uc.Logout(ctx, req.SessionId)
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (rep *pb.RegisterReply, err error) {
	rep = new(pb.RegisterReply)
	rep.Uid, err = s.uc.Register(ctx, req.Account, s.rsaDecrypt(req.Password))
	if err != nil {
		if errors.Is(err, define.ErrAccountRegistered) {
			err = pb.ErrorAccountRegistered("账号: %s 已经被注册", req.Account)
			return
		}
	}
	return
}
func (s *UserService) WechatRegister(ctx context.Context, req *pb.WechatRegisterRequest) (rep *pb.RegisterReply, err error) {
	rep = new(pb.RegisterReply)
	rep.Uid, err = s.uc.WechatRegister(ctx, req.Openid)
	if err != nil {
		if errors.Is(err, define.ErrWechatRegistered) {
			err = pb.ErrorWechatRegistered("微信用户: %s 已经被注册", req.Openid)
		}
	}
	return
}
func (s *UserService) PhoneNumberRegister(ctx context.Context, req *pb.PhoneNumberRegisterRequest) (rep *pb.RegisterReply, err error) {
	rep = new(pb.RegisterReply)
	rep.Uid, err = s.uc.PhoneNumberRegister(ctx, req.PhoneNumber)
	if err != nil {
		if errors.Is(err, define.ErrPhoneNumberRegistered) {
			err = pb.ErrorPhoneNumberRegistered("手机号: %s 已经被注册", req.PhoneNumber)
		}
	}
	return
}
func (s *UserService) WechatBind(ctx context.Context, req *pb.WechatBindRequest) (rep *emptypb.Empty, err error) {
	return new(emptypb.Empty), s.uc.BindWechat(ctx, req.Uid, req.Openid)
}
func (s *UserService) PhoneNumberBind(ctx context.Context, req *pb.PhoneNumberBindRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	return rep, s.uc.BindPhoneNumber(ctx, req.Uid, req.PhoneNumber)
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (rep *pb.GetUserReply, err error) {
	rep = new(pb.GetUserReply)
	user, err := s.uc.GetUser(ctx, req.Uid)
	if err != nil {
		if errors.Is(err, define.ErrUserNotFound) {
			err = pb.ErrorUserNotFound("没有找到用户: %d", req.Uid)
		}
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
	users, err := s.uc.ListUser(ctx, req.Uids)
	for _, u := range users {
		rep.Users = append(rep.Users, &pb.GetUserReply_User{
			Uid:      u.UID,
			Username: u.Name,
		})
	}
	return
}

// 私钥解密密文
func (s *UserService) rsaDecrypt(password string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return ""
	}
	decryption, err := rsa.DecryptPKCS1v15(rand.Reader, s.privKey, ciphertext)
	if err != nil {
		return ""
	}
	return string(decryption)
}
