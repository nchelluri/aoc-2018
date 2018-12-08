package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Step C must be finished before step A can begin.
	stepRE := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin\.`)

	digraph := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		matches := stepRE.FindStringSubmatch(line)

		if len(matches) != 3 {
			panic("error parsing step line")
		}

		digraph[matches[2]] = append(digraph[matches[2]], matches[1])
		if digraph[matches[1]] == nil {
			digraph[matches[1]] = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	order := ""

	for len(digraph) > 0 {
		next := findNext(digraph)
		order += next
		delete(digraph, next)

		for node, edges := range digraph {
			for i, edge := range edges {
				if edge == next {
					digraph[node] = append(edges[:i], edges[i+1:]...)
				}
			}
		}
	}

	fmt.Println(order)
}

func findNext(digraph map[string][]string) string {
	candidates := []string{}

	for node, edges := range digraph {
		if len(edges) == 0 {
			candidates = append(candidates, node)
		}
	}

	sort.Strings(candidates)

	return candidates[0]
}
