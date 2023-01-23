package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
)

type TotalRepo interface {
	IncrBy(ctx context.Context, key string, value int64) (result int64, err error)
	DecrBy(ctx context.Context, key string, value int64) (result int64, err error)
	MSet(ctx context.Context, data map[string]interface{}) (err error)
	MGet(ctx context.Context, keys []string) (values []interface{}, err error)
	Del(ctx context.Context, keys []string) (err error)
}

type TotalUseCase struct {
	repo TotalRepo
	log  *log.Helper
}

func NewTotalUseCase(repo TotalRepo, logger log.Logger) *TotalUseCase {
	return &TotalUseCase{
		repo: repo,
		log:  log.NewHelper(logger)}
}

// redis method

func (uc *TotalUseCase) Incr(ctx context.Context, data map[uint32]int32) (result int64, err error) {
	for pid, num := range data {
		if result, err = uc.repo.IncrBy(ctx, strconv.Itoa(int(pid)), int64(num)); err != nil {
			uc.log.Errorf("failed to incr key, error: %v", err)
		}
	}
	return
}

func (uc *TotalUseCase) Decr(ctx context.Context, data map[uint32]int32) (result int64, err error) {
	for pid, num := range data {
		if result, err = uc.repo.DecrBy(ctx, strconv.Itoa(int(pid)), int64(num)); err != nil {
			uc.log.Errorf("failed to decr key, error: %v", err)
		}
	}
	return
}

func (uc *TotalUseCase) MSet(ctx context.Context, pointTotals map[uint32]uint32) (err error) {
	var d = make(map[string]interface{})
	for k, v := range pointTotals {
		d[strconv.Itoa(int(k))] = v
	}
	if err = uc.repo.MSet(ctx, d); err != nil {
		uc.log.Errorf("failed to set keys, error: %v", err)
	}
	return
}

func (uc *TotalUseCase) MGet(ctx context.Context, pids []uint32) (data map[uint32]int32, err error) {
	data = make(map[uint32]int32)
	var keys []string
	var values []interface{}
	for _, pid := range pids {
		keys = append(keys, strconv.Itoa(int(pid)))
	}
	if values, err = uc.repo.MGet(ctx, keys); err != nil {
		uc.log.Errorf("failed to get keys, error: %v", err)
	}
	for i := range values {
		value, ok := values[i].(string)
		if ok {
			var num int
			num, err = strconv.Atoi(value)
			if err != nil {
				uc.log.Errorf("failed to get int form string, error: %v", err)
			}
			data[pids[i]] = int32(num)
		} else {
			uc.log.Errorf("failed to get string type of value")
		}
	}
	return
}

func (uc *TotalUseCase) Del(ctx context.Context, pids []uint32) (err error) {
	var keys []string
	for _, pid := range pids {
		keys = append(keys, strconv.Itoa(int(pid)))
	}
	if err = uc.repo.Del(ctx, keys); err != nil {
		uc.log.Errorf("failed to del keys, error: %v", err)
	}
	return
}
