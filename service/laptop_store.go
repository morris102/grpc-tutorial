package service

import (
	"errors"
	"grpc-go/pb"
	"sync"
)

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	data   map[string]*pb.Laptop
	mutext sync.RWMutex
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (s *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	s.mutext.RLock()
	defer s.mutext.RUnlock()

	if s.data[laptop.Id] != nil {
		return errors.New("laptop is existed")
	}

	s.data[laptop.Id] = laptop

	return nil
}
