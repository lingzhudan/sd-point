package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type WechatUseCase struct {
	repo WechatRepo
	log  *log.Helper
}

func NewWechatUseCase(repo WechatRepo, logger log.Logger) *WechatUseCase {
	return &WechatUseCase{repo: repo, log: log.NewHelper(logger)}
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type OpenidResponse struct {
	// 用户唯一标识
	Openid string `json:"openid"`
	// 会话密钥
	SessionKey string `json:"session_key"`
	// 用户在开放平台的唯一标识符
	Unionid string `json:"unionid"`
	// errcode: 40029	code 无效	js_code无效
	// errcode: 45011	api minute-quota reach limit  mustslower  retry next minute	API 调用太频繁，请稍候再试
	// errcode: 40226	code blocked	高风险等级用户，小程序登录拦截 。风险等级详见用户安全解方案
	// errcode: -1		system error	系统繁忙，此时请开发者稍候再试
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type PhoneNumberResponse struct {
	PhoneInfo struct {
		// 用户绑定的手机号（国外手机号会有区号）
		PhoneNumber string `json:"phoneNumber"`
		// 没有区号的手机号
		PurePhoneNumber string `json:"purePhoneNumber"`
		// 区号
		CountryCode string `json:"countryCode"`
	} `json:"phone_info"`
	// 错误码	错误码取值	解决方案
	// -1	system error	系统繁忙，此时请开发者稍候再试
	// 40029	code 无效	js_code无效
	Errcode int
	Errmsg  string
}

func (r *PhoneNumberResponse) GetPhoneNumber() (phoneNumber string) {
	phoneNumber = r.PhoneInfo.PhoneNumber
	return
}

type WechatRepo interface {
	GetOpenid(ctx context.Context, c *WechatAccountCode) (a *WechatAccount, err error)
	GetPhoneNumber(ctx context.Context, c *WechatAccountCode) (a *WechatAccount, err error)
}

func (uc *WechatUseCase) GetOpenid(ctx context.Context, c *WechatAccountCode) (a *WechatAccount, err error) {
	if a, err = uc.repo.GetOpenid(ctx, c); err != nil {
		uc.log.Errorf("failed to get openid, error: %v", err)
	}
	return
}
func (uc *WechatUseCase) GetPhoneNumber(ctx context.Context, c *WechatAccountCode) (a *WechatAccount, err error) {
	if a, err = uc.repo.GetPhoneNumber(ctx, c); err != nil {
		uc.log.Errorf("failed to get openid, error: %v", err)
	}
	return
}
