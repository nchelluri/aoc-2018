package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ids := []string{}

	for scanner.Scan() {
		id := scanner.Text()
		ids = append(ids, id)
	}

	var id1, id2 string

OuterLoop:
	for i, candidateID1 := range ids {
	InnerLoop:
		for j, candidateID2 := range ids {
			if i == j {
				continue
			}

			numDiffChars := uint(0)

			for k := range candidateID1 {
				if candidateID1[k] != candidateID2[k] {
					numDiffChars++

					if numDiffChars > 1 {
						continue InnerLoop
					}
				}
			}

			if numDiffChars == 1 {
				id1 = candidateID1
				id2 = candidateID2
				break OuterLoop
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := range id1 {
		if id1[i] == id2[i] {
			fmt.Printf("%s", string(id1[i]))
		}
	}

	fmt.Println()
}
