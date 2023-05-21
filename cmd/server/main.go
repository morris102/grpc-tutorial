package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/pb"
	"grpc-go/service"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
		//return
	}

	log.Println("server started")
}
