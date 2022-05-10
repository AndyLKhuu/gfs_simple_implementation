package client


import (
	"context"
	"fmt"
	"gfs/master/protos"
	"log"
	// "time"
	// "reflect"

	"google.golang.org/grpc"
)

type Client struct {
	masterConn *grpc.ClientConn // used to later close connection
	masterClient *protos.MasterClient // used to invoke RPCs
}


// Initializes a new Client. Pass in master's identifier to link Client to master
func InitClient() *Client {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	// defer conn.Close() // TODO: remove

	c := protos.NewMasterClient(conn)


	client := new(Client);
	client.masterConn = conn;
	client.masterClient = &c;

	return client;
}


/* 	methods to implement:
	- create
	- delete
	- open
	- close
	- read
	- write
	- snapshot*
	- record append*
 */


func (client *Client) Create(filename string) int {
	// invoke  master's 'Create' rpc

	/* 
	1) Invoke master's create-RPC 
		- master will handle updating metadata and allocating chunk
	*/


	fmt.Println("TODO: invoke Master's RPC to create new file with name 'filename'.")
	return -1;
}


func (client *Client) Delete(filename string) int {
	// invoke  master's 'Delete' rpc

	/* 
	Notes:
	1) Invoke master's create-RPC 
		- master will handle updating metadata and allocating chunk
	*/


	// masterClient := *(client.masterClient);
	fmt.Println("TODO: invoke Master's RPC to delete filew with name 'filename'.")
	return -1;
}


func (client *Client) Open(filename string) int {
	// What does this do exactly?? 
	return -1;
}


func (client *Client) Read(filename string, offset int, data []byte) (bytes_read int, status int) {
	// invoke  master's 'Read' rpc

	/* 
	Notes:

	1) Retrieve chunk size with Master RPC
	1) Calculate chunkIndex
	2) masterReadRPC(filename, chunkIndex) => (chunkHandle, chunkLocation)
		- how to handle file being distributed across multiple chunk servers? 
	3) chunkServerRPC(chunkHandle, byteRange) => data 
	[done.]
	*/

	masterClient := *(client.masterClient);
	response, err := masterClient.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetSystemChunkSize: %s", err)
	}

	fmt.Println(response, err);


	return 0, -1;

}


func (client *Client) Write(filename string, offset int, data []byte) (bytes_written int, status int) {
	// invoke  master's 'Write' rpc

	/* 
	Notes:	
	1) Calculate chunkIndex
	2) masterWriteRPC(filename, chunkIndex) => (chunkHandle, chunkLocation, data)
	3) chunkServerRPC(chunkHandle, byteRange) => status? 
	[done.]
	*/

	// masterClient := *(client.masterClient);

	return 0, -1;

}
