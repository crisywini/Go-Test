package main

import (
	"fmt"
)

//Declarar un arreglo var identifier [len]type
func main() {
	numbers := NumbersOneToTen()
	fmt.Println(numbers)
	showElements(numbers)
}
func showElements(numbers [10]int) {
	for i := range numbers {
		fmt.Println(numbers[i])
	}
}
func NumbersOneToTen() [10]int {
	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}
	return numbers
}
