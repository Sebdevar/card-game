package entities

import (
	"github.com/card-game/pkg/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewDeck(test *testing.T) {
	expectedCards := []entities.Card{{Value: 2, Suit: entities.Clubs}, {Value: 5, Suit: entities.Spades}, {Value: 17, Suit: entities.Diamonds}}
	actualCards := entities.NewCustomDeck(expectedCards).GetCards()
	test.Run("Should create deck with passed cards", func(subTest *testing.T) {
		assert.Equal(test, expectedCards, actualCards)
	})

}

func Test_Shuffle(test *testing.T) {
	deck := entities.NewRegularDeckWithoutJokers()
	cards := deck.GetCards()
	deck.Shuffle()

	test.Run("Should change cards order", func(subTest *testing.T) {
		assert.NotEqual(subTest, cards, deck.GetCards())
	})
	test.Run("Should not change card amount", func(subTest *testing.T) {
		assert.Equal(subTest, len(cards), len(deck.GetCards()))
	})
}

func Test_GetCards(test *testing.T) {
	expectedCards := []entities.Card{{Value: 2, Suit: entities.Clubs}, {Value: 5, Suit: entities.Spades}, {Value: 17, Suit: entities.Diamonds}}
	deck := entities.NewCustomDeck(expectedCards)

	test.Run("Should return the cards from the deck", func(subTest *testing.T) {
		assert.Equal(subTest, expectedCards, deck.GetCards())
	})

	test.Run("Should return an independent copy of the cards", func(subTest *testing.T) {
		cards := deck.GetCards()
		cards[0].Value += 1
		assert.NotEqual(subTest, cards, deck.GetCards())
	})
}
