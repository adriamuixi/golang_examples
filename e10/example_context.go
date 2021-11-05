package main

//https://www.youtube.com/watch?v=RH_lcxKMgN4

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	log.Println(ctx.Value("mykey"))

	ctx2 := context.WithValue(ctx, "mykey", "123")
	log.Println(ctx2.Value("mykey"))
}
