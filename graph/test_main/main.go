package main

import (
	"fmt"

	"github.com/relvox/iridescence_go/graph"
)

func main() {
	g := graph.DGraph[int, int]{}
	fmt.Println(g)
}
