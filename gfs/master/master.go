package master

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	cs "gfs/chunkserver/protos"
	"gfs/master/protos"
	"gfs/master/services"

	"google.golang.org/grpc"
)

// TO:DO Better name + Propagate as global variable
var shared_file_path = "../temp_dfs_storage/shared"

type ChunkServerConfig struct {
	cs_addr   string
	cs_client cs.ChunkServerClient
}

func InitMasterServer(mAddr string, numChunkServers int, chunkServerPortBase int) {
	fmt.Println("starting up master server.")

	err := os.MkdirAll(shared_file_path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", mAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.NewMasterServer()
	grpcServer := grpc.NewServer()
	protos.RegisterMasterServer(grpcServer, s)
	chunkserver_chan := make(chan ChunkServerConfig)

	// Serve Master Routine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			// TO:DO Propagate this error out to fail server
			log.Fatalf("failed to serve: %s.", err)
		}
	}()

	// Start connections to chunkservers
	go func() {
		for i := 0; i < int(numChunkServers); i++ {
			cs_addr := ":" + strconv.Itoa(chunkServerPortBase+i)
			log.Printf("connecting to chunkserver %s.", cs_addr)
			var conn *grpc.ClientConn

			conn, err := grpc.Dial(cs_addr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
			if err != nil {
				log.Printf("did not connect to chunkserver %s.", cs_addr)
				continue
			}
			log.Printf("successfully connected to chunkserver %s.", cs_addr)

			c := cs.NewChunkServerClient(conn)

			chunkserver_chan <- ChunkServerConfig{cs_addr: cs_addr, cs_client: c}
		}
	}()

	// TO:DO Fix the subtle bug where system will crash if the number of
	// successful chunkservers created != numChunkServers variable, then
	// system will crash

	// Store connections to chunkservers
	// TO:DO Restructure connections to chunkservers so that they are connected
	// asynchronously and not in increasing order. (Connection to CS 2 shouldn't depend on Connection to CS 1)
	for i := 0; i < int(numChunkServers); i++ {
		config := <-chunkserver_chan
		s.ChunkServerClients[config.cs_addr] = config.cs_client
		log.Printf("storing chunkserver client at address %s", config.cs_addr)
	}
	close(chunkserver_chan)
}
