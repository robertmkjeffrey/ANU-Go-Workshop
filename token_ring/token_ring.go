package main

import (
	"fmt"
	"log"
	"sync"
)

const TOTAL_NODES int = 5

type token int

func sendToken(t token, out chan<- token) {
	// TODO: Do something with this.
	go func() {
		out <- t
	}()
}

func token_node(id int, in <-chan token, out chan<- token, wg *sync.WaitGroup) {

	defer wg.Done()

	var currentLeader token = token(id)

	// TODO - send your token (currentLeader) to initiate an election bid.
	sendToken(currentLeader, out)

	for {
		recieved_token := <-in

		if recieved_token > currentLeader {
			currentLeader = recieved_token
			// TODO - forward token to share election bid.
			sendToken(recieved_token, out)
		} else if recieved_token == token(id) {
			log.Printf("Node [%d] claims victory!\n", id)
			// TODO - send token to claim victory.
			sendToken(recieved_token, out)
			// Recieve your own token and then shut down.
			<-in
			return
		} else if recieved_token == currentLeader {
			log.Printf("Node [%d] acknowledges node [%d] as leader.", id, currentLeader)
			// TODO - send token to share victory message.
			sendToken(recieved_token, out)
			return
		}

	}

}

func main() {

	var wg sync.WaitGroup

	// Create a slice (list) of token channels with inital length TOTAL_NODES.
	channels := make([](chan token), TOTAL_NODES)

	// Initialise these channels.
	for i := range channels {
		// No arguments mean these channels will be unbuffered.
		channels[i] = make(chan token)
	}

	for id := range channels {
		var last_node (<-chan token) // Reciever-only channel
		var next_node (chan<- token) // Sender-only channel

		// Ring topology
		// This would be much easier if we had modulo types!
		if id == len(channels)-1 {
			last_node = channels[id]
			next_node = channels[0]
		} else {
			last_node = channels[id]
			next_node = channels[id+1]
		}

		// TODO: create each node here.
		wg.Add(1)
		go token_node(id, last_node, next_node, &wg)
	}

	// TODO: This currently blocks infinitely.
	// Find a way to wait until the token election is complete and then halt.
	wg.Wait()
	fmt.Println("Shutdown.")
}
