package main

import "fmt"

//Задание для продвинутых (необязательное). Написать приложение,
//которое ищет все простые числа от 0 до N включительно.
//Число N должно быть задано из стандартного потока ввода.

func main() {
	var number, a, q int
	var f bool

	fmt.Print("Введите число: ")

	fmt.Scanln(&number)
	for a = 2; a < number; a++ {
		f = false
		for q = 2; q <= a/2; q++ {

			if a%q == 0 {
				f = true
			}
		}
		if f == false {
			fmt.Println("Простые числа: ", a)
		}
	}

}
