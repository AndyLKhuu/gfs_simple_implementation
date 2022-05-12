package services

import (
	"context"
	"gfs/chunkserver/protos"
	"log"
	"os"
)

var chunkServerTempDirectoryPath = "../temp_dfs_storage/"

type ChunkServer struct {
	protos.UnimplementedChunkServerServer
	ChunkHandleToFile map[uint64]string
	rootpath          string
}

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
	// TO:DO What's the safe way to do these string concatenations?
	filepath := s.rootpath + "chunk"
	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return &protos.Ack{Msg: "successfully replicated chunk."}, nil
}
