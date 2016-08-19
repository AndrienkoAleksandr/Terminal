package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func main() {
	file, err := os.Open("source")
	if err != nil {
		panic("Failed to open file: " + err.Error())
	}
	reader := bufio.NewReader(file)
	var buff = make([]byte, 8192)
	var amountBytes int;
	for {
		amountBytes, err = reader.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
			return
		}
		if (amountBytes == 0) {
			fmt.Println()
			break
		}
		fmt.Print(string(buff))
	}
	errClose := file.Close()
	if (errClose != nil) {
		fmt.Println("Failed to close file: " + errClose.Error())
	}
	fmt.Println("File successfully read")
}
