package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	sample_data "github.com/psinthorn/gostore/data_generator"
	"github.com/psinthorn/gostore/pb"
	serializer "github.com/psinthorn/gostore/serializer"
	"github.com/stretchr/testify/require"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	// filename and path
	binaryFile := "./../tmp/laptop.bin"
	jsonfile := "./../tmp/laptop.json"

	// test write protobuffer message to binary file
	laptop1 := sample_data.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	// test read from binary and wrote to new protobuffer message
	laptop2 := &pb.Laptop{}
	err = serializer.ReadBinaryfileToProtoMessage(binaryFile, laptop2)
	require.NoError(t, err)

	// test compare 2 files is equal
	require.True(t, proto.Equal(laptop1, laptop2))

	// test convert and write protobuffer message to json file
	err = serializer.WrtiteProtobufToJsonFile(laptop2, jsonfile)
	require.NoError(t, err)
}

func TestReadBinaryfileToProtoMessage(t *testing.T) {
	binaryFile := "./../tmp/laptop.bin"

	laptop2 := &pb.Laptop{}
	err := serializer.ReadBinaryfileToProtoMessage(binaryFile, laptop2)
	require.NoError(t, err)
}
