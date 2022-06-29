package main

import "fmt"

func countArea(str string, d ...int) int {
	var total = 1
	switch str != "" {
	case str == "Rectangle":
		for _, v := range d {
			total *= v
		}
	default:
		for _, v := range d {
			total = v * v
		}
	}
	return total
}

func main() {
	fmt.Println(countArea("Rectangle", 50, 90))
	fmt.Println(countArea("Square", 30))
}
