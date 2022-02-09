package examplejson

import (
	"bufio"
	"fmt"
	"os"
)

func InitExampleFileWrite() {

	data := []byte("hello golang!\n")
	// data := ReadFile("./example_json/person_data.json")
	err := os.WriteFile("./example_json/text1", data, 0644)
	checkError(err)

	f, err := os.Create("./example_json/text2")
	checkError(err)
	defer f.Close()

	// data2 := []byte{11, 12, 13, 45, 64, 77, 55, 97, 111, 109}
	data2 := "hello there\n"
	n, err := f.Write([]byte(data2))
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n)

	n2, err := f.WriteString("string here\n")
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n2)
	f.Sync()

	w := bufio.NewWriter(f)
	n3, err := w.WriteString("writing using bufio\n")
	checkError(err)
	fmt.Printf("wrote using bufio with %d bytes", n3)
	w.Flush()
}
