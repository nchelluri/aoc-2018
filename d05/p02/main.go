package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	polymer := scanner.Text()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	unitsMap := make(map[string]struct{})
	for _, unit := range polymer {
		if unit < 'a' {
			unit += 32
		}
		unitsMap[strings.ToUpper(string(unit))] = struct{}{}
	}

	minLength := len(polymer)
	for unit := range unitsMap {
		candidatePolymer := strings.Replace(polymer, unit, "", -1)
		candidatePolymer = strings.Replace(candidatePolymer, strings.ToLower(unit), "", -1)

		len := len(react(candidatePolymer))
		if len < minLength {
			minLength = len
		}
	}

	fmt.Println(minLength)
}

func react(polymer string) string {
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

	return polymer
}
