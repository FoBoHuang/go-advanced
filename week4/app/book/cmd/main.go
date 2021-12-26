package main

import (
	"context"
	"os"
	"week4/app/book/internal/conf"
	"week4/pkg/appmanage"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf, customerConf := conf.GenConf()
	app := initApp(dbConf, httpConf, grpcConf, customerConf)
	app.Register(&appmanage.RegisterInfo{
		Appid:   "book:v1",
		AppName: "book manager service",
	})
	app.Run(ctx, os.Interrupt)
	defer cancel()
}
