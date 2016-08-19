package main

import (
	"fmt"
	"time"
	"sync"
)

func first(done chan bool, waitGr *sync.WaitGroup)  {
	defer func() {
		fmt.Println("first close")
		waitGr.Done()
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

func second(done chan bool, waitGr *sync.WaitGroup)  {
	defer func() {
		fmt.Println("second close")
		waitGr.Done()
	}()

	var i int
	for {
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

func main() {
	done := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)

	go first(done, &wg)
	go second(done, &wg)

	wg.Wait()
	fmt.Println("main complete")
}
