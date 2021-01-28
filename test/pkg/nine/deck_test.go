package nine

import (
	"github.com/card-game/pkg/nine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Shuffle(test *testing.T) {
	deck := nine.NewDeck()
	cards := deck.GetCards()
	deck.Shuffle()

	test.Run("Should change cards order", func(subTest *testing.T) {
		assert.NotEqual(subTest, cards, deck.GetCards())
	})
	test.Run("Should not change card amount", func(subTest *testing.T) {
		assert.Equal(subTest, len(cards), len(deck.GetCards()))
	})

	test.Run("Should still contain the same cards", func(subTest *testing.T) {
		for _, card := range cards {
			assert.Contains(subTest, deck.GetCards(), card)
		}
	})
}

func Test_GetCards(test *testing.T) {
	deck := nine.NewDeck()

	test.Run("Should return an independent copy of the cards", func(subTest *testing.T) {
		cards := deck.GetCards()
		cards[0].Value += 1
		assert.NotEqual(subTest, cards, deck.GetCards())
	})

	test.Run("Should return empty array when no cards", func(subTest *testing.T) {
		emptyDeck := nine.NewDeck()
		_, _ = emptyDeck.TakeMultipleCards(nine.MaxAmountOfCardsInDeck)
		assert.Empty(subTest, emptyDeck.GetCards())
	})
}

func Test_TakeACard(test *testing.T) {
	deck := nine.NewDeck()
	expectedCards := deck.GetCards()

	test.Run("Should return the first card of the deck", func(subTest *testing.T) {
		removedCard, err := deck.TakeACard()
		if assert.NoError(subTest, err) {
			assert.Equal(subTest, expectedCards[0], removedCard)
		}
	})

	test.Run("Should remove returned card from deck", func(subTest *testing.T) {
		removedCard, err := deck.TakeACard()
		if assert.NoError(subTest, err) {
			assert.NotContains(subTest, deck.GetCards(), removedCard)
		}
	})

	test.Run("Should be able to take the last card", func(subTest *testing.T) {
		singleCardDeck := nine.NewDeck()
		_, _ = singleCardDeck.TakeMultipleCards(nine.MaxAmountOfCardsInDeck - 1)
		_, err := singleCardDeck.TakeACard()
		if assert.NoError(subTest, err) {
			assert.Empty(subTest, singleCardDeck.GetCards())
		}
	})

	test.Run("Should return an error when the deck is empty", func(subTest *testing.T) {
		emptyDeck := nine.NewDeck()
		_, _ = emptyDeck.TakeMultipleCards(nine.MaxAmountOfCardsInDeck)
		_, err := emptyDeck.TakeACard()
		assert.Error(subTest, err)
	})
}

func Test_TakeMultipleCards(test *testing.T) {
	deck := nine.NewDeck()

	test.Run("Should return the amount of cards taken", func(subTest *testing.T) {
		const amountOfCardsToTake = 17
		cards, err := deck.TakeMultipleCards(amountOfCardsToTake)
		if assert.NoError(subTest, err) {
			assert.Equal(subTest, amountOfCardsToTake, len(cards))
		}
	})

	test.Run("Should remove cards taken from the deck", func(subTest *testing.T) {
		cards, err := deck.TakeMultipleCards(3)
		if assert.NoError(subTest, err) {
			for _, card := range cards {
				assert.NotContains(subTest, deck.GetCards(), card)
			}
		}
	})

	test.Run("Should return an error if less cards in deck than requested amount", func(subTest *testing.T) {
		_, err := deck.TakeMultipleCards(55)
		assert.Error(subTest, err)
	})
}
