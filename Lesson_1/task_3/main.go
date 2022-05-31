package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	directoryPath := "C:\\Users\\pnedo\\GolandProjects\\Golang_Level_2\\Lesson_1\\folder\\"

	for i := 0; i < 1; i++ {
		_, err := os.Create("./folder/sample" + strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("created: " + "sample" + strconv.Itoa(i))
	}

	defer func() {
		if v := recover(); v != nil {
			v = fmt.Sprintf("Во время выполения функции произошла критическая ошибка: %s", v)
		}
		err := os.RemoveAll(directoryPath)
		if err != nil {
			fmt.Println(err)
		}
	}()
}
