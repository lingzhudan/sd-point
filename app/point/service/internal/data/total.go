package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/point/service/internal/biz"
)

type totalRepo struct {
	data *Data
	log  *log.Helper
}

// NewTotalRepo .
func NewTotalRepo(data *Data, logger log.Logger) biz.TotalRepo {
	return &totalRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// redis method

// IncrBy 新增点数总数
func (r *totalRepo) IncrBy(ctx context.Context, key string, value int64) (result int64, err error) {
	cmd := r.data.rdb.IncrBy(ctx, key, value)
	if result, err = cmd.Result(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

// DecrBy 新增点数总数
func (r *totalRepo) DecrBy(ctx context.Context, key string, value int64) (result int64, err error) {
	cmd := r.data.rdb.DecrBy(ctx, key, value)
	if result, err = cmd.Result(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

// MSet 批量设置redis数据
//
// 例：
//
//	err := MSet(context.Background(), map[string]interface{"key": 100, "key2": "100"})
func (r *totalRepo) MSet(ctx context.Context, data map[string]interface{}) (err error) {
	if err = r.data.rdb.MSet(ctx, data).Err(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

// MGet 批量获取redis数据
//
// 例：
//
//	values, err := MGet(context.Background(), []string{"key", "key2"})
//
// values: []interface{100, "100"}, err: nil
func (r *totalRepo) MGet(ctx context.Context, keys []string) (values []interface{}, err error) {
	if values, err = r.data.rdb.MGet(ctx, keys...).Result(); err != nil {
		r.log.Errorf("rdb error: %v", err)
	}
	return
}

// Del 删除redis数据
func (r *totalRepo) Del(ctx context.Context, keys []string) (err error) {
	if _, err = r.data.rdb.Del(ctx, keys...).Result(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}
