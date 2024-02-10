package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	address := "50051"
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", address))
	if err != nil {
		log.Fatalf("failed to listen on port %s : %v", address, err)
	}
	server := grpc.NewServer()

	log.Printf("server running on port %s", address)
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve on port %s : %v", address, err)
	}
}