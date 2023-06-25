package main

import "fmt"

func main() {
	upperBound := 1000000
	isPrime := make([]bool, upperBound)
	primes := make([]int, 0)

	for i := 2; i < upperBound; i++ {
		if !isPrime[i] {
			primes = append(primes, i)
			for j := i; j < upperBound; j += i {
				isPrime[j] = true
			}
		}
	}
	fmt.Println(primes[10000])
}
