package main

import "fmt"

func ProcessMessage(c chan func() (int, string, bool)) {
	c <- func() (int, string, bool) {
		//some logic here
		return 0, "s", true
	}
}

func main() {
	c := make(chan func() (int, string, bool))
	go ProcessMessage(c)
	y, z, n := (<-c)()
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(n)
}
