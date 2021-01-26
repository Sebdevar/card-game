package entities

import (
	"math/rand"
	"time"
)

const (
	amountOfCardsInRegularDeckWithoutJokers = 52
	amountOfCardsPerSuitInRegularDeck       = 13
)

type Deck struct {
	cards []Card
}

func NewCustomDeck(cards []Card) *Deck {
	cardsCopy := make([]Card, len(cards))
	copy(cardsCopy, cards)
	return &Deck{
		cards: cardsCopy,
	}
}

func NewRegularDeckWithoutJokers() *Deck {
	cards := make([]Card, amountOfCardsInRegularDeckWithoutJokers)
	for i, suit := range GetSuitIterable() {
		for j := 0; j < amountOfCardsPerSuitInRegularDeck; j++ {
			cards[amountOfCardsPerSuitInRegularDeck*i+j] = Card{Value: j + 1, Suit: suit}
		}
	}
	return &Deck{
		cards: cards,
	}
}

func (deck *Deck) GetCards() []Card {
	cardsCopy := make([]Card, len(deck.cards))
	copy(cardsCopy, deck.cards)
	return cardsCopy
}

func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}
