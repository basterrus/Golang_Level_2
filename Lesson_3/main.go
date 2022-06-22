package main

import (
	"fmt"
	"github.com/basterrus/Golang_Level_2/Lesson_3/myPackage/myPackage"
	_ "github.com/basterrus/Golang_Level_2/Lesson_3/v3"
)

func main() {
	fmt.Println("HelloWorld")
	myPackage.SayHello()

}
