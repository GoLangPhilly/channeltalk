package main

import (
	"fmt"
)

func main() {
	s1 := counter("Eric")
	s2 := counter("Chris")

	var s1Complete, s2Complete bool
	for !s1Complete || !s2Complete {
		select {
		case v := <-s1:
			fmt.Println(v)
			s1Complete = true
			s1 = nil
		case v := <-s2:
			fmt.Println(v)
			s2Complete = true
			s2 = nil
		}
	}
}

func counter(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
		}
	}()
	return c
}
