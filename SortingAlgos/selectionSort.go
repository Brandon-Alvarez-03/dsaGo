// selectionSort.go
package sortingAlgos

import "fmt"

// SelectionSort implements the selection sort algorithm
// Time Complexity: O(n²)
// Space Complexity: O(1)
func SelectionSort(arr []int) []int {
    n := len(arr)
    
    for i := 0; i < n-1; i++ {
        // Find minimum element in unsorted portion
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Swap minimum element with first element of unsorted portion
        if minIdx != i {
            arr[i], arr[minIdx] = arr[minIdx], arr[i]
        }
    }
    fmt.Println("SelectionSort Result:", arr)
    return arr
}