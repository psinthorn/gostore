package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/psinthorn/gostore/pb"
	"github.com/psinthorn/gostore/service"
	"google.golang.org/grpc"
)

func main() {

	// สร้าง server และสิ่งที่ต้องการในการ สร้างและ run server
	// 1. สร้าง port สำหรับ server
	// 2. สร้าง server instance object
	// 3. สร้าง gRPC server ขึ้นมาใหม่โดยเพิ่ม server object เข้าไปด้วย
	// 4. register gRPC server

	// 1. สร้าง port สำหรับ server โดยการรับค่ามาจาก commad line
	port := flag.Int("port", 0, "Server port")
	flag.Parse()
	log.Printf("gRPC server start on port %d ", *port)

	serverAddr := fmt.Sprintf("0.0.0.0:%d", *port)

	// 2. สร้าง server instance object
	laptopServer := service.NewLaptopServer(service.NewInmemoryLaptopStore())
	gRPCServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(gRPCServer, laptopServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal("Can't start server  ", err)
	}
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal("Can't start server  ", err)
	}

}

// func gRPCStart() string {
// 	startServer := "gRPC Server start"
// 	fmt.Println("gRPC Server start")
// 	return startServer
// }
