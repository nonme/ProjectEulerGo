package main

import (
	"fmt"
)

func main() {
	const lMin, lMax = 100, 1000
	maxPalindrome := 0
	var a, b int
	for i := lMin; i < lMax; i++ {
		for j := lMin; j < lMax; j++ {
			number := i * j
			sNumber := fmt.Sprint(number)

			isPalindrome := true
			n := len(sNumber)
			for c := 0; c < n/2; c++ {
				if sNumber[c] != sNumber[n-c-1] {
					isPalindrome = false
				}
			}
			if isPalindrome && number > maxPalindrome {
				maxPalindrome = number
				a = i
				b = j
			}
		}
	}
	fmt.Println(maxPalindrome, a, b)
}
