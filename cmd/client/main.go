package main

import (
	"context"
	"flag"
	"fmt"
	"log"

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

	laptop := sample_data.NewLaptop()
	fmt.Println(laptop)
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Fatal("Laptop is already exist")
		} else {
			log.Fatal("can't create laptop")
		}
		return
	}

	log.Printf("create laptop with id %s ", res.Id)

}
