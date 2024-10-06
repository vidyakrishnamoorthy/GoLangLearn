package main

import (
	"GOLANGLEARN/searchingAlgorithms/search"
	"fmt"
)

func main() {
	arr := []int{2, 4, 6, 8, 10, 12}
	target := 8

	fmt.Println("Linear Search:", search.LinearSearch(arr, target))
	fmt.Println("Binary Search:", search.BinarySearch(arr, target))
}
