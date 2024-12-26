package searchAlgos

import "fmt"

// LinearSearch performs a linear search on an array
func LinearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			fmt.Println("\nTarget found at index:", i)
			return i
		}
	}
	fmt.Println("\nTarget not found")
	return -1
}
