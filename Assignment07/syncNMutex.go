package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0 || n == 0
}

func main() {
	n := 0
	var m sync.Mutex

	go func(m *sync.Mutex) {
		m.Lock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, "is even")
			return
		}
		m.Unlock()
		fmt.Println(n, "is odd")
	}(&m)

	go func(m *sync.Mutex) {
		m.Lock()
		n++
		m.Unlock()
	}(&m)

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
