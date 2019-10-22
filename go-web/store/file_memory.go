package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hola Mundo!\n")
	//* write to file using WriteFile
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	//* read from file using ReadFile
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	//* write to file using File struct
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	//* read from file using File struct
	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
