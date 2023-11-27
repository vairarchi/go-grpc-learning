package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vairarchi/go-grpc-learning/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes() was invoked")

	req := &pb.PrimeRequest{
		OneNumber: 123455840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error whilecalling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error receiving the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)

	}

}
