package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"sd-point/api/sd-point/interface/v1"
	"sd-point/app/sd-point/interface/internal/biz"
	"sd-point/app/sd-point/interface/internal/conf"
	"sd-point/app/sd-point/interface/internal/server/middleware"
	"sd-point/app/sd-point/interface/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, s *service.SdPointInterfaceService, uc *biz.UserUseCase, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			// 请求详细信息日志打印提前于其他操作
			logging.Server(logger),
			// 参数合法校验
			validate.Validator(),
			selector.Server(
				middleware.CheckAuthorization(uc),
			).Match(NewWhiteListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	v1.RegisterSdPointInterfaceHTTPServer(srv, s)

	return srv
}

// NewWhiteListMatcher 不需要认证的白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := map[string]struct{}{
		"/api.sd_point.interface.v1.SdPointInterface/GetPublicKey": {}, // 公钥
		"/api.sd_point.interface.v1.SdPointInterface/Login":        {}, // 用户登录
		//"/v1/user/register/*": {}, // 用户注册
		"/api.sd_point.interface.v1.SdPointInterface/Register": {}, // 用户登出
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
