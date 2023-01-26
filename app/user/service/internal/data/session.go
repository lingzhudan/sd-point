package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/user/service/internal/biz"
)

type sessionRepo struct {
	data *Data
	log  *log.Helper
}

// NewSessionRepo .
func NewSessionRepo(data *Data, logger log.Logger) biz.SessionRepo {
	return &sessionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Set 设置redis数据
func (r *sessionRepo) Set(ctx context.Context, key string, value interface{}) (err error) {
	if err = r.data.rdb.MSet(ctx, key, value).Err(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

// Get 获取redis数据
func (r *sessionRepo) Get(ctx context.Context, key string) (value interface{}, err error) {
	if value, err = r.data.rdb.Get(ctx, key).Result(); err != nil {
		r.log.Errorf("rdb error: %v", err)
	}
	return
}

// Del 删除redis数据
func (r *sessionRepo) Del(ctx context.Context, keys []string) (err error) {
	if _, err = r.data.rdb.Del(ctx, keys...).Result(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}
