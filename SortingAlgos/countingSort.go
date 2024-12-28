// countingSort.go
package sortingAlgos

import "fmt"

// CountingSort implements the counting sort algorithm
// Time Complexity: O(n + k) where k is the range of input
// Space Complexity: O(k)
// Best for small range of integers
func CountingSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    // Find the range of input array
    min, max := findMinMax(arr)
    
    // Create counting array
    range_ := max - min + 1
    count := make([]int, range_)
    output := make([]int, len(arr))

    // Store count of each element
    for _, num := range arr {
        count[num-min]++
    }

    // Modify count array to store actual position
    for i := 1; i < range_; i++ {
        count[i] += count[i-1]
    }

    // Build output array
    for i := len(arr) - 1; i >= 0; i-- {
        output[count[arr[i]-min]-1] = arr[i]
        count[arr[i]-min]--
    }

    // Copy output array to input array
    copy(arr, output)
    
    fmt.Println("CountingSort Result:", arr)
    return arr
}