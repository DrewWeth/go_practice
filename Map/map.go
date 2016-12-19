// Attempting to force Golang to do functional mapping over an array. Basic tests show it's about 8X slower than iterative approach.

// Output:
// Map  2000000	       634 ns/op
// Iterative 20000000	        76.4 ns/op
// [Finished in 3.958s]

package main

import (
	"fmt"
	"testing"
)

type any interface{}

var print = fmt.Println
var square = func(n int) int {
	return n * n
}

func FunctionalMap(operation func(p ...any) any, is []any, n int) []any {
	r := make([]any, n)
	for index, value := range is {
		r[index] = operation(value)
	}
	return r
}

func Square(vals ...any) any {
	val := (vals[0]).(any) // Only one input
	return (val).(int) * (val).(int)
}

func doSquare(input []int) []int {
	output := make([]int, len(input))
	for i, _ := range input {
		output[i] = input[i] * input[i]
	}
	return output
}

func main() {
	print(square(5))

	a := testing.Benchmark(bmMap)
	print("Map", a)
	b := testing.Benchmark(bmIterative)
	print("Iterative", b)
}

func bmMap(b *testing.B) {
	arr := []any{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		FunctionalMap(Square, arr, len(arr))
	}
}

func bmIterative(b *testing.B) {
	arr := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		doSquare(arr)
	}
}
