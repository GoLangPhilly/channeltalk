package main

import (
	"fmt"
	"time"
)

func main() {
	c := longTask("hello go")
	fmt.Println(<-c)

	v, ok := <-c

	fmt.Printf("Value: '%v', Still Open?: %t", v, ok)
}

func longTask(s string) <-chan string {
	c := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		c <- s
		close(c)
	}()

	return c
}
