package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x, y, velX, velY int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	points := []point{}

	// position=< 9,  1> velocity=< 0,  2>
	pointRE := regexp.MustCompile(`^position=<\s*(-?\d+),\s*(-?\d+)>\s+velocity=<\s*(-?\d+),\s*(-?\d+)>$`)

	for scanner.Scan() {
		matches := pointRE.FindStringSubmatch(scanner.Text())

		if len(matches) != 5 {
			panic("error parsing point line")
		}

		x, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		velX, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}

		velY, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}

		points = append(points, point{x: x, y: y, velX: velX, velY: velY})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var xRange, yRange *int
	for i := 0; ; i++ {
		newXRange, newYRange := update(&points, i)

		if xRange == nil && yRange == nil {
			xRange = &newXRange
			yRange = &newYRange
			continue
		}

		if *xRange < newXRange && *yRange < newYRange {
			// grid is getting bigger, not smaller: we are done
			break
		}
	}
}

func print(points []point, minX, maxX, minY, maxY, currentTime int) {
	fmt.Println("Time:", currentTime+1)
	grid := [][]bool{}
	for i := maxY; i >= minY; i-- {
		row := []bool{}
		for j := maxX; j >= minX; j-- {
			row = append(row, false)
		}
		grid = append(grid, row)
	}

	for _, point := range points {
		grid[point.y-minY][point.x-minX] = true
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("\n\n\n")

}

func update(points *[]point, currentTime int) (int, int) {
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32

	for i, point := range *points {
		point.x += point.velX
		point.y += point.velY
		(*points)[i] = point

		if point.x < minX {
			minX = point.x
		}

		if point.y < minY {
			minY = point.y
		}

		if point.x > maxX {
			maxX = point.x
		}

		if point.y > maxY {
			maxY = point.y
		}
	}

	if maxX-minX < 100 && maxY-minY < 100 {
		print(*points, minX, maxX, minY, maxY, currentTime)
	}

	return maxX - minX, maxY - minY
}
