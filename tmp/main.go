package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	total := 10
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go printout(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
	fmt.Println(i)
	wg.Done()
}

// https://golang.org/pkg/sync/atomic/#pkg-index
