package nine

import (
	"github.com/card-game/pkg/nine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PlayCard(test *testing.T) {
	test.Run("Should pass play through lobby channel", func(subTest *testing.T) {
		deck := nine.NewDeck()
		channel := make(chan nine.Play)
		player := nine.NewPlayer()
		player.SetLobbyChannel(channel)

		hand, err := deck.TakeMultipleCards(nine.MaxNumberOfCardsInHand)
		assert.NoError(subTest, err)

		err = player.GiveNewHand(hand)
		assert.NoError(subTest, err)

		expectedCard := hand[0]
		go func() {
			err = player.PlayCard(expectedCard)
			assert.NoError(subTest, err)
		}()

		expectedPlay := nine.Play{
			PlayerId: player.GetId(),
			Card:     expectedCard,
		}
		receivedPlay := <-channel
		assert.Equal(subTest, expectedPlay, receivedPlay)
	})
}
