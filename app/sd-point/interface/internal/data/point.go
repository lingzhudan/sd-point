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
	rep, err := pr.data.pc.GetPoint(ctx, &pointv1.GetPointRequest{Pid: pid})
	if err != nil {
		pr.log.Errorf("grpc client error: %v", err)
		return
	}
	p := rep.Point
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

func (pr *pointRepo) ListPoint(ctx context.Context, begin, count, uid uint32) (points *biz.ListPoint, err error) {
	rep, err := pr.data.pc.ListPoint(ctx, &pointv1.ListPointRequest{
		Begin: begin,
		Count: count,
		Uid:   uid,
	})
	if err != nil {
		pr.log.Errorf("grpc client error: %v", err)
		return
	}
	points = &biz.ListPoint{Finished: rep.Finished}
	for _, p := range rep.Points {
		points.Points = append(points.Points, &biz.Point{
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

func (pr *pointRepo) CreatePoint(ctx context.Context, uid uint32, name, desc string) (err error) {
	if _, err = pr.data.pc.CreatePoint(ctx, &pointv1.CreatePointRequest{
		Uid:  uid,
		Name: name,
		Desc: desc,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdatePoint(ctx context.Context, pid uint32, name, desc string) (err error) {
	if _, err = pr.data.pc.UpdatePoint(ctx, &pointv1.UpdatePointRequest{
		Pid:  pid,
		Name: name,
		Desc: desc,
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
	rep, err := pr.data.pc.GetRecord(ctx, &pointv1.GetRecordRequest{
		Rid: rid,
	})
	if err != nil {
		pr.log.Errorf("grpc client error: %v", err)
		return
	}
	r := rep.Record
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

func (pr *pointRepo) ListRecord(ctx context.Context, begin, count, pid uint32) (records *biz.ListRecord, err error) {
	rep, err := pr.data.pc.ListRecord(ctx, &pointv1.ListRecordRequest{
		Begin: begin,
		Count: count,
		Pid:   pid,
	})
	if err != nil {
		pr.log.Errorf("grpc client error: %v", err)
		return
	}
	records = &biz.ListRecord{Finished: rep.Finished}
	for _, r := range rep.Records {
		records.Records = append(records.Records, &biz.Record{
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

func (pr *pointRepo) CreateRecord(ctx context.Context, pid uint32, num int32, desc string, clickedAt uint64) (err error) {
	if _, err = pr.data.pc.CreateRecord(ctx, &pointv1.CreateRecordRequest{
		Pid:       pid,
		Num:       num,
		ClickedAt: clickedAt,
		Desc:      desc,
	}); err != nil {
		pr.log.Errorf("grpc client error: %v", err)
	}
	return
}

func (pr *pointRepo) UpdateRecord(ctx context.Context, rid uint32, num int32, desc string, clickedAt uint64) (err error) {
	if _, err = pr.data.pc.UpdateRecord(ctx, &pointv1.UpdateRecordRequest{
		Rid:       rid,
		Num:       num,
		ClickedAt: clickedAt,
		Desc:      desc,
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
