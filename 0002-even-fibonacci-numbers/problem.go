package main

import "fmt"

func main() {
	a, b, sum := 0, 1, 0
	for a < 4000000 {
		if a%2 == 0 {
			sum += a
		}
		t := a
		a = b
		b = t + b
	}

	fmt.Println(sum)
}
