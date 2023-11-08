package main

import (
	"context"
	"log"

	pb "github.com/vairarchi/go-grpc-learning/calculator/proto"
)

func (s Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum() invoked by: %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
