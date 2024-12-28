// interfaces.go
package dataStructures

// Graph defines the interface for graph operations
type Graph interface {
    AddVertex(key interface{})
    AddEdge(from, to interface{}, weight float64)
    RemoveVertex(key interface{})
    RemoveEdge(from, to interface{})
    GetVertices() []interface{}
    GetNeighbors(vertex interface{}) []interface{}
}