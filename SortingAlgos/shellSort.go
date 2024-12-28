// shellSort.go
package sortingAlgos

import "fmt"

// ShellSort implements the Shell sort algorithm
// Time Complexity: O(n²) worst case, but performs better in practice
// Space Complexity: O(1)
func ShellSort(arr []int) []int {
    n := len(arr)
    // Start with a large gap, then reduce it
    for gap := n/2; gap > 0; gap /= 2 {
        // Do a gapped insertion sort
        for i := gap; i < n; i++ {
            temp := arr[i]
            // Shift earlier gap-sorted elements up until the correct location
            j := i
            for j >= gap && arr[j-gap] > temp {
                arr[j] = arr[j-gap]
                j -= gap
            }
            arr[j] = temp
        }
    }
    fmt.Println("ShellSort Result:", arr)
    return arr
}