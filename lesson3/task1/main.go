package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

// Калькулятор

func main() {
	var op string
	f := new(big.Int)
	var d int64
	var a, b, c, res float64

	fmt.Print("Какую арифметическую операцию (+, -, *, /, √, ^, !) вы хотите выполнить? ")
	fmt.Scanln(&op)

	switch op {
	case "!":
		fmt.Print("Введите число: ")
		fmt.Scanln(&d)
		f.MulRange(1, d)
		fmt.Print("Результат выполнения операции: ", f)
	case "+":
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		res = a + b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case "-":
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		res = a - b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case "*":
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		if a == 0 {
			fmt.Println("На 0 умножать нельзя")
		} else if b == 0 {
			fmt.Println("На 0 умножать нельзя")
		}
		res = a * b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case "/":
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		if a == 0 {
			fmt.Println("На 0 делить нельзя")
		} else if b == 0 {
			fmt.Println("На 0 делить нельзя")
		}
		res = a / b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case "√":
		fmt.Print("Введите число: ")
		fmt.Scanln(&c)
		res = math.Sqrt(c)
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case "^":
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		res = math.Pow(a, b)
		fmt.Printf("Результат выполнения операции: %.2f\n", res)

	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

}
