package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	p "sd-point/api/point/service/v1"
	"sd-point/app/point/service/internal/conf"
	"sd-point/app/point/service/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, ac *conf.Auth, point *service.PointService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			// 请求详细信息日志打印提前于其他操作
			logging.Server(logger),
			// 参数合法校验
			validate.Validator(),
			// 身份认证 服务器端 jwt
			jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(ac.JwtKey), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	p.RegisterPointServer(srv, point)

	return srv
}
