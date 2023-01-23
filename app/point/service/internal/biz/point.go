package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	ListPoint(ctx context.Context, cond *PointCond) (points []*Point, err error)
	CreatePoint(ctx context.Context, point *Point) (err error)
	UpdatePoint(ctx context.Context, point *Point) (err error)
	DeletePoint(ctx context.Context, pid uint32) error
}

type Point struct {
	PID  uint32 `gorm:"column:pid;primaryKey;comment:自律点编号;"`
	UID  uint32 `gorm:"column:uid;index;comment:用户编号;"`
	Name string `gorm:"column:name;size:32;comment:自律点名称;"`
	Desc string `gorm:"column:desc;size:1024;comment:点数描述;"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;"`
}

func (p *Point) TableName() string {
	return "sd-point"
}

// point method

func (uc *PointUseCase) Get(ctx context.Context, pid uint32) (point *Point, err error) {
	if point, err = uc.repo.GetPoint(ctx, pid); err != nil {
		uc.log.Debugf("failed to get point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) List(ctx context.Context, cond *PointCond) (points []*Point, err error) {
	if points, err = uc.repo.ListPoint(ctx, cond); err != nil {
		uc.log.Debugf("failed to get points, error: %v", err)
	}
	return
}

func (uc *PointUseCase) Create(ctx context.Context, point *Point) (err error) {
	if err = uc.repo.CreatePoint(ctx, point); err != nil {
		uc.log.Debugf("failed to create point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) Update(ctx context.Context, point *Point) (err error) {
	if err = uc.repo.UpdatePoint(ctx, point); err != nil {
		uc.log.Debugf("failed to update point, error: %v", err)
	}
	return
}

func (uc *PointUseCase) Delete(ctx context.Context, pid uint32) (err error) {
	if err = uc.repo.DeletePoint(ctx, pid); err != nil {
		uc.log.Errorf("failed to delete point, error: %v", err)
	}
	return
}
