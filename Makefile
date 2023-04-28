protoc_file=./datamodel.proto
proto_go_file=./
target_file=guanyu


protoc_clean:
	rm ./server/grpc_datamodel.go
	rm ./server/grpc_server.go

protoc:
	protoc --go_out=${proto_go_file} --go-grpc_out=${proto_go_file} ${protoc_file}
	mv ./server/datamodel.pb.go ./server/grpc_datamodel.go
	mv ./server/datamodel_grpc.pb.go ./server/grpc_server.go 

linux:
	GOARCH=amd64 GOOS=linux go build -o ${target_file} ./