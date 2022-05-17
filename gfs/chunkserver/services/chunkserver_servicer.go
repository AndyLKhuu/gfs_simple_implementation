package services

import (
	"context"
	"gfs/chunkserver/protos"
	cs "gfs/chunkserver/protos"

	"log"
	"os"
	"sync"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"

)

// TO:DO Find better naming than ch for chunkhandle messages

var chunkServerTempDirectoryPath = "../temp_dfs_storage/"

type ChunkServer struct {
	protos.UnimplementedChunkServerServer
	ChunkHandleToFile map[uint64]string // Chunkhandle to filepath of chunk
	Rootpath          string            // Root directory path for chunkserver
	Address           string            // Address of chunkserver
}

// TO:DO There should be an init function for the ChunkServer Struct

func (s *ChunkServer) Read(ctx context.Context, readReq *protos.ReadRequest) (*protos.ReadReply, error) {
	return &protos.ReadReply{Data: "chicken"}, nil
}

func (s *ChunkServer) ReceiveWriteData(ctx context.Context, writeBundle *protos.WriteDataBundle) (*protos.Ack, error) {

	data := writeBundle.Data
	// size := writeBundle.Size
	chunkHandle := writeBundle.Ch
	chunkServers := writeBundle.ChunkServers

	log.Println(chunkServers)
	log.Printf("THIS CS ADDR: %S", chunkServers[0])

	// var chunkServerChannels [5]chan int
	chunkServerChannels := make(map[int]chan int)
	for i := 1; i < len(chunkServers); i++ {
		chunkServerChannels[i] = make(chan int)

	}

	if (len(chunkServers) > 1) {
		log.Println("we are in priamry")
		// log.Println("TODO: propogate data to seconday chunk servers")
		var wg sync.WaitGroup



		log.Printf("len of chunkServers: %d", len(chunkServers))

		for i := 1; i < len(chunkServers); i++ {
			wg.Add(1)
			log.Println("adding 1 to workign rgoup")
			go func (id int, done chan int) { 
				defer wg.Done()
				log.Println("---")
				log.Println(id)

				log.Println("---")

				secondaryChunkServerAddr := chunkServers[id]
				conn, err := grpc.Dial(secondaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to chunk server
				log.Printf("%s here A", secondaryChunkServerAddr)

				if err != nil {
					log.Printf("error when primaryCS connecting to secondaryCS: %s", err)
					done <- -1
					return
				}
				log.Printf("%s here B", secondaryChunkServerAddr)
				secondaryChunkServerClient := cs.NewChunkServerClient(conn);
				_, err = secondaryChunkServerClient.ReceiveWriteData(context.Background(), 
					&cs.WriteDataBundle{Data: data, Size: int64(len(data)), Ch: chunkHandle, ChunkServers: []string{secondaryChunkServerAddr}})
				if err != nil {
					log.Printf("error when primaryCS sending write data to secondaryCS: %s", err)
					done <- -1
					return
				}
				log.Printf("%s here C", secondaryChunkServerAddr)

				// done <- 0 // buggy

				// here we do secondary commit
				log.Println("here we do 2nd commit")
			}(i, chunkServerChannels[i])

		}
		log.Println("%s WAITING", chunkServers[0])
		wg.Wait()
	}

	// check that all secondaries committed before proceeding 
	log.Printf("%s here", chunkServers[0])
	log.Println(chunkServerChannels)

	thisChunkServerAddr := chunkServers[0][1:]
	path := chunkServerTempDirectoryPath + thisChunkServerAddr + "/" + strconv.FormatUint(chunkHandle, 10)
	file, err := os.OpenFile(path, os.O_WRONLY, 0644) 
	if err != nil {
		return &protos.Ack{}, status.Errorf(codes.InvalidArgument, "Error opening file to write.") // change
	}
	file.WriteAt(data, 0) // Todo: change offset 

	file.Close()


	// here, commit primary 


	return &protos.Ack{}, nil
}

func (s *ChunkServer) PrimaryCommitMutate(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *ChunkServer) SecondaryCommitMutate(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *ChunkServer) CreateNewChunk(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	chunkHandle := ch.Ch
	filepath := s.Rootpath + "/" + strconv.Itoa(int(chunkHandle))

	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	s.ChunkHandleToFile[chunkHandle] = filepath

	return &protos.Ack{Msg: "successfully replicated chunk on " + s.Address}, nil
}
