package graph_test

// import (
// 	"os"
// 	"reflect"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/relvox/iridescence_go/graph"
// )

// func Test_Graph_AddEdge(t *testing.T) {
// 	g := make(graph.Graph[string, int])
// 	g.AddEdge("a", "b", 1, true)

// 	if _, ok := g["a"]["b"]; !ok {
// 		t.Errorf("Expected edge from a to b, but it was not found")
// 	}

// 	if _, ok := g["b"]["a"]; !ok {
// 		t.Errorf("Expected edge from b to a, but it was not found")
// 	}
// }

// func Test_Graph_FromDotFile(t *testing.T) {
// 	g, err := graph.FromDot[int]("./test_data/test.dot")
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}

// 	expected := make(graph.Graph[string, int])
// 	expected.AddEdge("a", "b", 1, true)

// 	if !reflect.DeepEqual(g, expected) {
// 		t.Errorf("Graph was incorrect, got: %v, want: %v.", g, expected)
// 	}
// }

// func Test_Graph_ToDotFile(t *testing.T) {
// 	g := make(graph.Graph[string, int])
// 	g.AddEdge("a", "b", 1, true)

// 	err := g.ToDot("./test_data/output.dot")
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}

// 	// Read the file and check its contents
// 	content, err := os.ReadFile("./test_data/output.dot")
// 	if err != nil {
// 		t.Errorf("Expected no error, but got: %v", err)
// 	}

// 	expected := "graph {\n  \"a\" -- \"b\" [label=\"1\"];\n  \"b\" -- \"a\" [label=\"1\"];\n}\n"
// 	if string(content) != expected {
// 		t.Errorf("File content was incorrect, got: %s, want: %s.", content, expected)
// 	}
// }

// type Challenge struct {
// 	Challenge string `json:"challenge"`
// 	Stat      string `json:"stat"`
// 	Value     int    `json:"value"`
// }

// func Test_Graph_SaveAndLoadGraph(t *testing.T) {
// 	g := &graph.Graph[string, Challenge]{
// 		"Forest Entrance": {
// 			"Deep Forest": {"Fight off a pack of wolves", "Strength", 5},
// 		},
// 		"Deep Forest": {
// 			"Abandoned Hut": {"Solve a riddle carved into a tree", "Intelligence", 7},
// 		},
// 	}

// 	filename := "test_graph.json"
// 	defer os.Remove(filename) // clean up file after test

// 	err := graph.SaveGraphToJSON(g, filename)
// 	assert.NoError(t, err, "SaveGraphToFile should not return an error")

// 	loadedGraph, err := graph.LoadGraphFromJSON[string, Challenge](filename)
// 	assert.NoError(t, err, "LoadGraphFromFile should not return an error")

// 	assert.Equal(t, g, loadedGraph, "The loaded graph should be equal to the original graph")
// }

// func Test_Graph_LoadGraphFromFile_NonexistentFile(t *testing.T) {
// 	_, err := graph.LoadGraphFromJSON[string, Challenge]("nonexistent_file.json")
// 	assert.Error(t, err, "LoadGraphFromFile should return an error for a nonexistent file")
// }
