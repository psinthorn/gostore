package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WrtiteProtobufToJsonFile(message proto.Message, filename string) error {

	data, err := ConvertProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("can't convert protobuffer message to json file %w ", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("can't json file %w ", err)
	}

	return nil
}

// WriteProtobufToBinaryFile wil wrtie protobuffer message to binary file
// return error (if succesess will return nil error)
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	// write protobuffer to binary file
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("can't marshal proto message to binary file %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("can't write binary to file %w", err)
	}
	return nil
}

// ReadBinaryfileToProtoMessage will read data from binary file and Umarshal to proto message
// return error (if succesess will return nil error)
func ReadBinaryfileToProtoMessage(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can't read binary file %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("can't unmarshal data to proto message %w", err)
	}
	return nil
}
