package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	clientv3 "go.etcd.io/etcd/client/v3"

	pointv1 "sd-point/api/point/service/v1"
	userv1 "sd-point/api/user/service/v1"
	"sd-point/app/sd-point/interface/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistrar,
	NewDiscovery,
	NewUserServiceClient,
	NewPointServiceClient,
	NewWechatRepo,
	NewUserRepo,
	NewPointRepo,
)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
	pc  pointv1.PointClient
}

// NewData .
func NewData(
	uc userv1.UserClient,
	pc pointv1.PointClient,
	logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(logger)
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		uc:  uc,
		pc:  pc,
		log: l,
	}, cleanup, nil
}

// NewDiscovery 服务注册发现的提供方法
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Etcd.Address},
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

// NewRegistrar 服务注册发现的提供方法
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Etcd.Address},
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

// NewUserServiceClient 用户模块服务grpc客户端的提供方法
func NewUserServiceClient(ac *conf.Auth, r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// 指定需要发现的服务名
		grpc.WithEndpoint("discovery:///sd-point.user.service"),
		// 服务发现者
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			// 捕获恐慌
			recovery.Recovery(),
			// 身份认证 客户端 jwt
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(ac.JwtKey), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

// NewPointServiceClient 点数模块服务grpc客户端的提供方法
func NewPointServiceClient(ac *conf.Auth, r registry.Discovery) pointv1.PointClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// 指定需要发现的服务名
		grpc.WithEndpoint("discovery:///sd-point.point.service"),
		// 服务发现者
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			// 捕获恐慌
			recovery.Recovery(),
			// 身份认证 客户端 jwt
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(ac.JwtKey), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := pointv1.NewPointClient(conn)
	return c
}
