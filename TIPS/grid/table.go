package grid

import (
	"fmt"
	"io"
	"quarto-semestri/TIPS/graph"
)

const (
	start  = "\033[1;34m%s\033[0m"
	end    = "\033[1;36m%s\033[0m"
	node   = "\033[1;33m%s\033[0m"
	banned = "\033[1;31m%s\033[0m"
	path   = "\033[1;235m%s\033[0m"
)

// Grid grid implementation
type Grid struct {
	graph.ItemGraph
	size       int
	start, end [2]int
	banned     [][2]int
	path       [][2]int
}

// New create new grid
func New(size int, start, end [2]int, banned [][2]int) (grid *Grid, startNode, endNode *graph.Node) {
	grid = new(Grid)
	grid.end = end
	grid.start = start
	grid.size = size
	grid.banned = banned
	startNode, endNode = new(graph.Node), new(graph.Node)
	Table := new([][]*graph.Node)
	for i := 0; i < size; i++ {
		*Table = append(*Table, []*graph.Node{})
		for j := 0; j < size; j++ {
			n := &graph.Node{Value: graph.Item{Cord: [2]int{j, i}, Access: !grid.IsBanned([2]int{j, i})}}
			if [2]int{j, i} == start {
				startNode = n
			}
			if [2]int{j, i} == end {
				endNode = n
			}
			grid.AddNode(n)
			(*Table)[i] = append((*Table)[i], n)
			if i == 0 && j == 0 {
				continue
			}
			if j == 0 {
				grid.AddEdge((*Table)[i-1][j], n)
				continue
			}
			if i == 0 {
				grid.AddEdge((*Table)[i][j-1], n)
				continue
			}
			grid.AddEdge((*Table)[i][j-1], n)
			grid.AddEdge((*Table)[i-1][j], n)
		}
	}
	return grid, startNode, endNode
}

// Print print grid in stdout
func (g *Grid) Print(w io.Writer) error {
	fmt.Print("\n")
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			cord := [2]int{j, i}
			switch {
			case cord == g.start:
				if _, err := fmt.Fprintf(w, start, "Q0"); err != nil {
					return err
				}
			case cord == g.end:
				if _, err := fmt.Fprintf(w, end, "Q1"); err != nil {
					return err
				}
			case g.IsBanned(cord):
				if _, err := fmt.Fprintf(w, banned, "x "); err != nil {
					return err
				}
			case g.IsPath(cord):
				if _, err := fmt.Fprintf(w, path, "+ "); err != nil {
					return err
				}
			default:
				if _, err := fmt.Fprintf(w, node, "+ "); err != nil {
					return err
				}
			}
		}
		fmt.Print("\n")
	}
	return nil
}

// IsBanned return true if node in this cord banned
func (g *Grid) IsBanned(cord [2]int) bool {
	for _, b := range g.banned {
		if b == cord {
			return true
		}
	}
	return false
}

// IsPath return true if node in this cord path
func (g *Grid) IsPath(cord [2]int) bool {
	for _, b := range g.path {
		if b == cord {
			return true
		}
	}
	return false
}

// AddPath path on node
func (g *Grid) AddPath(cord [2]int) {
	g.path = append(g.path, cord)
}
