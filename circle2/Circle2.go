package main

import (
	"fmt"
	"time"
)

func first(done chan bool, fullDone chan bool)  {
	defer func() {
		fullDone <- true
		fmt.Println("first close")
	}()

	for {
		select {
		case <- done:
			return
		default :
			fmt.Println("first")
			time.Sleep(time.Second)
		}
	}
}

func second(done chan bool, fullDone chan bool)  {
	defer func() {
		fullDone <- true
		fmt.Println("second close")
	}()

	var i int
	for {
		select {
		case <- done:
			return
		default :
			fmt.Println("second")
			time.Sleep(time.Second)
			if i == 3 {
				done <- true
				close(done)
				return
			}
			i++
		}
	}
}


func main() {
	done := make(chan bool, 1)
	fullDone := make(chan bool, 2)

	go first(done, fullDone)
	go second(done, fullDone)


	<- fullDone
	<- fullDone
	close(fullDone)
	fmt.Println("main complete")
}

