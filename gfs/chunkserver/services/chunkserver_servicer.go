package services

import (
	"context"
	"fmt"
	"gfs/chunkserver/protos"
	"log"
	"os"
	"strconv"
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
	filepath := s.Rootpath + "/" + strconv.Itoa(int(chunkHandle)) + ".txt"

	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	s.ChunkHandleToFile[chunkHandle] = filepath

	return &protos.Ack{Msg: "successfully replicated chunk on " + s.Address}, nil
}

func (s *ChunkServer) ReceiveLease(ctx context.Context, l *protos.LeaseBundle) (*protos.Ack, error) {
	return &protos.Ack{Msg: fmt.Sprintf("successfully received lease for chunk %d", l.Ch)}, nil
}
