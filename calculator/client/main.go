package main

import (
	"log"

	pb "github.com/vairarchi/go-grpc-learning/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	s := pb.NewCalculatorServiceClient(conn)

	// doSum(s)
	doPrimes(s)
}
