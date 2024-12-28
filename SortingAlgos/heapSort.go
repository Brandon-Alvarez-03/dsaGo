// heapSort.go
package sortingAlgos

import "fmt"

// HeapSort implements the heap sort algorithm
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func HeapSort(arr []int) []int {
    n := len(arr)

    // Build max heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    // Extract elements from heap one by one
    for i := n - 1; i > 0; i-- {
        // Move current root to end
        arr[0], arr[i] = arr[i], arr[0]
        // Call heapify on reduced heap
        heapify(arr, i, 0)
    }
    fmt.Println("HeapSort Result:", arr)
    return arr
}

// heapify maintains the heap property
func heapify(arr []int, n, i int) {
    largest := i     // Initialize largest as root
    left := 2*i + 1  // Left child
    right := 2*i + 2 // Right child

    // If left child is larger than root
    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    // If right child is larger than largest so far
    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    // If largest is not root
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        // Recursively heapify the affected sub-tree
        heapify(arr, n, largest)
    }
}