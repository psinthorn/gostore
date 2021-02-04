#-----------------------------------------------------------
# Mac OS, Unix, Linux
#-----------------------------------------------------------
run: 
	go run main.go

# Generate Server 
# consume we run command in the *.proto file directory 
gen-server:
	protoc  --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative --proto_path=./protos ./protos/*.proto

# Generate Client 
gen-client:
	protoc --go_out=pb --go_opt=paths=source_relative     --go-grpc_out=pb --go-grpc_opt=paths=source_relative    client.proto 

# for macOS
clean: 
	rm pb/*.go

# to all test
test:
	go test -cover -race ./..

#individual file test	
testfile:
	go test -cover -race file_test.go 



# -------------------------------------------------------------------------
# Windows section
# -------------------------------------------------------------------------

# Generate protobuffer 
gen-win: 
	protoc --proto_path=./protos ./protos/*.proto --go_out=./pb

# Delete command line 
del:
	del pb\*.go



