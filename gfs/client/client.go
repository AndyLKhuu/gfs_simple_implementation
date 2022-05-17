package client

import (
	"context"
	"gfs/master/protos" // alias this import to 'm' to match 'cs' 
	"log"
	cs "gfs/chunkserver/protos"
	"time"
	"math/rand"
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
		log.Printf("did not connect: %s", err)
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
		log.Println("error when calling CreateFile %s.", err);
		return -1;
	} 
	log.Printf("Succesfully created file: %s", path);
	return 0;
}

func (client *Client) Remove(path string) int {
	masterClient := *(client.MasterClient);
	_, err := masterClient.RemoveFile(context.Background(), &protos.FileRemoveRequest{Path: path})
	if err != nil {
		log.Println("error when calling RemoveFile %s", err);
		return -1;
	} 
	log.Printf("Succesfully removed file: %s", path);
	return 0;
}

func (client *Client) Read(path string, offset int64, data []byte) int {
	masterClient := *(client.MasterClient);
	getSystemChunkSizeReply, err := masterClient.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Printf("error when calling GetSystemChunkSize: %s", err)
		return -1; 
	}

	chunkSize := getSystemChunkSizeReply.Size;
 	chunkIdx := int32(offset/chunkSize);

	getChunkLocationReply, err := masterClient.GetChunkLocation(context.Background(), &protos.ChunkLocationRequest{Path: path, ChunkIdx: chunkIdx})
	if err != nil {
		log.Printf("error when calling GetChunkLocation: %s", err);
		return -1;
	}

	chunkLocations := getChunkLocationReply.ChunkServerIds
	chunkHandle := getChunkLocationReply.ChunkHandle

	chunkServerAddr := chunkLocations[rand.Intn(len(chunkLocations))]; // Current readRequest load balancing is Random
	conn, err := grpc.Dial(chunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to chunk server
	if err != nil {
		log.Printf("error when client connecting to chunk server: %s", err)
		return -1
	}

	chunkServerClient := cs.NewChunkServerClient(conn);
	readReply, err := chunkServerClient.Read(context.Background(), &cs.ReadRequest{Ch: chunkHandle, L: 0, R: 0})

	copy(data, []byte(readReply.Data) )
	return 0;
}

func (client *Client) Write(path string, offset int64, data []byte) int {
	masterClient := *(client.MasterClient);
	getSystemChunkSizeReply, err := masterClient.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Printf("error when calling GetSystemChunkSize: %s", err)
		return -1
	}

	chunkSize := getSystemChunkSizeReply.Size;

	// iterate from chunkStart to chunkStart + len(data) and update chunkIdx
 	chunkIdx := int32(offset/chunkSize);

	getChunkLocationReply, err := masterClient.GetChunkLocation(context.Background(), &protos.ChunkLocationRequest{Path: path, ChunkIdx: chunkIdx})
	if err != nil {
		log.Printf("error when calling GetChunkLocation: %s", err);
		return -1;
	}

	// create new chunks if needed

	chunkLocations := getChunkLocationReply.ChunkServerIds
	chunkHandle := getChunkLocationReply.ChunkHandle
	log.Printf("Obtained (chunkLocations, chunkHandle): (%d, %d)", chunkLocations, chunkHandle)

	primaryChunkServerAddr := chunkLocations[0];
	conn, err := grpc.Dial(primaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to chunk server
	if err != nil {
		log.Printf("error when client connecting to chunk server: %s", err)
		return -1
	}

	primaryChunkServerClient := cs.NewChunkServerClient(conn);
	log.Printf("Client connected to chunk server: %s", primaryChunkServerClient);
	log.Printf("TODO: build and invoke chunkserverWriteRPC(chunkHandle, byteRange) => chunkData") 

 	//TO:DO We also have to remove the file meta data from the in-memory structures on the master
	
	 // Here, we can pass the secondaryChunkServerAddr over the RPC so primaryCS can relay the writeReq to secondaries. 
	// The RPC can check for nil secondaryChunkServerAddr. If nil, don't relay bc we are in the case of secondary. If !- nil, relay bc it is primary.
	// That way, we can just use 1 single chunkServerWrite RPC handler.

	/** 
		
	primaryChunkServerClient.ReceiveWriteData(context.background(), &protos.WriteDataBundle{Data: data, Size: len(data), Ch: chunkHandle});

	
	
	
	*/


	return 0;
}
