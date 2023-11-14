package godeck

import (
	"fmt"
	"math/rand"

	"github.com/emirpasic/gods/stacks/arraystack"
)

type PlayerHand struct {
	Cards []Card
}

func (h *PlayerHand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
}

func (h *PlayerHand) RemoveCard(i int) Card {
	card := h.Cards[i]
	h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
	return card
}

func (h *PlayerHand) AddCards(cards []Card) {
	h.Cards = append(h.Cards, cards...)
}

type Card struct {
	Suit  string
	Value string
}

func (c Card) String() string {
	if c.Suit == "Joker" {
		return c.Suit
	}
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}

type Deck struct {
	Cards     []Card
	NumDecks  int
	Jokers    bool
	NumJokers int
}

func (d *Deck) populate() {
	suits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for i := 0; i < d.NumDecks; i++ {

		for _, suit := range suits {
			for _, value := range values {
				d.Cards[i] = Card{suit, value}
				i++
			}
			if d.Jokers {
				for j := 0; j < d.NumJokers; j++ {
					d.Cards[i] = Card{"Joker", "Joker"}
					i++
				}
			}
		}
	}
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Deal() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func NewDeck(NumDecks int, Jokers bool, NumJokers int) *Deck {
	d := new(Deck)
	d.NumDecks = NumDecks
	numCards := 52 * NumDecks
	if Jokers {
		numCards += NumJokers * NumDecks
	}
	d.Cards = make([]Card, numCards)
	d.populate()
	d.shuffle()
	return d
}

type DiscardPile struct {
	Cards *arraystack.Stack
}

func (d *DiscardPile) AddCard(c Card) {
	d.Cards.Push(c)
}

func (d *DiscardPile) RemoveCard() Card {
	card, _ := d.Cards.Pop()
	return card.(Card)
}

// if nil is returned, the discard pile is empty
func (d *DiscardPile) Peek() Card {
	card, _ := d.Cards.Peek()
	return card.(Card)
}

func (d *DiscardPile) Size() int {
	return d.Cards.Size()
}

func (d *DiscardPile) TakeAll() []Card {
	cards := make([]Card, d.Cards.Size())
	for i := 0; i < len(cards); i++ {
		cards[i] = d.RemoveCard()
	}
	return cards
}

func NewDiscardPile() *DiscardPile {
	d := new(DiscardPile)
	d.Cards = arraystack.New()
	return d
}
