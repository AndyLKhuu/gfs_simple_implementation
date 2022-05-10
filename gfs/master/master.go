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
	fmt.Println("Starting up master server.")
	lis, err := net.Listen("tcp", mAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := services.MasterServer{}

	grpcServer := grpc.NewServer()

	protos.RegisterMasterServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			// TO:DO Propagate this error out to fail server
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	go func() {
		// TO:DO Implement some type of repolling mechanism until chunkservers are up. Add a timeout to avoid long waits
		time.Sleep(2 * time.Second) //Arbitrary Number
		// Start and store connections to chunkservers
		for i := 0; i < int(numChunkServers); i++ {
			cs_addr := ":" + strconv.Itoa(chunkServerPortBase+i)
			fmt.Println("Connecting to chunkserver " + cs_addr)
			var conn *grpc.ClientConn
			defer conn.Close()
			conn, err := grpc.Dial(cs_addr, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			fmt.Println("Successfully connected to chunkserver " + cs_addr)

			c := cs.NewChunkServerClient(conn)

			response, err := c.Read(context.Background(), &cs.ReadRequest{})
			if err != nil {
				log.Fatalf("Error when calling Read: %s", err)
			}
			log.Printf("Read reply is : %s", response.Data)
		}
	}()
}
