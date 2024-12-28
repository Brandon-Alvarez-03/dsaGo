// queues.go
package dataStructures

import "fmt"

// ArrayQueue implements the Queue interface using a slice
type ArrayQueue struct {
    items []interface{}
}

func NewQueue() *ArrayQueue {
    return &ArrayQueue{
        items: make([]interface{}, 0),
    }
}

func (q *ArrayQueue) Enqueue(value interface{}) error {
    q.items = append(q.items, value)
    return nil
}

func (q *ArrayQueue) Dequeue() (interface{}, error) {
    if q.IsEmpty() {
        return nil, fmt.Errorf("queue is empty")
    }
    
    value := q.items[0]
    q.items = q.items[1:]
    return value, nil
}

func (q *ArrayQueue) Peek() (interface{}, error) {
    if q.IsEmpty() {
        return nil, fmt.Errorf("queue is empty")
    }
    return q.items[0], nil
}

func (q *ArrayQueue) IsEmpty() bool {
    return len(q.items) == 0
}

func (q *ArrayQueue) PrintQueue() {
    fmt.Print("Queue (front to back): ")
    for _, item := range q.items {
        fmt.Printf("%v -> ", item)
    }
    fmt.Println("back")
}