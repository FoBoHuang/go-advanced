package main

import (
	"context"
	"os"
	"week4/app/customer/internal/conf"
	"week4/pkg/appmanage"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf := conf.GenConf()
	app := initApp(dbConf, httpConf, grpcConf)
	app.Register(&appmanage.RegisterInfo{
		Appid:   "customer:v1",
		AppName: "customer manager service",
	})
	app.Run(ctx, os.Interrupt)
	defer cancel()
}
