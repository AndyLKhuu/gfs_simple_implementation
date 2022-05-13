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

func (client *Client) Create(path string) int {
	masterClient := *(client.MasterClient);
	_, err := masterClient.CreateFile(context.Background(), &protos.FileCreateRequest{Path: path, RepFactor: 1})
	if err != nil {
		log.Println("Error creating file.");
		return -1;
	} 
	log.Printf("Succesfully created file: %s", path);
	return 0;
}

func (client *Client) Remove(path string) int {
	masterClient := *(client.MasterClient);
	_, err := masterClient.RemoveFile(context.Background(), &protos.FileRemoveRequest{Path: path})
	if err != nil {
		log.Println("Error deleting file.");
		return -1;
	} 
	log.Printf("Succesfully deleted file: %s", path);
	return 0;
}

func (client *Client) Read(path string, offset int64, data []byte) int {
	masterClient := *(client.MasterClient);
	getSystemChunkSizeReply, err := masterClient.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetSystemChunkSize: %s", err)
		return -1; 
	}

	chunkSize := getSystemChunkSizeReply.Size;
 	chunkIdx := int32(offset/chunkSize);

	 getChunkLocationReply, err := masterClient.GetChunkLocation(context.Background(), &protos.ChunkLocationRequest{Path: path, ChunkIdx: chunkIdx})
	if err != nil {
		log.Fatalf("Error when calling GetChunkLocation: %s", err);
		return -1;
	}

	log.Println(getChunkLocationReply);

	chunkLocation := getChunkLocationReply.ChunkServerIds
	chunkHandle := getChunkLocationReply.ChunkHandle
	log.Printf("Obtained (chunkLocation, chunkHandle): (%s, %d)", chunkLocation, chunkHandle)
	log.Printf("TODO: build and invoke chunkserverReadRPC(chunkHandle, byteRange) => chunkData")
	return 0;
}

func (client *Client) Write(path string, offset int64, data []byte) int {
	masterClient := *(client.MasterClient);
	getSystemChunkSizeReply, err := masterClient.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetSystemChunkSize: %s", err)
		return -1
	}

	chunkSize := getSystemChunkSizeReply.Size;
 	chunkIdx := int32(offset/chunkSize);

	getChunkLocationReply, err := masterClient.GetChunkLocation(context.Background(), &protos.ChunkLocationRequest{Path: path, ChunkIdx: chunkIdx})
	if err != nil {
		log.Fatalf("Error when calling GetChunkLocation: %s", err);
		return -1;
	}
	log.Println(getChunkLocationReply);

	chunkLocation := getChunkLocationReply.ChunkServerIds
	chunkHandle := getChunkLocationReply.ChunkHandle
	log.Printf("Obtained (chunkLocation, chunkHandle): (%d, %d)", chunkLocation, chunkHandle)
	log.Printf("TODO: build and invoke chunkserverWriteRPC(chunkHandle, byteRange) => chunkData")
	return 0;
}
