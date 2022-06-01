package main

import (
	"fmt"
	"math"
)

//2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
//Площадь круга должна вводиться пользователем с клавиатуры.

func main() {
	var S float64 // Площадь
	var C float64 // Длина
	var D float64 // Диаметр
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&S)

	//Длина окружности
	C = 2 * math.Pi * math.Sqrt(S/math.Pi)
	fmt.Println("Длина окружности: ", C)

	//Диаметр окружности
	D = 2 * math.Sqrt(S/math.Pi)
	fmt.Println("Диаметр окружности: ", D)
}
