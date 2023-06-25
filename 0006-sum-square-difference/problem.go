package main

import "fmt"

func main() {
	sumOfSquares, squareSum := 0, 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		squareSum += i
	}
	squareSum *= squareSum

	difference := squareSum - sumOfSquares
	fmt.Println(difference)
}
