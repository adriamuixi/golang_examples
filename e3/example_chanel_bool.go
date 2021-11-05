package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	queues := []string{"rabbitMQ", "kafka", "redis", "SNS"}
	results := make(chan bool)
	//uncomment this part to show dead lock
	//<-results
	for _, queue := range queues {
		queue := queue
		go func(f string) {
			processQueue(queue)
			println(queue)
			results <- true
		}(queue)
	}

	// remove that part to understand block channels
	for i := 0; i < len(queues); i++ {
		time.Sleep(2 * time.Second)
		<-results
	}

	fmt.Printf("done in %s\n", time.Since(started))
}

func processQueue(queue string) {
	fmt.Printf("processing message %s.....\n", queue)
	time.Sleep(2 * time.Second)
	fmt.Printf("done processing message %s\n", queue)
	fmt.Printf("")
}
