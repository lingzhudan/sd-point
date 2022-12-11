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

func (pr *pointRepo) ListPoint(ctx context.Context, cond *biz.PointCond) (points []*biz.Point, err error) {
	whereStage, args := cond.ParseCond()
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Where(whereStage, args...).Limit(cond.Count).Offset(cond.Begin).Find(&points).Error; err != nil {
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

func (pr *pointRepo) CreatePoint(ctx context.Context, point *biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Create(point).Error; err != nil {
		pr.log.Debugf("failed to create point, db error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdatePoint(ctx context.Context, point *biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Unscoped().Model(&biz.Point{PID: point.PID}).Updates(point).Error; err != nil {
		pr.log.Debugf("failed to update point, db error: %v", err)
	}
	return
}

func (pr *pointRepo) DeletePoint(ctx context.Context, pid int32) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Delete(&biz.Point{}, pid).Error; err != nil {
		pr.log.Debugf("failed to delete point, db error: %v", err)
	}
	return
}

func (pr *pointRepo) ListRecord(ctx context.Context, cond *biz.RecordCond) (records []*biz.PointRecord, err error) {
	whereStage, args := cond.ParseCond()
	if err = pr.data.db.WithContext(ctx).Model(&biz.PointRecord{}).Where(whereStage, args...).Limit(cond.Count).Offset(cond.Begin).Find(&records).Error; err != nil {
		pr.log.Debugf("failed to get records, db error: %v", err)
	}
	return
}

func (pr *pointRepo) CreateRecords(ctx context.Context, records []*biz.PointRecord) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.PointRecord{}).Create(&records).Error; err != nil {
		pr.log.Debugf("failed to create record, db error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdateRecord(ctx context.Context, record *biz.PointRecord) (err error) {
	if err = pr.data.db.WithContext(ctx).Unscoped().Model(&biz.PointRecord{RID: record.RID}).Updates(record).Error; err != nil {
		pr.log.Debugf("failed to update record, db error: %v", err)
	}
	return
}

func (pr *pointRepo) DeleteRecord(ctx context.Context, rid int32) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.PointRecord{}).Delete(&biz.PointRecord{}, rid).Error; err != nil {
		pr.log.Debugf("failed to delete records, db error: %v", err)
	}
	return
}
