package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numTwoLettersSame := uint(0)
	numThreeLettersSame := uint(0)

	for scanner.Scan() {
		id := scanner.Text()

		lettersSeen := make(map[rune]uint)

		for _, runeValue := range id {
			lettersSeen[runeValue]++
		}

		twoLettersSame := false
		threeLettersSame := false

		for _, count := range lettersSeen {
			if count == 2 && !twoLettersSame {
				numTwoLettersSame++
				twoLettersSame = true
				continue
			}

			if count == 3 && !threeLettersSame {
				numThreeLettersSame++
				threeLettersSame = true
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(numTwoLettersSame * numThreeLettersSame)
}
