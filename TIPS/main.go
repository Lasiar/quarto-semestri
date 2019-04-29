package main

import (
	"fmt"
	"quarto-semestri/TIPS/graph"
)

func main() {
	g := fillGraph()
	g.String()
	g.Traverse(func(n *graph.Node) {
		fmt.Printf("%v\n", n)
	})
}

func fillGraph() *graph.ItemGraph {
	g := new(graph.ItemGraph)
	nA := graph.Node{Value: "A"}
	nB := graph.Node{Value: "B"}
	nC := graph.Node{Value: "C"}
	nD := graph.Node{Value: "D"}
	nE := graph.Node{Value: "E"}
	nF := graph.Node{Value: "F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
	return g
}
