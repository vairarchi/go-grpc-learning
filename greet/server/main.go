package main

import (
	"log"
	"net"

	pb "github.com/vairarchi/go-grpc-learning/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v", err)
	}

	log.Printf("Listening on %v", addr)

	opts := []grpc.ServerOption{}
	tls := true // change that to false if needed

	if tls {
		certFile := "ssl/server.crt"
		key := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, key)
		if err != nil {
			log.Fatalf("Failed loading certificates: %f\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
