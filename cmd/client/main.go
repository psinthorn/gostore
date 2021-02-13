package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	sample_data "github.com/psinthorn/gostore/data_generator"
	"github.com/psinthorn/gostore/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	serverAddr := flag.String("port", "", "Server address")
	flag.Parse()
	log.Printf("Server address %s", *serverAddr)
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("can't connect to server", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)
	fmt.Println("gRPC Client Start...")

	// Prepare server connection context port
	// Register server

	// receive req from client and process business logic
	// return response to req
	// serve server

	// Create new laptop
	laptop := sample_data.NewLaptop()

	// // to test new laptop with no id
	// laptop.Id = ""

	// // to test new laptop with invalid id
	// laptop.Id = "invalid-uuid"

	// // to test new laptop with existing id
	// // you can copy and past to id value
	// laptop.Id = "put-existing-id"

	fmt.Println("------------------------------------------------------")
	fmt.Println(laptop)
	fmt.Println("------------------------------------------------------")

	fmt.Println("------------------------------------------------------")
	log.Printf("new laptop object with id %s ", laptop.Id)
	fmt.Println("------------------------------------------------------")

	// Create laptop request object
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// use context with timeout for make a connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// return data from server
	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Fatal("Laptop is already exist")
		} else {
			log.Fatal("can't create laptop: ", err)
		}
		return
	}
	fmt.Println("------------------------------------------------------")
	log.Printf("laptop created and store in-memory with id %s ", res.Id)
	fmt.Println("------------------------------------------------------")

}
