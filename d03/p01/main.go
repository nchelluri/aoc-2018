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

	claimRE := regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`) // #1 @ 49,222: 19x20
	for scanner.Scan() {
		claim := scanner.Text()
		valueStrs := claimRE.FindStringSubmatch(claim)

		if valueStrs == nil || len(valueStrs) != 5 {
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

		left := values[0]
		top := values[1]
		width := values[2]
		height := values[3]

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				fabric[i+top][j+left]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	numFabricPointsDoublyClaimed := 0
	for _, fabricRow := range fabric {
		for _, fabricPoint := range fabricRow {
			if fabricPoint > 1 {
				numFabricPointsDoublyClaimed++
			}
		}
	}

	fmt.Println(numFabricPointsDoublyClaimed)
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
