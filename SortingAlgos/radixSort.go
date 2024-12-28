// radixSort.go
package sortingAlgos

import "fmt"

// RadixSort implements the radix sort algorithm for non-negative integers
// Time Complexity: O(d * (n + k)) where d is the number of digits and k is the range of values
// Space Complexity: O(n + k)
func RadixSort(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }

    // Find the maximum number to know number of digits
    max := getMax(arr)

    // Do counting sort for every digit
    // Instead of passing digit number, exp is passed
    // exp is 10^i where i is current digit number
    for exp := 1; max/exp > 0; exp *= 10 {
        countingSortByDigit(arr, exp)
    }

    fmt.Println("RadixSort Result:", arr)
    return arr
}

func getMax(arr []int) int {
    max := arr[0]
    for _, value := range arr {
        if value > max {
            max = value
        }
    }
    return max
}

func countingSortByDigit(arr []int, exp int) {
    n := len(arr)
    output := make([]int, n)
    count := make([]int, 10) // 0-9 digits

    // Store count of occurrences in count[]
    for i := 0; i < n; i++ {
        digit := (arr[i] / exp) % 10
        count[digit]++
    }

    // Change count[i] so that count[i] now contains
    // actual position of this digit in output[]
    for i := 1; i < 10; i++ {
        count[i] += count[i-1]
    }

    // Build the output array
    for i := n - 1; i >= 0; i-- {
        digit := (arr[i] / exp) % 10
        output[count[digit]-1] = arr[i]
        count[digit]--
    }

    // Copy the output array to arr[]
    copy(arr, output)
}