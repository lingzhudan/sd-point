package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	v1 "sd-point/api/point/service/v1"
	"time"
)

// 点数记录方法

func (s *PointService) CreateRecord(ctx context.Context, req *v1.CreateRecordRequest) (rep *emptypb.Empty, err error) {
	return new(emptypb.Empty), s.uc.CreateRecord(ctx, req.Pid, req.Num, req.Desc, time.Unix(int64(req.ClickedAt), 0))
}

func (s *PointService) DeleteRecord(ctx context.Context, req *v1.DeleteRecordRequest) (rep *emptypb.Empty, err error) {
	return new(emptypb.Empty), s.uc.DeleteRecord(ctx, req.Rid)
}

func (s *PointService) UpdateRecord(ctx context.Context, req *v1.UpdateRecordRequest) (rep *emptypb.Empty, err error) {
	return new(emptypb.Empty), s.uc.UpdateRecord(ctx, req.Rid, req.Num, req.Desc)
}

func (s *PointService) GetRecord(ctx context.Context, req *v1.GetRecordRequest) (rep *v1.GetRecordReply, err error) {
	rep = new(v1.GetRecordReply)
	r, err := s.uc.GetRecord(ctx, req.Rid)
	if err != nil {
		return
	}
	rep.Record = &v1.Record{
		Rid:       r.RID,
		Pid:       r.PID,
		Num:       r.Num,
		ClickedAt: uint64(r.ClickedAt.Unix()),
		Desc:      r.Desc,
		CreatedAt: uint64(r.CreatedAt.Unix()),
		UpdatedAt: uint64(r.UpdatedAt.Unix()),
	}
	return
}

func (s *PointService) ListRecord(ctx context.Context, req *v1.ListRecordRequest) (rep *v1.ListRecordReply, err error) {
	rep = &v1.ListRecordReply{Finished: true}
	records, err := s.uc.ListRecord(ctx, int(req.Begin), int(req.Count+1), req.Pid)
	if err != nil {
		return
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
		})
	}
	rep.Count = uint32(len(rep.Records))
	return
}
