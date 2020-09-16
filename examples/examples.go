package main

import (
	"bytes"
	"fmt"
	"math"
)

var buff *bytes.Buffer
var stones []int
var memo []int

func main() {

	fmt.Print()
	//var flag bool = isPalindrome("aabaa")
	var word string = "aa3aa"

	println(minInsertions(word, 0, len(word)-1))

}

/*
This method allows to find the min insertions to get a string into a palindrome
*/
func minInsertions(word string, start int, end int) int {

	if start > end {
		return math.MaxInt64
	}

	if start == end {
		return 0
	}
	if start == end-1 {
		if word[start] == word[end] {
			return 0
		} else {
			return 1
		}
	}
	if word[start] == word[end] {
		return minInsertions(word, start+1, end-1)
	}
	return min(minInsertions(word, start, end-1), minInsertions(word, start+1, end)+1)
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	if x > y {
		return y
	}
	return x
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] != word[(len(word)-1)-i] {
			return false
		}
	}
	return true
}
