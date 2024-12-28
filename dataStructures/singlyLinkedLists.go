// singlyLinkedLists.go
package dataStructures

import "fmt"

type SinglyLinkedList struct {
    head *ListNode
    size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
    return &SinglyLinkedList{
        head: nil,
        size: 0,
    }
}

func (l *SinglyLinkedList) Insert(value interface{}) error {
    newNode := NewListNode(value)
    if l.head == nil {
        l.head = newNode
    } else {
        current := l.head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    l.size++
    return nil
}

func (l *SinglyLinkedList) Delete(value interface{}) error {
    if l.head == nil {
        return fmt.Errorf("list is empty")
    }

    if l.head.Value == value {
        l.head = l.head.Next
        l.size--
        return nil
    }

    current := l.head
    for current.Next != nil && current.Next.Value != value {
        current = current.Next
    }

    if current.Next == nil {
        return fmt.Errorf("value not found")
    }

    current.Next = current.Next.Next
    l.size--
    return nil
}

func (l *SinglyLinkedList) Search(value interface{}) bool {
    current := l.head
    for current != nil {
        if current.Value == value {
            return true
        }
        current = current.Next
    }
    return false
}

func (l *SinglyLinkedList) PrintList() {
    current := l.head
    fmt.Print("List: ")
    for current != nil {
        fmt.Printf("%v -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}