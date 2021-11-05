package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
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
	errGroup, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 40; i++ {
		fmt.Println(runtime.NumGoroutine())
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Fatal(err)
		}
		errGroup.Go(func() error {
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
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		log.Fatal(err)
	}
}
