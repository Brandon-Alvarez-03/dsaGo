// mergeSort.go
package sortingAlgos

import "fmt"

// MergeSort implements the merge sort algorithm
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func MergeSort(arr []int) []int {
    // Base case: arrays of size 0 or 1 are already sorted
    if len(arr) <= 1 {
        return arr
    }

    // Divide the array into two halves
    mid := len(arr) / 2
    left := MergeSort(arr[:mid])
    right := MergeSort(arr[mid:])

    // Merge the sorted halves
    result := merge(left, right)
    fmt.Println("MergeSort Result:", result)
    return result
}

func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    leftIdx, rightIdx := 0, 0

    // Compare elements from both arrays and merge them in sorted order
    for leftIdx < len(left) && rightIdx < len(right) {
        if left[leftIdx] <= right[rightIdx] {
            result = append(result, left[leftIdx])
            leftIdx++
        } else {
            result = append(result, right[rightIdx])
            rightIdx++
        }
    }

    // Append remaining elements
    result = append(result, left[leftIdx:]...)
    result = append(result, right[rightIdx:]...)
    return result
}