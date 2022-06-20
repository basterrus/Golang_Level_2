// Package errors implements functions to demonstrate godoc features
//
// The task_3 function
//
package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// function task_3 for calculate result

	var numOne, numTwo, result float32
	var operation string

	// Заправшиваем действие у пользователя
	fmt.Printf("Выберете нужное действие (+, -, *, /, ^, !, help, exit) " +
		"\nДля вывода подробной справки введите 'help'" +
		"\nДля выхода введите 'exit': \n")
	fmt.Scan(&operation)

	// Проверяем на команду "exit", если сделать в case, то до выхода просит ввести 2 значения, не нужно!
	if operation == "exit" {
		fmt.Println("Программа завершена")
		os.Exit(1)
	} else {
		fmt.Println("Введите первое число: ")
		fmt.Scan(&numOne)

		fmt.Println("Введите второе число: ")
		fmt.Scan(&numTwo)

		// Проверка на совпадение выбранного действия и выдача результата
		switch operation {
		case "help":
			fmt.Println("'+' операция сложения")
			fmt.Println("'-' операция вычитания")
			fmt.Println("'*' произведение 2х чисел")
			fmt.Println("'/' операция деления")
			fmt.Println("'^' возведение первого числа в степень второго числа")
			fmt.Println("'!' Высчитывает факториал каждого из введеных чисел")
		case "+":
			result = numOne + numTwo
		case "-":
			result = numOne - numTwo
		case "*":
			result = numOne * numTwo
		case "/":
			//Проверка деления на 0
			if numTwo == 0 {
				fmt.Println("ZeroDivisionError: На ноль делить нельзя! Программа закрыта!")
				os.Exit(1)
			} else {
				result = numOne / numTwo
			}
		case "^":
			result = float32(math.Pow(float64(numOne), float64(numTwo)))

		case "!":
			result = factorial(int(numOne))
			fmt.Printf("Результат выполнения: %.2f\n", result)
			result = factorial(int(numTwo))
			fmt.Printf("Результат выполнения: %.2f\n", result)
		default:
			fmt.Println("Вы выбрали не корректную операцию!")
			os.Exit(1)
		}

	}
	//Проверка для операции факториала, что бы не игнорировать второе введеное значение
	// выводим оба результата через printf в case и завершаем программу
	// иначе получучается дублирование результата, для красоты
	if operation == "!" {
		os.Exit(1)
	} else {
		//Форматированная строка результат выдается с точностью до 2х знаков после точки "10.00"
		fmt.Printf("Результат выполнения: %.2f\n", result)
	}
}

// Функция подсчета факториала
func factorial(n int) float32 {
	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return float32(result)
}
