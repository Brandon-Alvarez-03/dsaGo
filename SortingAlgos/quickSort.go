// quickSort.go
package sortingAlgos

import "fmt"

// QuickSort implements the quicksort algorithm using the Lomuto partition scheme
// Time Complexity: O(n log n) average case, O(n²) worst case
// Space Complexity: O(log n) due to recursion stack
func QuickSort(arr []int) []int {
    // Handle base cases
    if len(arr) <= 1 {
        return arr
    }
    quickSortHelper(arr, 0, len(arr)-1)
    fmt.Println("QuickSort Result:", arr)
    return arr
}

func quickSortHelper(arr []int, low, high int) {
    if low < high {
        // Partition the array and get the pivot position
        pivot := partition(arr, low, high)
        // Recursively sort the left and right portions
        quickSortHelper(arr, low, pivot-1)
        quickSortHelper(arr, pivot+1, high)
    }
}

func partition(arr []int, low, high int) int {
    // Choose the rightmost element as pivot
    pivot := arr[high]
    // Index of smaller element
    i := low - 1

    // Compare each element with pivot
    for j := low; j < high; j++ {
        if arr[j] <= pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    // Place pivot in its correct position
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}