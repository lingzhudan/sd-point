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
	return
}

func (pr *pointRepo) GetPoint(ctx context.Context, pid int64) (point *biz.Point, err error) {
	return
}

func (pr *pointRepo) CreatePoint(ctx context.Context, point *biz.Point) (err error) {
	return
}

func (pr *pointRepo) UpdatePoint(ctx context.Context, pid int64, point *biz.Point) (err error) {
	return
}

func (pr *pointRepo) DeletePoint(ctx context.Context, pid int64) (err error) {
	return
}
