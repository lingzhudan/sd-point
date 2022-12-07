// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"sd-point/internal/biz"
	"sd-point/internal/conf"
	"sd-point/internal/data"
	"sd-point/internal/server"
	"sd-point/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	pointRepo := data.NewPointRepo(dataData, logger)
	pointUsecase := biz.NewPointUsecase(pointRepo, logger)
	pointService := service.NewPointService(pointUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, pointService, logger)
	httpServer := server.NewHTTPServer(confServer, pointService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
