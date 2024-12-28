// insertionSort.go
package sortingAlgos

import "fmt"

// InsertionSort implements the insertion sort algorithm
// Time Complexity: O(n²)
// Space Complexity: O(1)
// Best for small arrays or nearly sorted arrays
func InsertionSort(arr []int) []int {
    for i := 1; i < len(arr); i++ {
        // Store current element as key
        key := arr[i]
        j := i - 1

        // Move elements greater than key ahead
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        // Place key in its correct position
        arr[j+1] = key
    }
    fmt.Println("InsertionSort Result:", arr)
    return arr
}
