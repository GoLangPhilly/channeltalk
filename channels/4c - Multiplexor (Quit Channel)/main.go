package main

import (
	"fmt"
	"sync"
)

func main() {
	q := make(chan struct{})
	m := merge(counter("odds", q), counter("ends", q))

	go func() {
		fmt.Scanln()
		close(q)
	}()

	for v := range m {
		fmt.Println(v)
	}
}

func counter(s string, q <-chan struct{}) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", s, i):
			case <-q:
				close(c)
				return
			}
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
