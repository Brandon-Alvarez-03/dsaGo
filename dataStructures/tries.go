// tries.go
package dataStructures

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{
        root: NewTrieNode(),
    }
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if _, exists := node.Children[ch]; !exists {
            node.Children[ch] = NewTrieNode()
        }
        node = node.Children[ch]
    }
    node.IsEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if _, exists := node.Children[ch]; !exists {
            return false
        }
        node = node.Children[ch]
    }
    return node.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if _, exists := node.Children[ch]; !exists {
            return false
        }
        node = node.Children[ch]
    }
    return true
}