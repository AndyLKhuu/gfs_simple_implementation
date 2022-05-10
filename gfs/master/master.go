package master

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	cs "gfs/chunkserver/protos"
	"gfs/master/protos"
	"gfs/master/services"

	"google.golang.org/grpc"
)

func InitMasterServer(mAddr string, numChunkServers int, chunkServerPortBase int) {
	fmt.Println("starting up master server.")
	lis, err := net.Listen("tcp", mAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.MasterServer{}

	grpcServer := grpc.NewServer()

	protos.RegisterMasterServer(grpcServer, &s)

	// Serve Master Routine
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			// TO:DO Propagate this error out to fail server
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// Start and store connections to chunkservers
	go func() {
		for i := 0; i < int(numChunkServers); i++ {
			cs_addr := ":" + strconv.Itoa(chunkServerPortBase+i)
			log.Printf("connecting to chunkserver %s", cs_addr)
			var conn *grpc.ClientConn

			conn, err := grpc.Dial(cs_addr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
			if err != nil {
				log.Printf("did not connect to chunkserver %s", cs_addr)
			}
			log.Printf("successfully connected to chunkserver %s", cs_addr)

			c := cs.NewChunkServerClient(conn)

			response, err := c.Read(context.Background(), &cs.ReadRequest{})
			if err != nil {
				log.Fatalf("error when calling Read: %s", err)
			}
			log.Printf("read reply is : %s", response.Data)
		}
	}()
}
