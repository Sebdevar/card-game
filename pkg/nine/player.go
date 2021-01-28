package nine

import (
	"errors"
	"math/rand"
)

const (
	MaxNumberOfCardsInHand = 13
)

type handType = []Card

type Play struct {
	PlayerId uint32
	Card     Card
}

type Player struct {
	id           uint32
	hand         handType
	roundPoints  uint
	score        int
	lobbyChannel chan<- Play
}

func NewPlayer() *Player {
	return &Player{
		id: rand.Uint32(),
	}
}

// getIndexOfCard returns the index of searchedCard in the player's hand.
// -1 is returned if the card isn't found.
func (player *Player) getIndexOfCard(searchedCard Card) int {
	for index, card := range player.hand {
		if card == searchedCard {
			return index
		}
	}
	return -1
}

func (player *Player) removeCardAtIndex(cardIndex int) (err error) {
	if cardIndex < 0 || cardIndex > len(player.hand) {
		err = errors.New("could not remove card: index out of range")
		return
	}
	player.hand = append(player.hand[:cardIndex], player.hand[cardIndex+1:]...)
	return
}

func (player *Player) SetLobbyChannel(channel chan Play) {
	player.lobbyChannel = channel
}

func (player *Player) GiveNewHand(hand handType) (err error) {
	if len(hand) > MaxNumberOfCardsInHand {
		err = errors.New("cannot give new hand: too many cards")
		return
	}
	player.hand = hand
	return
}

func (player *Player) PlayCard(card Card) (err error) {
	cardIndex := player.getIndexOfCard(card)
	if cardIndex == -1 {
		err = errors.New("cannot play card: not found in player's hand")
		return
	}

	err = player.removeCardAtIndex(cardIndex)
	if err != nil {
		return
	}

	player.lobbyChannel <- Play{
		PlayerId: player.id,
		Card:     card,
	}
	return
}

func (player *Player) GetId() uint32 {
	return player.id
}
