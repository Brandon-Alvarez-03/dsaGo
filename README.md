# Data Structures and Algorithms in Go (dsaGo)

A comprehensive implementation of common data structures and algorithms in Go, featuring searching and sorting algorithms with a focus on clean, efficient, and well-documented code.

## Features

- **Search Algorithms**

  - Linear Search: Simple sequential search implementation
  - Binary Search: Efficient search for sorted arrays

- **Sorting Algorithms**

  - Bubble Sort: Basic sorting algorithm with swap tracking

- **Utility Functions**
  - Random Array Generator: Creates test arrays of specified size

## Project Structure

dsaGo/
├── README.md
├── main.go # Main application entry point
├── go.mod # Go module file
├── makefile # Build and run automation
├── SortingAlgos/
│ └── bubbleSort.go # Bubble sort implementation
├── searchAlgos/
│ ├── binarySearch.go # Binary search implementation
│ └── linearSearch.go # Linear search implementation
└── utils/
└── array.go # Array generation utilities
Copy

## Prerequisites

- Go 1.16 or higher
- Make (optional, for using makefile commands)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/dsaGo.git
   cd dsaGo
   ```

Initialize the Go module (if not already initialized):
bashCopymake init
or manually:
bashCopygo mod init dsaGo
go mod tidy

Usage
Using Make Commands
The project includes a makefile with common commands:

Build the project:
bashCopymake build

Run the application:
bashCopymake run

Clean up binaries:
bashCopymake clean

Running Directly
You can also run the application directly using Go:
bashCopygo run main.go
Interactive Features
The application will prompt you to:

Enter the size of the array you want to generate
Input a target value to search for
The program will then:

Generate a random array
Perform a linear search for your target
Sort the array using bubble sort
Perform a binary search on the sorted array

Implementation Details
Search Algorithms
Linear Search

Time Complexity: O(n)
Implementation: Sequential search through array elements
Returns: Index of found element or -1 if not found

Binary Search

Time Complexity: O(log n)
Prerequisite: Sorted array
Implementation: Divide and conquer approach
Returns: 0 if found, -1 if not found

Sorting Algorithms
Bubble Sort

Time Complexity: O(n²)
Space Complexity: O(1)
Implementation: Swap-based sorting with optimization for already sorted arrays
Features: Tracks swap count to optimize performance

Utility Functions
Random Array Generator

Generates arrays of specified size
Random values range: 0-100
Uses Go's math/rand package with time-based seed

Contributing

Fork the repository
Create your feature branch (git checkout -b feature/amazing-feature)
Commit your changes (git commit -m 'Add some amazing feature')
Push to the branch (git push origin feature/amazing-feature)
Open a Pull Request

License
This project is licensed under the MIT License - see the LICENSE file for details.
Acknowledgments

Built as part of learning Go programming language
Inspired by classic computer science algorithms and data structures

Future Enhancements

Add more sorting algorithms (Quick Sort, Merge Sort, etc.)
Implement additional search algorithms
Add comprehensive unit tests
Include benchmarking tools
Add visualization for algorithm steps
