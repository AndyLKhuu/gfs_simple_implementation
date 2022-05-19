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
	"time"
)

type leaseInfo struct {
	primary string    // Chunkserver which holds the lease
	endTime time.Time // System Time to release lock
}

type MasterServer struct {
	protos.UnimplementedMasterServer
	ChunkServerClients map[string]cs.ChunkServerClient
	Files              map[string][]uint64 // Files to ChunkHandles
	Chunks             map[uint64][]string // ChunkHandles to Replica Locations
	leases             map[uint64]leaseInfo
	chunkHandleSet     map[uint64]bool // Set of used chunkhandles
}

func NewMasterServer() *MasterServer {
	return &MasterServer{
		ChunkServerClients: make(map[string]cs.ChunkServerClient),
		Files:              make(map[string][]uint64),
		Chunks:             make(map[uint64][]string),
		chunkHandleSet:     make(map[uint64]bool),
		leases:             make(map[uint64]leaseInfo)}
}

// Returns the address of the chunkserver which has a lease over the chunk.
func (s *MasterServer) checkOrCreateLease(ch uint64, chunkServers []string) (string, error) {
	l, ok := s.leases[ch]
	if !ok {
		// TO:DO Better heuristic to give lease to chunkserver under lower load, currently giving to first chunkserver
		lease := leaseInfo{primary: chunkServers[0], endTime: time.Now().Add(time.Second * 60)}
		s.leases[ch] = lease
		res, err := s.ChunkServerClients[lease.primary].ReceiveLease(context.Background(), &cs.LeaseBundle{Ch: ch, TimeEnd: lease.endTime.Unix()})
		if err != nil {
			return "", err
		}
		log.Print(res.Msg)
		return lease.primary, nil
	}
	return l.primary, nil
}

// Renew the lease for a chunkhandler at a chunkserver.
func (s *MasterServer) renewLease(ch uint64, addr string) error {
	l, ok := s.leases[ch]
	if !ok {
		return errors.New("no lease has been validated for this chunk.")
	}

	if l.primary != addr {
		return errors.New(fmt.Sprint("given chunkserver %s does not own a lease on chunk %d.", addr, ch))
	}

	l.endTime = time.Now().Add(60 * time.Second)
	s.leases[ch] = l
	return nil
}

func (s *MasterServer) SendHeartBeatMessage(ctx context.Context, cid *protos.ChunkServerID) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *MasterServer) GetChunkLocation(ctx context.Context, chunkLocReq *protos.ChunkLocationRequest) (*protos.ChunkLocationReply, error) {
	path := chunkLocReq.Path
	chunkIdx := chunkLocReq.ChunkIdx
	chunkHandle := s.Files[path][chunkIdx]
	chunkServerIds := s.Chunks[chunkHandle]

	primary, err := s.checkOrCreateLease(chunkHandle, chunkServerIds)
	if err != nil {
		return &protos.ChunkLocationReply{}, err
	}

	chunkServers := []string{primary}
	for _, addr := range chunkServerIds {
		if addr != primary {
			chunkServers = append(chunkServers, addr)
		}
	}

	return &protos.ChunkLocationReply{ChunkHandle: chunkHandle, Primary: primary, ChunkServerIds: chunkServers}, nil // TODO fix hard coded values
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
	//TO:DO We also have to remove the file meta data from the in-memory structures on the master

	return &protos.Ack{Message: fmt.Sprintf("successfuly deleted file at path %s", path)}, nil
}

func (s *MasterServer) RenewChunkLease(ctx context.Context, req *protos.RenewChunkLeaseRequest) (*protos.Ack, error) {
	err := s.renewLease(req.Ch, req.Addr)
	if err != nil {
		return &protos.Ack{}, err
	}
	return &protos.Ack{Message: fmt.Sprintf("successfully renewed lease for chunk %d", ch.Ch)}, nil
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
