package main

import (
	"container/list"
	"fmt"
)

const numPlayers = 432
const lastMarble = 7101900

type circle struct {
	marbles         *list.List
	currentMarbleEl *list.Element
}

func main() {
	scores := [numPlayers]int{}
	currentMarble := 1
	list := list.New()
	currentMarbleEl := list.PushFront(0)
	circle := circle{
		marbles:         list,
		currentMarbleEl: currentMarbleEl,
	}

	for turn := 0; currentMarble <= lastMarble; turn, currentMarble = (turn+1)%numPlayers, currentMarble+1 {
		if currentMarble%23 == 0 {
			scores[turn] += currentMarble
			var score int
			score = removeFromCircle(&circle)
			scores[turn] += score
		} else {
			placeInCircle(&circle, currentMarble)
		}
	}

	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}

func placeInCircle(circle *circle, currentMarble int) {
	for i := 0; i < 1; i++ {
		circle.currentMarbleEl = circle.currentMarbleEl.Next()
		if circle.currentMarbleEl == nil {
			circle.currentMarbleEl = circle.marbles.Front()
		}
	}
	circle.marbles.InsertAfter(currentMarble, circle.currentMarbleEl)
	circle.currentMarbleEl = circle.currentMarbleEl.Next()
}

func removeFromCircle(circle *circle) int {
	for i := 0; i < 7; i++ {
		circle.currentMarbleEl = circle.currentMarbleEl.Prev()
		if circle.currentMarbleEl == nil {
			circle.currentMarbleEl = circle.marbles.Back()
		}
	}
	newCurrent := circle.currentMarbleEl.Next()
	if circle.currentMarbleEl == nil {
		circle.currentMarbleEl = circle.marbles.Front()
	}

	score := circle.currentMarbleEl.Value.(int)

	circle.marbles.Remove(circle.currentMarbleEl)
	circle.currentMarbleEl = newCurrent

	return score
}
