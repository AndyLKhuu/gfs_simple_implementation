package test

import (
	"gfs/client"
	"log"
	"math/rand"
	"time"
)

var shared_file_path = "../temp_dfs_storage/shared/"

var testPool = []func(c *client.Client){
	SmallFileTest,
	SmallGappedFileTest,
	MediumFileTest,
	LargeFileTest,
}

// Run a specified test
func Run(fn func(c *client.Client), c *client.Client) {
	fn(c)
}

// Run a random test
func RunRandom(c *client.Client) {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	fn := testPool[(random.Intn(len(testPool)))]
	fn(c)
}

// Run all tests
func RunAll(c *client.Client) {
	for i := 0; i < len(testPool); i++ {
		fn := testPool[i]
		fn(c)
	}
}

// Run a random test
func SmallFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: SmallFileTest")
	var str = "hello"
	var smallFileName = shared_file_path + "smallFile.txt"
	c.Create(smallFileName)
	c.Write(smallFileName, 2, []byte(str))
}

func SmallGappedFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: SmallGappedFileTest")
	var str = "hello"
	var smallFileGapName = shared_file_path + "smallFileGap.txt"
	c.Create(smallFileGapName)
	c.Write(smallFileGapName, 70, []byte(str)) // offset of 70!
}

func MediumFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: MediumFileTest")
	medFileName := shared_file_path + "mediumFile.txt"
	var medStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c.Create(medFileName)
	c.Write(medFileName, 20, []byte(medStr))
}

func LargeFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: LargeFileTest")
	lgFileName := shared_file_path + "largeFile.txt"
	var longStr = "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of 'de Finibus Bonorum et Malorum' (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, 'Lorem ipsum dolor sit amet..', comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from 'de Finibus Bonorum et Malorum' by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham."
	c.Create(lgFileName)
	c.Write(lgFileName, 10, []byte(longStr))
}

// Other workflows
