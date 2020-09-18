package bytes

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

func main() {

	//fmt.Print("hola")
	//var buf *bytes.Buffer = NewBuffer()
	//fmt.Println(buf, " buf")
	arr := [4]float64{1, 2, 3, 4}
	slice := make([]float64, 10, 20)
	slice[0] = -20
	for i := len(slice) - 1; i >= 0; i-- {
		slice[i] = float64(i)
	}
	fmt.Println(SumNumbers(arr))
	fmt.Println(SumNumberSlice(slice))
	fmt.Println(SumAndAverage(slice))
	fmt.Println(MaxAndMin(slice))
	slice1 := StrToByte("Alejandro")
	slice2 := StrToByte("Cristian")
	fmt.Println(Compare(slice2, slice1))
	sort.Float64s(slice)
	fmt.Println(slice)
	fmt.Println(ReverseString("Google"))
}

func ReverseString(str string) (res string) {

	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return
}
func StrToByte(a string) []byte {
	slice := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		slice[i] = a[i]
	}
	return slice
}
func NewBuffer() *bytes.Buffer {
	return new(bytes.Buffer)
}
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}

func SumNumbers(arr [4]float64) float64 {

	var res float64
	for _, value := range arr {
		res += value
	}
	return res
}
func SumNumberSlice(arr []float64) float64 {
	var res float64
	for _, value := range arr {
		res += value
	}
	return res
}
func SumAndAverage(arr []float64) (float64, float64) {
	var res float64
	var counter int
	for _, value := range arr {
		res += value
		counter++
	}

	return res, res / float64(counter)
}
func minSlice(arr []float64) (min float64) {
	min = math.MaxFloat64
	for _, val := range arr {
		min = math.Min(min, val)
	}
	return
}
func maxSlice(arr []float64) (max float64) {
	max = minSlice(arr)
	for _, val := range arr {
		max = math.Max(val, max)
	}
	return
}
func MaxAndMin(arr []float64) (max float64, min float64) {
	return maxSlice(arr), minSlice(arr)
}
