package main

//https://www.youtube.com/watch?v=RH_lcxKMgN4

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	_ = context.WithValue(ctx, "mykey", 123)

	ctx2, cancelFunc := context.WithCancel(ctx)
	log.Println(ctx2.Err())
	cancelFunc()
	log.Println(ctx2.Err() == context.Canceled)
}
