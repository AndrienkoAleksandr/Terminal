package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
)

var (
	target string = "target"
	source string = "source"
)

func main()  {
	source, err := os.Open(source);
	if err != nil {
		panic("Failed to open source file: " + err.Error())
	}

	targetFile := getTargetFile()

	defer func() {
		fmt.Println("\nclosing files ")
		if errClose := source.Close(); errClose != nil {
			fmt.Println("Failed to close file: " + errClose.Error())
		}
		if errClose := targetFile.Close(); errClose != nil {
			fmt.Println("Failed to close file: " + errClose.Error())
		}
	}()

	done := make(chan bool)

	go write(source, targetFile, done)
	go read(source, done)

	<- done
}

func getTargetFile() *os.File {
	var targetFile *os.File
	var errTargetFile error
	if !Exists(target) {
		targetFile, errTargetFile = os.Create(target)
		if errTargetFile != nil {
			fmt.Println("failed to create file ", errTargetFile.Error())
		}
	} else {
		targetFile, errTargetFile = os.OpenFile(target, os.O_APPEND|os.O_WRONLY, os.ModeAppend);
		if  errTargetFile != nil {
			panic("Failed to open file: " + errTargetFile.Error())
		}
	}
	return targetFile
}

func read(source *os.File, done chan bool) {
	defer func() {
		fmt.Print("done read")
		done <- true
	}()

	reader := bufio.NewReader(source)
	var buff = make([]byte, 2)
	for {
		readBytes, err := reader.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if (readBytes == 0) {
			break
		}
		//fmt.Print(string(buff))
	}
}

func write(source *os.File, target *os.File, done chan bool) {
	defer func() {
		fmt.Println("done write")
		//done <- true
	}()

	fmt.Print("go")
	reader := bufio.NewReader(source)
	var buff = make([]byte, 2)
	for {
		//fmt.Println("go go")
		readBytes, err := reader.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if (readBytes == 0) {
			continue
		}
		if _, err := target.Write(buff); err != nil {
			panic(err)
		}
		fmt.Print(string(buff))
	}
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
