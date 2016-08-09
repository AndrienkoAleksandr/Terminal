package main

import (
"fmt"
"errors"
"time"
)

func run1() error {
	time.Sleep(time.Millisecond * 2000)

	fmt.Println("run1")

	return errors.New("some1")
}

func run2() error {
	time.Sleep(time.Millisecond * 1000)

	fmt.Println("run2")

	return errors.New("some2")
}

func main()  {
	fmt.Println("")
	fmt.Println("begin")

	er := make(chan error)

	go run1()
	go run2()

	<- er

}