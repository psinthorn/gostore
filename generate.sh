!#/bin/bash

# Generate command
protoc --proto_path=./protos ./protos/*.proto --go_out=plugins=grpc:pb