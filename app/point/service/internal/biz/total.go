package biz

import (
	"context"
)

type TotalRepo interface {
	Exists(ctx context.Context, pid uint32) (exists bool, err error)
	SetTotal(ctx context.Context, pid uint32, total int32) (err error)
	IncrTotal(ctx context.Context, pid uint32, increment int32) (total int32, err error)
	DecrTotal(ctx context.Context, pid uint32, decrement int32) (total int32, err error)
	GetTotal(ctx context.Context, pid uint32) (total int32, err error)
	DeleteTotal(ctx context.Context, pid uint32) (err error)
}

// redis method

func (uc *PointUseCase) IncrTotal(ctx context.Context, pid uint32, increment int32) (total int32, err error) {
	if total, err = uc.GetOrExistsTotal(ctx, pid); err != nil {
		uc.log.Errorf("failed to get or exists total, error: %v", err)
		return
	}
	if total, err = uc.repo.IncrTotal(ctx, pid, increment); err != nil {
		uc.log.Errorf("failed to incr total, error: %v", err)
	}
	return
}

func (uc *PointUseCase) DecrTotal(ctx context.Context, pid uint32, decrement int32) (total int32, err error) {
	if total, err = uc.GetOrExistsTotal(ctx, pid); err != nil {
		uc.log.Errorf("failed to get or exists total, error: %v", err)
		return
	}
	if total, err = uc.repo.IncrTotal(ctx, pid, decrement); err != nil {
		uc.log.Errorf("failed to incr total, error: %v", err)
	}
	return
}

func (uc *PointUseCase) SetTotal(ctx context.Context, pid uint32, total int32) (err error) {
	if err = uc.repo.SetTotal(ctx, pid, total); err != nil {
		uc.log.Errorf("failed to set total, error: %v", err)
	}
	return
}

func (uc *PointUseCase) GetTotal(ctx context.Context, pid uint32) (total int32, err error) {
	return uc.GetOrExistsTotal(ctx, pid)
}

func (uc *PointUseCase) DelTotal(ctx context.Context, pid uint32) (err error) {
	if err = uc.repo.DeleteTotal(ctx, pid); err != nil {
		uc.log.Errorf("failed to del total, error: %v", err)
	}
	return
}

// GetOrExistsTotal 获取或设置点数总数
//   - 存在则返回总数
//   - 不存在则查询关系数据库并设置点数总数
func (uc *PointUseCase) GetOrExistsTotal(ctx context.Context, pid uint32) (total int32, err error) {
	ok, err := uc.repo.Exists(ctx, pid)
	if err != nil {
		uc.log.Errorf("failed to get total exists, error: %v", err)
		return
	}
	if !ok {
		rs, err := uc.ListRecord(ctx, 0, 0, pid)
		if err != nil {
			uc.log.Errorf("internal error: %v", err)
			return 0, err
		}
		for _, r := range rs {
			total += r.Num
		}
		if err = uc.SetTotal(ctx, pid, total); err != nil {
			uc.log.Errorf("internal error: %v", err)
			return 0, err
		}
	}
	if total, err = uc.repo.GetTotal(ctx, pid); err != nil {
		uc.log.Errorf("failed to get total, error: %v", err)
	}
	return
}
