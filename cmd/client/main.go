package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-go/pb"
	"log"
)

func main() {
	address := flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// create a stub
	conn, err := grpc.Dial(*address, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	// create the client
	client := pb.NewLaptopServiceClient(conn)

	log.Println("after client")

	// call service
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	laptop := pb.Laptop{
		Id:  id.String(),
		Cpu: "i9",
		Ram: "16gb",
		Ssd: "256gb",
		Gpu: "Nvidia GTX 1080",
	}
	res, err := client.CreateLaptop(context.Background(), &pb.CreateLaptopRequest{Laptop: &laptop})
	if err != nil {

		log.Fatal(err)
		fmt.Println("failed to create laptop")
	}

	log.Printf("created laptop with id :%v", res)
}
