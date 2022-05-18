package services

import (
	"context"
	"gfs/chunkserver/protos"
	"log"
	"os"
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
	WriteCache 				map[uint64]*protos.WriteDataBundle // Internal "LRU" 
}

func (s *ChunkServer) Read(ctx context.Context, readReq *protos.ReadRequest) (*protos.ReadReply, error) {
	return &protos.ReadReply{Data: "chicken"}, nil
}

func (s *ChunkServer) ReceiveWriteData(ctx context.Context, writeBundle *protos.WriteDataBundle) (*protos.Ack, error) {
	// ChunkServer stores write data in internal LRU
	chunkHandle := writeBundle.Ch
	s.WriteCache[chunkHandle] = writeBundle
	return &protos.Ack{}, nil
}

func (s *ChunkServer) PrimaryCommitMutate(ctx context.Context, primaryCommitMutateRequest *protos.PrimaryCommitMutateRequest) (*protos.Ack, error) {
	chunkHandle := primaryCommitMutateRequest.Ch
	secondaryChunkServerAddresses := primaryCommitMutateRequest.SecondaryChunkServerAddresses
	for i := 0; i < len(secondaryChunkServerAddresses); i++ { // Should we async this instead of sequential?
		secondaryChunkServerAddr := secondaryChunkServerAddresses[i]
		conn, err := grpc.Dial(secondaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to secondary chunk server
		defer conn.Close()
		if err != nil {
			log.Printf("error occured when primaryCS diaing to secondaryCS: %s", err)
			return &protos.Ack{}, status.Errorf(codes.Unavailable, "error occured when primaryCS diaing to secondaryCS")
		}

		// Primary forwards commit request to secondary
		secondaryChunkServerClient := protos.NewChunkServerClient(conn);
		_, err = secondaryChunkServerClient.SecondaryCommitMutate(context.Background(), &protos.ChunkHandle{Ch: chunkHandle})
		if err != nil {
			log.Printf("error occured on secondaryCommitMutate %s", err)
			return &protos.Ack{}, status.Errorf(codes.Unavailable, "error occured on secondaryCommitMutate")
		}
		// TODO: If there is an error, do we want to roll back the secondary's commit?
	}

	// Primary commits 
	path := chunkServerTempDirectoryPath + s.Address + "/" + strconv.FormatUint(chunkHandle, 10)
	file, err := os.OpenFile(path, os.O_WRONLY, 0644) 
	if err != nil {
		return &protos.Ack{}, status.Errorf(codes.InvalidArgument, "Error opening file to write in primaryCS")
	}
	data := s.WriteCache[chunkHandle].Data
	offset :=  s.WriteCache[chunkHandle].Offset

	file.WriteAt(data, offset)
	delete(s.WriteCache, chunkHandle)
	file.Close()
	return &protos.Ack{}, nil
}

func (s *ChunkServer) SecondaryCommitMutate(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	chunkHandle := ch.Ch
	path := chunkServerTempDirectoryPath + s.Address + "/" + strconv.FormatUint(chunkHandle, 10)
	file, err := os.OpenFile(path, os.O_WRONLY, 0644) 
	if err != nil {
		return &protos.Ack{}, status.Errorf(codes.InvalidArgument, "Error opening file to write in secondaryCS") // change
	}
	data := s.WriteCache[chunkHandle].Data
	offset := s.WriteCache[chunkHandle].Offset
	file.WriteAt(data, offset)
	delete(s.WriteCache, chunkHandle)
	file.Close()
	return &protos.Ack{}, nil
}

func (s *ChunkServer) CreateNewChunk(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	chunkHandle := ch.Ch
	filepath := s.Rootpath + "/" + strconv.FormatUint(uint64(chunkHandle), 10)

	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	s.ChunkHandleToFile[chunkHandle] = filepath

	return &protos.Ack{Msg: "successfully replicated chunk on " + s.Address}, nil
}
