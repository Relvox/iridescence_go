package graph

import "github.com/relvox/iridescence_go/utils"

type UDGraph[Node comparable, Edge any] AbstractGraph[Node, Edge]

func NewUDGraph[Node comparable, Edge any]() UDGraph[Node, Edge] {
	return make(UDGraph[Node, Edge])
}

func (g UDGraph[Node, Edge]) AddEdge(a, b Node, edge Edge) {
	if _, ok := g[a]; !ok {
		g[a] = make(map[Node]Edge)
	}
	g[a][b] = edge
}

func (g UDGraph[Node, Edge]) AddBothEdges(a, b Node, edge Edge) {
	g.AddEdge(a, b, edge)
	g.AddEdge(b, a, edge)
}

func (g UDGraph[Node, Edge]) AllNodes() []Node {
	return utils.Keys(g)
}

func (g UDGraph[Node, Edge]) EdgesNeighbors(a Node) ([]Edge, []Node) {
	var resultEs []Edge
	var resultNs []Node
	for n, e := range g[a] {
		resultEs = append(resultEs, e)
		resultNs = append(resultNs, n)
	}
	return resultEs, resultNs
}
