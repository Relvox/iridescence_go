package graph

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Graph[Node comparable, Edge any] interface {
	AddEdge(a, b Node, e Edge)
	AddBothEdges(a, b Node, e Edge)
	AllNodes() []Node
	EdgesNeighbors(a Node) ([]Edge, []Node)
}

func toDot[Node comparable, Edge any](g Graph[Node, Edge], header, edgeStr string) (string, error) {
	w := strings.Builder{}
	w.WriteString(fmt.Sprintf("%s {\n", header))
	for _, a := range g.AllNodes() {
		edges, nodes := g.EdgesNeighbors(a)
		for i, edge := range edges {
			b := nodes[i]
			edgeJson, err := json.Marshal(edge)
			if err != nil {
				return "", err
			}
			w.WriteString(fmt.Sprintf("  \"%v\" %s \"%v\" [label=\"%s\"];\n", a, edgeStr, b, string(edgeJson)))
		}
	}
	w.WriteString("\n}")
	return w.String(), nil
}

func ToDotDirected[Node comparable, Edge any](g Graph[Node, Edge]) (string, error) {
	return toDot(g, "digraph", "->")
}

func ToDotUndirected[Node comparable, Edge any](g Graph[Node, Edge]) (string, error) {
	return toDot(g, "graph", "--")
}
