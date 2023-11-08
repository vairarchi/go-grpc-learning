package main

import (
	"context"
	"log"

	pb "github.com/vairarchi/go-grpc-learning/calculator/proto"
)

func doSum(s pb.CalculatorServiceClient) {
	log.Printf("doSum() is invoked")

	res, err := s.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 3,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %v", res.Result)
}
