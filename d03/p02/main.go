package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fabric := [][]int{}

	const fabricDims = 1000

	for i := 0; i < fabricDims; i++ {
		fabricRow := []int{}

		for j := 0; j < fabricDims; j++ {
			fabricRow = append(fabricRow, 0)
		}

		fabric = append(fabric, fabricRow)
	}

	scanner := bufio.NewScanner(os.Stdin)

	initiallyGoodClaims := make(map[int][]int)
	claimRE := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`) // #1 @ 49,222: 19x20
	for scanner.Scan() {
		claim := scanner.Text()
		valueStrs := claimRE.FindStringSubmatch(claim)

		if valueStrs == nil || len(valueStrs) != 6 {
			panic("error parsing claim")
		}

		values := []int{}

		for _, valueStr := range valueStrs[1:] {
			value, err := strconv.Atoi(valueStr)

			if err != nil {
				panic(err)
			}

			values = append(values, value)
		}

		claimNum := values[0]

		left := values[1]
		top := values[2]
		width := values[3]
		height := values[4]

		claimAllGood := true

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if fabric[i+top][j+left] > 0 {
					claimAllGood = false
				}

				fabric[i+top][j+left]++
			}
		}

		if claimAllGood {
			initiallyGoodClaims[claimNum] = values[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

FinalClaimCheck:
	for claimNum, claim := range initiallyGoodClaims {
		left := claim[0]
		top := claim[1]
		width := claim[2]
		height := claim[3]

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if fabric[i+top][j+left] > 1 {
					continue FinalClaimCheck
				}
			}
		}

		fmt.Println(claimNum)
	}
}

func printFabric(fabric [][]int) {
	for _, fabricRow := range fabric {
		for _, fabricPoint := range fabricRow {
			output := "."
			if fabricPoint > 1 {
				output = "X"
			} else if fabricPoint > 0 {
				output = "#"
			}
			fmt.Printf("%s", output)
		}
		fmt.Println()
	}
}
