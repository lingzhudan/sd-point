package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// redis method

// Exists 获取数据是否存在
func (r *pointRepo) Exists(ctx context.Context, pid uint32) (exists bool, err error) {
	cmd := r.data.rdb.Exists(ctx, strconv.Itoa(int(pid)))
	result, err := cmd.Result()
	if err != nil {
		r.log.Errorf("db error: %v", err)
	}
	exists = result == 1
	return
}

// IncrTotal 新增点数总数
func (r *pointRepo) IncrTotal(ctx context.Context, pid uint32, increment int32) (total int32, err error) {
	cmd := r.data.rdb.DecrBy(ctx, strconv.Itoa(int(pid)), int64(increment))
	i, err := cmd.Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("rdb error: %v", err)
		}
	}
	total = int32(i)
	return
}

// DecrTotal 新增点数总数
func (r *pointRepo) DecrTotal(ctx context.Context, pid uint32, decrement int32) (total int32, err error) {
	cmd := r.data.rdb.DecrBy(ctx, strconv.Itoa(int(pid)), int64(decrement))
	i, err := cmd.Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("rdb error: %v", err)
		}
	}
	total = int32(i)
	return
}

// SetTotal 设置点数总数
func (r *pointRepo) SetTotal(ctx context.Context, pid uint32, total int32) (err error) {
	if err = r.data.rdb.Set(ctx, strconv.Itoa(int(pid)), total, 0).Err(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

// GetTotal 获取点数总数
func (r *pointRepo) GetTotal(ctx context.Context, pid uint32) (total int32, err error) {
	value, err := r.data.rdb.Get(ctx, strconv.Itoa(int(pid))).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("rdb error: %v", err)
			return
		}
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		r.log.Error("total not number")
	}
	total = int32(i)
	return
}

// DelTotal 删除点数总数
func (r *pointRepo) DeleteTotal(ctx context.Context, pid uint32) (err error) {
	if _, err = r.data.rdb.Del(ctx, strconv.Itoa(int(pid))).Result(); err != nil {
		if !errors.Is(err, redis.Nil) {
			r.log.Errorf("rdb error: %v", err)
		}
	}
	return
}
