package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type worker struct {
	workingOn string
	freeAt    int
}

const numWorkers = 5

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

	var time int
	workers := [numWorkers]worker{}
	order := ""

	for time = 0; ; time++ {
		for w, worker := range workers {
			if worker.freeAt == time {
				if worker.workingOn != "" {
					delete(digraph, worker.workingOn)
					for node, edges := range digraph {
						for i, edge := range edges {
							if edge == worker.workingOn {
								digraph[node] = append(edges[:i], edges[i+1:]...)
							}
						}
					}
					order += worker.workingOn
				}
			}

			if worker.freeAt <= time {
				next := findNext(digraph, workers)

				if next != "" {
					workers[w].workingOn = next
					workers[w].freeAt = time + amountOfTimeForStep(next)
				} else {
					workers[w].workingOn = ""
				}
			}
		}

		if len(digraph) == 0 {
			break
		}
	}

	fmt.Printf("%d\t\t\t%s\n", time, order)
}

func findNext(digraph map[string][]string, workers [numWorkers]worker) string {
	nexts := []string{}

Outer:
	for node, edges := range digraph {
		if len(edges) == 0 {
			for _, worker := range workers {
				if worker.workingOn == node {
					continue Outer
				}
			}

			nexts = append(nexts, node)
		}
	}

	if len(nexts) < 1 {
		return ""
	}

	sort.Strings(nexts)

	return nexts[0]
}

func amountOfTimeForStep(step string) int {
	runes := []rune(step)

	if len(runes) != 1 {
		panic("more runes than expected in step string")
	}

	return int(runes[0]) - 65 + 1 + 60
}
