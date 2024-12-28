// structs.go
package dataStructures

// ListNode represents a node in a linked list
type ListNode struct {
    Value interface{}
    Next  *ListNode
    Prev  *ListNode // Used for doubly-linked lists
}

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

// HashEntry represents an entry in the hash table
type HashEntry struct {
    Key   interface{}
    Value interface{}
}

// GraphVertex represents a vertex in a graph
type GraphVertex struct {
    Key   interface{}
    Edges map[interface{}]float64 // Maps to weight of edge
}

// TrieNode represents a node in a trie
type TrieNode struct {
    Children map[rune]*TrieNode
    IsEnd    bool
}

// GetValue implements the Node interface for ListNode
func (n *ListNode) GetValue() interface{} {
    return n.Value
}

// SetValue implements the Node interface for ListNode
func (n *ListNode) SetValue(v interface{}) {
    n.Value = v
}

// NewListNode creates a new ListNode
func NewListNode(value interface{}) *ListNode {
    return &ListNode{
        Value: value,
        Next:  nil,
        Prev:  nil,
    }
}

// NewTreeNode creates a new TreeNode
func NewTreeNode(value int) *TreeNode {
    return &TreeNode{
        Value: value,
        Left:  nil,
        Right: nil,
    }
}

// NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
    return &TrieNode{
        Children: make(map[rune]*TrieNode),
        IsEnd:    false,
    }
}

// NewGraphVertex creates a new GraphVertex
func NewGraphVertex(key interface{}) *GraphVertex {
    return &GraphVertex{
        Key:   key,
        Edges: make(map[interface{}]float64),
    }
}