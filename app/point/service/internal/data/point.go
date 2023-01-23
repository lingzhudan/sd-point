package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/point/service/internal/biz"
)

type pointRepo struct {
	data *Data
	log  *log.Helper
}

// NewPointRepo .
func NewPointRepo(data *Data, logger log.Logger) biz.PointRepo {
	return &pointRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// point method

func (pr *pointRepo) ListPoint(ctx context.Context, cond *biz.PointCond) (points []*biz.Point, err error) {
	whereStage, args := cond.ParseCond()
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Where(whereStage, args...).
		Limit(int(cond.Count)).Offset(int(cond.Begin)).Find(&points).Error; err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) GetPoint(ctx context.Context, pid uint32) (point *biz.Point, err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).First(&point, pid).Error; err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) CreatePoint(ctx context.Context, point *biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Create(point).Error; err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdatePoint(ctx context.Context, point *biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{PID: point.PID}).Updates(point).Error; err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) DeletePoint(ctx context.Context, pid uint32) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Delete(&biz.Point{}, pid).Error; err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) GetRecord(ctx context.Context, rid uint32) (record *biz.Record, err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Record{RID: rid}).First(&record).Error; err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}
