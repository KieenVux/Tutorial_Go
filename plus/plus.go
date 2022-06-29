package main

import "fmt"

func sum(x, y, z int) int {
	return x + y*z
}
func partialSum(x int) func(y, z int) int {
	return func(y, z int) int {
		return sum(x, y, z)
	}
}
func main() {
	partial, partial2 := partialSum(3), partialSum(5)
	fmt.Println(partial(4, 5))
	fmt.Println(partial2(6, -9))
	fmt.Println(partialSum(5)(8, 9))
}
