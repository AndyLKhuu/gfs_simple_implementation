package services

import (
	"context"
	"fmt"
	cs "gfs/chunkserver/protos"
	"gfs/master/protos"
	"os"
)

type MasterServer struct {
	protos.UnimplementedMasterServer
	ChunkServerClients map[string]cs.ChunkServerClient
	Files              map[string][]int
}

func (s *MasterServer) SendHeartBeatMessage(ctx context.Context, cid *protos.ChunkServerID) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *MasterServer) GetFileLocation(ctx context.Context, chunkLocReq *protos.ChunkLocationRequest) (*protos.ChunkLocationReply, error) {
	return &protos.ChunkLocationReply{}, nil
}

func (s *MasterServer) GetSystemChunkSize(ctx context.Context, sysChunkSizeReq *protos.SystemChunkSizeRequest) (*protos.ChunkSize, error) {
	return &protos.ChunkSize{Size: 64}, nil
}

func (s *MasterServer) ReceiveClientWriteRequest(ctx context.Context, clientWriteReq *protos.ClientWriteRequest) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *MasterServer) CreateFile(ctx context.Context, createReq *protos.FileCreateRequest) (*protos.Ack, error) {
	path := createReq.Path
	repFactor := createReq.RepFactor
	// Create File
	_, e := os.Create(path)
	if e != nil {
		return &protos.Ack{Message: fmt.Sprintf("failed to create file at path %s", path)}, nil
	}

	// Replicate empty chunks
	numChunkServers := len(s.ChunkServerClients)
	if repFactor > numChunkServers || repFactor < 1 {
		return &protos.Ack{Message: "Replication Factor is Invalid"}, nil
	}

	// Choose first REPFACTOR serers
	for k, v := range s.ChunkServerClients {
		v.CreateNewChunk(context.Background(), cs.ChunkHandler{})
	}

	return &protos.Ack{Message: fmt.Sprintf("successfuly created file at path %s", path)}, nil
}
