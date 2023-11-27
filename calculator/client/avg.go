package main

import (
	"context"
	"log"

	pb "github.com/vairarchi/go-grpc-learning/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int64{3, 5, 8, 54, 23}

	for _, num := range numbers {
		log.Printf("Sending Number: %d\n", num)

		stream.Send(&pb.AverageRequest{
			Number: num,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
