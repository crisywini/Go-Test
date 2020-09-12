package main

import "fmt"

func main() {
	fmt.Println(isMajor(1, 1))

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
