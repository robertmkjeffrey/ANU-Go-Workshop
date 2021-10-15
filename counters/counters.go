package main

import (
	"log"
	"time"
)

// If the values seem too orderly, try setting this to zero!
const DELAY time.Duration = 500 * time.Millisecond
const TIMEOUT time.Duration = 500 * time.Millisecond

func counter(target int, update_chan chan<- int) {
	current_value := 50
	for {
		if current_value < target {
			update_chan <- 1
		} else if current_value > target {
			update_chan <- -1
		} else {
			log.Printf("Counter [%d] reads target and halts.\n", target)
			return
		}
	}
}

func count_manager(update_chan <-chan int, result_chan chan<- int) {
	var count int = 50

	for {
		select {
		case update := <-update_chan:
			count = count + update
			log.Printf("Manager sets value to %d.\n", count)
			time.Sleep(DELAY)
		// time.After returns a channel that is sent to once the delay has passed.
		case <-time.After(TIMEOUT):
			result_chan <- count
			return
		}
	}
}

func main() {

	// Unbuffered channel.
	update_chan := make(chan int)

	results := make(chan int)
	go count_manager(update_chan, results)

	// Helper function for creating counters.
	create_counter := func(target int) {
		go counter(target, update_chan)
	}

	create_counter(30)
	create_counter(40)
	create_counter(50)
	create_counter(60)

	log.Printf("Final count is %d", <-results)
}
