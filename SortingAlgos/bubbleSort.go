package sortingAlgos

import (
	"fmt"
)

func BubbleSort(arr []int) []int {
	// test_arr := []int{3, 2, 6, 5, 9, 1, 11, 10}

	// compare the first element to the following and if lower, swap
	// keep track of "swaps" only once the swap tracker is ZERO at the END of the iteration can we confirm the arr has been sorted
	
	for swap_count :=  1; swap_count != 0; {
		swap_count = 0
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap_count += 1
				}
			}
		}
	fmt.Println("\nArray Sorted: ", arr)
	return arr
}
