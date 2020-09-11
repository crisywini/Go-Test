package main

import (
	"fmt"
	//Ojo con eso manito

	as "./counterS"
	"./piTest"
)

const name2 = "HOLA"

const (
	MALE = iota //Enumerations
	FEMALE
	SHEMALE
)
const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type (
	i   int
	s   string
	f32 float32
	b   bool
)

var twoPi = 2 * piTest.Pi
var a = "G"
var copy = as.StrTest

func main() {
	/**
	var (
		name3   s
		name4   s
		HOME      = os.Getenv("HOME")
		USER      = os.Getenv("USER")
		GOROOT    = os.Getenv("GOROOT")
		path      = os.Getenv("PATH")
		name    s = "hola"
		is_name b = false
		age     i
	)
	var height f32
	fmt.Print(name + "\n")
	fmt.Println(is_name)
	if name == "hola" {
		age = 12
		height = f32(age) * 0.5
		name3 = "TEST"
		name4 = "s(age + 1)"
	} else {
		age = 13
	}
	height, name = 45, "hola"
	name3, name = name, name3 //swaping variables
	fmt.Println(name, name3, "Prueba que estoy ejecutando")

	fmt.Println(age, MALE, HOME, USER, GOROOT, path)
	fmt.Println(height)
	fmt.Println(SHEMALE)
	fmt.Println(name4)
	fmt.Println(name3)
	**/
	//fmt.Printf("2*pi = %g\n", twoPi)
	//n()
	//m()
	//n()
	/**a = "G"
	fmt.Println(a)
	f1()*/
	/*
		fmt.Println(KB)
		fmt.Println(MB)
		fmt.Println(YB)
		a := 0.0
		b := KB / a
		print(b)

		for i := 0; i < 10; i++ {
			a := rand.Int()
			fmt.Printf("%d / ", a)
		}
	*/
	/**
	var str string
	var str2 string
	str = "Alfonzo"
	str2 = "Augusto"
	if str > str2 {
		print("Augusto es mayor que alfonzo")
	}
	if str2 > str {
		println("Alfonzo es mayor que Augusto")
		println(len(str) + len(str2))
	}
	if str == str2 {
		print("Alfonzo y Augusto son iguales")
	}
	*/

}

func n() {
	fmt.Println(a)
}

func m() {
	a := "O" // METHOD SCOPE
	a = "O"  //GLOBAL SCOPE
	fmt.Println(a)
}

func f1() {
	a := "O"
	fmt.Println(a)
	f2()
}
func f2() {
	print(a)
}

func factorial_recursive(number int) int {
	if number == 0 {
		return 1
	} else {
		return number * factorial_recursive(number-1)
	}
}
