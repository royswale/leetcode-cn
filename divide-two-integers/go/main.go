package main

import "fmt"

func main() {
	fmt.Println("divide-two-integers")

	divide(5, 2)
}
func divide(dividend int, divisor int) int {
	fmt.Printf("***** %d / %d\n", dividend, divisor)
	var quotient int = 0

	fmt.Printf("***** %d / %d = %d\n", dividend, divisor, quotient)
	return 0
}
