//package maps
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

	mapf := MapFunctions()

	fmt.Println(mapf)
	delete(mapf, "suma")

	mapSlices := MapSlices()
	mapSlices[0] = make([]int, 5)
	mapSlices[1] = make([]int, 10)
	mapSlices[2] = make([]int, 3)

	fmt.Println(mapSlices)

}
func isOnMap(m map[string]func(x, y int) int, key string) bool {
	_, res := m[key]
	return res
}
func MapSlices() map[int][]int {
	return make(map[int][]int)
}

func MapFunctions() map[string]func(x, y int) int {

	mapF := map[string]func(x, y int) int{
		"suma":           func(x, y int) int { return x + y },
		"resta":          func(x, y int) int { return x - y },
		"multiplicacion": func(x, y int) int { return x * y },
		"division":       func(x, y int) int { return x / y },
		"modulo":         func(x, y int) int { return x % y },
	}
	return mapF
}
