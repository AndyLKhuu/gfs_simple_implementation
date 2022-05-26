package test

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

	time.Sleep(2 * time.Second) //Arbitrary Number

	exitVal := m.Run()

	_ = os.RemoveAll("../temp_dfs_storage")

	os.Exit(exitVal)
}

func Test_MultipleClientsSimpleCreateWriteAndRead(t *testing.T) {
	c1, err := client.NewClient(masterServerPort)
	defer c1.MasterConn.Close()
	assert.NoError(t, err)

	c2, err := client.NewClient(masterServerPort)
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

func Test_MultipleClientsOverlappingWrites(t *testing.T) {
	c1, err := client.NewClient(masterServerPort)
	assert.NoError(t, err)
	defer c1.MasterConn.Close()

	c2, err := client.NewClient(masterServerPort)
	assert.NoError(t, err)
	defer c2.MasterConn.Close()

	str1 := "hello"
	str2 := " there"
	expectedStr := "hello there"
	smallFileName := shared_file_path + "testFile1.txt"
	success := c1.Create(smallFileName)
	assert.Equal(t, 0, success)

	done := make(chan bool)

	go func() {
		actualBytesWritten := c1.Write(smallFileName, 0, []byte(str1))
		assert.Equal(t, len(str1), actualBytesWritten)
		done <- true
	}()

	go func() {
		actualBytesWritten := c2.Write(smallFileName, uint64(len(str1)), []byte(str2))
		assert.Equal(t, len(str2), actualBytesWritten)
		done <- true
	}()

	firstWriteSuccess := <-done
	assert.True(t, firstWriteSuccess)
	secondWriteSuccess := <-done
	assert.True(t, secondWriteSuccess)

	readBuf := make([]byte, len(expectedStr))
	actualBytesRead := c2.Read(smallFileName, 0, readBuf)

	// This is a purposely redundant test to ensure correctness.
	assert.Equal(t, len(expectedStr), actualBytesRead)
	assert.Equal(t, expectedStr, string(readBuf))
}

func Test_ClientWriteRemove(t *testing.T) {
	c1, err := client.NewClient(masterServerPort)
	assert.NoError(t, err)
	defer c1.MasterConn.Close()

	c2, err := client.NewClient(masterServerPort)
	assert.NoError(t, err)
	defer c2.MasterConn.Close()

	longStr := "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of 'de Finibus Bonorum et Malorum' (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, 'Lorem ipsum dolor sit amet..', comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from 'de Finibus Bonorum et Malorum' by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham."
	largeFileName := shared_file_path + "largeFile.txt"
	success1 := c1.Create(largeFileName)
	assert.Equal(t, 0, success1)

	medStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	medFileName := shared_file_path + "mediumFile.txt"
	success2 := c2.Create(medFileName)
	assert.Equal(t, 0, success2)

	done := make(chan bool)

	go func() {
		actualBytesWritten := c1.Write(longStr, 0, []byte(longStr))
		assert.Equal(t, len(longStr), actualBytesWritten)
		assert.Equal(t, 0, c1.Remove(largeFileName))
		done <- true
	}()

	go func() {
		actualBytesWritten := c2.Write(medStr, 0, []byte(medStr))
		assert.Equal(t, len(medStr), actualBytesWritten)
		assert.Equal(t, 0, c2.Remove(medFileName))
		done <- true
	}()

	firstWriteRemoveSuccess := <-done
	assert.True(t, firstWriteRemoveSuccess)

	secondWriteRemoveSuccess := <-done
	assert.True(t, secondWriteRemoveSuccess)

	_, medFileDoesNotExist := os.Stat(medFileName)
	assert.Equal(t, true, medFileDoesNotExist != nil)

	_, largeFileDoesNotExist := os.Stat(largeFileName)
	assert.Equal(t, true, largeFileDoesNotExist != nil)

}
