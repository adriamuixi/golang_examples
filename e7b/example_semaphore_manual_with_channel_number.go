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
	sem := make(chan int, 10)
	for i := 0; i < 40; i++ {
		fmt.Println("Number of concurrent go routines", runtime.NumGoroutine())
		sem <- i
		go func(i int) {
			defer func() {
				num := <-sem
				fmt.Println("iteration number: ", num)
			}()

			defer wg.Done()

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
