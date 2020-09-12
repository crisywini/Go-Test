package statements

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	number, err := MySqrt(-16.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(number)
	Greeting("Hola familia ", "Papi")
	DatabaseOperations()

}
func Greeting(message string, people ...string) {

	if len(people) == 1 {
		fmt.Println(message, " ", people[0])
	} else {
		for _, val := range people {
			fmt.Println(message, val)
		}
	}

}
func CloseDatabaseConnection() {
	fmt.Println("Closing connections")
}
func DatabaseOperations() {
	defer CloseDatabaseConnection()
	fmt.Println("Connections to Database")
	fmt.Println("Connected to Database succesfull")
	fmt.Println("Using Database ")
	return
	//Defer is executed before the return statement
}

func isMajor(a int, b int) bool {
LABEL1:
	switch {
	case a > b:
		return true
	case a == b:
		a++
		goto LABEL1
	}

	return false
}
func MySqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, errors.New("Negative number, this should not be")
	}
	var half float64 = 1.0 / 2.0
	return math.Pow(number, half), nil
}

func triangle(n int, pointer *string) {
	aux := n

	for i := 0; i < n; i++ {
		for j := 0; j < aux; j++ {
			*pointer += "G"
		}
		*pointer += "\n"
		aux--
	}

}

func figure(n int) string {

	var figure string
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			figure += "G"
		}
		figure += "\n"
	}
	return figure
}
