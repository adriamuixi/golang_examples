package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	queues := []string{"rabbitMQ", "kafka", "redis", "SNS"}
	for _, queue := range queues {
		processQueue(queue)
	}
	fmt.Printf("done in %s\n", time.Since(started))
}

func processQueue(queue string) {
	fmt.Printf("processing message %s.....\n",queue)
	time.Sleep(2 * time.Second)
	fmt.Printf("done processing message %s\n", queue)
	fmt.Printf("")
}


