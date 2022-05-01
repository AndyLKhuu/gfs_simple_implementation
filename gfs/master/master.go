package master

import (
	"log"
	"net"

	"gfs/master/protos"
	"gfs/master/services"

	"google.golang.org/grpc"
)

func InitMasterServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.MasterServer{}

	grpcServer := grpc.NewServer()

	protos.RegisterMasterServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
