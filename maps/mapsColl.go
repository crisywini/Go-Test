package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println()
	var table map[int]string = make(map[int]string)

	table[0] = "hola"
	for i := 1; i <= 10; i++ {
		table[i] = strconv.Itoa(i) + " Hola"
	}
	fmt.Println(table)
}
