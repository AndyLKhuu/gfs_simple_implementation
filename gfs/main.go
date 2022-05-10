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

// func initClientConnection() {
// 	var conn *grpc.ClientConn
// 	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %s", err)
// 	}
// 	defer conn.Close()

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 3



// 	c := protos.NewMasterClient(conn)
// 	response, err := c.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
// 	if err != nil {
// 		log.Fatalf("Error when calling GetSystemChunkSize: %s", err)
// 	}

// 	log.Printf("The ChunkSize of the system is : %d", response.Size)
// }

func main() {
	// Start up Master Server
	go master.InitMasterServer(masterServerPort, NUM_CHUNK_SERVERS, chunkServerPortBase)

	log.Printf("Master server initialization complete.");


	// Start up Chunkservers
	for i := 0; i < NUM_CHUNK_SERVERS; i++ {
		go chunkserver.InitChunkServer(chunkServerPortBase + i)
	}

	time.Sleep(2 * time.Second) //Arbitrary Number

	log.Printf("Chunk server initialization complete.");
	

	// for i := 1; i < 5; i++ {
	c := client.InitClient()
	// todo: close client connections 
	// }
	fmt.Println(c);

	c.Create("test");

	select {}
}
