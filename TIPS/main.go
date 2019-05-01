package main

import (
	"fmt"
	"quarto-semestri/TIPS/graph"
	"quarto-semestri/TIPS/grid"
)

func main() {
	g, start, end := grid.New(11, [2]int{2, 2}, [2]int{3, 3}, [][2]int{{2, 3}, {2, 9}, {3, 2}})
	find := false
	g.TraverseWithStart(start, func(node *graph.Node) {
		if find {
			return
		}

		if *end == *node {
			find = true
			fmt.Println("FIND")
		}
		fmt.Println(node, end)
		g.AddPath(node.Value.Cord)
		printLine()
		g.Print()
	})
}

func printLine() { fmt.Println("----------------------------------") }
