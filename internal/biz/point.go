package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Point struct {
	PID       int64  `gorm:"pid;primaryKey;comment:自律点编号;"`
	UID       int64  `gorm:"uid;index;comment:用户编号;"`
	ClickedAt int64  `gorm:"clicked_at;comment:点击时间，由用户上传;"`
	Num       int64  `gorm:"num;comment:点数次数;"`
	Desc      string `gorm:"desc;comment:点数描述;"`
}

func (p *Point) TableName() string {
	return "sd_point"
}

type PointUsecase struct {
	repo   PointRepo
	logger log.Logger
}

func NewPointUsecase(repo PointRepo, logger log.Logger) *PointUsecase {
	return &PointUsecase{repo: repo, logger: logger}
}

type PointRepo interface {
	ListPoint(ctx context.Context) (points []*Point, err error)
	GetPoint(ctx context.Context, pid int64) (point *Point, err error)
	CreatePoint(ctx context.Context, point *Point) (err error)
	UpdatePoint(ctx context.Context, pid int64, point *Point) (err error)
	DeletePoint(ctx context.Context, pid int64) error
}

func (uc *PointUsecase) Get(ctx context.Context, pid int64) (point *Point, err error) {
	return uc.repo.GetPoint(ctx, pid)
}

func (uc *PointUsecase) List(ctx context.Context) (points []*Point, err error) {
	return uc.repo.ListPoint(ctx)
}

func (uc *PointUsecase) Create(ctx context.Context, point *Point) error {
	return uc.repo.CreatePoint(ctx, point)
}

func (uc *PointUsecase) Update(ctx context.Context, id int64, point *Point) error {
	return uc.repo.UpdatePoint(ctx, id, point)
}

func (uc *PointUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeletePoint(ctx, id)
}
