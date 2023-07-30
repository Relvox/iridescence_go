package graph

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type UDGraph[Node comparable, Edge any] map[Node]map[Node]Edge

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

func (g UDGraph[Node, Edge]) ToDot(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.WriteString("graph {\n")
	if err != nil {
		return err
	}

	for a, edges := range g {
		for b, edge := range edges {
			edgeJson, err := json.Marshal(edge)
			if err != nil {
				return err
			}
			_, err = w.WriteString(fmt.Sprintf("  \"%v\" -- \"%v\" [label=\"%s\"];\n", a, b, string(edgeJson)))
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
