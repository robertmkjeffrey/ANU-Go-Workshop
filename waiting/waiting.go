package main

import (
	"log"
	"time"
)

func delay_print(message string) {
	time.Sleep(1 * time.Second)
	log.Println(message)
}

func main() {
	// TODO: Use a WaitGroup to ensure the second message prints.
	// Make sure it's still concurrent!

	delay_print("Sequentially, the main function waits for us to finish...")
	go delay_print("Did you see me?")
	log.Println("...but it might not concurrently.")
	log.Println("Unlike tasks in Ada, goroutines halt when their parent goroutine returns.")
	log.Println("How can you ensure the concurrent print gets to finish?")
}
