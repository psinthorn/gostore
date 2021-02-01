gen: 
	protoc --proto_path=./protos ./protos/*.proto --go_out=plugins=grpc:pb

# clean: 
# 	rm pb/*.go
clean:
	del pb\*.go

run: 
	go run main.go