package main

import (

	"fmt"
	// "context"
	"gfs/chunkserver"
	"gfs/master"
	// "gfs/master/protos"
	"gfs/client"
	"log"
	"time"

	// "google.golang.org/grpc"
)

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 3
var NUM_CLIENTS = 1
var shared_file_path = "../temp_dfs_storage/shared/"

func main() {
	// Start up Master Server
	go master.InitMasterServer(masterServerPort, NUM_CHUNK_SERVERS, chunkServerPortBase)

	// Start up Chunkservers
	for i := 0; i < NUM_CHUNK_SERVERS; i++ {
		go chunkserver.InitChunkServer(chunkServerPortBase + i)
	}

	time.Sleep(2 * time.Second) //Arbitrary Number

	fmt.Println("-------");

	// Start up Clients
	for i := 0; i < NUM_CLIENTS; i++ {
		go func () {
			c := client.InitClient(masterServerPort) // Idea: go func() this so we can run clients in parallel. In those funcs, we can run different workloads 
			defer c.MasterConn.Close();
			log.Println("Initialized a client");
			fname := shared_file_path + "test.txt";
			c.Create(fname);
			c.Read(fname, 0, nil);
			var str = "hello";
			c.Write(fname, 0, []byte(str));
			// c.Remove(fname); // careful using this when implementing chunk server repl.
		}()
	}

	select {}
}
