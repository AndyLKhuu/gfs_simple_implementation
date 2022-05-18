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

func (s *MasterServer) GetChunkLocation(ctx context.Context, chunkLocReq *protos.ChunkLocationRequest) (*protos.ChunkLocationReply, error) {
	path := chunkLocReq.Path
	chunkIdx := chunkLocReq.ChunkIdx;

	// Generate new chunks up to the chunkIdx we want. Question: do we want our FS to support files with gaps?
	for (int32(len(s.Files[path])) < chunkIdx + 1) {
		s.createNewChunk(path)
	}
	chunkHandle := s.Files[path][chunkIdx] 
	chunkServerIds := s.Chunks[chunkHandle]

	return &protos.ChunkLocationReply{ChunkHandle: chunkHandle, ChunkServerIds: chunkServerIds}, nil 
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

	// ch := s.generateChunkHandle()

	numChunkServers := len(s.ChunkServerClients)
	// TO:DO Have better replication factor check (> n/2?) // I feel like this should be a boot up checker and doesnt need to be performed every time we create a file. 
	if int(repFactor) > numChunkServers || repFactor < 1 {
		return &protos.Ack{Message: "Replication Factor is Invalid"}, errors.New("Invalid Replication Factor Value")
	}

	createNewChunkStatus := s.createNewChunk(path);
	if (createNewChunkStatus == -1) {
		return &protos.Ack{Message: "Error creating new chunk"}, errors.New("Error creating new chunk")
	}

	return &protos.Ack{Message: fmt.Sprintf("successfuly created file at path %s", path)}, nil
}

func (s *MasterServer) RemoveFile(ctx context.Context, removeReq *protos.FileRemoveRequest) (*protos.Ack, error) {	
	path := removeReq.Path
	e := os.Remove(path)
	if e != nil {
		log.Fatal("couldn't Delete File \n")
	}
	log.Println("TODO: Implement deleting file across chunk servers.")
 	//TO:DO We also have to remove the file meta data from the in-memory structures on the master

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

func (s *MasterServer) createNewChunk(path string)  int {
	ch := s.generateChunkHandle()
	chunks := append(s.Files[path], ch)
	// Choose REPFACTOR servers
	// TO:DO Choose replication servers to maximize throughoupt and ensure load balancing
	for k, v := range s.ChunkServerClients {
		res, err := v.CreateNewChunk(context.Background(), &cs.ChunkHandle{Ch: ch})
		if err != nil {
			s.Chunks[ch] = []string{}
			return -1
		}
		log.Printf(res.Msg)

		s.Chunks[ch] = append(s.Chunks[ch], k)
	}
	s.Files[path] = chunks
	return 0
}

