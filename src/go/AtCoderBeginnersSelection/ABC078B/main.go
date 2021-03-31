package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)

	if x%50 != 0 {
		panic("X must be multiple 50")
	}

	count := findExactCoin(x, a, b, c)
	fmt.Println(count)
}

func findExactCoin(price, amount_a, amount_b, amount_c int) int {
	result := 0
	for i := 0; i <= amount_a; i++ {
		for j := 0; j <= amount_b; j++ {
			for k := 0; k <= amount_c; k++ {
				amount := ((500 * i) + (100 * j) + (50 * k))
				if amount == 0 {
					continue
				}

				rate := float64(price) / float64(amount)
				if rate == 1.0 {
					//fmt.Println(i, j, k, amount, rate)
					result++
				}
			}
		}
	}
	return result
}
