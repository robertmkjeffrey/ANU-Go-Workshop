package main

import (
	"log"
	"sync"
	"time"
)

// If the values seem too orderly, try setting this to zero!
const DELAY time.Duration = 500 * time.Millisecond
const TIMEOUT time.Duration = 500 * time.Millisecond

func counter(target int, count *ProtectedInt, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		count.RLock()
		log.Printf("Counter [%d] reads %d.\n", target, count.n)
		current_value := count.n
		// time.Sleep(DELAY)
		count.RUnlock()

		if current_value < target {
			count.Lock()
			count.n += 1
			count.Unlock()
		} else if current_value > target {
			count.Lock()
			count.n -= 1
			count.Unlock()
		} else {
			log.Printf("Counter [%d] reads target and halts.\n", target)
			return
		}
	}
}

type ProtectedInt struct {
	sync.RWMutex
	n int
}

func main() {

	var count ProtectedInt
	var wg sync.WaitGroup

	// Helper function for creating counters.
	create_counter := func(target int) {
		wg.Add(1)
		// Passing a pointer to share the value.
		go counter(target, &count, &wg)
	}

	create_counter(30)
	create_counter(40)
	create_counter(50)
	create_counter(60)

	wg.Wait()

}
