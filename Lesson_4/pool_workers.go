package main

import (
	"fmt"
	"sync"
)

func main() {

	var waitGroup sync.WaitGroup
	var count int
	var workers = make(chan int, 100)

	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)

		go func(ch chan int) {
			defer waitGroup.Done()
			ch <- count
			count++
			fmt.Println(<-ch)
		}(workers)
	}
	waitGroup.Wait()
	fmt.Printf("final = %d\n", count)
}
