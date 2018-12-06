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
	var leftMost, rightMost, topMost, bottomMost point

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
			rightMost = point{x, y}
		}

		if minX == -1 || x < minX {
			minX = x
			leftMost = point{x, y}
		}

		if y > maxY {
			maxY = y
			bottomMost = point{x, y}
		}

		if minY == -1 || y < minY {
			minY = y
			topMost = point{x, y}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	numClosestPoints := make(map[point]int)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			currentPoint := point{x, y}

			minManhattanDistance := -1
			var closestPoint point
			manhattanDistancesSeenCount := make(map[int]int)

			for _, point := range points {
				distance := manhattanDistance(currentPoint, point)
				if minManhattanDistance == -1 || distance < minManhattanDistance {
					minManhattanDistance = distance
					closestPoint = point
				}

				manhattanDistancesSeenCount[distance]++
			}

			if minManhattanDistance == -1 {
				panic("minManhattanDistance seen was -1, the sentinel value")
			}

			if manhattanDistancesSeenCount[minManhattanDistance] == 1 {
				numClosestPoints[closestPoint]++
			}
		}
	}

	boundingRectTopLeft := point{x: leftMost.x, y: topMost.y}
	boundingRectTopRight := point{x: rightMost.x, y: topMost.y}
	boundingRectBottomLeft := point{x: leftMost.x, y: bottomMost.y}
	boundingRectBottomRight := point{x: rightMost.x, y: bottomMost.y}

	boundingPoints := []point{}
	for p := boundingRectTopLeft; p.x <= boundingRectTopRight.x; p.x++ {
		boundingPoints = append(boundingPoints, p)
	}
	for p := boundingRectTopLeft; p.y <= boundingRectBottomLeft.y; p.y++ {
		boundingPoints = append(boundingPoints, p)
	}
	for p := boundingRectBottomLeft; p.x <= boundingRectBottomRight.x; p.x++ {
		boundingPoints = append(boundingPoints, p)
	}
	for p := boundingRectTopRight; p.y <= boundingRectBottomRight.y; p.y++ {
		boundingPoints = append(boundingPoints, p)
	}

	infiniteAreaPoints := infiniteAreaPoints(points, boundingPoints)

	maxArea := -1
MaxAreaCalculation:
	for point, area := range numClosestPoints {
		for _, infiniteAreaPoint := range infiniteAreaPoints {
			if point == infiniteAreaPoint {
				continue MaxAreaCalculation
			}
		}

		if area > maxArea {
			maxArea = area
		}
	}

	fmt.Println(maxArea)
}

// Sadly, I could not do this on my own. I needed this reddit comment to get me going:
// https://www.reddit.com/r/adventofcode/comments/a3nfvy/2018_day_6_criteria_for_excluding_infinite_area/eb837ka/
func infiniteAreaPoints(points, boundingPoints []point) []point {
	infiniteAreaPoints := []point{}

	for _, boundingPoint := range boundingPoints {
		minDistance := -1
		distanceSightings := make(map[int]int)
		var minDistancePoint point
		for _, point := range points {
			distance := manhattanDistance(point, boundingPoint)
			distanceSightings[distance]++

			if minDistance == -1 || distance < minDistance {
				minDistance = distance
				minDistancePoint = point
			}
		}

		if minDistance >= 0 && distanceSightings[minDistance] == 1 {
			infiniteAreaPoints = append(infiniteAreaPoints, minDistancePoint)
		}
	}

	return infiniteAreaPoints
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
