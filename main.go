package main

import (
	"dsaGo/dataStructures"
	"dsaGo/searchAlgos"
	"dsaGo/sortingAlgos"
	"dsaGo/utils"
	"fmt"
)

func main() {
    fmt.Println("\n=== Data Structures and Algorithms in Go ===")
    
    // Demonstrate array operations and searching
    demonstrateSearching()
    
    // Demonstrate sorting algorithms
    demonstrateSorting()
    
    // Demonstrate data structures
    demonstrateDataStructures()
    
    // Demonstrate recursive algorithms
    demonstrateRecursion()
}

func demonstrateSearching() {
    fmt.Println("\n=== Searching Algorithms ===")
    
    arr := utils.GenerateRandomArray(10)
    fmt.Printf("\nGenerated array: %v\n", arr)
    
    var target int
    fmt.Print("\nEnter target value to search: ")
    fmt.Scanln(&target)
    
    // Linear Search
    fmt.Println("\nPerforming Linear Search...")
    if idx := searchAlgos.LinearSearch(arr, target); idx != -1 {
        fmt.Printf("Linear Search: Found %d at index %d\n", target, idx)
    } else {
        fmt.Printf("Linear Search: %d not found\n", target)
    }
    
    // Sort array for Binary Search
    sortedArr := sortingAlgos.QuickSort(arr)
    
    // Binary Search
    fmt.Println("\nPerforming Binary Search...")
    if idx := searchAlgos.BinarySearch(sortedArr, target); idx != -1 {
        fmt.Printf("Binary Search: Found %d at index %d\n", target, idx)
    } else {
        fmt.Printf("Binary Search: %d not found\n", target)
    }
}

func demonstrateSorting() {
    fmt.Println("\n=== Sorting Algorithms ===")
    
    // Generate random array
    arr := utils.GenerateRandomArray(10)
    fmt.Printf("\nOriginal array: %v\n", arr)
    
    // Create copies for different sorts
    arr1 := make([]int, len(arr))
    arr2 := make([]int, len(arr))
    arr3 := make([]int, len(arr))
    arr4 := make([]int, len(arr))
    arr5 := make([]int, len(arr))
    
    copy(arr1, arr)
    copy(arr2, arr)
    copy(arr3, arr)
    copy(arr4, arr)
    copy(arr5, arr)
    
    // Demonstrate different sorting algorithms
    fmt.Println("\nQuick Sort:")
    sortingAlgos.QuickSort(arr1)
    
    fmt.Println("\nMerge Sort:")
    sortingAlgos.MergeSort(arr2)
    
    fmt.Println("\nHeap Sort:")
    sortingAlgos.HeapSort(arr3)
    
    fmt.Println("\nInsertion Sort:")
    sortingAlgos.InsertionSort(arr4)
    
    fmt.Println("\nSelection Sort:")
    sortingAlgos.SelectionSort(arr5)
}

func demonstrateDataStructures() {
    fmt.Println("\n=== Data Structures ===")
    
    // Demonstrate Linked List
    fmt.Println("\n--- Linked List Operations ---")
    list := dataStructures.NewSinglyLinkedList()
    list.Insert(1)
    list.Insert(2)
    list.Insert(3)
    list.PrintList()
    
    // Demonstrate Stack
    fmt.Println("\n--- Stack Operations ---")
    stack := dataStructures.NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    stack.PrintStack()
    
    // Demonstrate Queue
    fmt.Println("\n--- Queue Operations ---")
    queue := dataStructures.NewQueue()
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    queue.PrintQueue()
    
    // Demonstrate Binary Search Tree
    fmt.Println("\n--- Binary Search Tree Operations ---")
    bst := dataStructures.NewBinarySearchTree()
    bst.Insert(5)
    bst.Insert(3)
    bst.Insert(7)
    bst.PrintTree()
}

func demonstrateRecursion() {
    fmt.Println("\n=== Recursive Algorithms ===")
    
    // Factorial
    n := 5
    if result, err := utils.Factorial(n); err == nil {
        fmt.Printf("\nFactorial of %d = %d\n", n, result)
    }
    
    // Fibonacci
    if result, err := utils.FibonacciMemoized(n); err == nil {
        fmt.Printf("Fibonacci(%d) = %d\n", n, result)
    }
    
    // Tower of Hanoi
    fmt.Println("\nSolving Tower of Hanoi with 3 disks:")
    utils.TowerOfHanoi(3, "A", "B", "C")
}