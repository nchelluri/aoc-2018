package main

import (
	"container/list"
	"flag"
	"fmt"
)

// MarbleCircle is an AoC 2018 day 9 marble circle.
type MarbleCircle struct {
	CurrentMarble int

	currentMarbleEl *list.Element
	lastMarble      int
	marbles         *list.List
	scores          []int
	turn            int
	verbose         bool
}

func main() {
	numPlayers := flag.Int("numPlayers", 432, "number of players")
	lastMarble := flag.Int("lastMarble", 7101900, "last marble")
	verbose := flag.Bool("verbose", false, "sets verbose mode on")
	flag.Parse()
	m := NewMarbleCircle(*numPlayers, *lastMarble, *verbose)

	for m.Process() {
	}

	fmt.Println(m.HighScore())
}

// NewMarbleCircle creates a new MarbleCircle.
func NewMarbleCircle(numPlayers, lastMarble int, verbose bool) *MarbleCircle {
	list := list.New()
	currentMarbleEl := list.PushFront(0)
	scores := []int{}
	for i := 0; i < numPlayers; i++ {
		scores = append(scores, 0)
	}
	return &MarbleCircle{
		CurrentMarble: 1,

		currentMarbleEl: currentMarbleEl,
		lastMarble:      lastMarble,
		marbles:         list,
		scores:          scores,
		turn:            0,
		verbose:         verbose,
	}
}

// Process processes the turn for the next player. Returns whether there is more processing required or not.
func (m *MarbleCircle) Process() bool {
	if m.CurrentMarble%23 == 0 {
		m.scores[m.turn] += m.CurrentMarble + m.remove()
	} else {
		m.place()
	}

	if m.verbose {
		m.print()
	}

	done := false
	if m.CurrentMarble == m.lastMarble {
		done = true
	}
	m.CurrentMarble++
	m.turn = (m.turn + 1) % len(m.scores)

	return !done
}

// HighScore returns the highest score for the MarbleCircle. Only makes sense once processing the game is complete.
func (m *MarbleCircle) HighScore() int {
	maxScore := 0
	for _, score := range m.scores {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

// Print prints out a marble circle, as found in the Advent of Code example.
func (m *MarbleCircle) print() {
	fmt.Printf("[%d] ", m.turn+1)
	for n := m.marbles.Front(); n != nil; n = n.Next() {
		if n.Value.(int) == m.CurrentMarble {
			fmt.Printf("(%d) ", n.Value)
			continue
		}
		fmt.Printf(" %d  ", n.Value)
	}
	fmt.Println()
}

// place places the next marble in the marble circle.
func (m *MarbleCircle) place() *list.Element {
	m.marbles.InsertAfter(m.CurrentMarble, m.getNext())
	return m.getNext()
}

// remove removes the 7th counter-clockwise marble to the current marble and returns its score.
func (m *MarbleCircle) remove() int {
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
