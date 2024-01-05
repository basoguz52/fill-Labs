package main

import "fmt"

func divideByTwoRecursive(n int) {
	if n >= 2 {
		// If n is even and greater than or equal to 2, call the function again and then print n
		divideByTwoRecursive(n / 2)
		fmt.Printf("Processing with even number: %d\n", n)
	} else {
		// If n is less than or equal to 2, end the recursion
		fmt.Println("End of recursion")
	}
}

func main() {
	// Example usage
	divideByTwoRecursive(9)
}

