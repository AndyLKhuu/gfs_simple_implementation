package main

import (
	"context"
	"gfs/chunkserver"
	"gfs/master"
	"gfs/master/protos"
	"log"

	"google.golang.org/grpc"
)

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 5

func initClientConnection() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := protos.NewMasterClient(conn)

	response, err := c.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Printf("error when calling GetSystemChunkSize: %s", err)
	}

	log.Printf("the ChunkSize of the system is : %d", response.Size)
}

func main() {
	// Start up Master Server
	master.InitMasterServer(masterServerPort, NUM_CHUNK_SERVERS, chunkServerPortBase)

	// Start up Chunkservers
	for i := 0; i < NUM_CHUNK_SERVERS; i++ {
		go chunkserver.InitChunkServer(chunkServerPortBase + i)
	}

	// time.Sleep(2 * time.Second) //Arbitrary Number

	// initClientConnection()

	select {}
}
