//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"week4/app/customer/internal/biz"
	"week4/app/customer/internal/conf"
	"week4/app/customer/internal/data"
	"week4/app/customer/internal/server"
	"week4/app/customer/internal/service"
	"week4/pkg/appmanage"

	"github.com/google/wire"
)

func initApp(
	db *conf.ConfDB,
	http *conf.HttpConf,
	grpc *conf.GrpcConf,
) *appmanage.AppManage {
	panic(wire.Build(
		server.ProvideSet,
		data.ProvideSet,
		service.ProvideSet,
		biz.ProvideSet,
		appmanage.NewAppManage,
	))
}
