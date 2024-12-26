package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomArray creates and returns a randomized array of specified size
func GenerateRandomArray(size int) []int {
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(101)  // Values between 0-100
	}
	fmt.Println("Generated array:", arr)
	return arr
}
