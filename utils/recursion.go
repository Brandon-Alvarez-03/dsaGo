// recursion.go
package utils

import "fmt"

// RecursionExamples provides various recursive algorithms and their implementations
// Each function demonstrates different aspects of recursive problem-solving

// Factorial calculates n! using recursion
// Example: Factorial(5) = 5 * 4 * 3 * 2 * 1 = 120
// Time Complexity: O(n)
// Space Complexity: O(n) due to recursive call stack
func Factorial(n int) (int, error) {
    // Input validation
    if n < 0 {
        return 0, fmt.Errorf("factorial is not defined for negative numbers")
    }
    
    // Base case: factorial of 0 or 1 is 1
    if n <= 1 {
        return 1, nil
    }
    
    // Recursive case: n! = n * (n-1)!
    smaller, err := Factorial(n - 1)
    if err != nil {
        return 0, err
    }
    
    result := n * smaller
    fmt.Printf("Calculated factorial(%d) = %d\n", n, result)
    return result, nil
}

// Fibonacci calculates the nth number in the Fibonacci sequence
// Example: Fibonacci(6) = 8 (sequence: 0, 1, 1, 2, 3, 5, 8)
// Time Complexity: O(2^n) - exponential without memoization
// Space Complexity: O(n) due to recursive call stack
func Fibonacci(n int) (int, error) {
    if n < 0 {
        return 0, fmt.Errorf("fibonacci is not defined for negative numbers")
    }
    
    // Base cases
    if n <= 1 {
        return n, nil
    }
    
    // Recursive case: F(n) = F(n-1) + F(n-2)
    n1, err := Fibonacci(n - 1)
    if err != nil {
        return 0, err
    }
    
    n2, err := Fibonacci(n - 2)
    if err != nil {
        return 0, err
    }
    
    result := n1 + n2
    fmt.Printf("Calculated fibonacci(%d) = %d\n", n, result)
    return result, nil
}

// FibonacciMemoized provides an optimized version of Fibonacci using memoization
// Time Complexity: O(n) - linear with memoization
// Space Complexity: O(n) for memoization cache
func FibonacciMemoized(n int) (int, error) {
    if n < 0 {
        return 0, fmt.Errorf("fibonacci is not defined for negative numbers")
    }
    
    // Initialize memoization cache
    memo := make(map[int]int)
    result, err := fibMemo(n, memo)
    
    if err != nil {
        return 0, err
    }
    
    fmt.Printf("Calculated memoized fibonacci(%d) = %d\n", n, result)
    return result, nil
}

// Helper function for memoized Fibonacci calculation
func fibMemo(n int, memo map[int]int) (int, error) {
    // Check if value exists in cache
    if val, exists := memo[n]; exists {
        return val, nil
    }
    
    // Base cases
    if n <= 1 {
        memo[n] = n
        return n, nil
    }
    
    // Recursive case with memoization
    n1, err := fibMemo(n-1, memo)
    if err != nil {
        return 0, err
    }
    
    n2, err := fibMemo(n-2, memo)
    if err != nil {
        return 0, err
    }
    
    memo[n] = n1 + n2
    return memo[n], nil
}

// BinarySearch performs a recursive binary search on a sorted array
// Time Complexity: O(log n)
// Space Complexity: O(log n) due to recursive call stack
func BinarySearchRecursive(arr []int, target, left, right int) (int, error) {
    if left > right {
        return -1, fmt.Errorf("element not found")
    }
    
    mid := left + (right-left)/2
    
    if arr[mid] == target {
        return mid, nil
    }
    
    if arr[mid] > target {
        return BinarySearchRecursive(arr, target, left, mid-1)
    }
    
    return BinarySearchRecursive(arr, target, mid+1, right)
}

// TowerOfHanoi solves the Tower of Hanoi puzzle recursively
// Time Complexity: O(2^n)
// Space Complexity: O(n) due to recursive call stack
func TowerOfHanoi(n int, source, auxiliary, target string) {
    if n > 0 {
        // Move n-1 disks from source to auxiliary peg
        TowerOfHanoi(n-1, source, target, auxiliary)
        
        // Move the largest disk from source to target peg
        fmt.Printf("Move disk %d from %s to %s\n", n, source, target)
        
        // Move n-1 disks from auxiliary to target peg
        TowerOfHanoi(n-1, auxiliary, source, target)
    }
}

// RecursivePowerFunction calculates x^n using recursion
// Example: Power(2, 3) = 8
// Time Complexity: O(log n)
// Space Complexity: O(log n)
func Power(x float64, n int) (float64, error) {
    // Handle negative exponents
    if n < 0 {
        x = 1 / x
        n = -n
    }
    
    return powerHelper(x, n), nil
}

func powerHelper(x float64, n int) float64 {
    // Base cases
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x
    }
    
    // Recursive case with optimization for even/odd exponents
    result := powerHelper(x, n/2)
    result = result * result
    
    if n%2 == 1 {
        result = result * x
    }
    
    return result
}

// GCD calculates the Greatest Common Divisor using recursion (Euclidean algorithm)
// Time Complexity: O(log min(a,b))
// Space Complexity: O(log min(a,b))
func GCD(a, b int) (int, error) {
    if a < 0 || b < 0 {
        return 0, fmt.Errorf("GCD is not defined for negative numbers")
    }
    
    if b == 0 {
        return a, nil
    }
    
    return GCD(b, a%b)
}