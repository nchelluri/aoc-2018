package main

import (
	"container/list"
	"fmt"
)

const numPlayers = 432
const lastMarble = 7101900

// MarbleCircle is an AoC 2018 day 9 marble circle.
type MarbleCircle struct {
	CurrentMarble   int
	marbles         *list.List
	currentMarbleEl *list.Element
	lastMarble      int
}

// NewMarbleCircle creates a new MarbleCircle.
func NewMarbleCircle(lastMarble int) *MarbleCircle {
	list := list.New()
	currentMarbleEl := list.PushFront(0)
	return &MarbleCircle{
		marbles:         list,
		currentMarbleEl: currentMarbleEl,
		CurrentMarble:   1,
		lastMarble:      lastMarble,
	}
}

func main() {
	scores := [numPlayers]int{}
	m := NewMarbleCircle(lastMarble)

	for turn := 0; m.CurrentMarble <= m.lastMarble; turn, m.CurrentMarble = (turn+1)%numPlayers, m.CurrentMarble+1 {
		if m.CurrentMarble%23 == 0 {
			scores[turn] += m.CurrentMarble + m.Remove()
		} else {
			m.Place()
		}

		// m.Print(turn)
	}

	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}

// Print prints out a marble circle, as found in the Advent of Code example.
func (m *MarbleCircle) Print(turn int) {
	fmt.Printf("[%02d] ", turn+1)
	for n := m.marbles.Front(); n != nil; n = n.Next() {
		if n.Value.(int) == m.CurrentMarble {
			fmt.Printf("(%d) ", n.Value)
			continue
		}
		fmt.Printf(" %d  ", n.Value)
	}
	fmt.Println()
}

// Place places the next marble in the marble circle.
func (m *MarbleCircle) Place() *list.Element {
	m.marbles.InsertAfter(m.CurrentMarble, m.getNext())
	return m.getNext()
}

// Remove removes the 7th counter-clockwise marble to the current marble and returns its score.
func (m *MarbleCircle) Remove() int {
	for i := 0; i < 7; i++ {
		m.getPrev()
	}
	toBeRemoved := m.currentMarbleEl
	score := toBeRemoved.Value.(int)
	m.getNext()
	m.marbles.Remove(toBeRemoved)
	return score
}

func (m *MarbleCircle) getNext() *list.Element {
	m.currentMarbleEl = m.currentMarbleEl.Next()
	if m.currentMarbleEl == nil {
		m.currentMarbleEl = m.marbles.Front()
	}
	return m.currentMarbleEl
}

func (m *MarbleCircle) getPrev() *list.Element {
	m.currentMarbleEl = m.currentMarbleEl.Prev()
	if m.currentMarbleEl == nil {
		m.currentMarbleEl = m.marbles.Back()
	}
	return m.currentMarbleEl
}
