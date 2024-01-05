package main

import "fmt"

func mostRepeatedElement(arr []string) string {
	// Create a map to store the count of each element
	elementCount := make(map[string]int)

	// Iterate through the array and update the count in the map
	for _, element := range arr {
		elementCount[element]++
	}

	// Find the element with the maximum count
	var mostRepeated string
	maxCount := 0
	for element, count := range elementCount {
		if count > maxCount {
			maxCount = count
			mostRepeated = element
		}
	}

	return mostRepeated
}

func test() {
	inputData := []string{"apple", "pie", "apple", "red", "red", "red"}
	output := mostRepeatedElement(inputData)
	fmt.Println(output)
}

func main() {
	// Test with the given input
	test()
}
