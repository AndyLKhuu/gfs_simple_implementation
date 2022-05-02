package chunkserver

//go:generate bash -c "protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative chunkserver.proto"
