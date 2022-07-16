package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
)

//1. С какими интерфейсами мы уже сталкивались в предыдущих уроках?
//Функции ввода/вывода из пакета fmt: Print(f),(ln), Scan(ln)

//2. Посмотрите примеры кода в своём портфолио.
//Везде ли ошибки обрабатываются грамотно?
//Хотите ли вы переписать какие-либо функции?

type Operator string

const (
	OpPlus      Operator = "+"
	OpMinus     Operator = "-"
	OpMultiply  Operator = "*"
	OpDivide    Operator = "/"
	OpRoot      Operator = "√"
	OpDegree    Operator = "^"
	OpFactorial Operator = "!"
)

func Calculate(a, b, c float64, f big.Int, d int64, op Operator) (res float64, err error) {
	switch op {
	case OpFactorial:
		f.MulRange(1, d)
		fmt.Print("Результат выполнения операции: ", f)
	case OpPlus:
		res = a + b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case OpMinus:
		res = a - b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case OpMultiply:
		if a == 0 {
			fmt.Println("На 0 умножать нельзя")
		} else if b == 0 {
			fmt.Println("На 0 умножать нельзя")
		}
		res = a * b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case OpDivide:
		if a == 0 {
			fmt.Println("На 0 делить нельзя")
			err = errors.New("cant divide by zero")
		} else if b == 0 {
			fmt.Println("На 0 делить нельзя")
			err = errors.New("cant divide by zero")
		}
		res = a / b
		fmt.Printf("Результат выполнения операции: %.2f\n", res)

	case OpRoot:
		res = math.Sqrt(c)
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
	case OpDegree:
		res = math.Pow(a, b)
		fmt.Printf("Результат выполнения операции: %.2f\n", res)

	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	return res, err
}

func main() {
	var op Operator
	/*	f := new(big.Int)*/
	var d int64
	var a, b, c, res float64

	fmt.Print("Какую арифметическую операцию (+, -, *, /, √, ^, !) вы хотите выполнить? ")
	fmt.Scanln(&op)

	if op == OpFactorial {
		fmt.Print("Введите число: ")
		fmt.Scanln(&d)
	} else if op == OpPlus || op == OpMinus || op == OpMultiply || op == OpDivide || op == OpDegree {
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	} else if op == OpRoot {
		fmt.Print("Введите число: ")
		fmt.Scanln(&c)
	}

	res, err := Calculate(a, b, c, big.Int{}, d, op)
	if err == nil {
		fmt.Println("Result ", res)
	} else {
		fmt.Println(err.Error())
	}
}
