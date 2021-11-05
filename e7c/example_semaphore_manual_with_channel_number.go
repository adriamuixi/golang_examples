package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Title    string `json:"title"`
	Complete int    `json:"complete"`
}

func main() {
	wg := sync.WaitGroup{}
	now := time.Now()
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		//fmt.Println("Number of concurrent go routines", runtime.NumGoroutine())
		go func(i int) {
			defer wg.Done()
			//fmt.Println("iteration: ", i)
		}(i)
	}
	wg.Wait()
	end := now.Sub(now)
	fmt.Println("finish in:", end.String())
}
