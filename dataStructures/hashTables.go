// hashTables.go
package dataStructures

import (
    "fmt"
    "hash/fnv"
)

type HashTable struct {
    buckets    []*HashEntry
    size       int
    capacity   int
}

func NewHashTable(capacity int) *HashTable {
    return &HashTable{
        buckets:  make([]*HashEntry, capacity),
        size:     0,
        capacity: capacity,
    }
}

func (h *HashTable) hash(key interface{}) uint32 {
    hasher := fnv.New32()
    hasher.Write([]byte(fmt.Sprintf("%v", key)))
    return hasher.Sum32() % uint32(h.capacity)
}

func (h *HashTable) Insert(key, value interface{}) error {
    index := h.hash(key)
    entry := &HashEntry{Key: key, Value: value}
    
    // Linear probing for collision resolution
    for i := 0; i < h.capacity; i++ {
        probeIndex := (int(index) + i) % h.capacity
        if h.buckets[probeIndex] == nil {
            h.buckets[probeIndex] = entry
            h.size++
            return nil
        }
        if h.buckets[probeIndex].Key == key {
            h.buckets[probeIndex].Value = value
            return nil
        }
    }
    return fmt.Errorf("hash table is full")
}

func (h *HashTable) Get(key interface{}) (interface{}, error) {
    index := h.hash(key)
    
    for i := 0; i < h.capacity; i++ {
        probeIndex := (int(index) + i) % h.capacity
        if h.buckets[probeIndex] == nil {
            break
        }
        if h.buckets[probeIndex].Key == key {
            return h.buckets[probeIndex].Value, nil
        }
    }
    return nil, fmt.Errorf("key not found")
}

func (h *HashTable) Delete(key interface{}) error {
    index := h.hash(key)
    
    for i := 0; i < h.capacity; i++ {
        probeIndex := (int(index) + i) % h.capacity
        if h.buckets[probeIndex] == nil {
            break
        }
        if h.buckets[probeIndex].Key == key {
            h.buckets[probeIndex] = nil
            h.size--
            return nil
        }
    }
    return fmt.Errorf("key not found")
}

func (h *HashTable) PrintTable() {
    for i, entry := range h.buckets {
        if entry != nil {
            fmt.Printf("Bucket[%d]: %v -> %v\n", i, entry.Key, entry.Value)
        }
    }
}