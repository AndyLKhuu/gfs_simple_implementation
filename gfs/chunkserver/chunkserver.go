package chunkserver

import (
	"fmt"
	"gfs/chunkserver/protos"
	"gfs/chunkserver/services"
	"strconv"

	"google.golang.org/grpc"

	"log"
	"net"
	"os"
)

// TO:DO Have these variables be global instead of replicating them
var chunkServerTempDirectoryPath = "../temp_dfs_storage/"

func InitChunkServer(csAddr int) {
	// TO:DO Refactor so that we aren't calling strconv so many times
	chunkserverRootDir := chunkServerTempDirectoryPath + strconv.Itoa(csAddr)
	fmt.Println("starting up chunkserver " + strconv.Itoa(csAddr) + ".")
	if err := os.MkdirAll(chunkserverRootDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(csAddr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.ChunkServer{
		ChunkHandleToFile: make(map[uint64]string),
		Rootpath:          chunkserverRootDir,
		Address:           strconv.Itoa(csAddr),
		WriteCache:        make(map[string]*protos.WriteDataBundle)}

	grpcServer := grpc.NewServer()

	protos.RegisterChunkServerServer(grpcServer, &s)

	// Start Chunkserver
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()
}
