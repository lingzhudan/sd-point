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

// point method

func (pr *pointRepo) ListPoint(ctx context.Context, cond *biz.PointCond) (points []*biz.Point, err error) {
	whereStage, args := cond.ParseCond()
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Where(whereStage, args...).Limit(cond.Count).Offset(cond.Begin).Find(&points).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) GetPoint(ctx context.Context, pid int32) (point *biz.Point, err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{PID: pid}).First(&point).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) CreatePoint(ctx context.Context, point *biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Create(point).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdatePoint(ctx context.Context, point *biz.Point) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{PID: point.PID}).Updates(point).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) DeletePoint(ctx context.Context, pid int32) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.Point{}).Delete(&biz.Point{}, pid).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) GetRecord(ctx context.Context, rid int32) (record *biz.PointRecord, err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.PointRecord{RID: rid}).First(&record).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
	}
	return
}

// record method

func (pr *pointRepo) ListRecord(ctx context.Context, cond *biz.RecordCond) (records []*biz.PointRecord, err error) {
	whereStage, args := cond.ParseCond()
	if err = pr.data.db.WithContext(ctx).Model(&biz.PointRecord{}).Where(whereStage, args...).Limit(cond.Count).Offset(cond.Begin).Find(&records).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
	}
	return
}

func (pr *pointRepo) CreateRecords(ctx context.Context, records []*biz.PointRecord) (err error) {
	if err = pr.data.db.WithContext(ctx).Model(&biz.PointRecord{}).Create(&records).Error; err != nil {
		pr.log.Debugf("db error: %v", err)
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

// redis method

// IncrBy 新增点数总数
func (pr *pointRepo) IncrBy(ctx context.Context, key string, value int64) (result int64, err error) {
	cmd := pr.data.rdb.IncrBy(ctx, key, value)
	if result, err = cmd.Result(); err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

// DecrBy 新增点数总数
func (pr *pointRepo) DecrBy(ctx context.Context, key string, value int64) (result int64, err error) {
	cmd := pr.data.rdb.DecrBy(ctx, key, value)
	if result, err = cmd.Result(); err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

// MSet 批量设置redis数据
//
// 例：
//
//	err := MSet(context.Background(), map[string]interface{"key": 100, "key2": "100"})
func (pr *pointRepo) MSet(ctx context.Context, data map[string]interface{}) (err error) {
	if err = pr.data.rdb.MSet(ctx, data).Err(); err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}

// MGet 批量获取redis数据
//
// 例：
//
//	values, err := MGet(context.Background(), []string{"key", "key2"})
//
// values: []interface{100, "100"}, err: nil
func (pr *pointRepo) MGet(ctx context.Context, keys []string) (values []interface{}, err error) {
	if values, err = pr.data.rdb.MGet(ctx, keys...).Result(); err != nil {
		pr.log.Errorf("rdb error: %v", err)
	}
	return
}

// Del 删除redis数据
func (pr *pointRepo) Del(ctx context.Context, keys []string) (err error) {
	if _, err = pr.data.rdb.Del(ctx, keys...).Result(); err != nil {
		pr.log.Errorf("db error: %v", err)
	}
	return
}
