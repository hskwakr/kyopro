package main

import "fmt"

func main() {
	var amount int
	fmt.Scan(&amount)

	var numbers [200]int
	for i := 0; i < amount; i++ {
		fmt.Scan(&numbers[i])
	}

	count := 0
	for isThereOdd(numbers) {
		count++
		numbers = divide(numbers)
	}

	fmt.Printf("%d\n", count)
}

func isThereOdd(numbers [200]int) bool {
	result := true
	for _, num := range numbers {
		if num%2 != 0 {
			result = false
		}
	}
	return result
}

func divide(numbers [200]int) [200]int {
	var result [200]int
	for i, num := range numbers {
		result[i] = num / 2
	}
	return result
}
