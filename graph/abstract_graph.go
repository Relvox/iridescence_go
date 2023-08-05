package graph

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/relvox/iridescence_go/utils"
)

type AbstractGraph[Node comparable, Edge any] map[Node]map[Node]Edge

func NewAbstractGraph[Node comparable, Edge any]() AbstractGraph[Node, Edge] {
	return make(AbstractGraph[Node, Edge])
}

func (g AbstractGraph[Node, Edge]) AddEdge(a, b Node, edge Edge) {
	if _, ok := g[a]; !ok {
		g[a] = make(map[Node]Edge)
	}
	g[a][b] = edge
}

func (g AbstractGraph[Node, Edge]) AddBothEdges(a, b Node, edge Edge) {
	g.AddEdge(a, b, edge)
	g.AddEdge(b, a, edge)
}

func (g AbstractGraph[Node, Edge]) AllNodes() []Node {
	return utils.Keys(g)
}

func (g AbstractGraph[Node, Edge]) EdgesNeighbors(a Node) ([]Edge, []Node) {
	var resultEs []Edge
	var resultNs []Node
	for n, e := range g[a] {
		resultEs = append(resultEs, e)
		resultNs = append(resultNs, n)
	}
	return resultEs, resultNs
}

func FromDGDot[Node comparable, Edge any](path string, convert func(string) Node) (AbstractGraph[Node, Edge], error) {
	return FromDot[Node, Edge](path, "->", convert)
}

func FromUDGDot[Node comparable, Edge any](path string, convert func(string) Node) (AbstractGraph[Node, Edge], error) {
	return FromDot[Node, Edge](path, "--", convert)
}

func FromDot[Node comparable, Edge any](path, edgeString string, convert func(string) Node) (AbstractGraph[Node, Edge], error) {
	g := make(AbstractGraph[Node, Edge])
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
	graph AbstractGraph[Node, Edge],
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
) (AbstractGraph[Node, Edge], error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var graph AbstractGraph[Node, Edge]
	err = json.Unmarshal(data, &graph)
	return graph, err
}
