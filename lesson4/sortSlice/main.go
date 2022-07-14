package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Познакомьтесь с алгоритмом сортировки вставками.
//Напишите приложение, которое принимает на вход набор целых чисел и отдаёт на выходе его же в отсортированном виде.
//Сохраните код, он понадобится нам в дальнейших уроках.

func main() {

	slice := generateSlice(6)
	fmt.Println("\n--- Не отсортирован --- \n", slice)
	sort(slice)
	fmt.Println("\n--- Отсортирован ---\n", slice)
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func sort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
