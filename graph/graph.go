package graph

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/relvox/iridescence/utils"
)

type Graph map[string]utils.Set[string]

func (m Graph) Edge(a, b string, bidirectional bool) Graph {
	var result Graph = make(Graph)

	for k, v := range m {
		result[k] = make(utils.Set[string])
		for p := range v {
			result[k][p] = utils.U
		}
	}
	if _, ok := result[a]; !ok {
		result[a] = make(utils.Set[string])
	}
	result[a][b] = utils.U
	if bidirectional {
		if _, ok := result[b]; !ok {
			result[b] = make(utils.Set[string])
		}
		result[b][a] = utils.U
	}
	return result
}

func (m Graph) Path(from, to string, dist int) (bool, []string) {
	if from == to {
		return true, []string{from}
	}
	var queue []string = []string{from}
	var prevs map[string]string = make(map[string]string)
	dists := map[string]int{
		from: 0,
	}
	visited := utils.Set[string]{
		from: utils.U,
	}

	for len(queue) > 0 {
		item := queue[0]
		if item == to {
			break
		}
		queue = queue[1:]
		if dists[item] >= dist {
			continue
		}
		for neb := range m[item] {
			if _, ok := visited[neb]; !ok {
				visited[neb] = utils.U
				prevs[neb] = item
				dists[neb] = dists[item] + 1
				queue = append(queue, neb)
			}
		}
	}
	if _, ok := prevs[to]; !ok {
		return false, nil
	}
	var path []string = make([]string, dists[to]+1)
	path[len(path)-1] = to

	for i := len(path) - 2; i >= 0; i-- {
		path[i] = prevs[path[i+1]]
	}
	return true, path
}

func FromDot(path string) Graph {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(`^  "([\w' ]*)" -- "([\w' ]*)".*$`) // this can also be a regex

	if err != nil {
		log.Fatal(err)
	}

	var result Graph = make(Graph)
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			matches := r.FindStringSubmatch(scanner.Text())
			result = result.Edge(matches[1], matches[2], true)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
