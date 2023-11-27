package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vairarchi/go-grpc-learning/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet() invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Vaibhav"},
		{FirstName: "Rajat"},
		{FirstName: "Dhawal"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
