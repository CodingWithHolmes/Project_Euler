package main

import "fmt"

func sumEvenFibonacci(limit int) int {
	sum := 0
	a, b := 0, 1 
	fmt.Println("a", a)
	fmt.Println("b", b)
	for b < limit {
		if b%2 == 0 { 
			sum += b
		}
		a, b = b, a+b 
		fmt.Println("a", a)
		fmt.Println("b", b)
	}
	return sum
}

func main() {
	limit := 4000000
	result := sumEvenFibonacci(limit)
	fmt.Println("Sum of even Fibonacci numbers below 4 million:", result)
