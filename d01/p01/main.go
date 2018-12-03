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
		freqChange, err := strconv.Atoi(scanner.Text())

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
