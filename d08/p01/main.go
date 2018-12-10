package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	children []node
	metadata []int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	valueStrs := strings.Split(input, " ")

	values := []int{}
	for _, valueStr := range valueStrs {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	root, _ := parseNodes(1, 0, values)
	fmt.Println(metadataSum(root[0]))
}

func parseNodes(numNodes, startIndex int, values []int) ([]node, int) {
	pos := startIndex

	var foundNodes []node

	for foundNodes = []node{}; len(foundNodes) < numNodes; {
		numChildren := values[pos]
		numMetadata := values[pos+1]

		var children []node
		var metadata []int
		if numChildren == 0 {
			metadata = values[pos+2 : pos+2+numMetadata]
			pos += 2 + numMetadata
		} else {
			newChildren, metadataIndex := parseNodes(numChildren, pos+2, values)
			children = newChildren
			pos = metadataIndex + numMetadata
			metadata = values[metadataIndex:pos]
		}

		n := node{
			children: children,
			metadata: metadata,
		}

		foundNodes = append(foundNodes, n)
	}

	return foundNodes, pos
}

func metadataSum(root node) int {
	sum := 0
	for _, metadata := range root.metadata {
		sum += metadata
	}

	for _, child := range root.children {
		sum += metadataSum(child)
	}

	return sum
}
