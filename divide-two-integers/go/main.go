package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("divide-two-integers")

	divide(5, 2)
	divide(5, -2)
	divide(-5, 2)
	divide(-5, -2)
	divide(math.MinInt32, -1)
	divide(math.MinInt32, -2)
}

func divide(dividend int, divisor int) int {
	// fmt.Printf("***** %d / %d\n", dividend, divisor)

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	flag_diff := (dividend > 0) != (divisor > 0)

	var up = dividend
	if up > 0 {
		up = -up
	}
	var down = divisor
	if down > 0 {
		down = -down
	}

	var quotient int = 0

	for up <= down {
		var i = 0
		var tmp = down

		for (up - tmp) <= tmp {
			tmp <<= 1
			i += 1
		}

		up -= tmp
		quotient += 1 << i
	}

	if flag_diff {
		quotient = -quotient
	}

	return quotient
	// fmt.Printf("***** %d / %d = %d\n", dividend, divisor, quotient)
}
