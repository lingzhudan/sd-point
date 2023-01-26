package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/sd-point/interface/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "sd-point/api/sd-point/interface/v1"
)

type SdPointInterfaceService struct {
	pb.UnimplementedSdPointInterfaceServer

	uc *biz.UserUseCase
	pc *biz.PointUseCase

	log *log.Helper
}

func NewSdPointInterfaceService(uc *biz.UserUseCase, pc *biz.PointUseCase, logger log.Logger) *SdPointInterfaceService {
	return &SdPointInterfaceService{
		uc:  uc,
		pc:  pc,
		log: log.NewHelper(logger),
	}
}

func (s *SdPointInterfaceService) CreatePoint(ctx context.Context, req *pb.CreatePointRequest) (_ *emptypb.Empty, err error) {
	// TODO 用户模块待完善
	if err = s.pc.CreatePoint(ctx, &biz.Point{
		//UID:  uid,
		Name: req.Point.Name,
		Desc: req.Point.Desc,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) UpdatePoint(ctx context.Context, req *pb.UpdatePointRequest) (_ *emptypb.Empty, err error) {
	p := req.Point
	if err = s.pc.UpdatePoint(ctx, &biz.Point{
		PID: p.Pid,
		//UID:       uid,
		Name: p.Name,
		Desc: p.Desc,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (_ *emptypb.Empty, err error) {
	if err = s.pc.DeletePoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (rep *pb.GetPointReply, err error) {
	var p *biz.Point
	if p, err = s.pc.GetPoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	rep.Point = &pb.GetPointReply_Point{
		Pid:       p.PID,
		Total:     p.Total,
		Name:      p.Name,
		Desc:      p.Desc,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
	return
}
func (s *SdPointInterfaceService) ListPoint(ctx context.Context, req *pb.ListPointRequest) (rep *pb.ListPointReply, err error) {
	var ps []*biz.Point
	if ps, err = s.pc.ListPoint(ctx, &biz.PointCond{
		Begin: req.Begin,
		Count: req.Count,
		PIDs:  req.Pids,
		//UIDs:  []uint32{uid},
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	for _, p := range ps {
		rep.Points = append(rep.Points, &pb.GetPointReply_Point{
			Pid:       p.PID,
			Total:     p.Total,
			Name:      p.Name,
			Desc:      p.Desc,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			DeletedAt: p.DeletedAt,
		})
	}
	return
}
func (s *SdPointInterfaceService) CreateRecord(ctx context.Context, req *pb.CreateRecordRequest) (_ *emptypb.Empty, err error) {
	r := req.Record
	if err = s.pc.CreatRecord(ctx, &biz.Record{
		PID:       r.Pid,
		Num:       r.Num,
		Desc:      r.Desc,
		ClickedAt: r.ClickedAt,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) DeleteRecord(ctx context.Context, req *pb.DeleteRecordRequest) (_ *emptypb.Empty, err error) {
	if err = s.pc.DeleteRecord(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) UpdateRecord(ctx context.Context, req *pb.UpdateRecordRequest) (_ *emptypb.Empty, err error) {
	r := req.Record
	if err = s.pc.UpdateRecord(ctx, &biz.Record{
		RID:       r.Rid,
		Num:       r.Num,
		Desc:      r.Desc,
		ClickedAt: r.ClickedAt,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) ListRecord(ctx context.Context, req *pb.ListRecordRequest) (rep *pb.ListRecordReply, err error) {
	rep = new(pb.ListRecordReply)
	var rs []*biz.Record
	if rs, err = s.pc.ListRecord(ctx, &biz.RecordCond{
		Begin:        req.Begin,
		Count:        req.Count + 1,
		RIDs:         req.Rids,
		PIDs:         req.Pids,
		MinClickedAt: req.MinClickedAt,
		MaxClickedAt: req.MaxClickedAt,
	}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	for _, r := range rs {
		if uint32(len(rep.Records)) >= req.Count {
			break
		}
		rep.Records = append(rep.Records, &pb.Record{
			Rid:       r.RID,
			Pid:       r.PID,
			Num:       r.Num,
			ClickedAt: r.ClickedAt,
			Desc:      r.Desc,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
			DeletedAt: r.DeletedAt,
		})
	}
	rep.Finished = uint32(len(rs)) <= req.Count
	rep.Count = uint32(len(rep.Records))
	return
}
func (s *SdPointInterfaceService) Login(ctx context.Context, req *pb.LoginRequest) (rep *pb.LoginReply, err error) {
	rep = new(pb.LoginReply)
	switch req.LoginType {
	case 0:
		// 原生账号登录
		a := req.OriginAccount
		if rep.SessionId, err = s.uc.Login(ctx, &biz.OriginAccount{
			Account:  a.Account,
			Password: a.Password,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 1:
		// 微信账号登录
		a := req.WechatAccount
		if rep.SessionId, err = s.uc.WechatLogin(ctx, &biz.WechatAccount{
			OpenIDCode: a.OpenIdCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 2:
		// 微信手机号登录
		a := req.WechatAccount
		if rep.SessionId, err = s.uc.WechatPhoneNumberLogin(ctx, &biz.WechatAccount{
			PhoneNumberCode: a.PhoneNumberCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 4:
		// TODO 手机号登录暂不实现
	}
	return
}
func (s *SdPointInterfaceService) Logout(ctx context.Context, _ *emptypb.Empty) (_ *emptypb.Empty, err error) {
	// TODO 此处应该由header中提取sessionID
	sessionId := ""
	if err = s.uc.Logout(ctx, sessionId); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) Register(ctx context.Context, req *pb.RegisterRequest) (rep *pb.RegisterReply, err error) {
	rep = new(pb.RegisterReply)
	switch req.RegisterType {
	case 0:
		// 原生账号注册
		a := req.OriginAccount
		if rep.Uid, err = s.uc.Register(ctx, &biz.OriginAccount{
			Account:  a.Account,
			Password: a.Password,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 1:
		// 微信账号注册
		a := req.WechatAccount
		if rep.Uid, err = s.uc.WechatRegister(ctx, &biz.WechatAccount{
			OpenIDCode: a.OpenIdCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 2:
		// 微信手机号注册
		a := req.WechatAccount
		if rep.Uid, err = s.uc.WechatPhoneNumberRegister(ctx, &biz.WechatAccount{
			PhoneNumberCode: a.PhoneNumberCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 4:
		// TODO 手机号注册
	}
	return
}
func (s *SdPointInterfaceService) BindAccount(ctx context.Context, req *pb.BindAccountRequest) (_ *emptypb.Empty, err error) {
	switch req.BindType {
	case 0:
		// 微信账号绑定
		if err = s.uc.WechatBind(ctx, &biz.WechatAccount{OpenIDCode: req.WechatAccount.OpenIdCode}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 1:
		// 微信手机号绑定
		if err = s.uc.WechatPhoneNumberBind(ctx, &biz.WechatAccount{OpenIDCode: req.WechatAccount.OpenIdCode}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 2:
		// TODO 手机号绑定
	}
	return
}
func (s *SdPointInterfaceService) GetUser(ctx context.Context, req *pb.GetUserRequest) (rep *pb.GetUserReply, err error) {
	rep = new(pb.GetUserReply)
	var u *biz.User
	if u, err = s.uc.Get(ctx, req.Uid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	rep.User = &pb.GetUserReply_User{
		Uid:      u.UID,
		Username: u.Username,
	}
	return
}
func (s *SdPointInterfaceService) ListUser(ctx context.Context, req *pb.ListUserRequest) (rep *pb.ListUserReply, err error) {
	rep = new(pb.ListUserReply)
	var us []*biz.User
	if us, err = s.uc.List(ctx, &biz.UserCond{UIDs: req.Uids}); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	for _, u := range us {
		rep.Users = append(rep.Users, &pb.GetUserReply_User{
			Uid:      u.UID,
			Username: u.Username,
		})
	}
	return
}
