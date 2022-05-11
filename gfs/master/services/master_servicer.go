package services

import (
	"context"
	"fmt"
	cs "gfs/chunkserver/protos"
	"gfs/master/protos"
	"log"
	"os"
)

type MasterServer struct {
	protos.UnimplementedMasterServer
	Cs_clients map[string]cs.ChunkServerClient
}

func (s *MasterServer) SendHeartBeatMessage(ctx context.Context, cid *protos.ChunkServerID) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *MasterServer) GetFileLocation(ctx context.Context, chunkLocReq *protos.ChunkLocationRequest) (*protos.ChunkLocationReply, error) {
	
	filepath := chunkLocReq.Filepath
	chunkIdx := chunkLocReq.ChunkIdx;
	log.Printf("Master.GetFileLocation invoked with %s, %d", filepath, chunkIdx);
	
	return &protos.ChunkLocationReply{ChunkHandler: 0, ChunkServerIds: []int64{0, 0}}, nil // TODO fix hard coded values 
}

func (s *MasterServer) GetSystemChunkSize(ctx context.Context, sysChunkSizeReq *protos.SystemChunkSizeRequest) (*protos.ChunkSize, error) {
	return &protos.ChunkSize{Size: 64}, nil
}

func (s *MasterServer) ReceiveClientWriteRequest(ctx context.Context, clientWriteReq *protos.ClientWriteRequest) (*protos.Ack, error) {
	return &protos.Ack{}, nil
}

func (s *MasterServer) ReceiveClientCreateRequest(ctx context.Context, clientCreateReq *protos.ClientCreateRequest) (*protos.Ack, error) {
	err := createFile(clientCreateReq.Filepath);
	return &protos.Ack{}, err
}

func (s *MasterServer) ReceiveClientDeleteRequest(ctx context.Context, clientDeleteReq *protos.ClientDeleteRequest) (*protos.Ack, error) {
	err := deleteFile(clientDeleteReq.Filepath);
	return &protos.Ack{}, err
}

func createFile(filepath string) error {
	_, e := os.Create(filepath)
	if e != nil {
		log.Fatal("couldn't Create File \n")
	}
	log.Println("TODO: Implement creating file across chunk servers.")
	return e
}

func deleteFile(filepath string) error {
	e := os.Remove(filepath)
	if e != nil {
		log.Fatal("couldn't Delete File \n")
	}
	log.Println("TODO: Implement deleting file across chunk servers.")

	return e
}

func DummyFunction() {
	fmt.Println("hi")
	return
}
