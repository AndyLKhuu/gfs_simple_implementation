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
	cs_clients []cs.ChunkServerClient
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

func createFile(filepath string) error {
	_, e := os.Create(filepath)
	if e != nil {
		log.Fatal("couldn't Create File \n")
	}
	return e
}

func DummyFunction() {
	fmt.Println("hi")
	return
}
