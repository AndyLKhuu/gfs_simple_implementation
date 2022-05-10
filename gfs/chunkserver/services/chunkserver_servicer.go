package services

import (
	"context"
	"gfs/chunkserver/protos"
)

type ChunkServer struct {
	protos.UnimplementedChunkServerServer
}

func (s *ChunkServer) Read(ctx context.Context, readReq *protos.ReadRequest) (*protos.ReadReply, error) {
	return &protos.ReadReply{Data: "chicken"}, nil
}

func (s *ChunkServer) ReceiveWriteData(ctx context.Context, writeBundle *protos.WriteDataBundle) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *ChunkServer) PrimaryCommitMutate(ctx context.Context, ch *protos.ChunkHandler) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *ChunkServer) SecondaryCommitMutate(ctx context.Context, ch *protos.ChunkHandler) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *ChunkServer) CreateNewChunk(ctx context.Context, ch *protos.ChunkHandler) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}
