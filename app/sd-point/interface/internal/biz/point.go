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
	ListPoint(ctx context.Context, begin, count, uid uint32) (points *ListPoint, err error)
	CreatePoint(ctx context.Context, uid uint32, name, desc string) (err error)
	UpdatePoint(ctx context.Context, pid uint32, name, desc string) (err error)
	DeletePoint(ctx context.Context, pid uint32) error

	GetRecord(ctx context.Context, rid uint32) (record *Record, err error)
	ListRecord(ctx context.Context, begin, count, pid uint32) (records *ListRecord, err error)
	CreateRecord(ctx context.Context, pid uint32, num int32, desc string, clickedAt uint64) (err error)
	UpdateRecord(ctx context.Context, rid uint32, num int32, desc string, clickedAt uint64) (err error)
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

type ListPoint struct {
	Finished bool
	Points   []*Point
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

type ListRecord struct {
	Finished bool
	Records  []*Record
}

// 点数方法

func (uc *PointUseCase) GetPoint(ctx context.Context, pid uint32) (point *Point, err error) {
	return uc.repo.GetPoint(ctx, pid)
}

func (uc *PointUseCase) ListPoint(ctx context.Context, begin, count, uid uint32) (points *ListPoint, err error) {
	return uc.repo.ListPoint(ctx, begin, count, uid)
}

func (uc *PointUseCase) CreatePoint(ctx context.Context, uid uint32, name, desc string) (err error) {
	return uc.repo.CreatePoint(ctx, uid, name, desc)
}

func (uc *PointUseCase) UpdatePoint(ctx context.Context, pid uint32, name, desc string) (err error) {
	return uc.repo.UpdatePoint(ctx, pid, name, desc)
}

func (uc *PointUseCase) DeletePoint(ctx context.Context, pid uint32) (err error) {
	return uc.repo.DeletePoint(ctx, pid)
}

// 点数记录方法

func (uc *PointUseCase) GetRecord(ctx context.Context, rid uint32) (record *Record, err error) {
	return uc.repo.GetRecord(ctx, rid)
}

func (uc *PointUseCase) ListRecord(ctx context.Context, begin, count, pid uint32) (records *ListRecord, err error) {
	return uc.repo.ListRecord(ctx, begin, count, pid)
}

func (uc *PointUseCase) CreatRecord(ctx context.Context, pid uint32, num int32, desc string, clickedAt uint64) (err error) {
	return uc.repo.CreateRecord(ctx, pid, num, desc, clickedAt)
}

func (uc *PointUseCase) UpdateRecord(ctx context.Context, rid uint32, num int32, desc string, clickedAt uint64) (err error) {
	return uc.repo.UpdateRecord(ctx, rid, num, desc, clickedAt)
}

func (uc *PointUseCase) DeleteRecord(ctx context.Context, rid uint32) (err error) {
	return uc.repo.DeleteRecord(ctx, rid)
}
