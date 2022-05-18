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


	totalBytesToWrite := len(data)
	log.Printf("number of bytes to write: %d", totalBytesToWrite)

	chunkSize := getSystemChunkSizeReply.Size; // for testing, treat this as number of bytes 



	// for dataOffset := 0; dataOffset < len(data) {

		chunkIdx := int32(offset/chunkSize);
	// 	chunkOffset := dataOffset % chunkSize
		
		
	// 	nBytesToWrite := min 

		getChunkLocationReply, err := masterClient.GetChunkLocation(context.Background(), &protos.ChunkLocationRequest{Path: path, ChunkIdx: chunkIdx})
		if err != nil {
			log.Printf("error when calling GetChunkLocation: %s", err);
			return -1;
		}

		chunkLocations := getChunkLocationReply.ChunkServerIds
		chunkHandle := getChunkLocationReply.ChunkHandle

		// Client pushses data to all replicas 
		replicaReceiveStatus := make([]bool, len(chunkLocations))
		for i := 0; i < len(chunkLocations); i++ { // Should we async this instead of sequential?
			chunkServerAddr := chunkLocations[i]
			conn, err := grpc.Dial(chunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
			defer conn.Close()
			if err != nil {
				log.Printf("error when client connecting to chunk server: %s", err)
				replicaReceiveStatus[i] = false
				continue
			}
			chunkServerClient := cs.NewChunkServerClient(conn);
			_, err = chunkServerClient.ReceiveWriteData(context.Background(), 
				&cs.WriteDataBundle{Data: data, Size: int64(len(data)), Ch: chunkHandle}) // TODO: chunkify
			if err != nil {
				log.Printf("error when client sending write data to chunk server: %s", err)
				replicaReceiveStatus[i] = false
				continue
			}
			replicaReceiveStatus[i] = true
		}
		// Do we need to resend writeData to failed nodes? 

		// Client tells primary to commit 
		primaryChunkServerAddr := chunkLocations[0]	
		conn, err := grpc.Dial(primaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			log.Printf("error when client connecting to primary chunk server: %s", err)
			return -1
		}
		primaryChunkServerClient := cs.NewChunkServerClient(conn);
		_, err = primaryChunkServerClient.PrimaryCommitMutate(context.Background(), 
			&cs.PrimaryCommitMutateRequest{Ch: chunkHandle, SecondaryChunkServerAddresses: chunkLocations[1:]})



	// 	dataOffset += 64 // TODO
		
	// }



	return 0;
}
