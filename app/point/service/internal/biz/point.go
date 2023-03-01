package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type PointUseCase struct {
	repo PointRepo
	log  *log.Helper
}

func NewPointUseCase(repo PointRepo, logger log.Logger) *PointUseCase {
	return &PointUseCase{
		repo: repo,
		log:  log.NewHelper(logger)}
}

type PointRepo interface {
	GetPoint(ctx context.Context, pid uint32) (point *Point, err error)
	ListPoint(ctx context.Context, begin int, count int, uid uint32) (points []*Point, err error)
	CreatePoint(ctx context.Context, uid uint32, name string, desc string) (err error)
	UpdatePoint(ctx context.Context, pid uint32, name string, desc string) (err error)
	DeletePoint(ctx context.Context, pid uint32) error

	TotalRepo
	RecordRepo
}

type Point struct {
	PID   uint32
	UID   uint32
	Name  string
	Total int32
	Desc  string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// point method

func (uc *PointUseCase) GetPoint(ctx context.Context, pid uint32) (point *Point, err error) {
	if point, err = uc.repo.GetPoint(ctx, pid); err != nil {
		uc.log.Errorf("failed to get point, error: %v", err)
	}
	total, err := uc.GetTotal(ctx, pid)
	if err != nil {
		uc.log.Errorf("failed to get total, error: %v", err)
	}
	point.Total = total
	return
}

func (uc *PointUseCase) ListPoint(ctx context.Context, begin, count int, uid uint32) (points []*Point, err error) {
	if points, err = uc.repo.ListPoint(ctx, begin, count, uid); err != nil {
		uc.log.Errorf("failed to get points, error: %v", err)
	}
	for _, p := range points {
		total, err := uc.GetTotal(ctx, p.PID)
		if err != nil {
			uc.log.Errorf("failed to get total, error: %v", err)
		}
		p.Total = total
	}
	return
}

func (uc *PointUseCase) CreatePoint(ctx context.Context, uid uint32, name string, desc string) (err error) {
	if err = uc.repo.CreatePoint(ctx, uid, name, desc); err != nil {
		uc.log.Errorf("failed to create point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) UpdatePoint(ctx context.Context, uid uint32, name string, desc string) (err error) {
	if err = uc.repo.UpdatePoint(ctx, uid, name, desc); err != nil {
		uc.log.Errorf("failed to update point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) DeletePoint(ctx context.Context, pid uint32) (err error) {
	if err = uc.repo.DeletePoint(ctx, pid); err != nil {
		uc.log.Errorf("failed to delete point, error: %v", err)
	}
	if err = uc.repo.DeleteRecords(ctx, pid); err != nil {
		uc.log.Errorf("failed to delete record, error: %v", err)
	}
	if err = uc.repo.DeleteTotal(ctx, pid); err != nil {
		uc.log.Errorf("failed to delete total, error: %v", err)
	}
	return
}
