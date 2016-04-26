package main

import (
	"fmt"
	"sync"
)

func main() {
	m := merge(counter("odds"), counter("ends"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-m)
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

func merge(chans ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	mergeReceiver := func(c <-chan string) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	wg.Add(len(chans))
	for _, c := range chans {
		go mergeReceiver(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
