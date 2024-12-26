package main

import (
	"dsaGo/searchAlgos"
	"dsaGo/utils"
	"fmt"
)

func main() {
	var size int
	fmt.Print("Enter array size: ")
	fmt.Scanln(&size)

	arr := utils.GenerateRandomArray(size)

	var target int
	fmt.Print("Enter target value: ")
	fmt.Scanln(&target)

	searchAlgos.LinearSearch(arr, target)
}
