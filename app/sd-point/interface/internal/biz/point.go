package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type PointUseCase struct {
	repo PointRepo
	log  *log.Helper
}

func NewPointUsecase(repo PointRepo, logger log.Logger) *PointUseCase {
	return &PointUseCase{
		repo: repo,
		log:  log.NewHelper(logger)}
}

type PointRepo interface {
	GetPoint(ctx context.Context, pid uint32) (point *Point, err error)
	ListPoint(ctx context.Context, cond *PointCond) (points []*Point, err error)
	CreatePoint(ctx context.Context, point *Point) (err error)
	UpdatePoint(ctx context.Context, point *Point) (err error)
	DeletePoint(ctx context.Context, pid uint32) error

	GetRecord(ctx context.Context, rid uint32) (record *Record, err error)
	ListRecord(ctx context.Context, cond *RecordCond) (records []*Record, err error)
	CreateRecord(ctx context.Context, record *Record) (err error)
	UpdateRecord(ctx context.Context, record *Record) (err error)
	DeleteRecord(ctx context.Context, rid uint32) (err error)
}

type Point struct {
	// 自律点编号
	PID uint32
	// 用户编号
	UID uint32
	// 自律点名称
	Name string
	// 点数总数
	Total int32
	// 点数描述
	Desc string

	// 创建时间
	CreatedAt uint64
	// 更新时间
	UpdatedAt uint64
	// 删除时间
	DeletedAt uint64
}

type Record struct {
	// 点数记录编号
	RID uint32
	// 自律点编号
	PID uint32
	// 点数次数
	Num int32
	// 记录描述
	Desc string
	// 点击时间，由用户上传
	ClickedAt uint64

	// 创建时间
	CreatedAt uint64
	// 更新时间
	UpdatedAt uint64
	// 删除时间
	DeletedAt uint64
}

// 点数方法

func (uc *PointUseCase) GetPoint(ctx context.Context, pid uint32) (point *Point, err error) {
	if point, err = uc.repo.GetPoint(ctx, pid); err != nil {
		uc.log.Errorf("failed to get point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) ListPoint(ctx context.Context, cond *PointCond) (points []*Point, err error) {
	if points, err = uc.repo.ListPoint(ctx, cond); err != nil {
		uc.log.Errorf("failed to get point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) CreatePoint(ctx context.Context, point *Point) (err error) {
	if err = uc.repo.CreatePoint(ctx, point); err != nil {
		uc.log.Errorf("failed to get point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) UpdatePoint(ctx context.Context, point *Point) (err error) {
	if err = uc.repo.UpdatePoint(ctx, point); err != nil {
		uc.log.Errorf("failed to get point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) DeletePoint(ctx context.Context, pid uint32) (err error) {
	if err = uc.repo.DeletePoint(ctx, pid); err != nil {
		uc.log.Errorf("failed to get point, error: %v", err)
	}
	return
}

// 点数记录方法

func (uc *PointUseCase) GetRecord(ctx context.Context, rid uint32) (record *Record, err error) {
	if record, err = uc.repo.GetRecord(ctx, rid); err != nil {
		uc.log.Debugf("failed to get record, error: %v", err)
	}
	return
}

func (uc *PointUseCase) ListRecord(ctx context.Context, cond *RecordCond) (records []*Record, err error) {
	if records, err = uc.repo.ListRecord(ctx, cond); err != nil {
		uc.log.Debugf("failed to get records, error: %v", err)
	}
	return
}

func (uc *PointUseCase) CreatRecorde(ctx context.Context, record *Record) (err error) {
	if err = uc.repo.CreateRecord(ctx, record); err != nil {
		uc.log.Debugf("failed to create records, error: %v", err)
	}
	return
}

func (uc *PointUseCase) UpdateRecord(ctx context.Context, record *Record) (err error) {
	if err = uc.repo.UpdateRecord(ctx, record); err != nil {
		uc.log.Debugf("failed to update record, error: %v", err)
	}
	return
}

func (uc *PointUseCase) DeleteRecord(ctx context.Context, rid uint32) (err error) {
	if err = uc.repo.DeleteRecord(ctx, rid); err != nil {
		uc.log.Debugf("failed to delete record, error: %v", err)
	}
	return
}
