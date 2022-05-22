package test

import (
	"gfs/client"
	"log"
	"math/rand"
	"time"
)

var shared_file_path = "../temp_dfs_storage/shared/"

var testPool = []func(c *client.Client){
	WriteSmallFileTest,
	WriteSmallGappedFileTest,
	WriteMediumFileTest,
	WriteLargeFileTest,
	WriteReadSmallFileTest,
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
func WriteSmallFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteSmallFileTest")
	var str = "hello"
	var smallFileName = shared_file_path + "smallFile.txt"
	c.Create(smallFileName)
	c.Write(smallFileName, 2, []byte(str))
}

func WriteSmallGappedFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteSmallGappedFileTest")
	var str = "hello"
	var smallFileGapName = shared_file_path + "smallFileGap.txt"
	c.Create(smallFileGapName)
	c.Write(smallFileGapName, 70, []byte(str)) // offset of 70!
}

func WriteMediumFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteMediumFileTest")
	medFileName := shared_file_path + "mediumFile.txt"
	var medStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c.Create(medFileName)
	c.Write(medFileName, 20, []byte(medStr))
}

func WriteLargeFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteLargeFileTest")
	lgFileName := shared_file_path + "largeFile.txt"
	var longStr = "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of 'de Finibus Bonorum et Malorum' (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, 'Lorem ipsum dolor sit amet..', comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from 'de Finibus Bonorum et Malorum' by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham."
	c.Create(lgFileName)
	c.Write(lgFileName, 10, []byte(longStr))
}

func WriteReadSmallFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteReadSmallFileTest")
	var str = "hello"
	var smallFileName = shared_file_path + "smallFile.txt"
	c.Create(smallFileName)
	c.Write(smallFileName, 0, []byte(str))

	readBuffer := make([]byte, 5)
	c.Read(smallFileName, 0, readBuffer)

	correct := string(readBuffer) == str

	log.Printf("input string: %s", str)
	log.Printf("output string: %s", string(readBuffer))

	if correct {
		log.Printf("write input and read output match!")
	} else {
		log.Printf("write input and read output do not match!")
	}

}

func WriteReadMediumFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteReadMediumFileTest")
	medFileName := shared_file_path + "mediumFile.txt"
	var medStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c.Create(medFileName)
	c.Write(medFileName, 0, []byte(medStr))

	length := len(medStr)

	// Read into buffer
	readBuffer := make([]byte, length)
	c.Read(medFileName, 0, readBuffer)

	correct := string(readBuffer) == medStr

	log.Printf("input string: %s", medStr)
	log.Printf("output string: %s", string(readBuffer))

	if correct {
		log.Printf("write input and read output match!")
	} else {
		log.Printf("write input and read output do not match!")
	}
}

func WriteReadLargeFileTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteReadLargeFileTest")
	largeFileName := shared_file_path + "largeFile.txt"
	var longStr = "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of 'de Finibus Bonorum et Malorum' (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, 'Lorem ipsum dolor sit amet..', comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from 'de Finibus Bonorum et Malorum' by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham."
	c.Create(largeFileName)
	c.Write(largeFileName, 0, []byte(longStr))

	length := len(longStr)

	// Read into buffer
	readBuffer := make([]byte, length)
	c.Read(largeFileName, 0, readBuffer)

	correct := string(readBuffer) == longStr

	log.Printf("input string: %s \n", longStr)
	log.Printf("output string: %s \n", string(readBuffer))

	if correct {
		log.Printf("write input and read output match!")
	} else {
		log.Printf("write input and read output do not match!")
	}
}

func WriteReadLargeFileOffsettedTest(c *client.Client) {
	log.Printf("RUNNING TEST: WriteReadLargeFileOffsettedTest")
	largeFileName := shared_file_path + "largeFileOffsetted.txt"
	var longStr = "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of 'de Finibus Bonorum et Malorum' (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, 'Lorem ipsum dolor sit amet..', comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from 'de Finibus Bonorum et Malorum' by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham."
	c.Create(largeFileName)
	c.Write(largeFileName, 5, []byte(longStr)) // Write at offset 5

	length := len(longStr)

	// Read into buffer
	readBuffer := make([]byte, length)
	c.Read(largeFileName, 5, readBuffer) // Read at offset 5

	correct := string(readBuffer) == longStr

	log.Printf("input len, output len: %d, %d", len(longStr), len(readBuffer))

	log.Printf("input string: %s \n", longStr)
	log.Printf("output string: %s \n", string(readBuffer))

	if correct {
		log.Printf("write input and read output match!")
	} else {
		log.Printf("write input and read output do not match!")
	}
}

// Other workflows
var masterServerPort = ":9000"

func MultipleClients_SimpleCreateWriteAndRead() {
	c1, err := client.NewClient(masterServerPort)
	if err != nil {
		log.Printf("failed to initialize client %s", err)
	}
	defer c1.MasterConn.Close()

	c2, err := client.NewClient(masterServerPort)
	if err != nil {
		log.Printf("failed to initialize client %s", err)
	}
	defer c2.MasterConn.Close()

	str := "hello"
	smallFileName := shared_file_path + "smallFile.txt"
	c1.Create(smallFileName)
	c1.Write(smallFileName, 0, []byte(str))

	readBuf := make([]byte, len(str))
	c2.Read(smallFileName, 0, readBuf)

	correct := string(readBuf) == str

	log.Print(string(readBuf))

	if correct {
		log.Printf("works")
	} else {
		log.Printf("nope")
	}
}

func MultipleClients_OverlappingWrites() {
	c1, err := client.NewClient(masterServerPort)
	if err != nil {
		log.Printf("failed to initialize client %s", err)
	}
	defer c1.MasterConn.Close()

	c2, err := client.NewClient(masterServerPort)
	if err != nil {
		log.Printf("failed to initialize client %s", err)
	}
	defer c2.MasterConn.Close()

	str1 := "hello"
	str2 := " there"
	expectedStr := "hello there"
	smallFileName := shared_file_path + "smallFile.txt"
	c1.Create(smallFileName)

	go func() {
		c1.Write(smallFileName, 0, []byte(str1))
	}()

	go func() {
		c2.Write(smallFileName, int64(len(str1)), []byte(str2))
	}()

	// TO:DO Wait for both to finish and then check
	time.Sleep(time.Second * 5)

	readBuf := make([]byte, len(expectedStr))
	c2.Read(smallFileName, 0, readBuf)

	correct := string(readBuf) == expectedStr
	log.Print(string(readBuf))

	if correct {
		log.Printf("works")
	} else {
		log.Printf("nope")
	}
}
