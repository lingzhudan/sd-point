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

func (s *PointService) CreatePoints(ctx context.Context, req *pb.CreatePointRequest) (reply *emptypb.Empty, err error) {
	point := &biz.Point{
		PID:   req.Point.Pid,
		UID:   0,
		Name:  req.Point.Name,
		Total: req.Point.Total,
		Desc:  req.Point.Desc,
	}

	if err = s.uc.CreatePoint(ctx, point); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, pb.ErrorPointNotFound("", "")
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
		return nil, pb.ErrorUserNotFound("", "")
	}
	return &emptypb.Empty{}, nil
}

func (s *PointService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (reply *emptypb.Empty, err error) {
	if err = s.uc.DeletePoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, pb.ErrorUserNotFound("", "")
	}
	return &emptypb.Empty{}, nil
}

func (s *PointService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (reply *pb.GetPointReply, err error) {
	var point *biz.Point
	if point, err = s.uc.GetPoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, pb.ErrorPointNotFound("", "")
	}
	newPoint := &pb.PointInfo{
		Pid:       point.PID,
		Name:      point.Name,
		Total:     point.Total,
		Desc:      point.Desc,
		CreatedAt: point.CreatedAt.Unix(),
		UpdatedAt: point.UpdatedAt.Unix(),
		DeletedAt: point.DeletedAt.Time.Unix(),
	}

	return &pb.GetPointReply{Point: newPoint}, nil
}

func (s *PointService) ListPoint(ctx context.Context, req *pb.ListPointRequest) (reply *pb.ListPointReply, err error) {
	var infos []*pb.PointInfo
	var points []*biz.Point
	if points, err = s.uc.ListPint(ctx, &biz.PointCond{
		Begin: int(req.Begin),
		Count: int(req.Count),
		PIDs:  req.Pids,
		UIDs:  []int32{0},
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, pb.ErrorPointNotFound("", "")
	}
	for _, point := range points {
		newPoint := &pb.PointInfo{
			Pid:       point.PID,
			Name:      point.Name,
			Total:     point.Total,
			Desc:      point.Desc,
			CreatedAt: point.CreatedAt.Unix(),
			UpdatedAt: point.UpdatedAt.Unix(),
			DeletedAt: point.DeletedAt.Time.Unix(),
		}
		infos = append(infos, newPoint)
	}
	return &pb.ListPointReply{Points: infos}, nil
}

// record service

func (s *PointService) CreateRecords(ctx context.Context, req *pb.CreateRecordsRequest) (reply *emptypb.Empty, err error) {
	var records []*biz.PointRecord
	for _, record := range req.Records {
		newRecord := &biz.PointRecord{
			RID:       0,
			PID:       record.Pid,
			ClickedAt: time.Unix(record.ClickedAt, 0),
			Num:       int16(record.Num),
			Desc:      record.Desc,
		}
		records = append(records, newRecord)
	}
	if err = s.uc.CreateRecords(ctx, records); err != nil {
		s.log.Errorf("internal error: %v", err)
		return nil, pb.ErrorUserNotFound("", "")
	}
	return
}

func (s *PointService) UpdateRecord(ctx context.Context, req *pb.UpdateRecordRequest) (reply *emptypb.Empty, err error) {
	newRecord := &biz.PointRecord{
		RID:       req.Record.Rid,
		PID:       0,
		ClickedAt: time.Unix(req.Record.ClickedAt, 0),
		Num:       int16(req.Record.Num),
		Desc:      req.Record.Desc,
	}
	if err = s.uc.UpdateRecord(ctx, newRecord); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) DeleteRecord(ctx context.Context, req *pb.DeleteRecordRequest) (reply *emptypb.Empty, err error) {
	if err = s.uc.DeleteRecord(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}

func (s *PointService) ListRecord(ctx context.Context, req *pb.ListRecordRequest) (reply *pb.ListRecordReply, err error) {
	reply = &pb.ListRecordReply{}
	var records []*biz.PointRecord
	if records, err = s.uc.ListRecord(ctx, &biz.RecordCond{
		Begin:        int(req.Begin),
		Count:        int(req.Count),
		RIDs:         req.Rids,
		PIDs:         req.Pids,
		MinClickedAt: req.MinClickedAt,
		MaxClickedAt: req.MaxClickedAt,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	for _, r := range records {
		newRecord := &pb.RecordInfo{
			Rid:       r.RID,
			Pid:       r.PID,
			Num:       int32(r.Num),
			ClickedAt: r.ClickedAt.Unix(),
			Desc:      r.Desc,
			CreatedAt: r.CreatedAt.Unix(),
			UpdatedAt: r.UpdatedAt.Unix(),
			DeletedAt: r.DeletedAt.Time.Unix(),
		}
		reply.Records = append(reply.Records, newRecord)
	}
	return
}
