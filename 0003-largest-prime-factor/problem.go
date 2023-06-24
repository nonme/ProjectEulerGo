package main

import "fmt"

func findLargestPrimeFactor(number int) int {
	currentNumber := 2
	lastPrime := 0

	for number > 1 {
		for number%currentNumber == 0 {
			lastPrime = currentNumber
			number /= currentNumber
		}
		currentNumber++
	}
	return lastPrime
}

func main() {
	const number = 600851475143

	fmt.Println(findLargestPrimeFactor(number))
}
