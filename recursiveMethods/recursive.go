package recursiveMethods

import (
	"fmt"
	"math"
)

func main() {
	anonimus := func(x, y int) int { return x + y }
	number := anonimus(1, 1)
	fmt.Println(number)
	number2 := func(x, y float64) float64 { return math.Pow(x, y) }(2, 3)
	fmt.Println(number2)
	fv := func() { fmt.Println("Hello functional World") }
	fmt.Println(fv)     //It gives the address
	fmt.Println(test()) //It return 2 because yes, well actually because de defer is executed after the return value so 1+1 = 2
	fun := returningLambdas()
	another := fun(3)
	fmt.Println(another, " function inside of a function")
	for i := 0; i < 10; i++ {
		fmt.Println(i, " ", fibonacci(i))
	}
	fmt.Println("Now with lambdas")
	fibo := fibonacciLambda()
	for i := 0; i < 10; i++ {
		fmt.Println(i, " ", fibo(i))
	}

}

func fibonacciLambda() func(number int) int {
	return func(number int) int {
		res := 1
		if number > 1 {
			res = 0
			a := 1
			b := 1
			for i := 2; i <= number; i++ {
				res = a + b
				a = b
				b = res
			}
		}
		return res
	}
}

func test() (ret int) {
	defer func() { ret++ }()
	return 1
}
func returningLambdas() func(b int) int {
	return func(b int) int { return b * b }
}

func Fibonacci(i int) (int, int) {
	if i <= 1 {
		return i, 1
	} else {
		number, _ := Fibonacci(i - 2)
		number2, _ := Fibonacci(i - 1)
		return number + number2, i
	}
}

func fibonacci(i int) (res int) {
	if i <= 1 {
		res = 1
	} else {
		res = fibonacci(i-2) + fibonacci(i-1)
	}
	return
}

func Factorial(number int64) (res int64) {
	if number == 0 {
		res = 1
	} else {
		res = number * Factorial(number-1)
	}
	return
}
