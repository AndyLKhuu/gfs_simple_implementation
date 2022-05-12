package main

import (
	"context"
	"gfs/chunkserver"
	"gfs/master"
	"gfs/master/protos"
	"log"
	"time"

	"google.golang.org/grpc"
)

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 3
var shared_file_path = "../temp_dfs_storage/shared"

// Used as a rudimentary sanity check for new features
func initClientConnection() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := protos.NewMasterClient(conn)

	res, err := c.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Printf("error when calling GetSystemChunkSize: %s", err)
	}

	log.Printf("the ChunkSize of the system is : %d", res.Size)

	file := shared_file_path + "/a.txt"
	createReq := &protos.FileCreateRequest{Path: file, RepFactor: 1}

	response, err := c.CreateFile(context.Background(), createReq)
	if err != nil {
		log.Printf("error when calling CreateFile: %s", err)
	}

	log.Printf("response from calling createFile: %s", response.Message)
}

func main() {
	// Start up Master Server
	go master.InitMasterServer(masterServerPort, NUM_CHUNK_SERVERS, chunkServerPortBase)

	// Start up Chunkservers
	for i := 0; i < NUM_CHUNK_SERVERS; i++ {
		go chunkserver.InitChunkServer(chunkServerPortBase + i)
	}

	time.Sleep(2 * time.Second)

	go initClientConnection()

	select {}
}
