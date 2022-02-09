package examplejson

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func InitExampleFile() {
	// ReadFile("./example_json/person_data.json")

	f, err := os.Open("./example_json/text")
	checkError(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n, err := f.Read(b1)
	checkError(err)
	fmt.Printf("%d bytes : %s\n", n, string(b1[:n]))

	o2, err := f.Seek(5, 0)
	checkError(err)
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	checkError(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2]))

	o3, err := f.Seek(2, 1)
	checkError(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)
	checkError(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3[:n3]))

	_, err = f.Seek(0, 0)
	checkError(err)

	newReader := bufio.NewReader(f)
	b4, err := newReader.Peek(5)
	checkError(err)
	fmt.Printf("bytes : %s\n", string(b4))

}

func ReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	checkError(err)
	fmt.Println(string(data))
	return data
}
