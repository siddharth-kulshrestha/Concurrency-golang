package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func printsum(v1, v2 *value) {
	defer wg.Done()
	v1.mu.Lock()
	defer v1.mu.Unlock()

	time.Sleep(2 * time.Second)
	v2.mu.Lock()
	defer v2.mu.Unlock()

	fmt.Printf("Sum= %v\n, ", v1.value+v2.value)

}

func main() {

	var a, b value
	wg.Add(2)
	go printsum(&a, &b)
	go printsum(&b, &a)

	wg.Wait()
}
