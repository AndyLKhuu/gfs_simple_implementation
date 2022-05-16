package services

import (
	"context"
	"errors"
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
	chunkHandleSet     map[uint64]bool     // Set of used chunkhandles
}

func NewMasterServer() *MasterServer {
	return &MasterServer{
		ChunkServerClients: make(map[string]cs.ChunkServerClient),
		Files:              make(map[string][]uint64),
		Chunks:             make(map[uint64][]string),
		chunkHandleSet:     make(map[uint64]bool)}
}

func (s *MasterServer) SendHeartBeatMessage(ctx context.Context, cid *protos.ChunkServerID) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

// TODO: Rename to get chunkLocation for more accurate description
func (s *MasterServer) GetChunkLocation(ctx context.Context, chunkLocReq *protos.ChunkLocationRequest) (*protos.ChunkLocationReply, error) {
	path := chunkLocReq.Path
	chunkIdx := chunkLocReq.ChunkIdx;
	chunkHandle := s.Files[path][chunkIdx]
	chunkServerIds := s.Chunks[chunkHandle]
	return &protos.ChunkLocationReply{ChunkHandle: chunkHandle, ChunkServerIds: chunkServerIds}, nil // TODO fix hard coded values 
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

	ch := s.generateChunkHandle()
	fmt.Printf("created chunk with chunk handle %d \n", ch)
	chunks := append(s.Files[path], ch)

	// Replicate empty chunks
	numChunkServers := len(s.ChunkServerClients)
	// TO:DO Have better replication factor check (> n/2?)
	if int(repFactor) > numChunkServers || repFactor < 1 {
		return &protos.Ack{Message: "Replication Factor is Invalid"}, errors.New("Invalid Replication Factor Value")
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
	// log.Println("TODO: Implement creating file across chunk servers.")

	// At this point, file is created in the shared directory and all chunk servers are allocated

	s.Files[path] = chunks
	return &protos.Ack{Message: fmt.Sprintf("successfuly created file at path %s", path)}, nil
}

func (s *MasterServer) RemoveFile(ctx context.Context, removeReq *protos.FileRemoveRequest) (*protos.Ack, error) {	
	path := removeReq.Path
	e := os.Remove(path)
	if e != nil {
		log.Fatal("couldn't Delete File \n")
	}
	log.Println("TODO: Implement deleting file across chunk servers.")

	return &protos.Ack{Message: fmt.Sprintf("successfuly deleted file at path %s", path)}, nil
}


// Generate a unique chunkhandle.
func (s *MasterServer) generateChunkHandle() uint64 {
	ch := rand.Uint64()
	for s.chunkHandleSet[ch] {
		ch = rand.Uint64()
	}
	s.chunkHandleSet[ch] = true
	return ch
}
