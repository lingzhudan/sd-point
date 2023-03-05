package data

import (
	"context"
	"gorm.io/gorm"
	"sd-point/app/point/service/internal/biz"
	"time"
)

type Record struct {
	RID       uint32    `gorm:"column:rid;primaryKey;comment:点数记录编号;"`
	PID       uint32    `gorm:"column:pid;index;comment:自律点编号;"`
	Num       int32     `gorm:"column:num;comment:点数次数;"`
	Desc      string    `gorm:"column:desc;size:1024;comment:记录描述;"`
	ClickedAt time.Time `gorm:"column:clicked_at;comment:点击时间，由用户上传;"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;"`
}

func (r *Record) TableName() string {
	return "point_record"
}

func NewBizRecord(dR *Record) (bR *biz.Record) {
	bR = &biz.Record{
		RID:       dR.RID,
		PID:       dR.PID,
		Num:       dR.Num,
		Desc:      dR.Desc,
		ClickedAt: dR.ClickedAt,
		CreatedAt: dR.CreatedAt,
		UpdatedAt: dR.UpdatedAt,
	}
	return
}

// record method

func (r *pointRepo) GetRecord(ctx context.Context, rid uint32) (bR *biz.Record, err error) {
	dR := new(Record)
	if err = r.data.db.
		WithContext(ctx).
		First(&dR, rid).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	bR = NewBizRecord(dR)
	return
}

func (r *pointRepo) ListRecord(ctx context.Context, begin int, count int, pid uint32) (bRs []*biz.Record, err error) {
	dRs := make([]*Record, 0)
	if err = r.data.db.
		WithContext(ctx).
		Where(&Record{PID: pid}).
		Limit(count).
		Offset(begin).
		Find(&dRs).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	for _, dR := range dRs {
		bRs = append(bRs, NewBizRecord(dR))
	}
	return
}

func (r *pointRepo) CreateRecord(
	ctx context.Context,
	pid uint32,
	num int32,
	desc string,
	clickedAt time.Time,
) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Create(&Record{
			PID:       pid,
			Num:       num,
			Desc:      desc,
			ClickedAt: clickedAt,
		}).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *pointRepo) UpdateRecord(
	ctx context.Context,
	rid uint32,
	num int32,
	desc string,
	clickedAt time.Time,
) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Model(&Record{RID: rid}).
		Updates(&Record{
			Num:       num,
			Desc:      desc,
			ClickedAt: clickedAt,
		}).Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *pointRepo) DeleteRecord(ctx context.Context, rid uint32) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Delete(&Record{}, rid).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *pointRepo) DeleteRecords(ctx context.Context, pid uint32) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Where(&Record{PID: pid}).
		Delete(&Record{}).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}
