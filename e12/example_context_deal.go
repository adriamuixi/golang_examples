package main

//https://www.youtube.com/watch?v=RH_lcxKMgN4

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	_ = context.WithValue(ctx, "mykey", 123)

	t := time.Now().Add(1 * time.Second)
	ctx2, _ := context.WithDeadline(ctx, t)
	log.Print(ctx2.Err())
	time.Sleep(2 * time.Second)
	log.Println(ctx2.Err() == context.DeadlineExceeded)
}
