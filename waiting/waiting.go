package main

import (
<<<<<<< HEAD
	"log"
=======
	"fmt"
	"sync"
>>>>>>> Added solution for waiting
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

<<<<<<< HEAD
	delay_print("Sequentially, the main function waits for us to finish...")
	go delay_print("Did you see me?")
	log.Println("...but it might not concurrently.")
	log.Println("Unlike tasks in Ada, goroutines halt when their parent goroutine returns.")
	log.Println("How can you ensure the concurrent print gets to finish?")
=======
	var print_wg sync.WaitGroup

	print_wg.Add(1)
	delay_print("Sequentially, the main function waits for us to finish...", &print_wg)

	print_wg.Add(1)
	go delay_print("Did you see me?", &print_wg)

	fmt.Println("...but it might not concurrently.")
	fmt.Println("Unlike tasks in Ada, goroutines halt when their parent goroutine returns.")
	fmt.Println("How can you ensure the concurrent print gets to finish?")

	print_wg.Wait()
>>>>>>> Added solution for waiting
}
