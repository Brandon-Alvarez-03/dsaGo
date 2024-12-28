// binaryTrees.go
package dataStructures

import "fmt"

type BinarySearchTree struct {
    root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
    return &BinarySearchTree{root: nil}
}

func (t *BinarySearchTree) Insert(value int) error {
    if t.root == nil {
        t.root = NewTreeNode(value)
        return nil
    }
    return t.insertNode(t.root, value)
}

func (t *BinarySearchTree) insertNode(node *TreeNode, value int) error {
    if value < node.Value {
        if node.Left == nil {
            node.Left = NewTreeNode(value)
            return nil
        }
        return t.insertNode(node.Left, value)
    }
    if value > node.Value {
        if node.Right == nil {
            node.Right = NewTreeNode(value)
            return nil
        }
        return t.insertNode(node.Right, value)
    }
    return fmt.Errorf("value already exists")
}

func (t *BinarySearchTree) Search(value int) bool {
    return t.searchNode(t.root, value)
}

func (t *BinarySearchTree) searchNode(node *TreeNode, value int) bool {
    if node == nil {
        return false
    }
    if value == node.Value {
        return true
    }
    if value < node.Value {
        return t.searchNode(node.Left, value)
    }
    return t.searchNode(node.Right, value)
}

func (t *BinarySearchTree) Delete(value int) error {
    var err error
    t.root, err = t.deleteNode(t.root, value)
    return err
}

func (t *BinarySearchTree) deleteNode(node *TreeNode, value int) (*TreeNode, error) {
    if node == nil {
        return nil, fmt.Errorf("value not found")
    }

    if value < node.Value {
        var err error
        node.Left, err = t.deleteNode(node.Left, value)
        return node, err
    }
    if value > node.Value {
        var err error
        node.Right, err = t.deleteNode(node.Right, value)
        return node, err
    }

    // Case 1: Leaf node
    if node.Left == nil && node.Right == nil {
        return nil, nil
    }

    // Case 2: Single child
    if node.Left == nil {
        return node.Right, nil
    }
    if node.Right == nil {
        return node.Left, nil
    }

    // Case 3: Two children
    minRight := t.findMin(node.Right)
    node.Value = minRight.Value
    var err error
    node.Right, err = t.deleteNode(node.Right, minRight.Value)
    return node, err
}

func (t *BinarySearchTree) findMin(node *TreeNode) *TreeNode {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}

func (t *BinarySearchTree) PrintTree() {
    t.printInOrder(t.root)
    fmt.Println()
}

func (t *BinarySearchTree) printInOrder(node *TreeNode) {
    if node != nil {
        t.printInOrder(node.Left)
        fmt.Printf("%d -> ", node.Value)
        t.printInOrder(node.Right)
    }
}