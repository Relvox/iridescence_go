package graph

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type DGraph[Node comparable, Edge any] AbstractGraph[Node, Edge]

func NewDGraph[Node comparable, Edge any]() DGraph[Node, Edge] {
	return make(DGraph[Node, Edge])
}

func (g DGraph[Node, Edge]) AddEdge(a, b Node, edge Edge) {
	if _, ok := g[a]; !ok {
		g[a] = make(map[Node]Edge)
	}
	g[a][b] = edge
}

func (g DGraph[Node, Edge]) AddBothEdges(a, b Node, edge Edge) {
	g.AddEdge(a, b, edge)
	g.AddEdge(b, a, edge)
}

func (g DGraph[Node, Edge]) ToDot(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString("digraph {\n")
	if err != nil {
		return err
	}
	for a, edges := range g {
		for b, edge := range edges {
			edgeJson, err := json.Marshal(edge)
			if err != nil {
				return err
			}
			_, err = w.WriteString(fmt.Sprintf("  \"%v\" -> \"%v\" [label=\"%s\"];\n", a, b, string(edgeJson)))
			if err != nil {
				return err
			}
		}
	}

	_, err = w.WriteString("}\n")
	if err != nil {
		return err
	}

	return w.Flush()
}
