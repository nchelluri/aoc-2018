package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

func main() {
	points := []point{}
	scanner := bufio.NewScanner(os.Stdin)
	coordLineRE := regexp.MustCompile(`(\d+), (\d+)`)

	maxX, maxY := 0, 0
	minX, minY := -1, -1

	for scanner.Scan() {
		line := scanner.Text()

		matches := coordLineRE.FindStringSubmatch(line)

		if len(matches) != 3 {
			panic("error parsing coordinate line")
		}

		x, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		x--
		y--

		points = append(points, point{x, y})

		if x > maxX {
			maxX = x
		}

		if minX == -1 || x < minX {
			minX = x
		}

		if y > maxY {
			maxY = y
		}

		if minY == -1 || y < minY {
			minY = y
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	goodPoints := []point{}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			distanceToAllPointsSum := 0

			for _, pt := range points {
				distanceToAllPointsSum += manhattanDistance(point{x, y}, pt)
			}

			if distanceToAllPointsSum < 10000 {
				goodPoints = append(goodPoints, point{x, y})
			}
		}
	}

	fmt.Println(len(goodPoints))
}

func manhattanDistance(a, b point) int {
	firstSummand := a.x - b.x
	if firstSummand < 0 {
		firstSummand *= -1
	}
	secondSummand := a.y - b.y
	if secondSummand < 0 {
		secondSummand *= -1
	}

	return firstSummand + secondSummand
}
