package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Title    string `json:"title"`
	Complete int    `json:"complete"`
}

func main() {
	var t Task
	//wg := sync.WaitGroup{}
	//wg.Add(40)
	for i := 0; i < 10000; i++ {
		//1 - add value into routine to not repeat
		//fmt.Println(runtime.NumGoroutine()) -> number of routines
		//go func(i int) {
		//defer wg.Done()
		res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", i))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			log.Fatal(err)
		}
		fmt.Println(t.Title)
		//}(i)
	}
	//wg.Wait()
}
