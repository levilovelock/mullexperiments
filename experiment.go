package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	total   int = 60
	desired int = 4
	draw    int = 7
	mull    int = 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type card bool

func main() {
	deck := initializeDeck()
	experimentPureDraw(100000, deck)
	experimentMullDownTo(100000, mull, deck)
}

func initializeDeck() []card {
	var deck []card

	for i := 0; i < total; i++ {
		deck = append(deck, *new(card))
	}

	for i := 0; i < desired; i++ {
		deck[i] = true
	}

	return deck
}

func experimentMullDownTo(num int, downTo int, deck []card) {
	success := 0
	fail := 0

	for i := 0; i < num; i++ {
		for d := 7; d >= downTo; d-- {
			s := drawCardsAndReturnNumSuccesses(d, &deck)
			if s > 0 {
				success++
				break
			} else if d == downTo {
				fail++
			}
		}
	}

	fmt.Printf("Success Rate: %f\nFailure Rate: %f\n", (float32(success) / float32(num)), (float32(fail) / float32(num)))
}

func experimentPureDraw(num int, deck []card) []int {
	var results []int

	// Initialise with zeroes
	for i := 0; i < draw; i++ {
		results = append(results, 0)
	}

	for i := 0; i < num; i++ {
		s := drawCardsAndReturnNumSuccesses(draw, &deck)
		results[s]++
	}

	return results
}

func drawCardsAndReturnNumSuccesses(x int, deck *[]card) int {
	d := shuffleDeck(deck)
	numSuccesses := 0
	for i := 0; i < x; i++ {
		if d[i] {
			numSuccesses++
		}
	}
	return numSuccesses
}

func shuffleDeck(deck *[]card) []card {
	randomDeck := make([]card, len(*deck))
	perm := rand.Perm(len(*deck))
	for i, v := range perm {
		randomDeck[v] = (*deck)[i]
	}
	return randomDeck
}
