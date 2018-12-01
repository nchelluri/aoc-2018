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

	for scanner.Scan() {
		candidate := scanner.Text()

		if candidate[0] == '+' {
			candidate = candidate[1:]
		}

		freqChange, err := strconv.Atoi(candidate)

		if err != nil {
			panic(err)
		}

		freq += freqChange
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(freq)
}
