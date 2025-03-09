package djakstra

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		graph    Node
		expected Costs
	}{
		{
			name: "Simple Graph",
			graph: Node{
				"A": {"B": 6, "C": 2},
				"B": {"D": 1},
				"C": {"B": 3, "D": 5},
				"D": {},
			},
			expected: Costs{"A": 0, "B": 5, "C": 2, "D": 6},
		},
		{
			name: "Single Node",
			graph: Node{
				"A": {},
			},
			expected: Costs{"A": 0},
		},
		{
			name: "Disconnected Nodes",
			graph: Node{
				"A": {"B": 1},
				"C": {"D": 2},
			},
			expected: Costs{"A": 0, "B": 1}, // C e D não são alcançáveis de A
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Process(tt.graph)
			for node, cost := range tt.expected {
				if result[node] != cost {
					t.Errorf("For node %s, expected %d but got %d", node, cost, result[node])
				}
			}
		})
	}
}
