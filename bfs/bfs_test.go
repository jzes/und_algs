package bfs

import (
	"testing"
)

func TestBfs(t *testing.T) {
	// Criando um grafo de exemplo
	nodeE := &Graph{Item: "E", Nodes: []*Graph{}}
	nodeD := &Graph{Item: "D", Nodes: []*Graph{nodeE}}
	nodeC := &Graph{Item: "C", Nodes: []*Graph{nodeD}}
	nodeB := &Graph{Item: "B", Nodes: []*Graph{nodeC}}
	nodeA := &Graph{Item: "A", Nodes: []*Graph{nodeB, nodeD}}

	tests := []struct {
		name     string
		start    *Graph
		target   any
		expected any
	}{
		{"Find existing node", nodeA, "C", "C"},
		{"Find root node", nodeA, "A", "A"},
		{"Find last node", nodeA, "E", "E"},
		{"Node not in graph", nodeA, "X", nil},
		{"Disconnected graph", nodeE, "D", nil},
		{"Graph with cycle", nodeD, "D", "D"}, // Testando busca no próprio nó
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Bfs(tt.start, tt.target)
			if result != tt.expected {
				t.Errorf("Bfs(%v, %v) = (%v), expected (%v)",
					tt.start.Item, tt.target, result, tt.expected)
			}
		})
	}
}
