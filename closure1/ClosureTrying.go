package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	file, err := os.Open("source")
	if err != nil {
		panic("Failed to open file: " + err.Error())
	}

	go read(file, done)

	time.Sleep(2 * time.Microsecond)
	fmt.Println("CLOSE***************************************************************************")
	errClose := file.Close()
	if (errClose != nil) {
		fmt.Println("Failed to close file: " + errClose.Error())
	}

	<-done
	fmt.Print("done main function")
}

func read(file *os.File, done chan bool)  {
	reader := bufio.NewReader(file)
	var buff = make([]byte, 2)
	var amountBytes int;
	var err error;
	for {
		time.Sleep(1 * time.Microsecond)
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
	//errClose := file.Close()
	//if (errClose != nil) {
	//	fmt.Println("Failed to close file: " + errClose.Error())
	//}
	fmt.Println("File successfully read")

	done <- true
	close(done)
	fmt.Println("done rutine")
}


