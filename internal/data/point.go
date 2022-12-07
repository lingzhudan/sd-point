package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/internal/biz"
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

func (pr *pointRepo) ListPoint(ctx context.Context) (points []*biz.Point, err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Find(&points).Error; err != nil {
		pr.log.Debugf("failed to find points, db error: %v", err)
	}
	return
}

func (pr *pointRepo) GetPoint(ctx context.Context, pid int32) (point *biz.Point, err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{PID: pid}).Find(&point).Error; err != nil {
		pr.log.Debugf("failed to find point, db error: %v", err)
	}
	return
}

func (pr *pointRepo) CreatePoints(ctx context.Context, points []*biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Create(points).Error; err != nil {
		pr.log.Debugf("failed to create point, db error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdatePoint(ctx context.Context, pid int32, point *biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{PID: pid}).Updates(point).Error; err != nil {
		pr.log.Debugf("failed to update point, db error: %v", err)
	}
	return
}

func (pr *pointRepo) DeletePoint(ctx context.Context, pid int32) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{PID: pid}).Delete(&biz.Point{}).Error; err != nil {
		pr.log.Debugf("failed to delete point, db error: %v", err)
	}
	return
}
