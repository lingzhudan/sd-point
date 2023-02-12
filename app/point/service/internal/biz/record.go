package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type RecordRepo interface {
	GetRecord(ctx context.Context, rid uint32) (record *Record, err error)
	ListRecord(ctx context.Context, cond *RecordCond) (records []*Record, err error)
	CreateRecord(ctx context.Context, record *Record) (err error)
	UpdateRecord(ctx context.Context, record *Record) (err error)
	DeleteRecord(ctx context.Context, rid uint32) (err error)
}

type RecordUseCase struct {
	repo RecordRepo
	TotalRepo
	log *log.Helper
}

func NewRecordUseCase(repo RecordRepo, logger log.Logger) *RecordUseCase {
	return &RecordUseCase{
		repo: repo,
		log:  log.NewHelper(logger)}
}

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
	return "sd_point_record"
}

// point record method

func (uc *RecordUseCase) Get(ctx context.Context, rid uint32) (record *Record, err error) {
	if record, err = uc.repo.GetRecord(ctx, rid); err != nil {
		uc.log.Errorf("failed to get record, error: %v", err)
	}
	return
}

func (uc *RecordUseCase) List(ctx context.Context, cond *RecordCond) (records []*Record, err error) {
	if records, err = uc.repo.ListRecord(ctx, cond); err != nil {
		uc.log.Errorf("failed to get records, error: %v", err)
	}
	return
}

func (uc *RecordUseCase) Create(ctx context.Context, record *Record) (err error) {
	if err = uc.repo.CreateRecord(ctx, record); err != nil {
		uc.log.Errorf("failed to create records, error: %v", err)
	}
	return
}

func (uc *RecordUseCase) Update(ctx context.Context, record *Record) (err error) {
	if err = uc.repo.UpdateRecord(ctx, record); err != nil {
		uc.log.Errorf("failed to update record, error: %v", err)
	}
	return
}

func (uc *RecordUseCase) Delete(ctx context.Context, rid uint32) (err error) {
	if err = uc.repo.DeleteRecord(ctx, rid); err != nil {
		uc.log.Errorf("failed to delete record, error: %v", err)
	}
	return
}
