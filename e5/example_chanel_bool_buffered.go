package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	queues := []string{"rabbitMQ", "kafka", "redis", "SNS"}
	results := make(chan bool)
	for _, queue := range queues {
		queue := queue
		go func(f string) {
			close(results)
			processQueue(queue)
			results <- true
		}(queue)
	}

	for i := 0; i < len(queues); i++ {
		//sent on the channel (true) or is a zero value returned because the channel is closed and empty (false)
		value, ok := <-results
		fmt.Println("is the channel open?", ok)
		fmt.Println("first result of channel", value)
	}

	fmt.Printf("done in %s\n", time.Since(started))
}

func processQueue(queue string) {
	fmt.Printf("processing message %s.....\n", queue)
	time.Sleep(2 * time.Second)
	fmt.Printf("done processing message %s\n", queue)
	fmt.Printf("")
}
