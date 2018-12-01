package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	freq := 0

	scanner := bufio.NewScanner(os.Stdin)

	freqChanges := []int{}

	foundFreqs := make(map[int]struct{})

	for scanner.Scan() {
		candidate := scanner.Text()

		if candidate[0] == '+' {
			candidate = candidate[1:]
		}

		freqChange, err := strconv.Atoi(candidate)

		if err != nil {
			panic(err)
		}

		freqChanges = append(freqChanges, freqChange)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := 0; i < len(freqChanges); i = (i + 1) % len(freqChanges) {
		freq += freqChanges[i]

		_, exists := foundFreqs[freq]
		if exists {
			break
		}

		foundFreqs[freq] = struct{}{}
	}

	fmt.Println(freq)
}
