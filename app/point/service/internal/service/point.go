package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"sd-point/api/point/service/v1"
	"sd-point/app/point/service/internal/biz"
	"time"
)

type PointService struct {
	v1.UnimplementedPointServer

	pc  *biz.PointUseCase
	rc  *biz.RecordUseCase
	tc  *biz.TotalUseCase
	log *log.Helper
}

func NewPointService(pc *biz.PointUseCase, rc *biz.RecordUseCase,
	tc *biz.TotalUseCase, logger log.Logger) *PointService {
	return &PointService{
		pc:  pc,
		rc:  rc,
		tc:  tc,
		log: log.NewHelper(logger)}
}

// 点数方法

func (s *PointService) CreatePoint(ctx context.Context, req *v1.CreatePointRequest) (reply *emptypb.Empty, err error) {
	point := &biz.Point{
		PID:  req.Point.Pid,
		UID:  req.Point.Uid,
		Name: req.Point.Name,
		Desc: req.Point.Desc,
	}

	if err = s.pc.Create(ctx, point); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}

	return &emptypb.Empty{}, nil
}

func (s *PointService) UpdatePoint(ctx context.Context, req *v1.UpdatePointRequest) (reply *emptypb.Empty, err error) {
	point := &biz.Point{
		PID:  req.Point.Pid,
		UID:  0,
		Name: req.Point.Name,
		Desc: req.Point.Desc,
	}
	if err = s.pc.Update(ctx, point); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeletePoint(ctx context.Context, req *v1.DeletePointRequest) (reply *emptypb.Empty, err error) {
	if err = s.pc.Delete(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if err = s.tc.Del(ctx, []uint32{req.Pid}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) GetPoint(ctx context.Context, req *v1.GetPointRequest) (reply *v1.GetPointReply, err error) {
	reply = &v1.GetPointReply{}
	var point *biz.Point
	if point, err = s.pc.Get(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	var ptMap map[uint32]int32
	if ptMap, err = s.tc.MGet(ctx, []uint32{req.Pid}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	s.log.Debugf("pt: %v", ptMap[req.Pid])
	reply.Point = &v1.GetPointReply_Point{
		Pid:       point.PID,
		Name:      point.Name,
		Total:     ptMap[point.PID],
		Desc:      point.Desc,
		CreatedAt: uint64(point.CreatedAt.Unix()),
		UpdatedAt: uint64(point.UpdatedAt.Unix()),
		DeletedAt: uint64(point.DeletedAt.Time.Unix()),
	}
	return
}

func (s *PointService) ListPoint(ctx context.Context, req *v1.ListPointRequest) (reply *v1.ListPointReply, err error) {
	reply = &v1.ListPointReply{Finished: true}
	var points []*biz.Point
	if points, err = s.pc.List(ctx, &biz.PointCond{
		Begin: req.Begin,
		Count: req.Count + 1,
		PIDs:  req.Pids,
		UIDs:  []uint32{0},
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	var pids []uint32
	for _, p := range points {
		pids = append(pids, p.PID)
	}
	var ptMap map[uint32]int32
	ptMap, err = s.tc.MGet(ctx, pids)
	if err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	for i, p := range points {
		if i >= int(req.Count) {
			reply.Finished = false
			break
		}
		reply.Points = append(reply.Points, &v1.GetPointReply_Point{
			Pid:       p.PID,
			Name:      p.Name,
			Total:     ptMap[p.PID],
			Desc:      p.Desc,
			CreatedAt: uint64(p.CreatedAt.Unix()),
			UpdatedAt: uint64(p.UpdatedAt.Unix()),
			DeletedAt: uint64(p.DeletedAt.Time.Unix()),
		})
	}
	reply.Count = uint32(len(reply.Points))
	return
}

// 点数记录方法

func (s *PointService) CreateRecord(ctx context.Context, req *v1.CreateRecordRequest) (reply *emptypb.Empty, err error) {
	var records []*biz.Record
	var ptMap = make(map[uint32]int32)
	r := req.Record
	newRecord := &biz.Record{
		PID:       r.Pid,
		ClickedAt: time.Unix(int64(r.ClickedAt), 0),
		Num:       r.Num,
		Desc:      r.Desc,
	}
	ptMap[r.Pid] = r.Num
	records = append(records, newRecord)
	if err = s.rc.Create(ctx, newRecord); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if _, err = s.tc.Incr(ctx, ptMap); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeleteRecord(ctx context.Context, req *v1.DeleteRecordRequest) (reply *emptypb.Empty, err error) {
	var record *biz.Record
	if record, err = s.rc.Get(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if err = s.rc.Delete(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if _, err = s.tc.Decr(ctx, map[uint32]int32{record.PID: record.Num}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) UpdateRecord(ctx context.Context, req *v1.UpdateRecordRequest) (reply *emptypb.Empty, err error) {
	var record *biz.Record
	if record, err = s.rc.Get(ctx, req.Record.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if err = s.rc.Update(ctx, &biz.Record{
		RID:  req.Record.Rid,
		Num:  req.Record.Num,
		Desc: req.Record.Desc,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	alterNum := req.Record.Num - record.Num
	if alterNum > 0 {
		if _, err = s.tc.Incr(ctx, map[uint32]int32{record.PID: record.Num}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	} else if alterNum < 0 {
		if _, err = s.tc.Decr(ctx, map[uint32]int32{record.PID: record.Num}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	}
	return
}

func (s *PointService) ListRecord(ctx context.Context, req *v1.ListRecordRequest) (reply *v1.ListRecordReply, err error) {
	reply = &v1.ListRecordReply{Finished: true}
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
			reply.Finished = false
			break
		}
		reply.Records = append(reply.Records, &v1.Record{
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
	reply.Count = uint32(len(reply.Records))
	return
}
