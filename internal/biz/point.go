package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type PointUsecase struct {
	repo PointRepo
	log  *log.Helper
}

func NewPointUsecase(repo PointRepo, logger log.Logger) *PointUsecase {
	return &PointUsecase{repo: repo, log: log.NewHelper(logger)}
}

type PointRepo interface {
	ListPoint(ctx context.Context, cond *PointCond) (points []*Point, err error)
	GetPoint(ctx context.Context, pid int32) (point *Point, err error)
	CreatePoint(ctx context.Context, point *Point) (err error)
	UpdatePoint(ctx context.Context, point *Point) (err error)
	DeletePoint(ctx context.Context, pid int32) error

	GetRecord(ctx context.Context, rid int32) (record *PointRecord, err error)
	ListRecord(ctx context.Context, cond *RecordCond) (records []*PointRecord, err error)
	CreateRecords(ctx context.Context, records []*PointRecord) (err error)
	UpdateRecord(ctx context.Context, record *PointRecord) (err error)
	DeleteRecord(ctx context.Context, rid int32) (err error)

	IncrBy(ctx context.Context, key string, value int64) (result int64, err error)
	DecrBy(ctx context.Context, key string, value int64) (result int64, err error)
	MSet(ctx context.Context, data map[string]interface{}) (err error)
	MGet(ctx context.Context, keys []string) (values []interface{}, err error)
	Del(ctx context.Context, keys []string) (err error)
}

type Point struct {
	PID  int32  `gorm:"column:pid;primaryKey;comment:自律点编号;"`
	UID  int32  `gorm:"column:uid;index;comment:用户编号;"`
	Name string `gorm:"column:name;size:32;comment:自律点名称;"`
	Desc string `gorm:"column:desc;size:1024;comment:点数描述;"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;"`
}

func (p *Point) TableName() string {
	return "sd_point"
}

type PointRecord struct {
	RID       int32     `gorm:"column:rid;primaryKey;comment:点数记录编号;"`
	PID       int32     `gorm:"column:pid;index;comment:自律点编号;"`
	Num       int16     `gorm:"column:num;comment:点数次数;"`
	Desc      string    `gorm:"column:desc;size:1024;comment:记录描述;"`
	ClickedAt time.Time `gorm:"column:clicked_at;comment:点击时间，由用户上传;"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;"`
}

func (p *PointRecord) TableName() string {
	return "sd_point_record"
}

// point method

func (uc *PointUsecase) GetPoint(ctx context.Context, pid int32) (point *Point, err error) {
	if point, err = uc.repo.GetPoint(ctx, pid); err != nil {
		uc.log.Debugf("failed to get point, error: %v", err)
	}
	return
}

func (uc *PointUsecase) ListPint(ctx context.Context, cond *PointCond) (points []*Point, err error) {
	if points, err = uc.repo.ListPoint(ctx, cond); err != nil {
		uc.log.Debugf("failed to get points, error: %v", err)
	}
	return
}

func (uc *PointUsecase) CreatePoint(ctx context.Context, point *Point) (err error) {
	if err = uc.repo.CreatePoint(ctx, point); err != nil {
		uc.log.Debugf("failed to create point, error: %v", err)
	}
	return
}

func (uc *PointUsecase) UpdatePoint(ctx context.Context, point *Point) (err error) {
	if err = uc.repo.UpdatePoint(ctx, point); err != nil {
		uc.log.Debugf("failed to update point, error: %v", err)
	}
	return
}

func (uc *PointUsecase) DeletePoint(ctx context.Context, pid int32) (err error) {
	if err = uc.repo.DeletePoint(ctx, pid); err != nil {
		uc.log.Errorf("failed to delete point, error: %v", err)
	}
	return
}

// point record method

func (uc *PointUsecase) GetRecord(ctx context.Context, rid int32) (record *PointRecord, err error) {
	if record, err = uc.repo.GetRecord(ctx, rid); err != nil {
		uc.log.Debugf("failed to get record, error: %v", err)
	}
	return
}

func (uc *PointUsecase) ListRecord(ctx context.Context, cond *RecordCond) (points []*PointRecord, err error) {
	if points, err = uc.repo.ListRecord(ctx, cond); err != nil {
		uc.log.Debugf("failed to get records, error: %v", err)
	}
	return
}

func (uc *PointUsecase) CreateRecords(ctx context.Context, records []*PointRecord) (err error) {
	if err = uc.repo.CreateRecords(ctx, records); err != nil {
		uc.log.Debugf("failed to create records, error: %v", err)
	}
	return
}

func (uc *PointUsecase) UpdateRecord(ctx context.Context, record *PointRecord) (err error) {
	if err = uc.repo.UpdateRecord(ctx, record); err != nil {
		uc.log.Debugf("failed to update record, error: %v", err)
	}
	return
}

func (uc *PointUsecase) DeleteRecord(ctx context.Context, rid int32) (err error) {
	if err = uc.repo.DeleteRecord(ctx, rid); err != nil {
		uc.log.Debugf("failed to delete record, error: %v", err)
	}
	return
}

// redis method

func (uc *PointUsecase) IncrPointTotal(ctx context.Context, data map[int32]int16) (result int64, err error) {
	for pid, num := range data {
		if result, err = uc.repo.IncrBy(ctx, strconv.Itoa(int(pid)), int64(num)); err != nil {
			uc.log.Errorf("failed to incr key, error: %v", err)
		}
	}
	return
}

func (uc *PointUsecase) DecrPointTotal(ctx context.Context, data map[int32]int16) (result int64, err error) {
	for pid, num := range data {
		if result, err = uc.repo.DecrBy(ctx, strconv.Itoa(int(pid)), int64(num)); err != nil {
			uc.log.Errorf("failed to decr key, error: %v", err)
		}
	}
	return
}

func (uc *PointUsecase) MSetPointTotal(ctx context.Context, pointTotals map[int32]int32) (err error) {
	var d = make(map[string]interface{})
	for k, v := range pointTotals {
		d[strconv.Itoa(int(k))] = v
	}
	if err = uc.repo.MSet(ctx, d); err != nil {
		uc.log.Errorf("failed to set keys, error: %v", err)
	}
	return
}

func (uc *PointUsecase) MGetPointTotal(ctx context.Context, pids []int32) (data map[int32]int32, err error) {
	var keys []string
	var values []interface{}
	for _, pid := range pids {
		keys = append(keys, strconv.Itoa(int(pid)))
	}
	if values, err = uc.repo.MGet(ctx, keys); err != nil {
		uc.log.Errorf("failed to get keys, error: %v", err)
	}
	for i := range values {
		data[pids[i]] = values[i].(int32)
	}
	return
}

func (uc *PointUsecase) DelPointTotal(ctx context.Context, pids []int32) (err error) {
	var keys []string
	for _, pid := range pids {
		keys = append(keys, strconv.Itoa(int(pid)))
	}
	if err = uc.repo.Del(ctx, keys); err != nil {
		uc.log.Errorf("failed to del keys, error: %v", err)
	}
	return
}
