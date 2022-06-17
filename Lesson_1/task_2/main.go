package main

import (
	"fmt"
	"os"
	"time"
)

type simpleError struct {
	text string
	time time.Time
}

func (se *simpleError) Error() string {
	return fmt.Sprintf("%v Ошибка: %s\n", se.text, se.time)
}

func New(text string) *simpleError {
	return &simpleError{
		text: text,
		time: time.Now(),
	}

}

func ErrorWithPanic() (err error) {
	defer func() {
		if v := recover(); v != nil {
			v = New(fmt.Sprintf("Во время выполения функции произошла критическая ошибка: %s", v))
		}
	}()

	_, err = os.Open("sample.txt")
	if err != nil {
		return err
	}
	return
}

func main() {

	err := ErrorWithPanic()
	if err != nil {
		fmt.Println(err)
	}

}
