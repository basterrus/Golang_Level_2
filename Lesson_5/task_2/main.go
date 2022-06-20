package main

import (
	"fmt"
	"sync"
)

func funcForUnlockMutex(mutex *sync.Mutex) {
	defer mutex.Unlock()
}

func main() {
	var mutex sync.Mutex
	mutex.Lock()
	funcForUnlockMutex(&mutex)
	mutex.Lock()
	funcForUnlockMutex(&mutex)
	mutex.Lock()
	fmt.Println("Готово")
}

//2. Реализуйте функцию для разблокировки мьютекса с помощью defer
