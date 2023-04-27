package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	respCh := make(chan any, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	userName := fetchUser()
	fetchUserLikes(userName, respCh, wg)
	fetchUserMatch(userName, respCh, wg)

	wg.Wait() //block until 0

	close(respCh)

	for resp := range respCh {
		fmt.Println("Respuesta:", resp)
	}

	fmt.Println("took:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func fetchUserLikes(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	respCh <- 11
	wg.Done()
}

func fetchUserMatch(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Microsecond * 100)
	respCh <- "ANA"
	wg.Done()
}
