// bucketSort.go
package sortingAlgos

import (
	"fmt"
	"math"
)

// BucketSort implements the bucket sort algorithm
// Time Complexity: O(n + k) average case, O(n²) worst case
// Space Complexity: O(n + k)
// Best for uniformly distributed data over a range
func BucketSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    // Find range of values
    min, max := findMinMax(arr)
    if min == max {
        return arr
    }

    // Create buckets
    bucketCount := int(math.Sqrt(float64(len(arr)))) + 1
    buckets := make([][]int, bucketCount)
    
    // Initialize buckets
    for i := range buckets {
        buckets[i] = make([]int, 0)
    }

    // Distribute elements into buckets
    for _, num := range arr {
        idx := (num - min) * (bucketCount - 1) / (max - min)
        buckets[idx] = append(buckets[idx], num)
    }

    // Sort individual buckets using insertion sort
    pos := 0
    for i := 0; i < bucketCount; i++ {
        if len(buckets[i]) > 0 {
            InsertionSort(buckets[i])
            copy(arr[pos:], buckets[i])
            pos += len(buckets[i])
        }
    }

    fmt.Println("BucketSort Result:", arr)
    return arr
}

func findMinMax(arr []int) (int, int) {
    if len(arr) == 0 {
        return 0, 0
    }
    min, max := arr[0], arr[0]
    for _, num := range arr {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    return min, max
}