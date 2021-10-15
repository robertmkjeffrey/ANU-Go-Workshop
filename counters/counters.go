package main

import (
	"log"
	"time"
)

// If the values seem too orderly, try setting this to zero!
const DELAY time.Duration = 500 * time.Millisecond
const TIMEOUT time.Duration = 500 * time.Millisecond

func counter(target int, update_chan chan<- int, count_viewer <-chan int) {
	for {
		current_value := <-count_viewer
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

func count_manager(update_chan <-chan int, result_chan chan<- int, count_view chan<- int) {
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
		case count_view <- count:
		}
	}
}

func main() {

	// Unbuffered channel.
	update_chan := make(chan int)

	results := make(chan int)
	count_viewer := make(chan int)
	go count_manager(update_chan, results, count_viewer)

	// Helper function for creating counters.
	create_counter := func(target int) {
		go counter(target, update_chan, count_viewer)
	}

	create_counter(30)
	create_counter(40)
	create_counter(50)
	create_counter(60)

	log.Printf("Final count is %d", <-results)
}
