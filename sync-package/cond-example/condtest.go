package main

import (
	"fmt"
	"sync"
	"time"
)

/*

Lets assume we have a requirement where one goroutine is getting data from external sources
and other goroutines (multiple goroutines) want this data to be available to execute further.

We cannot use channels because then only one goroutine would be able to read the data from channel

We will use sync Cond{} to Broadcast this to other goroutines.

Idea is that we will share map{} as a variable global to all goroutines.

Goroutines which are dependent on data will be using cond.Wait() to wait till data gets available in the map

And the goroutine which is fetching data will fetch the data and use cond.Broadcast() method allowing other goroutines to release Wait.
*/

func fetchAndBroadcast(data map[string]string, cond *sync.Cond) {

	time.Sleep(time.Second)

	cond.L.Lock()
	defer cond.L.Unlock()

	data["ABC"] = "Inserted"

	cond.Broadcast()

}

func consumeData(batchID string, data map[string]string, cond *sync.Cond, wg *sync.WaitGroup) {

	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Wait()
	fmt.Print("goroutine number: ", batchID)
	fmt.Println(data)
	wg.Done()
}

func main() {

	wg := sync.WaitGroup{}

	m := sync.Mutex{}
	cond := sync.NewCond(&m)

	data := make(map[string]string, 1)

	wg.Add(2)
	// consumer 1
	go consumeData("id001", data, cond, &wg)

	// consumer 2
	go consumeData("id002", data, cond, &wg)

	go fetchAndBroadcast(data, cond)

	wg.Wait()

}
