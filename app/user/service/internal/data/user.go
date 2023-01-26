package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/user/service/internal/biz"
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

func (r *userRepo) ListUser(ctx context.Context, cond *biz.UserCond) (users []*biz.User, err error) {
	whereStage, args := cond.ParseCond()
	if err = r.data.db.
		WithContext(ctx).
		Where(whereStage, args...).
		Limit(cond.Count).
		Offset(cond.Begin).
		Find(&users).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *userRepo) GetUser(ctx context.Context, cond *biz.UserCond) (user *biz.User, err error) {
	whereStage, args := cond.ParseCond()
	if err = r.data.db.
		WithContext(ctx).
		Where(whereStage, args...).
		First(&user).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Create(&user).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Model(&biz.User{UID: user.UID}).
		Updates(&user).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *userRepo) DeleteUser(ctx context.Context, uid uint32) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Delete(&biz.User{}, uid).Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}
