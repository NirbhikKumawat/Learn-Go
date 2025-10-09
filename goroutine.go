package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 1000 {
		fmt.Println(from, ":", i)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {
	go f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}
