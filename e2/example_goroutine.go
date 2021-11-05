package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	started := time.Now()
	queues := []string{"rabbitMQ", "kafka", "redis", "SNS"}
	var wg sync.WaitGroup
	wg.Add(len(queues))
	for _, queue := range queues {
		//queue := queue
		go func(f string) {
			processQueue(queue)
			wg.Done()
		}(queue)
	}
	//Uncomment to No wait
	wg.Wait()
	fmt.Printf("done in %s\n", time.Since(started))
}

func processQueue(queue string) {
	fmt.Printf("processing message %s.....\n", queue)
	time.Sleep(2 * time.Second)
	fmt.Printf("done processing message %s\n", queue)
	fmt.Printf("")
}
