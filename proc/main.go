package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"proc/pb"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	srv := &server{
		cfg:     cfg,
		storage: storage,
	}

	log.Println("proc configuration")

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		log.Fatal("failed to listen", err)
	}

	rpcSrv := grpc.NewServer(
		grpc.UnaryInterceptor(srv.logInterceptor),
	)

	pb.RegisterImageServiceServer(rpcSrv, srv)
}

func (srv *server) logInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[Storage Service Interceptor]", info.FullMethod)

	m, err := handler(ctx, req)

	log.Println("post proc message", m)

	return m, err
}
