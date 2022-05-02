package services

import (
	"context"
	"fmt"
	"gfs/master/protos"
	"log"
	"os"
)

type MasterServer struct {
	protos.UnimplementedMasterServer
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

func CreateFile(filepath string) error {
	_, e := os.Create(filepath)
	if e != nil {
		log.Fatal("Couldn't Create File \n")
	}
	return e
}

func DummyFunction() {
	fmt.Println("hi")
	return
}
