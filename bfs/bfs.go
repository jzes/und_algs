package bfs

import (
	"fmt"

	"github.com/jzes/und_algs/queue"
)

type Graph struct {
	Item  any
	Nodes []*Graph
}

func Bfs(g *Graph, s any) any {
	qSearch := queue.Queue{}
	qSearch.Enqueue(g)
	searched := map[any]bool{}

	for {
		rNode, mpt := qSearch.Dequeue()
		if mpt {
			break
		}

		node, ok := rNode.(*Graph)
		if !ok {
			fmt.Println("conversion error")
			continue
		}

		if searched[node.Item] {
			continue
		}

		if node.Item == s {
			return s
		}
		for _, subNode := range node.Nodes {
			qSearch.Enqueue(subNode)
		}

		searched[node.Item] = true
	}
	return nil
}
