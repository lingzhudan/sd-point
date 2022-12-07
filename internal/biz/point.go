package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type Point struct {
	PID       int32     `gorm:"pid;primaryKey;comment:自律点编号;"`
	UID       int32     `gorm:"uid;index;comment:用户编号;"`
	ClickedAt time.Time `gorm:"clicked_at;comment:点击时间，由用户上传;"`
	Num       int16     `gorm:"num;comment:点数次数;"`
	Desc      string    `gorm:"desc;size:1024;comment:点数描述;"`

	CreatedAt time.Time      `gorm:"created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at;comment:删除时间;"`
}

func (p *Point) TableName() string {
	return "sd_point"
}

type PointUsecase struct {
	repo PointRepo
	log  *log.Helper
}

func NewPointUsecase(repo PointRepo, logger log.Logger) *PointUsecase {
	return &PointUsecase{repo: repo, log: log.NewHelper(logger)}
}

type PointRepo interface {
	ListPoint(ctx context.Context) (points []*Point, err error)
	GetPoint(ctx context.Context, pid int32) (point *Point, err error)
	CreatePoints(ctx context.Context, point []*Point) (err error)
	UpdatePoint(ctx context.Context, pid int32, point *Point) (err error)
	DeletePoint(ctx context.Context, pid int32) error
}

func (uc *PointUsecase) Get(ctx context.Context, pid int32) (point *Point, err error) {
	if point, err = uc.repo.GetPoint(ctx, pid); err != nil {
		uc.log.Debugf("failed to get point, error: %v", err)
	}
	return
}

func (uc *PointUsecase) List(ctx context.Context) (points []*Point, err error) {
	if points, err = uc.repo.ListPoint(ctx); err != nil {
		uc.log.Debugf("failed to get points, error: %v", err)
	}
	return
}

func (uc *PointUsecase) Create(ctx context.Context, point []*Point) (err error) {
	if err = uc.repo.CreatePoints(ctx, point); err != nil {
		uc.log.Debugf("failed to create point, error: %v", err)
	}
	return
}

func (uc *PointUsecase) Update(ctx context.Context, pid int32, point *Point) (err error) {
	if err = uc.repo.UpdatePoint(ctx, pid, point); err != nil {
		uc.log.Debugf("failed to update point, error: %v", err)
	}
	return
}

func (uc *PointUsecase) Delete(ctx context.Context, pid int32) (err error) {
	if err = uc.repo.DeletePoint(ctx, pid); err != nil {
		uc.log.Debugf("failed to delete point, error: %v", err)
	}
	return
}
