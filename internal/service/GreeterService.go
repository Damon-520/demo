package service

import (
	"context"

	pb "demoapi/internal/conf"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}
