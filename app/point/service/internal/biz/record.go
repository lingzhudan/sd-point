package biz

import (
	"context"
	"time"
)

type RecordRepo interface {
	GetRecord(ctx context.Context, rid uint32) (record *Record, err error)
	ListRecord(ctx context.Context, begin int, count int, uid uint32) (records []*Record, err error)
	CreateRecord(ctx context.Context, pid uint32, num int32, desc string, clickedAt time.Time) (err error)
	UpdateRecord(ctx context.Context, rid uint32, num int32, desc string) (err error)
	DeleteRecord(ctx context.Context, rid uint32) (err error)
	DeleteRecords(ctx context.Context, pid uint32) (err error)
}

type Record struct {
	RID       uint32
	PID       uint32
	Num       int32
	Desc      string
	ClickedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

// point record method

func (uc *PointUseCase) GetRecord(ctx context.Context, rid uint32) (record *Record, err error) {
	if record, err = uc.repo.GetRecord(ctx, rid); err != nil {
		uc.log.Errorf("failed to get record, error: %v", err)
	}
	return
}

func (uc *PointUseCase) ListRecord(ctx context.Context, begin int, count int, pid uint32) (records []*Record, err error) {
	if records, err = uc.repo.ListRecord(ctx, begin, count, pid); err != nil {
		uc.log.Errorf("failed to get records, error: %v", err)
	}
	return
}

func (uc *PointUseCase) CreateRecord(ctx context.Context, pid uint32, num int32, desc string, clickedAt time.Time) (err error) {
	if err = uc.repo.CreateRecord(ctx, pid, num, desc, clickedAt); err != nil {
		uc.log.Errorf("failed to create records, error: %v", err)
		return
	}
	if _, err = uc.IncrTotal(ctx, pid, num); err != nil {
		uc.log.Errorf("internal error: %v", err)
	}
	return
}

func (uc *PointUseCase) UpdateRecord(ctx context.Context, rid uint32, num int32, desc string) (err error) {
	r, err := uc.GetRecord(ctx, rid)
	if err != nil {
		uc.log.Errorf("internal error: %v", err)
		return
	}
	if err = uc.repo.UpdateRecord(ctx, rid, num, desc); err != nil {
		uc.log.Errorf("failed to update record, error: %v", err)
	}
	n := num - r.Num
	if n != 0 {
		if _, err = uc.IncrTotal(ctx, r.PID, n); err != nil {
			uc.log.Errorf("internal error: %v", err)
		}
	}
	return
}

func (uc *PointUseCase) DeleteRecord(ctx context.Context, rid uint32) (err error) {
	var r *Record
	if r, err = uc.repo.GetRecord(ctx, rid); err != nil {
		uc.log.Errorf("internal error: %v", err)
		return
	}
	if err = uc.repo.DeleteRecord(ctx, rid); err != nil {
		uc.log.Errorf("failed to delete record, error: %v", err)
	}
	if _, err = uc.DecrTotal(ctx, r.PID, r.Num); err != nil {
		uc.log.Errorf("internal error: %v", err)
	}
	return
}
