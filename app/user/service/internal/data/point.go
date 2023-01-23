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

func (r userRepo) ListUser(ctx context.Context, cond *biz.UserCond) (users []*biz.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepo) GetUser(ctx context.Context, uid uint32) (user *biz.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepo) CreateUser(ctx context.Context, user *biz.User) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepo) UpdateUser(ctx context.Context, user *biz.User) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r userRepo) DeleteUser(ctx context.Context, uid uint32) error {
	//TODO implement me
	panic("implement me")
}
