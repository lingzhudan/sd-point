package biz

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
	"gorm.io/gorm"
	rand2 "math/rand"
	"strconv"
	"strings"
	"time"
)

//go:embed private.pem
var privKeyBytes []byte

//go:embed public.pem
var publicKeyBytes []byte

type UserUseCase struct {
	repo           UserRepo
	privKey        *rsa.PrivateKey
	publicKeyBytes []byte
	log            *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
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
	return &UserUseCase{
		repo:           repo,
		privKey:        privKey,
		publicKeyBytes: publicKeyBytes,
		log:            log.NewHelper(logger),
	}
}

type UserRepo interface {
	ListUser(ctx context.Context, cond *UserCond) (users []*User, err error)
	GetUser(ctx context.Context, cond *UserCond) (user *User, err error)
	CreateUser(ctx context.Context, user *User) (err error)
	UpdateUser(ctx context.Context, user *User) (err error)
	DeleteUser(ctx context.Context, uid uint32) error
}

type User struct {
	UID         uint32 `gorm:"column:uid;primaryKey;comment:用户编号;"`
	Name        string `gorm:"column:name;size:32;comment:用户名称;"`
	Account     string `gorm:"column:account;uniqueIndex;size:64;comment:系统账号;"`
	Password    string `gorm:"column:password;size:64;comment:加密系统密码;"`
	Openid      string `gorm:"column:openid;uniqueIndex;size:64;comment:微信openID;"`
	PhoneNumber string `gorm:"column:phone_number;uniqueIndex;size:16;comment:手机号;"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;"`
}

func (u *User) TableName() string {
	return "user"
}

func (uc *UserUseCase) Get(ctx context.Context, cond *UserCond) (user *User, err error) {
	if user, err = uc.repo.GetUser(ctx, cond); err != nil {
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

func (uc *UserUseCase) Create(ctx context.Context, user *User) (err error) {
	if err = uc.repo.CreateUser(ctx, user); err != nil {
		uc.log.Errorf("failed to create user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) (err error) {
	if err = uc.repo.UpdateUser(ctx, user); err != nil {
		uc.log.Errorf("failed to update user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Delete(ctx context.Context, uid uint32) (err error) {
	if err = uc.repo.DeleteUser(ctx, uid); err != nil {
		uc.log.Errorf("failed to delete user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Login(ctx context.Context, u *User) (err error) {
	// 私钥解密密码
	uc.RsaDecrypt(ctx, u)
	if u, err = uc.repo.GetUser(ctx, &UserCond{
		Accounts:  []string{u.Account},
		Passwords: []string{u.Password},
	}); err != nil {
		uc.log.Errorf("internal error: %v", err)
	}
	return
}
func (uc *UserUseCase) Register(ctx context.Context, u *User) (err error) {
	// 私钥解密密码
	uc.RsaDecrypt(ctx, u)
	u2 := &User{
		Account:  u.Account,
		Password: u.Password,
	}
	if err = uc.repo.CreateUser(ctx, u2); err != nil {
		uc.log.Errorf("failed to register user by origin account, error: %v", err)
	}
	return
}
func (uc *UserUseCase) WechatRegister(ctx context.Context, u *User) (err error) {
	u2 := uc.NewUserOriginAccount(ctx)
	u2.Openid = u.Openid
	if err = uc.repo.CreateUser(ctx, u); err != nil {
		uc.log.Errorf("failed to register user by wechat account, error: %v", err)
	}
	return
}
func (uc *UserUseCase) PhoneNumberRegister(ctx context.Context, u *User) (err error) {
	u2 := uc.NewUserOriginAccount(ctx)
	u2.PhoneNumber = u.PhoneNumber
	if err = uc.repo.CreateUser(ctx, u2); err != nil {
		uc.log.Errorf("failed to register user by wechat account, error: %v", err)
	}
	return
}
func (uc *UserUseCase) WechatBind(ctx context.Context, u *User) (err error) {
	if err = uc.repo.UpdateUser(ctx, &User{UID: u.UID, Openid: u.Openid}); err != nil {
		uc.log.Errorf("failed to update user, error: %v", err)
	}
	return
}
func (uc *UserUseCase) PhoneNumberBind(ctx context.Context, u *User) (err error) {
	if err = uc.repo.UpdateUser(ctx, &User{UID: u.UID, PhoneNumber: u.PhoneNumber}); err != nil {
		uc.log.Errorf("failed to update user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) GetPublicKeyBytes(ctx context.Context) []byte {
	return uc.publicKeyBytes
}

// EncryptionPassword 通过用户的某些属性和密码混合加密生成加密数据写回用户结构体中
func (uc *UserUseCase) EncryptionPassword(ctx context.Context, u *User) {
	// TODO 自定义生成相关机密
	if len(u.Password) == 0 {
		panic("empty password to encrypt")
	}
	u.Password = strconv.FormatInt(u.CreatedAt.Unix(), 36) + u.Password
	return
}

// NewUserOriginAccount 创建随机的原生账号密码 密码已经进行服务器端加密
func (uc *UserUseCase) NewUserOriginAccount(ctx context.Context) (u *User) {
	u = new(User)
	// TODO 自定义生成相关账号密码
	// 随机账号为`sd_`+`10位时间生成的36进制数`+`3位随机36进制数`
	// 随机密码为`10位随机36进制数`
	var ab strings.Builder
	var pb strings.Builder
	ab.WriteString("sd_")
	ab.WriteString(strconv.FormatInt(time.Now().Unix(), 36))
	pb.WriteString(strconv.FormatInt(rand2.Int63n(time.Now().UnixMicro()), 36))
	for ab.Len() < 13 {
		ab.WriteString("0")
	}
	for ab.Len() < 16 {
		ab.WriteString(strconv.FormatInt(rand2.Int63n(36), 36))
	}
	for pb.Len() < 10 {
		pb.WriteString(strconv.FormatInt(rand2.Int63n(36), 36))
	}
	u.Account = ab.String()
	u.Password = pb.String()
	uc.EncryptionPassword(ctx, u)
	return
}

// RsaDecrypt 私钥解密密文 服务器端加密密码
func (uc *UserUseCase) RsaDecrypt(ctx context.Context, u *User) {
	var err error
	var ciphertext []byte
	if ciphertext, err = base64.StdEncoding.DecodeString(u.Password); err != nil {
		panic(err)
	}
	var decryption []byte
	if decryption, err = rsa.DecryptPKCS1v15(rand.Reader, uc.privKey, ciphertext); err != nil {
		panic(err)
	}
	u.Password = string(decryption)
	uc.EncryptionPassword(ctx, u)
	return
}
