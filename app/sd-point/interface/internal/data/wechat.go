package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"net/http"
	"sd-point/app/sd-point/interface/internal/biz"
	"sd-point/app/sd-point/interface/internal/conf"
	"strings"
	"time"
)

type wechatRepo struct {
	// 获取微信令牌的url
	tokenURL string
	// 获取微信openid的url
	openidURL string
	// 获取微信手机号的
	phoneNumberURL string
	// 微信小程序编号
	appID string
	// 微信小程序密码
	appSecret string

	// 微信令牌
	token string
	// 微信令牌持续时间
	tokenEffectiveTime int64
	// 微信令牌非法时间
	tokenInvalidAt int64

	log *log.Helper
}

// NewWechatRepo .
func NewWechatRepo(wc *conf.Wechat, logger log.Logger) biz.WechatRepo {
	return &wechatRepo{
		appID:              wc.AppId,
		appSecret:          wc.AppSecret,
		tokenURL:           wc.TokenUrl,
		tokenEffectiveTime: wc.TokenEffectiveTime,
		openidURL:          wc.OpenidUrl,
		phoneNumberURL:     wc.PhoneNumberUrl,
		log:                log.NewHelper(logger),
	}
}

func (r *wechatRepo) RefreshToken(ctx context.Context) (err error) {
	nowUnix := time.Now().Unix()
	if len(r.token) != 0 && r.tokenInvalidAt < nowUnix {
		return
	}
	var t *biz.TokenResponse
	var resp *http.Response
	var body []byte
	if resp, err = http.Get(fmt.Sprintf(r.tokenURL, r.appID, r.appSecret)); err != nil {
		r.log.Errorf("failed to get token from wechat server, error: %v", err)
		return
	}
	defer resp.Body.Close()
	if body, err = io.ReadAll(resp.Body); err != nil {
		r.log.Errorf("failed to read http body, error: %v", err)
		return
	}
	if err = json.Unmarshal(body, t); err != nil {
		r.log.Errorf("failed to unmarshal token response, error: %v", err)
		return
	}
	r.tokenInvalidAt = nowUnix + r.tokenEffectiveTime
	r.token = t.AccessToken
	return
}

func (r *wechatRepo) GetOpenid(ctx context.Context, c *biz.WechatAccountCode) (a *biz.WechatAccount, err error) {
	a = new(biz.WechatAccount)
	var o *biz.OpenidResponse
	var resp *http.Response
	var body []byte
	if resp, err = http.Get(fmt.Sprintf(r.openidURL, r.appID, r.appSecret, c.OpenidCode)); err != nil {
		r.log.Errorf("failed to get openid from wechat server, error: %v", err)
		return
	}
	defer resp.Body.Close()
	if body, err = io.ReadAll(resp.Body); err != nil {
		r.log.Errorf("failed to read http response, error: %v", err)
		return
	}
	if err = json.Unmarshal(body, o); err != nil {
		r.log.Errorf("failed to unmarshal token response, error: %v", err)
		return
	}
	if o.Errcode == 0 {
		a.Openid = o.Openid
	}
	return
}

func (r *wechatRepo) GetPhoneNumber(ctx context.Context, c *biz.WechatAccountCode) (a *biz.WechatAccount, err error) {
	a = new(biz.WechatAccount)
	var p *biz.PhoneNumberResponse
	var resp *http.Response
	var body []byte
	if err = r.RefreshToken(ctx); err != nil {
		r.log.Errorf("failed to get token from wechat server, error: %v", err)
		return
	}
	if resp, err = http.Post(fmt.Sprintf(r.phoneNumberURL, r.token),
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("code=%s", c.PhoneNumberCode))); err != nil {
		r.log.Errorf("failed to get phone number from wechat server, error: %v", err)
		return
	}
	defer resp.Body.Close()
	if body, err = io.ReadAll(resp.Body); err != nil {
		r.log.Errorf("failed to read http response, error: %v", err)
		return
	}
	if err = json.Unmarshal(body, p); err != nil {
		r.log.Errorf("failed to unmarshal phone number response, error: %v", err)
		return
	}
	a.PhoneNumber = p.GetPhoneNumber()
	return
}
