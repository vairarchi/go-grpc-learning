package main

import (
	"context"
	"log"

	pb "github.com/vairarchi/go-grpc-learning/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt() called")

	resp, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: ", e.Message())
			log.Printf("Error code from server: ", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number")
				return
			}
		} else {
			log.Printf("a non-gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", resp.Result)
}
