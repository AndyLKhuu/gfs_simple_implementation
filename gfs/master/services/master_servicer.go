package services

import (
	"context"
	"fmt"
	cs "gfs/chunkserver/protos"
	"gfs/master/protos"
	"log"
	"math/rand"
	"os"
)

type MasterServer struct {
	protos.UnimplementedMasterServer
	ChunkServerClients map[string]cs.ChunkServerClient
	Files              map[string][]uint64 // Files to ChunkHandles
	Chunks             map[uint64][]string // ChunkHandles to Replica Locations
}

// TO:DO There should be an init function for the Master Server

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
		return &protos.Ack{Message: fmt.Sprintf("failed to create file at path %s", path)}, e
	}

	// Generate random chunk handle
	// TO:DO Ensure that chunkhandles are unique
	ch := rand.Uint64()
	fmt.Printf("created chunk with chunk handle %d \n", ch)
	chunks := append(s.Files[path], ch)

	// Replicate empty chunks
	numChunkServers := len(s.ChunkServerClients)
	// TO:DO Have better replication factor check (> n/2?)
	if int(repFactor) > numChunkServers || repFactor < 1 {
		// TO:DO Create error wrapper for rep factor error
		return &protos.Ack{Message: "Replication Factor is Invalid"}, nil
	}

	// Choose REPFACTOR servers
	// TO:DO Choose replication servers to maximize throughoupt and ensure load balancing
	for k, v := range s.ChunkServerClients {
		res, err := v.CreateNewChunk(context.Background(), &cs.ChunkHandle{Ch: ch})
		if err != nil {
			s.Chunks[ch] = []string{}
			return &protos.Ack{Message: fmt.Sprintf("failed to create file at path %s", path)}, err
		}
		log.Printf(res.Msg)

		s.Chunks[ch] = append(s.Chunks[ch], k)
	}

	s.Files[path] = chunks
	return &protos.Ack{Message: fmt.Sprintf("successfuly created file at path %s", path)}, nil
}
