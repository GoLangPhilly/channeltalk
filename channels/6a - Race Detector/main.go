package main

import (
	"fmt"
	"sync"
)

func main() {
	service = make(map[string]string)
	w := sync.WaitGroup{}

	w.Add(1)
	go func() {
		RegisterService("Service1", "http://golang.org")
		RegisterService("Service2", "http://apple.com")
		w.Done()
	}()

	w.Add(1)
	go func() {
		fmt.Println(LookupService("Service1"))
		w.Done()
	}()

	w.Add(1)
	go func() {
		fmt.Println(LookupService("Service2"))
		w.Done()
	}()

	w.Wait()
}

var service map[string]string

func RegisterService(name string, addr string) {
	service[name] = addr
}

func LookupService(name string) string {
	return service[name]
}
