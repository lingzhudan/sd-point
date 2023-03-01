package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pointv1 "sd-point/api/point/service/v1"
	"sd-point/app/sd-point/interface/internal/biz"
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

func (pr *pointRepo) GetPoint(ctx context.Context, pid uint32) (point *biz.Point, err error) {
	var reply *pointv1.GetPointReply
	if reply, err = pr.data.pc.GetPoint(ctx, &pointv1.GetPointRequest{Pid: pid}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	p := reply.Point
	point = &biz.Point{
		PID:       p.Pid,
		Name:      p.Name,
		Total:     p.Total,
		Desc:      p.Desc,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	return
}

func (pr *pointRepo) ListPoint(ctx context.Context, cond *biz.PointCond) (points []*biz.Point, err error) {
	var reply *pointv1.ListPointReply
	if reply, err = pr.data.pc.ListPoint(ctx, &pointv1.ListPointRequest{
		Begin: cond.Begin,
		Count: cond.Count + 1,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	for _, p := range reply.Points {
		points = append(points, &biz.Point{
			PID:       p.Pid,
			Name:      p.Name,
			Total:     p.Total,
			Desc:      p.Desc,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}
	return
}

func (pr *pointRepo) CreatePoint(ctx context.Context, point *biz.Point) (err error) {
	if _, err = pr.data.pc.CreatePoint(ctx, &pointv1.CreatePointRequest{
		Uid:  point.UID,
		Name: point.Name,
		Desc: point.Desc,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdatePoint(ctx context.Context, point *biz.Point) (err error) {
	if _, err = pr.data.pc.UpdatePoint(ctx, &pointv1.UpdatePointRequest{
		Pid:  point.PID,
		Name: point.Name,
		Desc: point.Desc,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (pr *pointRepo) DeletePoint(ctx context.Context, pid uint32) (err error) {
	if _, err = pr.data.pc.DeletePoint(ctx, &pointv1.DeletePointRequest{Pid: pid}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (pr *pointRepo) GetRecord(ctx context.Context, rid uint32) (record *biz.Record, err error) {
	var reply *pointv1.GetRecordReply
	if reply, err = pr.data.pc.GetRecord(ctx, &pointv1.GetRecordRequest{
		Rid: rid,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	r := reply.Record
	record = &biz.Record{
		RID:       r.Rid,
		PID:       r.Pid,
		Num:       r.Num,
		Desc:      r.Desc,
		ClickedAt: r.ClickedAt,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		DeletedAt: r.DeletedAt,
	}
	return record, nil
}

func (pr *pointRepo) ListRecord(ctx context.Context, cond *biz.RecordCond) (records []*biz.Record, err error) {
	var reply *pointv1.ListRecordReply
	if reply, err = pr.data.pc.ListRecord(ctx, &pointv1.ListRecordRequest{
		Begin: cond.Begin,
		Count: cond.Count,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	for _, r := range reply.Records {
		records = append(records, &biz.Record{
			RID:       r.Rid,
			PID:       r.Pid,
			Num:       r.Num,
			Desc:      r.Desc,
			ClickedAt: r.ClickedAt,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
			DeletedAt: r.DeletedAt,
		})
	}
	return
}

func (pr *pointRepo) CreateRecord(ctx context.Context, record *biz.Record) (err error) {
	if _, err = pr.data.pc.CreateRecord(ctx, &pointv1.CreateRecordRequest{
		Pid:       record.PID,
		Num:       record.Num,
		ClickedAt: record.ClickedAt,
		Desc:      record.Desc,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdateRecord(ctx context.Context, record *biz.Record) (err error) {
	if _, err = pr.data.pc.UpdateRecord(ctx, &pointv1.UpdateRecordRequest{
		Rid:       record.RID,
		Num:       record.Num,
		ClickedAt: record.ClickedAt,
		Desc:      record.Desc,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (pr *pointRepo) DeleteRecord(ctx context.Context, rid uint32) (err error) {
	if _, err = pr.data.pc.DeleteRecord(ctx, &pointv1.DeleteRecordRequest{Rid: rid}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}
