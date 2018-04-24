package service

import (
	"todo_proto/pb/proto"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(context context.Context, message *proto.PingMessage) (*proto.PingMessage, error) {
	return &proto.PingMessage{Greeting: "Ciao Ciao" + message.Greeting}, nil
}
