package main

import (
	"fmt"
)

func division(a, b int) {
	result := a / b
	fmt.Println(result)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Программа восстановлена", err)
		}
	}()

	division(10, 0)
}
