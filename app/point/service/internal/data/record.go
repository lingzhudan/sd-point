package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/point/service/internal/biz"
)

type recordRepo struct {
	data *Data
	log  *log.Helper
}

// NewRecordRepo .
func NewRecordRepo(data *Data, logger log.Logger) biz.RecordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// record method

func (r *recordRepo) GetRecord(ctx context.Context, rid uint32) (record *biz.Record, err error) {
	if err = r.data.db.WithContext(ctx).Model(&biz.Record{}).First(&record, rid).Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *recordRepo) ListRecord(ctx context.Context, cond *biz.RecordCond) (records []*biz.Record, err error) {
	whereStage, args := cond.ParseCond()
	if err = r.data.db.WithContext(ctx).Model(&biz.Record{}).Where(whereStage, args...).
		Limit(int(cond.Count)).Offset(int(cond.Begin)).Find(&records).Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *recordRepo) CreateRecord(ctx context.Context, record *biz.Record) (err error) {
	if err = r.data.db.WithContext(ctx).Model(&biz.Record{}).Create(&record).Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *recordRepo) UpdateRecord(ctx context.Context, record *biz.Record) (err error) {
	if err = r.data.db.WithContext(ctx).Unscoped().Model(&biz.Record{RID: record.RID}).Updates(record).Error; err != nil {
		r.log.Errorf("failed to update record, db error: %v", err)
	}
	return
}

func (r *recordRepo) DeleteRecord(ctx context.Context, rid uint32) (err error) {
	if err = r.data.db.WithContext(ctx).Model(&biz.Record{}).Delete(&biz.Record{}, rid).Error; err != nil {
		r.log.Errorf("failed to delete records, db error: %v", err)
	}
	return
}
