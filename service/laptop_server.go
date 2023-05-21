package service

import (
	"context"
	"github.com/google/uuid"
	"grpc-go/pb"
	"log"
)

type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	store *InMemoryLaptopStore
}

func NewLaptopServer(store *InMemoryLaptopStore) *LaptopServer {
	return &LaptopServer{
		store: store,
	}
}

func (s *LaptopServer) CreateLaptop(
	ctx context.Context, request *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("receive a create laptop request with id: %s", laptop.Id)
	log.Printf("receive a info laptop CPU %v , Ram  %s , SSD %v ,GPU %v", laptop.Cpu, laptop.Ram, laptop.Ssd, laptop.Gpu)

	if len(laptop.Id) > 0 {
		// check valid id
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			log.Println("laptop id request invalid ")
			return nil, err
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		laptop.Id = id.String()
	}

	// save id in mem storage
	err := s.store.Save(laptop)
	if err != nil {
		log.Panicf("cannot save laptop with id %s", laptop.Id)
		return nil, err
	}

	return &pb.CreateLaptopResponse{Id: laptop.Id}, nil
}
