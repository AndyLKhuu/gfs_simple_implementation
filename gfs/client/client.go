package client

import (
	"context"
	"encoding/csv"
	"fmt"
	cs "gfs/chunkserver/protos"
	"gfs/master/protos" // alias this import to 'm' to match 'cs'
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var BENCHMARKING = true
var BENCHMARKING_DIR = "../benchmarking/"

type Client struct {
	MasterConn      *grpc.ClientConn     // used to later close connection
	MasterClient    *protos.MasterClient // used to invoke RPCs
	ChunkSize       uint64               // The chunksize of the filesystem.
	BenchmarkConfig BenchmarkConfig
}

type BenchmarkConfig struct {
	OutputDir    string
	Architecture string // ASYNC OR SEQ
}

// Initializes a new Client. Pass in master's identifier to link Client to master
func NewClient(mAddr string) (*Client, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(mAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %s", err)
	}

	c := protos.NewMasterClient(conn)

	getSystemChunkSizeReply, err := c.GetSystemChunkSize(context.Background(), &protos.SystemChunkSizeRequest{})
	if err != nil {
		log.Printf("error when calling GetSystemChunkSize: %s", err)
		return &Client{}, err
	}

	client := new(Client)
	client.MasterConn = conn
	client.MasterClient = &c
	client.ChunkSize = getSystemChunkSizeReply.Size

	return client, nil
}

func NewBenchmarkingClient(mAddr string, config BenchmarkConfig) (*Client, error) {
	client, err := NewClient(mAddr)
	if err != nil {
		log.Printf("error when creating a benchmarking client: %s", err)
	}
	client.BenchmarkConfig = config

	return client, err
}

func (client *Client) Create(path string) int {

	startTime := time.Now().UnixMicro()
	masterClient := *(client.MasterClient)
	_, err := masterClient.CreateFile(context.Background(), &protos.FileCreateRequest{Path: path, RepFactor: 1})
	if err != nil {
		log.Printf("error when calling CreateFile %s.", err)
		return -1
	}
	log.Printf("Successfully created file: %s", path)
	endTime := time.Now().UnixMicro()
	if BENCHMARKING {
		latency := endTime - startTime
		client.recordPerformance("CREATE", latency, 0)
	}
	return 0
}

func (client *Client) Remove(path string) int {
	startTime := time.Now().UnixMicro()
	masterClient := *(client.MasterClient)
	_, err := masterClient.RemoveFile(context.Background(), &protos.FileRemoveRequest{Path: path})
	if err != nil {
		log.Printf("error when calling RemoveFile %s", err)
		return -1
	}
	log.Printf("Successfully removed file: %s", path)
	endTime := time.Now().UnixMicro()
	if BENCHMARKING {
		latency := endTime - startTime
		client.recordPerformance("REMOVE", latency, 0)
	}
	return 0
}

func (client *Client) Read(path string, offset uint64, data []byte) int {
	startTime := time.Now().UnixMicro()
	masterClient := *(client.MasterClient)
	chunkSize := client.ChunkSize
	totalBytesRead := uint64(0)
	remainingBytesToRead := uint64(len(data))
	for dataOffset := offset; dataOffset < offset+uint64(len(data)); {
		chunkIdx := uint32(dataOffset / chunkSize) // check this
		chunkOffset := uint64(dataOffset) % chunkSize
		remainingChunkSpace := chunkSize - chunkOffset

		// Calculate max number of bytes to read
		nBytesToRead := remainingBytesToRead
		if remainingChunkSpace < nBytesToRead {
			nBytesToRead = remainingChunkSpace
		}

		getChunkLocationReply, err := masterClient.GetChunkLocation(context.Background(), &protos.ChunkLocationRequest{Path: path, ChunkIdx: chunkIdx})
		if err != nil {
			log.Printf("error when calling GetChunkLocation: %s", err)
			return -1
		}

		chunkLocations := getChunkLocationReply.ChunkServerIds
		chunkHandle := getChunkLocationReply.ChunkHandle
		chunkServerAddr := chunkLocations[rand.Intn(len(chunkLocations))]                             // Current readRequest load balancing is Random
		conn, err := grpc.Dial(chunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to chunk server
		if err != nil {
			log.Printf("error when client connecting to chunk server: %s", err)
			return -1
		}
		chunkServerClient := cs.NewChunkServerClient(conn)
		readReply, err := chunkServerClient.Read(context.Background(), &cs.ReadRequest{Ch: chunkHandle, L: int32(chunkOffset), R: int32(chunkOffset + nBytesToRead)}) // not implemented yet. just reads back chicken
		conn.Close()

		copy(data[totalBytesRead:totalBytesRead+nBytesToRead], []byte(readReply.Data))

		dataOffset += nBytesToRead
		remainingBytesToRead -= nBytesToRead
		totalBytesRead += nBytesToRead
	}
	endTime := time.Now().UnixMicro()
	if BENCHMARKING {
		latency := endTime - startTime
		client.recordPerformance("READ", latency, totalBytesRead)
	}
	return int(totalBytesRead)
}

func (client *Client) Write(path string, offset uint64, data []byte) int {
	startTime := time.Now().UnixMicro()

	masterClient := *(client.MasterClient)
	chunkSize := client.ChunkSize
	totalBytesWritten := uint64(0)
	remainingBytesToWrite := uint64(len(data))
	for dataOffset := offset; dataOffset < offset+uint64(len(data)); {
		chunkIdx := uint32(dataOffset / chunkSize)
		chunkOffset := uint64(dataOffset) % chunkSize
		remainingChunkSpace := chunkSize - chunkOffset

		// Calculate max number of bytes to write
		nBytesToWrite := remainingBytesToWrite
		if remainingChunkSpace < nBytesToWrite {
			nBytesToWrite = remainingChunkSpace
		}

		getChunkLocationReply, err := masterClient.GetChunkLocation(context.Background(), &protos.ChunkLocationRequest{Path: path, ChunkIdx: chunkIdx})
		if err != nil {
			log.Printf("error when calling GetChunkLocation: %s", err)
			return -1
		}

		chunkLocations := getChunkLocationReply.ChunkServerIds
		chunkHandle := getChunkLocationReply.ChunkHandle

		// Client pushes data to all replicas
		transactionId := uuid.New().String()
		replicaReceiveStatus := make([]bool, len(chunkLocations))
		for i := 0; i < len(chunkLocations); i++ { // TODO: optimize to async
			chunkServerAddr := chunkLocations[i]
			conn, err := grpc.Dial(chunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
			if err != nil {
				log.Printf("error when client connecting to chunk server %s: %s", chunkServerAddr, err)
				replicaReceiveStatus[i] = false
				continue
			}
			chunkServerClient := cs.NewChunkServerClient(conn)

			// TO:DO Repush data on failure
			_, err = chunkServerClient.ReceiveWriteData(context.Background(),
				&cs.WriteDataBundle{TxId: transactionId, Data: data[totalBytesWritten : totalBytesWritten+nBytesToWrite], Size: nBytesToWrite, Ch: chunkHandle, Offset: chunkOffset})
			if err != nil {
				log.Printf("error when calling ReceiveWriteData: %s", err)
				replicaReceiveStatus[i] = false
				continue
			}
			replicaReceiveStatus[i] = true
			conn.Close()
		}

		// Client tells primary to commit
		primaryChunkServerAddr := getChunkLocationReply.Primary
		conn, err := grpc.Dial(primaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
		if err != nil {
			log.Printf("error when client connecting to primary chunk server: %s", err)
			return -1
		}
		primaryChunkServerClient := cs.NewChunkServerClient(conn)
		_, err = primaryChunkServerClient.PrimaryCommitMutate(context.Background(),
			&cs.PrimaryCommitMutateRequest{Ch: chunkHandle, SecondaryChunkServerAddresses: chunkLocations[1:], TxId: transactionId})

		conn.Close()

		dataOffset += nBytesToWrite
		remainingBytesToWrite -= nBytesToWrite
		totalBytesWritten += nBytesToWrite
	}

	endTime := time.Now().UnixMicro()
	if BENCHMARKING {
		latency := endTime - startTime
		client.recordPerformance("WRITE", latency, totalBytesWritten)
	}
	return int(totalBytesWritten)
}

type record struct {
	operationName string // WRITE, READ, CREATE, REMOTE
	latency       int64  // latency
	load          uint64 // number of bytes worked on
}

func (client *Client) recordPerformance(operationName string, latency int64, load uint64) {
	// pipe to files
	fmt.Println("Getting cwd")

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	fmt.Println("----\n")
	fname := client.BenchmarkConfig.OutputDir + operationName + ".csv"
	fmt.Printf("filename: %s\n", fname)

	csvFile, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer csvFile.Close()
	if err != nil {
		log.Printf("error opening benchmark file for %s", operationName)
		return
	}
	writer := csv.NewWriter(csvFile)

	// data := record{operationName: operationName, latency: latency, load: load}
	data := [][]string{
		{fmt.Sprint(latency), fmt.Sprint(load)},
	}

	err = writer.WriteAll(data)
	if err != nil {
		log.Printf("error writing to benchmark file for %s: %s", operationName, err)
		return
	}

	log.Printf("%s latency: %d", operationName, latency)

}
