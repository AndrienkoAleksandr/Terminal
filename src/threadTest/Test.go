package main

import (
	"fmt"
	"errors"
	"time"
)

func run1(errs chan error) {
	time.Sleep(time.Millisecond * 2000)

	fmt.Println("run1")

	errs <- errors.New("some1")
}

func run2(errs chan error) {
	time.Sleep(time.Millisecond * 1000)

	fmt.Println("run2")

	errs <- errors.New("some2")
}

func main()  {
	fmt.Println("")
	fmt.Println("begin")

	er := make(chan error)

	go run1(er)
	go run2(er)

	<- er

}
