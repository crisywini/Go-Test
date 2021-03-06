package collections

//package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	WIDTH  = 10
	HEIGHT = 10
)

var Fibo [50]int
var matrix [WIDTH][HEIGHT]int

//Declarar un arreglo var identifier [len]type
func main() {
	/*
			numbers := NumbersOneToTen()
			fmt.Println(numbers)
			//showElements(numbers)
			test := [...]string{"hola", "cómo", "estas"}
			fmt.Println(test)
			ar := [3]int{1, 2, 3}
			fmt.Println(ar)
			modifyValues(&ar)
			fmt.Println(ar)
			arr2 := [3]int{3, 2, 1}
			ar = arr2
			ar[1] = 10
			fmt.Println(ar)
			fmt.Println(arr2)
			fillFibo()
		for i := 0; i < 50; i++ {
			fmt.Println(i, " value: ", Fibonacci(i))
		}
	*/
	b := []byte{'g', 'o', 'l', 'a', 'n', 'd'}

	fmt.Println(b[:])
	c := Slice(12, 100)
	fmt.Println(c)
	s := []int{3, 4, 6, 8, -1, 654, 6, 9}
	fmt.Println(s)
	sort(s)
	fmt.Println(s)
	fmt.Println(search(s, -2, 0, 0, len(s)-1))
	f := func(x int) int { return int(math.Pow(float64(x), 2)) }
	s2 := MapFunction(f, 1, 2, 3, 4, 5)
	fmt.Println(s2)

}
func sort(slice []int) {

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
func search(slice []int, element int, i int, l, r int) int {
	if r >= l {
		i = (l + r) / 2

		if slice[i] == element {
			return i
		}
		if slice[i] > element {
			return search(slice, element, i, l, i-1)
		}
		if slice[i] < element {
			return search(slice, element, i, i+1, r)
		}
	}
	return -1
}
func test(x ...int) {
	fmt.Println(len(x))
}

func MapFunction(f func(e int) int, x ...int) []int {
	var res []int = make([]int, len(x))
	for i, val := range x {
		res[i] = f(val)
	}
	return res
}

func Slice(length int, capacity int) []string {
	return make([]string, length, capacity)
}

//func SliceInt(length int, capacity int) []int {
//	return make([capacity]int)[0:length]
//}

func fillFibo() {
	Fibo[0] = 1
	Fibo[1] = 1
	for i := 2; i < len(Fibo); i++ {
		Fibo[i] = Fibo[i-2] + Fibo[i-1]
	}
}
func Fibonacci(n int) (number int) {

	number = Fibo[n]
	return
}

func showMatrix() {
	info := ""
	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			info += strconv.Itoa(matrix[i][j]) + " "
		}
		info += "\n"
	}
	fmt.Println(info)
}

func InitMatrix() {

	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			if i == j || (i+j) == WIDTH-1 || i == 0 || j == 0 || i == WIDTH-1 || j == HEIGHT-1 {
				matrix[i][j] = 1
			}
		}
	}
}

/*
To modify values has to use the pointer because in the functions Go create a copy of the array
*/
func modifyValues(a *[3]int) {
	for i := range a {
		a[i] = i * 2
	}
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
