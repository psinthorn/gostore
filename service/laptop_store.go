package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/psinthorn/gostore/pb"
)

// ErrorAlraedyExists use for global error in case of data is already exists
var ErrorAlraedyExists = errors.New("Data Alearedy Exists")

// LaptopStore is an interface to store a laptop
type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore to store laptop to inmemory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInmemoryLaptopStore will return inmemory object
func NewInmemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save will save data to inmemory
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrorAlraedyExists
	}

	// Save to store in this save function we will copy to inmemory
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("can't copy object %w ", err)
	}
	store.data[other.Id] = other
	return nil
}

// Save will save data to inmemory
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, fmt.Errorf("Can't find laptop with ID %v", id)
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("Can't copy laptop to in memory %v", err)
	}

	return other, nil

}

// -------------------------------------------------
// Store data to Database Section
// -------------------------------------------------
