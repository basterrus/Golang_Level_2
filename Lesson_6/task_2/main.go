package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

const count = 1000

func main() {

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

	go log.Println("Working")
	for i := 0; i < 1e8; i++ {
		if i%1e8 == 0 {
			runtime.Gosched()
		}
	}
}
