package exer

import (
	"fmt"
	"math/rand"
)

type edges struct {
	src    *graphNode
	dst    *graphNode
	weight int
}

type graphNode struct {
	value int
}

type root struct {
	rootNode *graphNode
}

func BuildGraph() []*edges {
	allNodes := []*graphNode{}
	Graph := []*edges{}

	for i := 0; i < 7; i++ {
		node := createNodes(i)
		allNodes = append(allNodes, node)
	}

	// for now only allowing atmost 2 neighbors per node and atleast 0
	for j := range allNodes {
		neighbor1 := &graphNode{
			value: rand.Intn(7),
		}
		neighbor2 := &graphNode{
			value: rand.Intn(7),
		}
		neighbors := []*graphNode{neighbor1, neighbor2}
		Graph = append(Graph, createEdges(allNodes[j], neighbors)...)
	}
	return Graph
}

func createNodes(nodeValue int) *graphNode {
	allNodes := &graphNode{
		value: nodeValue,
	}
	return allNodes
}

func createEdges(node *graphNode, neighbors []*graphNode) []*edges {
	tempGraph := []*edges{}
	for item := range neighbors {
		edge := &edges{
			src:    node,
			dst:    neighbors[item],
			weight: rand.Intn(15),
		}
		tempGraph = append(tempGraph, edge)
	}
	return tempGraph
}

var visited []edges

func AStar(graph []*edges, dst *graphNode, curr *graphNode) {
	tempNode := &edges{}

	var min int
	fmt.Println("travelling from", curr.value)
	if curr.value == dst.value {
		fmt.Println("destination reached")
		return
	}
	// fmt.Println("val", curr.value)
	for edge := range graph {
		if graph[edge].src.value == curr.value {
			// fmt.Println("76:", graph[edge].src.value)
			if len(visited) == 0 {
				min = graph[edge].weight
				// fmt.Println("79", graph[edge].src.value)
				visited = append(visited, *graph[edge])
				// fmt.Println("82", visited)
				// fmt.Println("81", visited)
			} else if graph[edge].weight < min {
				// ignore edges whose weight is greater than the current minimum
				// fmt.Println("85", graph[edge])
				min = graph[edge].weight
				visited = append(visited, *graph[edge])
				// fmt.Println("88", min)
				tempNode = graph[edge]
			} else {
				tempNode = nil
				continue
			}
		}
	}
	// fmt.Println("98:", tempNode.dst.value)
	if tempNode == nil {
		fmt.Println("cant travel any further")
		return
	} else {
		curr = tempNode.dst
		fmt.Println("travelled to", curr.value)
		AStar(graph, &graphNode{value: 6}, curr)
	}
}

func view(graph []*edges) {
	for edge := range graph {
		fmt.Println("\t", graph[edge].weight)
		fmt.Println(graph[edge].src.value, "--------------------->", graph[edge].dst.value)
	}
}
func Build() {
	g := BuildGraph()
	// fmt.Println(g[0].src.value)
	// view(g)
	AStar(g, &graphNode{value: 6}, &graphNode{value: 1})
}
