package main

import (
	"fmt"
	"log"
	"os"
)

func handle_create_file(filepath string) error {
	_, e := os.Create(filepath)
	if e != nil {
		log.Fatal("Couldn't Create File \n")
	}
	return e
}

func main() {
	fmt.Println("Master server is running")
	handle_create_file("../temp_dfs_storage/hello.txt")
}
