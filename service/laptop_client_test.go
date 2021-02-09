package service

import (
	"context"
	"net"
	"testing"

	sample_data "github.com/psinthorn/gostore/data_generator"
	"github.com/psinthorn/gostore/pb"
	"github.com/psinthorn/gostore/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddr := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddr)

	laptop := sample_data.NewLaptop()
	expectedId := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedId, res.Id)

	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check laptop is the same object?
	requireSameLaptop(t, laptop, other)

}

func startTestLaptopServer(t *testing.T) (*LaptopServer, string) {

	laptopServer := NewLaptopServer(NewInmemoryLaptopStore())

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	server := grpc.NewServer()
	pb.RegisterLaptopServiceServer(server, laptopServer)

	go server.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, addr string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)

	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1, laptop2 *pb.Laptop) {
	laptop1Json, err := serializer.ConvertProtobufToJson(laptop1)
	require.NoError(t, err)

	laptop2Json, err := serializer.ConvertProtobufToJson(laptop2)
	require.NoError(t, err)

	require.Equal(t, laptop1Json, laptop2Json)
}
