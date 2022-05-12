package client

import (
	"context"
	// "fmt"
	"gfs/master/protos"
	"log"
	// "time"
	// "reflect"

	"google.golang.org/grpc"
)

type Client struct {
	MasterConn *grpc.ClientConn // used to later close connection
	MasterClient *protos.MasterClient // used to invoke RPCs
}

// Initializes a new Client. Pass in master's identifier to link Client to master
func InitClient(mAddr string) *Client {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(mAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	c := protos.NewMasterClient(conn)

	client := new(Client);
	client.MasterConn = conn;
	client.MasterClient = &c;

	return client;
}


/* 	methods to implement:
	- create
	- delete
	- open*
	- close*
	- read
	- write
	- snapshot*
	- record append*
 */


func (client *Client) Create(filepath string) int {
	masterClient := *(client.MasterClient);
	_, err := masterClient.ReceiveClientCreateRequest(context.Background(), &protos.ClientCreateRequest{Filepath: filepath})
	if err != nil {
		log.Println("Error creating file.");
		return -1;
	} 
	log.Printf("Succesfully created file: %s", filepath);
	return 0;
}

func (client *Client) Delete(filepath string) int {
	masterClient := *(client.MasterClient);
	_, err := masterClient.ReceiveClientDeleteRequest(context.Background(), &protos.ClientDeleteRequest{Filepath: filepath})
	if err != nil {
		log.Println("Error deleting file.");
		return -1;
	} 
	log.Printf("Succesfully deleted file: %s", filepath);
	return 0;
}

func (client *Client) Open(filepath string) int {
	return -1;
}

func (client *Client) Read(filepath string, offset int64, data []byte) int {
	masterClient := *(client.MasterClient);
	getSystemChunkSizeReply, err := masterClient.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetSystemChunkSize: %s", err)
		return -1; 
	}

	chunkSize := getSystemChunkSizeReply.Size;
 	chunkIdx := int32(offset/chunkSize);

	getFileLocationReply, err := masterClient.GetFileLocation(context.Background(), &protos.ChunkLocationRequest{Filepath: filepath, ChunkIdx: chunkIdx})
	if err != nil {
		log.Fatalf("Error when calling GetFileLocation: %s", err);
		return -1;
	}

	log.Println(getFileLocationReply);

	chunkLocation := getFileLocationReply.ChunkServerIds
	chunkHandle := getFileLocationReply.ChunkHandler
	log.Printf("Obtained (chunkLocation, chunkHandle): (%d, %d)", chunkLocation, chunkHandle)
	log.Printf("TODO: build and invoke chunkserverReadRPC(chunkHandle, byteRange) => chunkData")
	return 0;
}

func (client *Client) Write(filepath string, offset int64, data []byte) int {
	masterClient := *(client.MasterClient);
	getSystemChunkSizeReply, err := masterClient.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetSystemChunkSize: %s", err)
		return -1
	}

	chunkSize := getSystemChunkSizeReply.Size;
 	chunkIdx := int32(offset/chunkSize);

	getFileLocationReply, err := masterClient.GetFileLocation(context.Background(), &protos.ChunkLocationRequest{Filepath: filepath, ChunkIdx: chunkIdx})
	if err != nil {
		log.Fatalf("Error when calling GetFileLocation: %s", err);
		return -1;
	}
	log.Println(getFileLocationReply);

	chunkLocation := getFileLocationReply.ChunkServerIds
	chunkHandle := getFileLocationReply.ChunkHandler
	log.Printf("Obtained (chunkLocation, chunkHandle): (%d, %d)", chunkLocation, chunkHandle)
	log.Printf("TODO: build and invoke chunkserverWriteRPC(chunkHandle, byteRange) => chunkData")
	return 0;
}
