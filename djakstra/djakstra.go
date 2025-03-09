package djakstra

import (
	"fmt"
	"math"
)

type Node map[any]map[any]int

func (n Node) Adjacents(key any) map[any]int {
	return n[key]
}

type Costs map[any]int

type Processed map[any]bool

func Process(graf Node) Costs {
	startNode := GetStart(graf)
	costs := NewCosts(graf, startNode)
	processed := Processed{}

	cheaperNode := GetLower(costs, processed)
	for cheaperNode != nil {
		adjacents := graf[cheaperNode]
		for currentNode, currentCost := range adjacents {
			newCost := costs[cheaperNode] + currentCost
			if prevCost, ok := costs[currentNode]; !ok || newCost < prevCost {
				costs[currentNode] = newCost
			}
		}
		processed[cheaperNode] = true
		cheaperNode = GetLower(costs, processed)
		fmt.Println("cheaper Node", cheaperNode)
	}

	return costs
}

func GetStart(graff Node) any {
	childs := map[any]bool{}
	for _, adjacents := range graff {
		for childNode := range adjacents {
			childs[childNode] = true
		}
	}
	for k := range graff {
		isInChild := childs[k]
		if !isInChild {
			return k
		}
	}
	return nil
}

func NewCosts(graf Node, start any) Costs {
	costs := Costs{}
	for node, adj := range graf[start] {
		costs[node] = adj
	}
	costs[start] = 0
	return costs
}

func GetLower(costs Costs, processed Processed) any {
	lower := math.MaxInt
	var lowerNode any

	for node, cost := range costs {
		if cost < lower && !processed[node] {
			lower = cost
			lowerNode = node
		}
	}
	return lowerNode
}
