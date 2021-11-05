package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Task struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Title    string `json:"title"`
	Complete int    `json:"complete"`
}

func main() {
	var t Task
	wg := sync.WaitGroup{}
	wg.Add(40)
	sem := semaphore.NewWeighted(10)
	for i := 0; i < 40; i++ {
		fmt.Println(runtime.NumGoroutine())
		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Fatal(err)
		}
		go func(i int) {
			defer wg.Done()
			defer sem.Release(1)
			res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i))
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
				log.Fatal(err)
			}
			fmt.Println(t.Title)
		}(i)
	}
	wg.Wait()
}
