package services

import (
	"context"
	"errors"
	"fmt"
	"gfs/chunkserver/protos"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

// TO:DO Find better naming than ch for chunkhandle messages

var chunkServerTempDirectoryPath = "../temp_dfs_storage/"

type ChunkServer struct {
	protos.UnimplementedChunkServerServer
	ChunkHandleToFile map[uint64]string                  // Chunkhandle to filepath of chunk
	Rootpath          string                             // Root directory path for chunkserver
	Address           string                             // Address of chunkserver
	WriteCache        map[string]*protos.WriteDataBundle // Internal "LRU"
	leases            map[uint64]int64                   // Chunkhandle to end time of lease
}

func NewChunkServer(addr string) ChunkServer {
	chunkserverRootDir := chunkServerTempDirectoryPath + addr
	fmt.Println("starting up chunkserver " + addr + ".")
	if err := os.MkdirAll(chunkserverRootDir, os.ModePerm); err != nil {
		// TO:DO This shouldn't really fatally crash the program, it's just one program
		log.Fatal(err)
	}

	return ChunkServer{
		ChunkHandleToFile: make(map[uint64]string),
		Rootpath:          chunkserverRootDir,
		Address:           addr,
		WriteCache:        make(map[string]*protos.WriteDataBundle),
		leases:            make(map[uint64]int64)}
}

func (s *ChunkServer) Read(ctx context.Context, readReq *protos.ReadRequest) (*protos.ReadReply, error) {
	return &protos.ReadReply{Data: "chicken"}, nil
}

func (s *ChunkServer) ReceiveWriteData(ctx context.Context, writeBundle *protos.WriteDataBundle) (*protos.Ack, error) {
	// ChunkServer stores write data in internal LRU
	transactionId := writeBundle.TransactionId
	s.WriteCache[transactionId] = writeBundle
	log.Printf("Chunkserver %s successfully received write data", s.Address)
	return &protos.Ack{Message: "Chunkserver " + s.Address + " successfully received write data."}, nil
}

func (s *ChunkServer) PrimaryCommitMutate(ctx context.Context, primaryCommitMutateRequest *protos.PrimaryCommitMutateRequest) (*protos.Ack, error) {
	// Primary commits
	chunkHandle := primaryCommitMutateRequest.Ch
	transactionId := primaryCommitMutateRequest.TransactionId

	path := chunkServerTempDirectoryPath + s.Address + "/" + strconv.FormatUint(chunkHandle, 10) + ".txt"
	data := s.WriteCache[transactionId].Data
	offset := s.WriteCache[transactionId].Offset
	status := s.localWriteToFile(transactionId, path, data, offset)
	if status == -1 {
		return &protos.Ack{}, errors.New("Error opening file to write in primaryCS")
	}
	log.Printf("Primary chunkserver %s successfully committed", s.Address)

	// Forward request to secondaries
	secondaryChunkServerAddresses := primaryCommitMutateRequest.SecondaryChunkServerAddresses
	for i := 0; i < len(secondaryChunkServerAddresses); i++ { // TODO: optimize to async
		secondaryChunkServerAddr := secondaryChunkServerAddresses[i]
		conn, err := grpc.Dial(secondaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to secondary chunk server
		defer conn.Close()
		if err != nil {
			log.Printf("error occured when primaryCS dialing to secondaryCS: %s", err)
			return &protos.Ack{}, errors.New("error occured when primaryCS dialing to secondaryCS")
		}

		secondaryChunkServerClient := protos.NewChunkServerClient(conn)
		_, err = secondaryChunkServerClient.SecondaryCommitMutate(context.Background(), &protos.SecondaryCommitMutateRequest{Ch: chunkHandle, TransactionId: transactionId})
		if err != nil {
			log.Printf("error occured on secondaryCommitMutate %s", err)
			return &protos.Ack{}, errors.New("error occured on secondaryCommitMutate")
		}
		// TODO: If there is an error, do we want to roll back the secondary's commit?
	}
	log.Printf("Primary chunkserver %s successfully committed and forwarded to %s", s.Address, secondaryChunkServerAddresses)
	return &protos.Ack{Message: "Primary chunkserver " + s.Address + " successfully committed and forwarded"}, nil
}

func (s *ChunkServer) SecondaryCommitMutate(ctx context.Context, secondaryCommitMutateRequest *protos.SecondaryCommitMutateRequest) (*protos.Ack, error) {
	chunkHandle := secondaryCommitMutateRequest.Ch
	transactionId := secondaryCommitMutateRequest.TransactionId

	path := chunkServerTempDirectoryPath + s.Address + "/" + strconv.FormatUint(chunkHandle, 10) + ".txt"
	data := s.WriteCache[transactionId].Data
	offset := s.WriteCache[transactionId].Offset
	status := s.localWriteToFile(transactionId, path, data, offset)
	if status == -1 {
		return &protos.Ack{}, errors.New("Error opening file to write in secondaryCS")
	}
	log.Printf("Secondary chunkserver %s successfully committed", s.Address)
	return &protos.Ack{Message: "Secondary chunkserver " + s.Address + " successfully committed"}, nil
}

func (s *ChunkServer) CreateNewChunk(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	chunkHandle := ch.Ch
	filepath := s.Rootpath + "/" + strconv.FormatUint(uint64(chunkHandle), 10) + ".txt"

	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	s.ChunkHandleToFile[chunkHandle] = filepath

	return &protos.Ack{Message: "successfully replicated chunk on " + s.Address}, nil
}

func (s *ChunkServer) localWriteToFile(transactionId string, path string, data []byte, offset int64) int {
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return -1
	}
	file.WriteAt(data, offset)
	delete(s.WriteCache, transactionId)
	file.Close()
	return 0
}

func (s *ChunkServer) ReceiveLease(ctx context.Context, l *protos.LeaseBundle) (*protos.Ack, error) {
	s.leases[l.Ch] = l.TimeEnd
	return &protos.Ack{Message: fmt.Sprintf("successfully received lease for chunk %d", l.Ch)}, nil
}
