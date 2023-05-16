package main

import (
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	var incremento = 0
	w.Add(2)

	go func() {
		for i := 0; i < 100; i++ {
			m.Lock()
			incremento++
			m.Unlock()
		}
		defer w.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			m.Lock()
			incremento++
			m.Unlock()
		}
		defer w.Done()
	}()

	w.Wait()

	fmt.Println("Final value:", incremento)
}
