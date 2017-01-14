package graph

import "fmt"

type Edge struct {
	v, len int
}

type Node struct {
	id      int
	visited bool
	edges   []*Edge
}
type Graph struct {
	nodes map[int]*Node
}

func (g *Graph) ShowGraph() {
	for k, v := range g.nodes {
		fmt.Printf("Node %d (id: %d):\n", k, v.id)
		v.showEdges()
	}
}

func (n *Node) showEdges() {
	for _, v := range n.edges {
		fmt.Printf("\tHead ID: %d\tLength: %d\n", v.v, v.len)
	}
	fmt.Printf("\n")
}

func New() *Graph {
	var g Graph
	g.nodes = make(map[int]*Node)
	return &g
}

func newNode(id int) *Node {
	var n Node
	n.id = id
	n.edges = make([]*Edge, 0)
	return &n
}

func (n *Node) AddEdge(v, len int) {
	e := &Edge{v: v, len: len}
	n.edges = append(n.edges, e)
}

func (g *Graph) AddNode(id int) *Node {
	n := newNode(id)
	g.nodes[n.id] = n
	return n
}

func (g *Graph) resetVisited() {
	for _, n := range g.nodes {
		n.visited = false
	}
}
