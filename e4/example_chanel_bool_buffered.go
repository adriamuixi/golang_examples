package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	queues := []string{"rabbitMQ", "kafka", "redis", "SNS"}
	results := make(chan bool)
	// results ;= make(chan bool,1)
	//results <- true //is blocking here -> we need an unbuffered channel -> "make(chan bool, 1)"
	for _, queue := range queues {
		queue := queue
		go func(f string) {
			processQueue(queue)
			results <- true
		}(queue)
	}

	fmt.Printf("done in %s\n", time.Since(started))
}

func processQueue(queue string) {
	fmt.Printf("processing message %s.....\n", queue)
	time.Sleep(2 * time.Second)
	fmt.Printf("done processing message %s\n", queue)
	fmt.Printf("")
}
