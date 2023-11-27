package main

import (
	"log"

	pb "github.com/vairarchi/go-grpc-learning/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("primes() was invoked with %v\n", in)

	number := in.OneNumber
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
