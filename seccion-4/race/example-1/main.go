package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				m.Lock()
				counter++
				m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value:", counter)
}
