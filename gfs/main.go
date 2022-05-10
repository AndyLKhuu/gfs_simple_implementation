package main

import (
	// "context"
	"fmt"
	"gfs/master"
	// "gfs/master/protos"
	"gfs/client"
	// "log"
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

// 	c := protos.NewMasterClient(conn)

// 	response, err := c.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
// 	if err != nil {
// 		log.Fatalf("Error when calling GetSystemChunkSize: %s", err)
// 	}

// 	log.Printf("The ChunkSize of the system is : %d", response.Size)
// }

func main() {
	/* Start Master Node*/
	fmt.Println("Starting up Master Server")

	go master.InitMasterServer()

	// TO:DO Add more complex synchronization mechanism to check for completion of Master Server Initialization
	time.Sleep(2 * time.Second) //Arbitrary Number

	// for i := 1; i < 5; i++ {
	c := client.InitClient()
	// }
	fmt.Println(c);

	c.Create("test");



}
