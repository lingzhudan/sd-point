package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/app/sd-point/interface/internal/biz"
	"sd-point/app/sd-point/interface/internal/define"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "sd-point/api/sd-point/interface/v1"
)

type SdPointInterfaceService struct {
	pb.UnimplementedSdPointInterfaceServer

	uc *biz.UserUseCase
	pc *biz.PointUseCase
	wc *biz.WechatUseCase

	log *log.Helper
}

func NewSdPointInterfaceService(
	uc *biz.UserUseCase,
	pc *biz.PointUseCase,
	wc *biz.WechatUseCase,
	logger log.Logger) *SdPointInterfaceService {
	return &SdPointInterfaceService{
		uc:  uc,
		pc:  pc,
		wc:  wc,
		log: log.NewHelper(logger),
	}
}

func (s *SdPointInterfaceService) CreatePoint(ctx context.Context, req *pb.CreatePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	uid := GetSession(ctx).UID
	if err = s.pc.CreatePoint(ctx, uid, req.Name, req.Desc); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) UpdatePoint(ctx context.Context, req *pb.UpdatePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.pc.UpdatePoint(ctx, req.Pid, req.Name, req.Desc); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) DeletePoint(ctx context.Context, req *pb.DeletePointRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.pc.DeletePoint(ctx, req.Pid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) GetPoint(ctx context.Context, req *pb.GetPointRequest) (rep *pb.GetPointReply, err error) {
	rep = new(pb.GetPointReply)
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
	rep = new(pb.ListPointReply)
	uid := GetSession(ctx).UID
	ps, err := s.pc.ListPoint(ctx, req.Begin, req.Count, uid)
	if err != nil {
		s.log.Errorf("internal error: %v", err)
		return
	}
	for _, p := range ps.Points {
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
	rep.Finished = ps.Finished
	rep.Count = uint32(len(ps.Points))
	return
}
func (s *SdPointInterfaceService) CreateRecord(ctx context.Context, req *pb.CreateRecordRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.pc.CreatRecord(ctx, req.Pid, req.Num, req.Desc, req.ClickedAt); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) DeleteRecord(ctx context.Context, req *pb.DeleteRecordRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.pc.DeleteRecord(ctx, req.Rid); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) UpdateRecord(ctx context.Context, req *pb.UpdateRecordRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	if err = s.pc.UpdateRecord(ctx, req.Rid, req.Num, req.Desc, req.ClickedAt); err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	return
}
func (s *SdPointInterfaceService) ListRecord(ctx context.Context, req *pb.ListRecordRequest) (rep *pb.ListRecordReply, err error) {
	rep = new(pb.ListRecordReply)
	rs, err := s.pc.ListRecord(ctx, req.Begin, req.Count, req.Pid)
	if err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	for _, r := range rs.Records {
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
		})
	}
	rep.Finished = rs.Finished
	rep.Count = uint32(len(rep.Records))
	return
}
func (s *SdPointInterfaceService) GetPublicKey(ctx context.Context, _ *emptypb.Empty) (rep *pb.GetPublicKeyReply, err error) {
	rep = new(pb.GetPublicKeyReply)
	bs, err := s.uc.GetPublicKey(ctx)
	if err != nil {
		s.log.Errorf("internal error: %v", err)
	}
	rep.PublicKey = string(bs)
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
		c := req.WechatAccountCode
		var a *biz.WechatAccount
		if a, err = s.wc.GetOpenid(ctx, &biz.WechatAccountCode{
			OpenidCode: c.OpenidCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
			return
		}
		if rep.SessionId, err = s.uc.WechatLogin(ctx, a); err != nil {
			if errors.Is(err, define.ErrAccountNotFound) {
				_, _ = s.uc.WechatRegister(ctx, a)
				if rep.SessionId, err = s.uc.WechatLogin(ctx, a); err != nil {
					s.log.Errorf("internal error: %v", err)
				}
				return
			}
			s.log.Errorf("internal error: %v", err)
		}
	case 2:
		// 微信手机号登录
		c := req.WechatAccountCode
		var a *biz.WechatAccount
		if a, err = s.wc.GetPhoneNumber(ctx, &biz.WechatAccountCode{
			PhoneNumberCode: c.PhoneNumberCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
			return
		}
		if rep.SessionId, err = s.uc.WechatPhoneNumberLogin(ctx, a); err != nil {
			if errors.Is(err, define.ErrAccountNotFound) {
				_, _ = s.uc.WechatPhoneNumberRegister(ctx, a)
				if rep.SessionId, err = s.uc.WechatPhoneNumberLogin(ctx, a); err != nil {
					s.log.Errorf("internal error: %v", err)
				}
				return
			}
			s.log.Errorf("internal error: %v", err)
		}
	case 4:
		// TODO 手机号登录暂不实现
	}
	return
}
func (s *SdPointInterfaceService) Logout(ctx context.Context, _ *emptypb.Empty) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	sId := GetSessionID(ctx)
	if err = s.uc.Logout(ctx, sId); err != nil {
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
		var a *biz.WechatAccount
		c := req.WechatAccount
		if a, err = s.wc.GetOpenid(ctx, &biz.WechatAccountCode{
			OpenidCode: c.OpenidCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
			return
		}
		if rep.Uid, err = s.uc.WechatRegister(ctx, &biz.WechatAccount{
			Openid: a.Openid,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 2:
		// 微信手机号注册
		var a *biz.WechatAccount
		c := req.WechatAccount
		if a, err = s.wc.GetPhoneNumber(ctx, &biz.WechatAccountCode{
			PhoneNumberCode: c.PhoneNumberCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
			return
		}
		if rep.Uid, err = s.uc.WechatPhoneNumberRegister(ctx, &biz.WechatAccount{
			PhoneNumber: a.PhoneNumber,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 4:
		// TODO 手机号注册
	}
	return
}
func (s *SdPointInterfaceService) BindAccount(ctx context.Context, req *pb.BindAccountRequest) (rep *emptypb.Empty, err error) {
	rep = new(emptypb.Empty)
	var uid uint32
	switch req.BindType {
	case 0:
		var a *biz.WechatAccount
		c := req.WechatAccountCode
		if a, err = s.wc.GetOpenid(ctx, &biz.WechatAccountCode{
			OpenidCode: c.OpenidCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
			return
		}
		// 微信账号绑定
		if err = s.uc.WechatBind(ctx, uid, a); err != nil {
			s.log.Errorf("internal error: %v", err)
		}
	case 1:
		var a *biz.WechatAccount
		c := req.WechatAccountCode
		if a, err = s.wc.GetPhoneNumber(ctx, &biz.WechatAccountCode{
			PhoneNumberCode: c.PhoneNumberCode,
		}); err != nil {
			s.log.Errorf("internal error: %v", err)
			return
		}
		// 微信手机号绑定
		if err = s.uc.WechatPhoneNumberBind(ctx, &biz.WechatAccount{PhoneNumber: a.PhoneNumber}); err != nil {
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
		return
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
