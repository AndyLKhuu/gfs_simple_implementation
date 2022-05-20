package main

import (
	"fmt"
	"gfs/chunkserver"
	"gfs/client"
	"gfs/master"
	"gfs/test"
	"log"
	"os"
	"time"
)

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 3
var NUM_CLIENTS = 1
var shared_file_path = "../temp_dfs_storage/shared/"

func main() {
	if err := os.MkdirAll("../temp_dfs_storage", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Start up Master Server
	go master.InitMasterServer(masterServerPort, NUM_CHUNK_SERVERS, chunkServerPortBase)

	// Start up Chunkservers
	for i := 0; i < NUM_CHUNK_SERVERS; i++ {
		go chunkserver.InitChunkServer(chunkServerPortBase + i)
	}

	time.Sleep(2 * time.Second) //Arbitrary Number

	fmt.Println("-------")

	// Start up Clients
	for i := 0; i < NUM_CLIENTS; i++ {
		go func() {
			c := client.InitClient(masterServerPort)
			defer c.MasterConn.Close()
			log.Println("Initialized a client")

			test.Run(test.WriteRemoveSmallFileTest, c)
			test.Run(test.WriteRemoveLargeFileTest, c)

		}()
	}

	log.Println("Done.")
	select {}
}
