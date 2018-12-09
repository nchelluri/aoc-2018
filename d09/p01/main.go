package main

import (
	"fmt"
)

func main() {
	const numPlayers = 9
	const lastMarble = 25

	scores := [numPlayers]int{}
	currentMarble := 1
	currentMarbleIndex := 0
	circle := []int{0}

	for turn := 0; currentMarble <= lastMarble; turn, currentMarble = (turn+1)%numPlayers, currentMarble+1 {
		if currentMarble%23 == 0 {
			scores[turn] += currentMarble
			var score int
			currentMarbleIndex, score, circle = removeFromCircle(circle, currentMarbleIndex)
			scores[turn] += score
		} else {
			currentMarbleIndex, circle = placeInCircle(circle, currentMarbleIndex, currentMarble)
		}

		fmt.Printf("%d => %+v\n\n", currentMarble, circle)
	}

	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}

func placeInCircle(circle []int, currentMarbleIndex, currentMarble int) (int, []int) {
	newMarbleIndex := (currentMarbleIndex + 1) % len(circle)
	fmt.Println("Placing", currentMarble, "at", newMarbleIndex)

	newCircle := []int{}
	newCircle = append(newCircle, circle[:newMarbleIndex+1]...)
	newCircle = append(newCircle, currentMarble)
	newCircle = append(newCircle, circle[newMarbleIndex+1:]...)

	return newMarbleIndex + 1, newCircle
}

func removeFromCircle(circle []int, currentMarbleIndex int) (int, int, []int) {
	newMarbleIndex := (currentMarbleIndex - 7) % len(circle)

	if newMarbleIndex < 0 {
		newMarbleIndex = len(circle) + newMarbleIndex
	}

	score := circle[newMarbleIndex]
	circle = append(circle[:newMarbleIndex], circle[newMarbleIndex+1:]...)

	marbleIndex := (newMarbleIndex - 1) % len(circle)
	if marbleIndex < newMarbleIndex%len(circle) {
		marbleIndex = newMarbleIndex
	}

	return marbleIndex, score, circle
}
