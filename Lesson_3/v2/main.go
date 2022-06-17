package v2

import (
	"fmt"
	_ "github.com/basterrus/Golang_Level_2/Lesson_3/myPackage"
	_ "github.com/basterrus/Golang_Level_2/Lesson_3/v2"
)

func main() {
	//myPackage.SayHello().hello()
	fmt.Println("hello v2.0.0")
}
