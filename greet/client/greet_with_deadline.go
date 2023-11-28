package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vairarchi/go-grpc-learning/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline() called")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Vaibhav",
	}

	resp, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected GRPC error: %v\n", err)
			}
		} else {
			log.Printf("a non-gRPC error: %v", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", resp.Result)
}
