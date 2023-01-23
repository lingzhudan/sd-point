package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
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
	CreateUser(ctx context.Context, user *User) (err error)
	UpdateUser(ctx context.Context, user *User) (err error)
	DeleteUser(ctx context.Context, uid uint32) error
}

type User struct {
	UID         uint32 `gorm:"column:uid;primaryKey;comment:用户编号;"`
	Name        string `gorm:"column:name;size:32;comment:用户名称;"`
	OpenID      string `gorm:"column:open_id;size:64;comment:微信openID;"`
	PhoneNumber string `gorm:"column:phone_number;size:16;comment:手机号;"`
	Desc        string `gorm:"column:desc;size:1024;comment:用户描述;"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;"`
}

func (u *User) TableName() string {
	return "user"
}

func (uc *UserUseCase) Get(ctx context.Context, uid uint32) (user *User, err error) {
	if user, err = uc.repo.GetUser(ctx, uid); err != nil {
		uc.log.Debugf("failed to get user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) List(ctx context.Context, cond *UserCond) (users []*User, err error) {
	if users, err = uc.repo.ListUser(ctx, cond); err != nil {
		uc.log.Debugf("failed to get user list, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (err error) {
	if err = uc.repo.CreateUser(ctx, user); err != nil {
		uc.log.Debugf("failed to create user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Update(ctx context.Context, user *User) (err error) {
	if err = uc.repo.UpdateUser(ctx, user); err != nil {
		uc.log.Debugf("failed to update user, error: %v", err)
	}
	return
}

func (uc *UserUseCase) Delete(ctx context.Context, uid uint32) (err error) {
	if err = uc.repo.DeleteUser(ctx, uid); err != nil {
		uc.log.Errorf("failed to delete user, error: %v", err)
	}
	return
}

// EncryptionPassword 通过用户的某些属性和密码混合加密生成加密数据
func (uc *UserUseCase) EncryptionPassword(ctx context.Context, user *User) (pwd string) {
	return ""
}
