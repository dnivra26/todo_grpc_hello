package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"todo_proto/pb/proto"
	service2 "todo_grpc_hello/service"

)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 7778))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterPingServer(grpcServer, &service2.Server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
