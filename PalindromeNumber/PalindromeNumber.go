package main

import (
	"fmt"
)

func main() {
	x := 121 //Example 1
	//x := -121	//Example 2
	//x := 10	//Example 3
	fmt.Println(IsPalindrome(x))
}
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	reversed := 0
	for num != 0 {
		reversed = 10*reversed + num%10
		num /= 10
	}
	return reversed == x
}
