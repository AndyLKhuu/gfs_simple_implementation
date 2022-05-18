package main

import (
	"fmt"
	"gfs/chunkserver"
	"gfs/master"
	"gfs/client"
	"log"
	"os"
	"time"
)

var masterServerPort = ":9000"
var chunkServerPortBase = 10000
var NUM_CHUNK_SERVERS = 3
var NUM_CLIENTS = 1
var shared_file_path = "../temp_dfs_storage/shared/"

func main() {
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

	fmt.Println("-------");

	// Start up Clients
	for i := 0; i < NUM_CLIENTS; i++ {
		go func () {
			c := client.InitClient(masterServerPort) 
			defer c.MasterConn.Close();
			log.Println("Initialized a client");
			fname := shared_file_path + "test.txt";
			c.Create(fname);


			// e2e simple read test.
			// readBuffer := make([]byte, 7)
			// c.Read(fname, 2, readBuffer);
			// log.Printf("Successful read result: %s", string(readBuffer))

			// var str = "hello";
			// c.Write(fname, 2, []byte(str));
			// c.Remove(fname); // careful using this when implementing chunk server repl.

			// var longStr = "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of 'de Finibus Bonorum et Malorum' (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, 'Lorem ipsum dolor sit amet..', comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from 'de Finibus Bonorum et Malorum' by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham."
			var medStr = "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical."
			c.Write(fname, 10, []byte(medStr)); 

			// TODO: currently writing at a file offset is not supported. need to propogate offset to chunkserver's write

		
			}()
	}

	log.Println("Done.")
	select {}
}
