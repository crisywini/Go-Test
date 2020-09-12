package main3

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	month := t.Month()
	day := t.Day()
	year := t.Year()
	fmt.Printf("The day is %d the month %d the year %d\n", day, month, year)
	fmt.Println(t)

	fmt.Println(week)
}
