package graph

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
