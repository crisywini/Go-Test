package main3

import (
	"fmt"
)

func main() {
	var number int = 13
	var numberAddress *int = &number
	var copy int = *numberAddress
	fmt.Println(number, "<- the number, the location ->", numberAddress)
	fmt.Println(copy)
	var str string = "HOLE"
	var ptr *string = &str
	println(str, " ", ptr)
	*ptr = "GOAL"
	println(str, " ", ptr)
	println()
	println(number, " ", numberAddress)
	*numberAddress = 15
	println(number, " ", numberAddress)
	pr2 := &ptr
	aux := "VAR"
	fmt.Println(pr2, " ", ptr, " valor ", str, " last ", aux)
	*pr2 = &aux
	fmt.Println(pr2, " ", ptr, " valor ", str, " new ", aux)
	fmt.Println(*ptr)

}
