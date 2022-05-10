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

var chunkServerTempDirectoryPath = "../temp_dfs_storage/"

func InitChunkServer(csAddr int) {
	fmt.Println("starting up chunkserver " + strconv.Itoa(csAddr) + ".")
	if err := os.MkdirAll(chunkServerTempDirectoryPath+strconv.Itoa(csAddr), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(csAddr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.ChunkServer{}

	grpcServer := grpc.NewServer()

	protos.RegisterChunkServerServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
