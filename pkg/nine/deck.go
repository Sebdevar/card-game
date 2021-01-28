package nine

import (
	"errors"
	"math/rand"
	"time"
)

const (
	MaxAmountOfCardsInDeck = 52
	AmountOfCardsPerSuit   = 13
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	cards := make([]Card, MaxAmountOfCardsInDeck)
	for i, suit := range GetSuitIterable() {
		for j := 0; j < AmountOfCardsPerSuit; j++ {
			cards[AmountOfCardsPerSuit*i+j] = Card{Value: j + 1, Suit: suit}
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

func (deck *Deck) TakeACard() (card Card, err error) {
	if len(deck.cards) == 0 {
		err = errors.New("could not remove card: deck is empty")
		return
	}
	card, deck.cards = deck.cards[0], deck.cards[1:]
	return
}

func (deck *Deck) TakeMultipleCards(amount int) (cards []Card, err error) {
	if len(deck.cards) < amount {
		err = errors.New("could not remove cards: too many cards requested")
		return
	}
	cards, deck.cards = deck.cards[0:amount], deck.cards[amount:]
	return
}
