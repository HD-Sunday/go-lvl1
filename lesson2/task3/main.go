package main

import (
	"fmt"
)

func main() {
	//3. С клавиатуры вводится трехзначное число.
	//Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
	//Например, для введенного 345 должно вывестись 3 сотни, 4 десятка, 5 единиц.
	//Нужно сделать это через математические операции.
	var a int

	fmt.Println("Введите 3х значное число")
	fmt.Scanln(&a)
	q := a / 100
	w := (a % 100) / 10
	e := a % 10

	fmt.Println("Сотни", +q, "Десятки", +w, "Еденицы", +e)

}
