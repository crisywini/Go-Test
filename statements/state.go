package main

func main() {
	var f int = 1
	switch {
	case f < 0:
		print("is 0")
	case f > 0:
		print("is 1")
	case f > 0 && f < 100:
		print("Is 2")

	}

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
