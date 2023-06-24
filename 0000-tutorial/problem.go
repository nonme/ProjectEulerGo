package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"

	"rsc.io/quote"
)

func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

// 'Naked' return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python bool
var java string

/*
	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
		// represents a Unicode code point

	float32 float64

	complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println(quote.Go())

	fmt.Println(rand.Intn(10)) // [0, 10)

	fmt.Println(add(42, 13))
	fmt.Println(multiply(3, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(15))

	var variable int
	fmt.Println(variable, c, python, java)

	var var1, var2, var3 = true, 3, "no!"
	fmt.Println(var1, var2, var3)

	fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v\n", z, z)

	int_ := 42
	float_ := float64(int_)
	uint_ := uint(float_)
	fmt.Printf("Type: %T Value %v\n", uint_, uint_)

	const Pi = 3.1415
	fmt.Println("Happy", Pi, "Day")
}
