package main

import (
	"fmt"
	"log"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/scriptonist/todo/service/pkg/api"
	"github.com/scriptonist/todo/service/pkg/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := "8000"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	var opts []grpc.ServerOption
	logger := logrus.NewEntry(logrus.New())
	opts = append(
		opts,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_logrus.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logger),
		)),
	)
	grpcServer := grpc.NewServer(opts...)
	api.RegisterTodoAPIServer(grpcServer, server.New())
	logger.Infof("Server Listening on: %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal(err)
	}
}
