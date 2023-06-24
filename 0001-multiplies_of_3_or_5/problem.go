package main

import "fmt"

func main() {
	sum := 0
	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	sum3 := (3 + 999) * (999 / 3) / 2
	sum5 := (5 + 995) * (995 / 5) / 2
	sum15 := (15 + 990) * (990 / 15) / 2

	altSum := sum3 + sum5 - sum15
	fmt.Println(altSum, altSum == sum)
}
