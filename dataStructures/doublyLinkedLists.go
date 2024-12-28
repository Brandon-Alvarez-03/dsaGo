// doublyLinkedLists.go
package dataStructures

import "fmt"

type DoublyLinkedList struct {
    head *ListNode
    tail *ListNode
    size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{
        head: nil,
        tail: nil,
        size: 0,
    }
}

func (l *DoublyLinkedList) Insert(value interface{}) error {
    newNode := NewListNode(value)
    if l.head == nil {
        l.head = newNode
        l.tail = newNode
    } else {
        newNode.Prev = l.tail
        l.tail.Next = newNode
        l.tail = newNode
    }
    l.size++
    return nil
}

func (l *DoublyLinkedList) Delete(value interface{}) error {
    if l.head == nil {
        return fmt.Errorf("list is empty")
    }

    current := l.head
    for current != nil {
        if current.Value == value {
            if current.Prev == nil {
                l.head = current.Next
                if l.head != nil {
                    l.head.Prev = nil
                }
            } else {
                current.Prev.Next = current.Next
            }

            if current.Next == nil {
                l.tail = current.Prev
            } else {
                current.Next.Prev = current.Prev
            }

            l.size--
            return nil
        }
        current = current.Next
    }
    return fmt.Errorf("value not found")
}

func (l *DoublyLinkedList) Search(value interface{}) bool {
    current := l.head
    for current != nil {
        if current.Value == value {
            return true
        }
        current = current.Next
    }
    return false
}

func (l *DoublyLinkedList) PrintList() {
    current := l.head
    fmt.Print("List (forward): ")
    for current != nil {
        fmt.Printf("%v <-> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")

    current = l.tail
    fmt.Print("List (backward): ")
    for current != nil {
        fmt.Printf("%v <-> ", current.Value)
        current = current.Prev
    }
    fmt.Println("nil")
}