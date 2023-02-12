package service

import (
	"context"

	v1 "sd-point/api/point/service/v1"
	"sd-point/app/point/service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

// 点数记录方法

func (s *PointService) CreateRecord(ctx context.Context, req *v1.CreateRecordRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	var records []*biz.Record
	var ptMap = make(map[uint32]int32)
	r := req.Record
	nr := &biz.Record{
		PID:       r.Pid,
		ClickedAt: time.Unix(int64(r.ClickedAt), 0),
		Num:       r.Num,
		Desc:      r.Desc,
	}
	ptMap[r.Pid] = r.Num
	records = append(records, nr)
	if err = s.rc.Create(ctx, nr); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if _, err = s.tc.Incr(ctx, ptMap); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeleteRecord(ctx context.Context, req *v1.DeleteRecordRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	var r *biz.Record
	if r, err = s.rc.Get(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	if err = s.rc.Delete(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if _, err = s.tc.Decr(ctx, map[uint32]int32{r.PID: r.Num}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) UpdateRecord(ctx context.Context, req *v1.UpdateRecordRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	var r *biz.Record
	if r, err = s.rc.Get(ctx, req.Record.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	if err = s.rc.Update(ctx, &biz.Record{
		RID:  req.Record.Rid,
		Num:  req.Record.Num,
		Desc: req.Record.Desc,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	n := req.Record.Num - r.Num
	if n > 0 {
		if _, err = s.tc.Incr(ctx, map[uint32]int32{r.PID: r.Num}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	} else if n < 0 {
		if _, err = s.tc.Decr(ctx, map[uint32]int32{r.PID: r.Num}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	}
	return
}

func (s *PointService) ListRecord(ctx context.Context, req *v1.ListRecordRequest) (rep *v1.ListRecordReply, err error) {
	rep = &v1.ListRecordReply{Finished: true}
	var records []*biz.Record
	if records, err = s.rc.List(ctx, &biz.RecordCond{
		Begin:        req.Begin,
		Count:        req.Count + 1,
		RIDs:         req.Rids,
		PIDs:         req.Pids,
		MinClickedAt: req.MinClickedAt,
		MaxClickedAt: req.MaxClickedAt,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	for i, r := range records {
		if i >= int(req.Count) {
			rep.Finished = false
			break
		}
		rep.Records = append(rep.Records, &v1.Record{
			Rid:       r.RID,
			Pid:       r.PID,
			Num:       r.Num,
			ClickedAt: uint64(r.ClickedAt.Unix()),
			Desc:      r.Desc,
			CreatedAt: uint64(r.CreatedAt.Unix()),
			UpdatedAt: uint64(r.UpdatedAt.Unix()),
			DeletedAt: uint64(r.DeletedAt.Time.Unix()),
		})
	}
	rep.Count = uint32(len(rep.Records))
	return
}
