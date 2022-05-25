package benchmark

import (
	"gfs/chunkserver"
	"gfs/client"
	"gfs/master"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 3
var NUM_CLIENTS = 1

var shared_file_path = "../temp_dfs_storage/shared/"

var BENCHMARK_CONFIG = client.BenchmarkConfig{
	OutputDir:    "results/",
	Architecture: "SEQ",
}

func TestMain(m *testing.M) {
	// TO:DO How to handle errors on the set up
	if err := os.MkdirAll("../temp_dfs_storage", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Start up Master Server
	go master.InitMasterServer(masterServerPort, NUM_CHUNK_SERVERS, chunkServerPortBase)

	// Start up Chunkservers
	for i := 0; i < NUM_CHUNK_SERVERS; i++ {
		go chunkserver.InitChunkServer(chunkServerPortBase + i)
	}

	log.Println("Creating benchmark output files")
	err := os.MkdirAll(BENCHMARK_CONFIG.OutputDir, os.ModePerm)
	if err != nil {
		log.Printf("err: %s", err)
	}
	// operations := []string{"WRITE", "READ", "CREATE", "REMOVE"}
	// for _, op := range operations {
	// 	os.Create(BENCHMARK_CONFIG.OutputDir + op + ".csv")
	// }

	time.Sleep(2 * time.Second) //Arbitrary Number

	exitVal := m.Run()

	_ = os.RemoveAll("../temp_dfs_storage")

	os.Exit(exitVal)
}

func Test_MultipleClientsSimpleCreateWriteAndRead(t *testing.T) {
	c1, err := client.NewBenchmarkingClient(masterServerPort, BENCHMARK_CONFIG)
	defer c1.MasterConn.Close()
	assert.NoError(t, err)

	c2, err := client.NewBenchmarkingClient(masterServerPort, BENCHMARK_CONFIG)
	assert.NoError(t, err)
	defer c2.MasterConn.Close()

	str := "hello"
	smallFileName := shared_file_path + "testFile.txt"
	success := c1.Create(smallFileName)
	assert.Equal(t, 0, success)

	actualBytesWritten := c1.Write(smallFileName, 0, []byte(str))
	assert.Equal(t, len(str), actualBytesWritten)

	readBuf := make([]byte, len(str))
	actualBytesRead := c2.Read(smallFileName, 0, readBuf)

	// This is a purposely redundant test to ensure correctness.
	assert.Equal(t, len(readBuf), actualBytesRead)
	assert.Equal(t, string(readBuf), str)
}
