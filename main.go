package main

import (
	"dsaGo/searchAlgos"
	"dsaGo/sortingAlgos"
	"dsaGo/utils"
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter array size: ")
	fmt.Scanln(&size)

	arr := utils.GenerateRandomArray(size)

	var target int
	fmt.Println("\nEnter target value: ")
	fmt.Scanln(&target)

	fmt.Println("\nPerforming Linear Search...")
	searchAlgos.LinearSearch(arr, target)

	fmt.Println("\nPerforming Bubble Sort...")
	sortedArr := sortingAlgos.BubbleSort(arr)

	fmt.Println("\nPerforming Binary Search (Placeholder)...")
	searchAlgos.BinarySearch(sortedArr, target)  // Placeholder call


}
