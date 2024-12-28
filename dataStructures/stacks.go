// stacks.go
package dataStructures

import "fmt"

// ArrayStack implements our stack data structure using a slice
type ArrayStack struct {
    items []interface{}
}

// NewStack creates and returns a new stack instance
func NewStack() *ArrayStack {
    return &ArrayStack{
        items: make([]interface{}, 0),
    }
}

// Push adds a new element to the top of the stack
func (s *ArrayStack) Push(value interface{}) error {
    s.items = append(s.items, value)
    return nil
}

// Pop removes and returns the top element from the stack
func (s *ArrayStack) Pop() (interface{}, error) {
    if s.IsEmpty() {
        return nil, fmt.Errorf("stack is empty")
    }
    
    lastIdx := len(s.items) - 1
    value := s.items[lastIdx]
    s.items = s.items[:lastIdx]
    return value, nil
}

// Peek returns the top element without removing it
func (s *ArrayStack) Peek() (interface{}, error) {
    if s.IsEmpty() {
        return nil, fmt.Errorf("stack is empty")
    }
    return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack has no elements
func (s *ArrayStack) IsEmpty() bool {
    return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *ArrayStack) Size() int {
    return len(s.items)
}

// PrintStack displays the stack contents from top to bottom
func (s *ArrayStack) PrintStack() {
    fmt.Print("Stack (top to bottom): ")
    for i := len(s.items) - 1; i >= 0; i-- {
        fmt.Printf("%v -> ", s.items[i])
    }
    fmt.Println("bottom")
}