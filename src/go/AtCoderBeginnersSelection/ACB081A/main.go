package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	var count int
	for _, v := range input {
		if v == '1' {
			count++
		}
	}
	fmt.Printf("%d", count)
}
