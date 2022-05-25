package chunkserver

import (
	"fmt"
	"gfs/chunkserver/protos"
	"gfs/chunkserver/services"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"log"
	"net"
)

var shared_file_path = "../temp_dfs_storage/"

// TO:DO Add mechanisms to gracefully handle termination of a chunkserver if it fails to start
func InitChunkServer(csAddr int) {
	addr := strconv.Itoa(csAddr)
	chunkserverDir := shared_file_path + addr
	if err := os.MkdirAll(chunkserverDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	fmt.Println("starting up chunkserver " + addr + ".")

	lis, err := net.Listen("tcp", ":"+addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.NewChunkServer(addr, chunkserverDir)
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
