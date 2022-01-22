package main

import "fmt"

func main () {
	for i :=0; i < 7; i++ {
		fmt.Printf("%d\n", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	
	return fibonacci(n-1) + fibonacci(n-2)
}