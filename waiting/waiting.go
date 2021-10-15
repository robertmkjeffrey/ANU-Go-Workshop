package main

import (
	"log"
	"sync"
	"time"
)

func delay_print(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	log.Println(message)
}

func main() {
	// TODO: Use a WaitGroup to ensure the second message prints.
	// Make sure it's still concurrent!

	var wg sync.WaitGroup

	wg.Add(1)
	delay_print("Sequentially, the main function waits for us to finish...", &wg)

	wg.Add(1)
	go delay_print("Did you see me?", &wg)

	log.Println("...but it might not concurrently.")
	log.Println("Unlike tasks in Ada, goroutines halt when their parent goroutine returns.")
	log.Println("How can you ensure the concurrent print gets to finish?")

	wg.Wait()
}
