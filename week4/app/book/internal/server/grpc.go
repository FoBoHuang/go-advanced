package server

import (
	"context"
	"net"
	v1 "week4/api/book/v1"
	"week4/app/book/internal/conf"
	"week4/app/book/internal/service"
	"week4/pkg/appmanage"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func (g *GrpcServer) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		g.server.Stop()
	}()
	return g.server.Serve(g.listener)
}

func NewGrpcServer(service *service.BookService, config *conf.GrpcConf) appmanage.GrpcServer {
	server := new(GrpcServer)
	lis, err := net.Listen("tcp", config.Addr())
	server.listener = lis
	if err != nil {
		panic(err.Error())
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterBookServiceServer(grpcServer, service)
	server.server = grpcServer
	return server
}
