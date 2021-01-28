package nine

type Suit string

const (
	Diamonds Suit = "Diamonds"
	Clubs         = "Clubs"
	Hearts        = "Hearts"
	Spades        = "Spades"
)

func GetSuitIterable() []Suit {
	return []Suit{Clubs, Diamonds, Hearts, Spades}
}

type Card struct {
	Value int
	Suit  Suit
}
