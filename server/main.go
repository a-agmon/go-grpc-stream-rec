package main

import (
	"log"
	"net"

	pb "github.com/a-agmon/rcserver/proto"
	"google.golang.org/grpc"
)

var serverAddress string = "localhost:9090"

func main() {
	log.Println("Starting server...")
	// create a TCP net lstener
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening on " + serverAddress)

	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// associate the GRPC with our generated handler
	pb.RegisterRcServiceServer(grpcServer, NewRCServer()) //&RCServer{})
	// set the GRPC server to listen on the listerner
	log.Println("Starting GRPC server")
	grpcServer.Serve(listener)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
