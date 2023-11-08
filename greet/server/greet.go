package main

import (
	"context"
	"log"

	pb "github.com/vairarchi/go-grpc-learning/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet() invoked by %v", in)
	return &pb.GreetResponse{
		Result: "hello " + in.FirstName,
	}, nil
}
