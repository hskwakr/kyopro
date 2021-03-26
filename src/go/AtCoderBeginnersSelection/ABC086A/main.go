package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	product := a * b

	if product%2 == 0 {
		fmt.Printf("Even")
	} else {
		fmt.Printf("Odd")
	}
}
