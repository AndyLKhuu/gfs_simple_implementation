package services

import (
	"context"
	"fmt"
	"gfs/chunkserver/protos"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var chunkServerTempDirectoryPath = "../temp_dfs_storage/"

var ASYNC = true

type ChunkServer struct {
	protos.UnimplementedChunkServerServer
	ChunkHandleToFile map[uint64]string                  // Chunkhandle to filepath of chunk
	Rootpath          string                             // Root directory path for chunkserver
	Address           string                             // Address of chunkserver
	WriteCache        map[string]*protos.WriteDataBundle // Internal "LRU"
	leases            map[uint64]int64                   // Chunkhandle to end time of lease
	pendingChunkTxs   map[uint64][]string                // Chunkhandle to list of pending transactions
	completedTxs      map[string]bool                    // Set of completed transactions
	inProgressTxs     map[string]bool                    // Set of transactions currently in progress
	pendingTxsLock    sync.Mutex                         // Lock for pending Txs list
	writeCacheLock    sync.RWMutex                       // RWLock for Write Cache
}

func NewChunkServer(addr string) ChunkServer {
	fmt.Println("starting up chunkserver " + addr + ".")
	chunkserverRootDir := chunkServerTempDirectoryPath + addr
	if err := os.MkdirAll(chunkserverRootDir, os.ModePerm); err != nil {
		// TO:DO This shouldn't really fatally crash the program, it's just one program
		log.Fatal(err)
	}

	return ChunkServer{
		ChunkHandleToFile: make(map[uint64]string),
		Rootpath:          chunkserverRootDir,
		Address:           addr,
		WriteCache:        make(map[string]*protos.WriteDataBundle),
		leases:            make(map[uint64]int64),
		pendingChunkTxs:   make(map[uint64][]string),
		completedTxs:      make(map[string]bool),
		inProgressTxs:     make(map[string]bool),
		pendingTxsLock:    sync.Mutex{},
		writeCacheLock:    sync.RWMutex{}}
}

func (s *ChunkServer) Read(ctx context.Context, readReq *protos.ReadRequest) (*protos.ReadReply, error) {
	chunkHandle := readReq.Ch
	leftBound := readReq.L
	rightBound := readReq.R

	path := chunkServerTempDirectoryPath + s.Address + "/" + strconv.FormatUint(chunkHandle, 10) + ".txt"
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return &protos.ReadReply{}, err
	}
	length := rightBound - leftBound
	data := make([]byte, length)
	_, err = file.ReadAt(data, int64(leftBound))
	if err != nil {
		return &protos.ReadReply{}, err
	}

	return &protos.ReadReply{Data: string(data)}, nil
}

func (s *ChunkServer) ReceiveWriteData(ctx context.Context, writeBundle *protos.WriteDataBundle) (*protos.Ack, error) {
	transactionId := writeBundle.TxId
	ch := writeBundle.Ch
	s.writeCacheLock.Lock()
	s.WriteCache[transactionId] = writeBundle
	s.writeCacheLock.Unlock()

	s.pendingTxsLock.Lock()
	s.pendingChunkTxs[ch] = append(s.pendingChunkTxs[ch], transactionId)
	s.pendingTxsLock.Unlock()
	return &protos.Ack{Message: "Chunkserver " + s.Address + " successfully received write data."}, nil
}

func (s *ChunkServer) PrimaryCommitMutate(ctx context.Context, primaryCommitMutateRequest *protos.PrimaryCommitMutateRequest) (*protos.Ack, error) {
	chunkHandle := primaryCommitMutateRequest.Ch
	transactionId := primaryCommitMutateRequest.TxId

	// TO:DO Fix coarse locking.
	s.pendingTxsLock.Lock()
	if s.completedTxs[transactionId] {
		s.pendingTxsLock.Unlock()
		return &protos.Ack{Message: "Primary chunkserver " + s.Address + " successfully committed and forwarded"}, nil
	}

	_, ok := s.ChunkHandleToFile[chunkHandle]
	if !ok {
		s.pendingTxsLock.Unlock()
		return &protos.Ack{Message: fmt.Sprintf("chunk %d not longer exists", chunkHandle)}, nil
	}

	// Serialize all mutations in some order
	mutations := s.pendingChunkTxs[chunkHandle]
	serialOrder := []string{}
	for _, txId := range mutations {
		serialOrder = append(serialOrder, txId)
		delete(s.inProgressTxs, txId)
		s.completedTxs[txId] = true
	}

	// Primary Commits
	s.applyMutations(serialOrder, chunkHandle)

	// Forward request to secondaries
	secondaryChunkServerAddresses := primaryCommitMutateRequest.SecondaryChunkServerAddresses

	if !ASYNC {
		for i := 0; i < len(secondaryChunkServerAddresses); i++ {
			secondaryChunkServerAddr := secondaryChunkServerAddresses[i]
			conn, err := grpc.Dial(secondaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to secondary chunk server
			defer conn.Close()
			if err != nil {
				s.pendingTxsLock.Unlock()
				return &protos.Ack{}, err
			}

			secondaryChunkServerClient := protos.NewChunkServerClient(conn)
			_, err = secondaryChunkServerClient.SecondaryCommitMutate(context.Background(), &protos.SecondaryCommitMutateRequest{Ch: chunkHandle, TxIds: mutations})
			if err != nil {
				s.pendingTxsLock.Unlock()
				return &protos.Ack{}, err
			}
			// TODO: If forwarding fails, keep trying until success or reach some boundary
		}
	} else {

		type response struct {
			Ack *protos.Ack
			err error
		}

		done := make(chan response, len(secondaryChunkServerAddresses))

		var client_wg sync.WaitGroup
		for i := 0; i < len(secondaryChunkServerAddresses); i++ {
			client_wg.Add(1)
			go func(i int) {
				defer client_wg.Done()
				secondaryChunkServerAddr := secondaryChunkServerAddresses[i]
				conn, err := grpc.Dial(secondaryChunkServerAddr, grpc.WithTimeout(5*time.Second), grpc.WithInsecure()) // connecting to secondary chunk server
				defer conn.Close()
				if err != nil {
					s.pendingTxsLock.Unlock()
					done <- response{&protos.Ack{}, err}
					return
				}

				secondaryChunkServerClient := protos.NewChunkServerClient(conn)
				_, err = secondaryChunkServerClient.SecondaryCommitMutate(context.Background(), &protos.SecondaryCommitMutateRequest{Ch: chunkHandle, TxIds: mutations})
				if err != nil {
					s.pendingTxsLock.Unlock()
					done <- response{&protos.Ack{}, err}
					return
				}

			}(i)
		}
		client_wg.Wait()
		close(done)

		// Check that all async processes did their job; clean ack
		for response := range done {
			if response.err != nil {
				return response.Ack, response.err
			}
		}

	}

	s.pendingTxsLock.Unlock()

	return &protos.Ack{Message: "Primary chunkserver " + s.Address + " successfully committed and forwarded"}, nil
}

func (s *ChunkServer) SecondaryCommitMutate(ctx context.Context, secondaryCommitMutateRequest *protos.SecondaryCommitMutateRequest) (*protos.Ack, error) {
	chunkHandle := secondaryCommitMutateRequest.Ch
	txOrder := secondaryCommitMutateRequest.TxIds
	err := s.applyMutations(txOrder, chunkHandle)
	if err != nil {
		return &protos.Ack{}, err
	}
	return &protos.Ack{Message: "Secondary chunkserver " + s.Address + " successfully committed"}, nil
}

func (s *ChunkServer) CreateNewChunk(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	chunkHandle := ch.Ch
	filepath := s.Rootpath + "/" + strconv.FormatUint(uint64(chunkHandle), 10) + ".txt"

	s.pendingTxsLock.Lock()
	_, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	s.ChunkHandleToFile[chunkHandle] = filepath
	s.pendingTxsLock.Unlock()

	return &protos.Ack{Message: "successfully replicated chunk on " + s.Address}, nil
}

func (s *ChunkServer) RemoveChunk(ctx context.Context, ch *protos.ChunkHandle) (*protos.Ack, error) {
	chunkHandle := ch.Ch
	filepath := s.Rootpath + "/" + strconv.FormatUint(uint64(chunkHandle), 10) + ".txt"

	s.pendingTxsLock.Lock()
	err := os.Remove(filepath)
	if err != nil {
		s.pendingTxsLock.Unlock()
		log.Fatal(err)
	}
	delete(s.ChunkHandleToFile, chunkHandle)
	s.pendingTxsLock.Unlock()

	return &protos.Ack{Message: "successfully removed chunk on " + s.Address}, nil
}

func (s *ChunkServer) ReceiveLease(ctx context.Context, l *protos.LeaseBundle) (*protos.Ack, error) {
	s.leases[l.Ch] = l.TimeEnd
	return &protos.Ack{Message: fmt.Sprintf("successfully received lease for chunk %d", l.Ch)}, nil
}

func (s *ChunkServer) applyMutations(mutationOrder []string, chunkHandle uint64) error {
	path := chunkServerTempDirectoryPath + s.Address + "/" + strconv.FormatUint(chunkHandle, 10) + ".txt"
	for _, txId := range mutationOrder {
		s.writeCacheLock.RLock()
		bundle, ok := s.WriteCache[txId]
		// TO:DO Fix this really hacky way to wait for data to be transmitted to secondary.
		if !ok {
			time.Sleep(time.Second * 5)
			bundle, _ = s.WriteCache[txId]
		}
		s.writeCacheLock.RUnlock()
		data := bundle.Data
		offset := bundle.Offset
		err := s.localWriteToFile(txId, path, data, offset)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ChunkServer) localWriteToFile(transactionId string, path string, data []byte, offset uint64) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteAt(data, int64(offset))
	if err != nil {
		return err
	}

	s.writeCacheLock.Lock()
	delete(s.WriteCache, transactionId)
	s.writeCacheLock.Unlock()
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
