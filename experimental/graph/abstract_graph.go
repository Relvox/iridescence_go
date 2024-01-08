package graph

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"golang.org/x/exp/maps"
)

type AdjGraph[Node comparable, Edge any] map[Node]map[Node]Edge

func NewAdjGraph[Node comparable, Edge any]() AdjGraph[Node, Edge] {
	return make(AdjGraph[Node, Edge])
}

func (g AdjGraph[Node, Edge]) AddEdge(a, b Node, edge Edge) {
	if _, ok := g[a]; !ok {
		g[a] = make(map[Node]Edge)
	}
	g[a][b] = edge
}

func (g AdjGraph[Node, Edge]) AddBothEdges(a, b Node, edge Edge) {
	g.AddEdge(a, b, edge)
	g.AddEdge(b, a, edge)
}

func (g AdjGraph[Node, Edge]) AllNodes() []Node {
	return maps.Keys(g)
}

func (g AdjGraph[Node, Edge]) EdgesNeighbors(a Node) ([]Edge, []Node) {
	var resultEs []Edge
	var resultNs []Node
	for n, e := range g[a] {
		resultEs = append(resultEs, e)
		resultNs = append(resultNs, n)
	}
	return resultEs, resultNs
}

func FromDotDigraph[Node comparable, Edge any](path string, convert func(string) Node) (AdjGraph[Node, Edge], error) {
	return FromDot[Node, Edge](path, "->", convert)
}

func FromDotGraph[Node comparable, Edge any](path string, convert func(string) Node) (AdjGraph[Node, Edge], error) {
	return FromDot[Node, Edge](path, "--", convert)
}

func FromDot[Node comparable, Edge any](path, edgeString string, convert func(string) Node) (AdjGraph[Node, Edge], error) {
	g := make(AdjGraph[Node, Edge])
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(fmt.Sprintf(`^\s*"([\w' ]*)" %s "([\w' ]*)" \[label="(.*)"\]\s*$`, edgeString))
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			matches := r.FindStringSubmatch(scanner.Text())
			var edge Edge
			err := json.Unmarshal([]byte(matches[3]), &edge)
			if err != nil {
				return nil, err
			}

			g.AddEdge(convert(matches[1]), convert(matches[2]), edge)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return g, nil
}

func SaveGraphToJSON[Node comparable, Edge any](
	graph AdjGraph[Node, Edge],
	filename string,
) error {
	data, err := json.MarshalIndent(graph, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadGraphFromJSON[Node comparable, Edge any](
	filename string,
) (AdjGraph[Node, Edge], error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var graph AdjGraph[Node, Edge]
	err = json.Unmarshal(data, &graph)
	return graph, err
}
