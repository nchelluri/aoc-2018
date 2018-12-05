package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	polymer := scanner.Text()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	changed := true
	for changed {
		if len(polymer) < 2 {
			break
		}

		changed = false

		for i, j := 0, 1; j < len(polymer); i, j = i+1, j+1 {
			if polymer[i]+32 == polymer[j] || polymer[j]+32 == polymer[i] {
				suffix := ""
				if j+1 < len(polymer) {
					suffix = polymer[j+1:]
				}
				polymer = polymer[:i] + suffix
				changed = true
			}
		}
	}

	fmt.Println(len(polymer))
}
