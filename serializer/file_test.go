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

	// test write protobuffer message to binary
	binaryFile := "./../tmp/laptop.bin"
	laptop1 := sample_data.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	// test read from binary to protobuffer message
	laptop2 := &pb.Laptop{}
	err = serializer.ReadBinaryfileToProtoMessage(binaryFile, laptop2)
	require.NoError(t, err)

	// test compare 2 files is equal
	require.True(t, proto.Equal(laptop1, laptop2))

}

func TestReadBinaryfileToProtoMessage(t *testing.T) {
	binaryFile := "./../tmp/laptop.bin"

	laptop2 := &pb.Laptop{}
	err := serializer.ReadBinaryfileToProtoMessage(binaryFile, laptop2)
	require.NoError(t, err)
}
