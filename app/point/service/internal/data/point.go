package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"sd-point/app/point/service/internal/biz"
	"time"
)

type pointRepo struct {
	data *Data

	log *log.Helper
}

// NewPointRepo .
func NewPointRepo(data *Data, logger log.Logger) biz.PointRepo {
	return &pointRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
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
	return "point"
}

func NewBizPoint(dP *Point) (bP *biz.Point) {
	bP = &biz.Point{
		PID:       dP.PID,
		UID:       dP.UID,
		Name:      dP.Name,
		Desc:      dP.Desc,
		CreatedAt: dP.CreatedAt,
		UpdatedAt: dP.UpdatedAt,
		DeletedAt: dP.DeletedAt.Time,
	}
	return
}

// point method

func (r *pointRepo) ListPoint(ctx context.Context, begin int, count int, uid uint32) (bPs []*biz.Point, err error) {
	dPs := make([]*Point, 0)
	if err = r.data.db.
		WithContext(ctx).
		Where("`uid` = ?", uid).
		Limit(count).
		Offset(begin).
		Find(&dPs).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	for _, dP := range dPs {
		bPs = append(bPs, NewBizPoint(dP))
	}
	return
}

func (r *pointRepo) GetPoint(ctx context.Context, pid uint32) (bP *biz.Point, err error) {
	dP := new(Point)
	if err = r.data.db.
		WithContext(ctx).
		First(&dP, pid).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	bP = NewBizPoint(dP)
	return
}

func (r *pointRepo) CreatePoint(ctx context.Context, uid uint32, name string, desc string) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Create(&Point{
			UID:  uid,
			Name: name,
			Desc: desc,
		}).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *pointRepo) UpdatePoint(ctx context.Context, pid uint32, name string, desc string) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Model(&Point{PID: pid}).
		Updates(&Point{
			Name: name,
			Desc: desc,
		}).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *pointRepo) DeletePoint(ctx context.Context, pid uint32) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Delete(&Point{}, pid).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}
