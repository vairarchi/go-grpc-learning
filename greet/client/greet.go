package main

import (
	"context"
	"log"

	pb "github.com/vairarchi/go-grpc-learning/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet() was called")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Vaibhav",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
