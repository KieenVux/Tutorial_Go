package main

import "fmt"

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

func main() {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)

	var number = 20
	var p *int
	p = &number
	fmt.Println(*p)

	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
	func() {
		var hugeNumber int
		hugeNumber = *p * number
		println(hugeNumber)
	}()

	variadicExample("red", "blue")

}
func variadicExample(s ...string) {
	fmt.Println(s[0])
}
