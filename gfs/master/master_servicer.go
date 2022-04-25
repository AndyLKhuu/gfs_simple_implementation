package master

import (
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
