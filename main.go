package main

import (
	"log"
	"net"

	"github.com/PierreKieffer/grpc-sandbox/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := server.Server{}

	grpcServer := grpc.NewServer()

	server.RegisterEventServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
