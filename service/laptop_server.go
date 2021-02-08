package service

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/psinthorn/gostore/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// laptopServer is aprovide  laptop server service
type LaptopServer struct {
	Store LaptopStore
}

// NewLaptopServer will return new laptop server
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

// CreateLaptop is Unary RPC to create new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreatelaptopResponse, error) {

	// รับค่าจาก req
	newLaptop := req.GetLaptop()
	log.Printf("Receive a create-laptop request with ID: %s", newLaptop.Id)

	// ตรวจสอบคค่าที่รับมาว่ามี id มาด้วยหรือไม่

	if len(newLaptop.Id) > 0 {
		// หากมีให้ตรวจสอบว่าค่าที่ส่งมาถูกต้องหรือไม่ หากไม่ถูกต้องาม format ให้ return error
		_, err := uuid.Parse(newLaptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Error Laptop ID is Invalid UUID %v ", err)
		}
	} else {
		// หากไม่มีให้ทำการสร้างค่า id ขึ้นมาใหม่ หากสร้างไม่ได้ให้ return error
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "can't generate new Id by uuid: %v", err)
		}
		// หากสร้างได้ให้กำหนด newLaptop.id = id ที่สร้างขึ้นมาใหม่
		newLaptop.Id = id.String()
	}

	// Save new Laptop to in memory
	err := server.Store.Save(newLaptop)
	if err != nil {
		return nil, fmt.Errorf("Error occure on save new laptop to store %w ", err)
	}

	// Save data success return laptop id to caller
	res := &pb.CreatelaptopResponse{
		Id: newLaptop.Id,
	}

	// response กลับให้ client
	return res, nil

	// Next is find the solution to store data to database
	// // save newLaptop to database
	// // เก็บข้อมูลลงดาต้าเบส
	// err := store.data.Save(newLaptop)
	// if err != nil {
	// 	return
	// }

}
