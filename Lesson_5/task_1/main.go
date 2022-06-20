package main

import "sync"

func main() {
	task1(10000)
}
func task1(n int) {
	var wg = sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			wg.Done()
		}()
	}
	wg.Wait()
}

//1. Напишите программу, которая запускает 𝑛 потоков и дожидается завершения их всех
