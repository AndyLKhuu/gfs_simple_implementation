package chunkserver

import (
	"gfs/chunkserver/protos"
	"gfs/chunkserver/services"
	"strconv"

	"google.golang.org/grpc"

	"log"
	"net"
)

// TO:DO Add mechanisms to gracefully handle termination of a chunkserver if it fails to start
func InitChunkServer(csAddr int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(csAddr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.NewChunkServer(strconv.Itoa(csAddr))
	grpcServer := grpc.NewServer()
	protos.RegisterChunkServerServer(grpcServer, &s)

	// Start Chunkserver
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// TO:DO Constantly check for the end of any leases the chunkserver has. How to prevent TOCTOU bugs here?
}

func InitBenchmarkingChunkServer(csAddr int, config services.BenchmarkConfig) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(csAddr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.NewBenchmarkingChunkServer(strconv.Itoa(csAddr), config)
	grpcServer := grpc.NewServer()
	protos.RegisterChunkServerServer(grpcServer, &s)

	// Start Chunkserver
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// TO:DO Constantly check for the end of any leases the chunkserver has. How to prevent TOCTOU bugs here?
}
