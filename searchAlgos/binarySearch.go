package searchAlgos

import (
	"fmt"
)

// define binarySearch function and utilize the array generator to make the searchable array

func BinarySearch (arr []int, target int) int {
	fmt.Println("\n", arr, "for binary search")
	start, end := 0, len(arr) - 1

	for start <= end {
		middle := (start + end) / 2
		
		if target == arr[middle] {
			fmt.Println("\nTarget: ", target, "found")
			return 0
		} else if target < arr[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	fmt.Println("\nTarget: ", target, " not found in array")
	return -1
}
