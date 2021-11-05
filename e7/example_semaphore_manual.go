package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
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
	sem := make(chan bool, 10)
	for i := 0; i < 40; i++ {
		fmt.Println(runtime.NumGoroutine())
		sem <- true
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
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
