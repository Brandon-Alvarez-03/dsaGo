// graphs.go
package dataStructures

import "fmt"

// AdjListGraph implements the Graph interface using an adjacency list representation.
// We use this name to clearly distinguish it from the Graph interface.
type AdjListGraph struct {
    vertices map[interface{}]*GraphVertex
}

// NewGraph creates and returns a new instance of AdjListGraph.
// This is our constructor function.
func NewGraph() *AdjListGraph {
    return &AdjListGraph{
        vertices: make(map[interface{}]*GraphVertex),
    }
}

// AddVertex adds a new vertex to the graph with the given key.
// If the vertex already exists, it does nothing.
func (g *AdjListGraph) AddVertex(key interface{}) {
    if _, exists := g.vertices[key]; !exists {
        g.vertices[key] = NewGraphVertex(key)
    }
}

// AddEdge adds a directed edge from the 'from' vertex to the 'to' vertex with the given weight.
// It creates the vertices if they don't exist.
func (g *AdjListGraph) AddEdge(from, to interface{}, weight float64) {
    // Ensure both vertices exist
    g.AddVertex(from)
    g.AddVertex(to)
    // Add the edge
    g.vertices[from].Edges[to] = weight
}

// RemoveVertex removes the vertex with the given key and all its edges.
func (g *AdjListGraph) RemoveVertex(key interface{}) {
    delete(g.vertices, key)
    // Remove all edges pointing to this vertex
    for _, vertex := range g.vertices {
        delete(vertex.Edges, key)
    }
}

// RemoveEdge removes the edge from the 'from' vertex to the 'to' vertex.
func (g *AdjListGraph) RemoveEdge(from, to interface{}) {
    if vertex, exists := g.vertices[from]; exists {
        delete(vertex.Edges, to)
    }
}

// GetVertices returns a slice of all vertex keys in the graph.
func (g *AdjListGraph) GetVertices() []interface{} {
    keys := make([]interface{}, 0, len(g.vertices))
    for k := range g.vertices {
        keys = append(keys, k)
    }
    return keys
}

// GetNeighbors returns a slice of all vertices that are connected to the given vertex.
func (g *AdjListGraph) GetNeighbors(vertex interface{}) []interface{} {
    if v, exists := g.vertices[vertex]; exists {
        neighbors := make([]interface{}, 0, len(v.Edges))
        for k := range v.Edges {
            neighbors = append(neighbors, k)
        }
        return neighbors
    }
    return nil
}

// PrintGraph prints a representation of the graph for debugging purposes.
func (g *AdjListGraph) PrintGraph() {
    for key, vertex := range g.vertices {
        fmt.Printf("Vertex %v -> ", key)
        for neighbor, weight := range vertex.Edges {
            fmt.Printf("(%v: %v) ", neighbor, weight)
        }
        fmt.Println()
    }
}