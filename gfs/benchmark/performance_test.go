package benchmark

import (
	"bytes"
	"gfs/chunkserver"
	"gfs/chunkserver/services"
	"gfs/client"
	"gfs/master"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ----------- BENCHMARK CONFIGURATIONS HERE --------------------------

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 3

var NTRIALS = 1 // repeat test NTRIALS time to collect

var ARCHITECTURE = "ASYNC"           // SEQ for sequential; ASYNC for async
var INPUT_FILE = "textfile_500b.txt" // input files vary in size
var INPUT_FILES = []string{
	"textfile_500b.txt",
	"textfile_1kb.txt",
	"textfile_1.5kb.txt",
	"textfile_5kb.txt",
	"textfile_10kb.txt",
	"textfile_20kb.txt",
	"textfile_30kb.txt",
	"textfile_40kb.txt",
	"textfile_64kb.txt",
	"textfile_100kb.txt",
	"textfile_250kb.txt",
	"textfile_500kb.txt",
	"textfile_750kb.txt",

	// x1000 larger in magnitude...
	"textfile_1mb.txt",
	"textfile_2mb.txt",
	"textfile_3mb.txt",
	"textfile_5mb.txt",

	"textfile_6mb.txt",
	"textfile_9mb.txt",
	"textfile_10mb.txt",
	"textfile_15mb.txt",

	"textfile_18mb.txt",
	"textfile_20mb.txt",

	// any larger, we are approaching the uppbound of GRPC's message capacity.
	// ERROR with GRPC
}

// ---------------------------------------------------------------------

// Client's benchmarking configuration
var CLIENT_BENCHMARK_CONFIG = client.BenchmarkConfig{
	Benchmarking: true,
	OutputDir:    "results_" + ARCHITECTURE + "/",
	Architecture: ARCHITECTURE,
}

// Chunkserver's benchmarking configuration
var CS_BENCHMARK_CONFIG = services.BenchmarkConfig{
	Benchmarking: true,
	OutputDir:    "results_" + ARCHITECTURE + "/",
	Architecture: ARCHITECTURE,
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
		go chunkserver.InitBenchmarkingChunkServer(chunkServerPortBase+i, CS_BENCHMARK_CONFIG)
	}

	log.Println("Creating benchmark output files")
	err := os.MkdirAll(CLIENT_BENCHMARK_CONFIG.OutputDir, os.ModePerm)
	if err != nil {
		log.Printf("err: %s", err)
	}

	time.Sleep(2 * time.Second) //Arbitrary Number

	exitVal := m.Run()

	// _ = os.RemoveAll("../temp_dfs_storage")

	os.Exit(exitVal)
}

func Test_Single3000ByteCreateWriteReadRemove(t *testing.T) {

	for _, file := range INPUT_FILES {

		inputFile := file
		for i := 0; i < NTRIALS; i++ {

			c1, err := client.NewBenchmarkingClient(masterServerPort, CLIENT_BENCHMARK_CONFIG)
			defer c1.MasterConn.Close()
			assert.NoError(t, err)

			targetFileName := "files/" + inputFile

			data, err := ioutil.ReadFile(targetFileName)
			if err != nil {
				log.Printf("error reading file %s: %s", targetFileName, err)
			}

			filename := inputFile
			c1.Create(filename)
			c1.Write(filename, 0, data)

			readBuffer := make([]byte, len(data))

			c1.Read(filename, 0, readBuffer)

			c1.Remove(filename)

			match := bytes.Compare(data, readBuffer)
			assert.True(t, match == 0, "ERROR: input and outputs do not match. Invalidate benchmark results.")
		}
	}

}

// NOTE: this does not run successfully at the moment. possibly bugs...?

// func Test_Multi3000ByteCreateWriteReadRemove(t *testing.T) {

// 	var F = []string{"textfile_500b.txt", "textfile_64kb.txt", "textfile_5mb.txt"}

// 	for _, file := range F {

// 		inputFile := file
// 		for i := 0; i < NTRIALS; i++ {

// 			var client_wg sync.WaitGroup
// 			for clientNum := 0; clientNum < 2; clientNum++ {
// 				client_wg.Add(1)
// 				go func() {

// 					c1, err := client.NewBenchmarkingClient(masterServerPort, CLIENT_BENCHMARK_CONFIG)
// 					defer c1.MasterConn.Close()
// 					assert.NoError(t, err)

// 					targetFileName := "files/" + inputFile

// 					data, err := ioutil.ReadFile(targetFileName)
// 					if err != nil {
// 						log.Printf("error reading file %s: %s", targetFileName, err)
// 					}

// 					filename := inputFile
// 					c1.Create(filename)
// 					c1.Write(filename, 0, data)

// 					readBuffer := make([]byte, len(data))

// 					c1.Read(filename, 0, readBuffer)

// 					// c1.Remove(filename)

// 					match := bytes.Compare(data, readBuffer)
// 					assert.True(t, match == 0, "ERROR: input and outputs do not match. Invalidate benchmark results.")
// 					client_wg.Done()

// 				}()

// 			}
// 			client_wg.Wait()

// 		}
// 	}

// }
