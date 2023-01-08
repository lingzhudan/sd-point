package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "sd-point/api/point/v1"
	"sd-point/internal/biz"
	"time"
)

type PointService struct {
	pb.UnimplementedPointServer

	uc  *biz.PointUsecase
	log *log.Helper
}

func NewPointService(uc *biz.PointUsecase, logger log.Logger) *PointService {
	return &PointService{uc: uc, log: log.NewHelper(logger)}
}

// point service

func (s *PointService) CreatePoint(ctx context.Context, req *pb.CreatePointRequest) (reply *emptypb.Empty, err error) {
	point := &biz.Point{
		PID:  req.Point.Pid,
		UID:  0,
		Name: req.Point.Name,
		Desc: req.Point.Desc,
	}

	if err = s.uc.CreatePoint(ctx, point); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}

	return &emptypb.Empty{}, nil
}

func (s *PointService) UpdatePoint(ctx context.Context, req *pb.UpdatePointRequest) (reply *emptypb.Empty, err error) {
	point := &biz.Point{
		PID:  req.Point.Pid,
		UID:  0,
		Name: req.Point.Name,
		Desc: req.Point.Desc,
	}
	if err = s.uc.UpdatePoint(ctx, point); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (reply *emptypb.Empty, err error) {
	if err = s.uc.DeletePoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if err = s.uc.DelPointTotal(ctx, []int32{req.Pid}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (reply *pb.GetPointReply, err error) {
	reply = &pb.GetPointReply{}
	var point *biz.Point
	if point, err = s.uc.GetPoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	var ptMap map[int32]int32
	if ptMap, err = s.uc.MGetPointTotal(ctx, []int32{req.Pid}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, err
	}
	s.log.Debugf("pt: %v", ptMap[req.Pid])
	reply.Point = &pb.PointInfo{
		Pid:       point.PID,
		Name:      point.Name,
		Total:     ptMap[point.PID],
		Desc:      point.Desc,
		CreatedAt: point.CreatedAt.Unix(),
		UpdatedAt: point.UpdatedAt.Unix(),
		DeletedAt: point.DeletedAt.Time.Unix(),
	}
	return
}

func (s *PointService) ListPoint(ctx context.Context, req *pb.ListPointRequest) (reply *pb.ListPointReply, err error) {
	reply = &pb.ListPointReply{Finished: true}
	var points []*biz.Point
	if points, err = s.uc.ListPint(ctx, &biz.PointCond{
		Begin: int(req.Begin),
		Count: int(req.Count + 1),
		PIDs:  req.Pids,
		UIDs:  []int32{0},
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	var pids []int32
	for _, p := range points {
		pids = append(pids, p.PID)
	}
	var ptMap map[int32]int32
	ptMap, err = s.uc.MGetPointTotal(ctx, pids)
	if err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	for i, p := range points {
		if i >= int(req.Count) {
			reply.Finished = false
			break
		}
		reply.Points = append(reply.Points, &pb.PointInfo{
			Pid:       p.PID,
			Name:      p.Name,
			Total:     ptMap[p.PID],
			Desc:      p.Desc,
			CreatedAt: p.CreatedAt.Unix(),
			UpdatedAt: p.UpdatedAt.Unix(),
			DeletedAt: p.DeletedAt.Time.Unix(),
		})
	}
	reply.Count = int32(len(reply.Points))
	return
}

// record service

func (s *PointService) CreateRecords(ctx context.Context, req *pb.CreateRecordsRequest) (reply *emptypb.Empty, err error) {
	var records []*biz.PointRecord
	var ptMap = make(map[int32]int16)
	for _, record := range req.Records {
		newRecord := &biz.PointRecord{
			PID:       record.Pid,
			ClickedAt: time.Unix(record.ClickedAt, 0),
			Num:       int16(record.Num),
			Desc:      record.Desc,
		}
		ptMap[record.Pid] = int16(record.Num)
		records = append(records, newRecord)
	}
	if err = s.uc.CreateRecords(ctx, records); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if _, _ = s.uc.IncrPointTotal(ctx, ptMap); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeleteRecord(ctx context.Context, req *pb.DeleteRecordRequest) (reply *emptypb.Empty, err error) {
	var record *biz.PointRecord
	if record, err = s.uc.GetRecord(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if err = s.uc.DeleteRecord(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if _, err = s.uc.DecrPointTotal(ctx, map[int32]int16{record.PID: record.Num}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) UpdateRecord(ctx context.Context, req *pb.UpdateRecordRequest) (reply *emptypb.Empty, err error) {
	var record *biz.PointRecord
	if record, err = s.uc.GetRecord(ctx, req.Record.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	if err = s.uc.UpdateRecord(ctx, &biz.PointRecord{
		RID:  req.Record.Rid,
		Num:  int16(req.Record.Num),
		Desc: req.Record.Desc,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	alterNum := int16(req.Record.Num) - record.Num
	if alterNum > 0 {
		if _, err = s.uc.IncrPointTotal(ctx, map[int32]int16{record.PID: record.Num}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	} else if alterNum < 0 {
		if _, err = s.uc.DecrPointTotal(ctx, map[int32]int16{record.PID: record.Num}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	}
	return
}

func (s *PointService) ListRecord(ctx context.Context, req *pb.ListRecordRequest) (reply *pb.ListRecordReply, err error) {
	reply = &pb.ListRecordReply{Finished: true}
	var records []*biz.PointRecord
	if records, err = s.uc.ListRecord(ctx, &biz.RecordCond{
		Begin:        int(req.Begin),
		Count:        int(req.Count + 1),
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
		reply.Records = append(reply.Records, &pb.RecordInfo{
			Rid:       r.RID,
			Pid:       r.PID,
			Num:       int32(r.Num),
			ClickedAt: r.ClickedAt.Unix(),
			Desc:      r.Desc,
			CreatedAt: r.CreatedAt.Unix(),
			UpdatedAt: r.UpdatedAt.Unix(),
			DeletedAt: r.DeletedAt.Time.Unix(),
		})
	}
	reply.Count = int32(len(reply.Records))
	return
}
