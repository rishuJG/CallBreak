package models

import (
    "math/rand"
    "time"
)

type Deck struct {
    Cards []Card
}

func NewDeck() *Deck {
    deck := &Deck{
        Cards: make([]Card, 0, 52),
    }
    
    for suit := Hearts; suit <= Spades; suit++ {
        for value := 2; value <= 14; value++ {
            deck.Cards = append(deck.Cards, Card{
                Suit:  suit,
                Value: value,
            })
        }
    }
    
    return deck
}

func (d *Deck) Shuffle() {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := len(d.Cards) - 1; i > 0; i-- {
        j := r.Intn(i + 1)
        d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
    }
}

func (d *Deck) Deal(numCards int) []Card {
    if len(d.Cards) < numCards {
        return nil
    }
    dealt := d.Cards[:numCards]
    d.Cards = d.Cards[numCards:]
    return dealt
} 