package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

const count = 1000

func main() {
	var (
		counter int
		lock    sync.Mutex
		wg      sync.WaitGroup
	)

	tr, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer func(tr *os.File) {
		err := tr.Close()
		if err != nil {
			panic(err)
		}
	}(tr)

	err = trace.Start(tr)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter += 1
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
