package models

import "fmt"

type Suit int

const (
    Hearts Suit = iota
    Diamonds
    Clubs
    Spades
)

func (s Suit) String() string {
    switch s {
    case Hearts:
        return "♥"
    case Diamonds:
        return "♦"
    case Clubs:
        return "♣"
    case Spades:
        return "♠"
    default:
        return "?"
    }
}

type Card struct {
    Suit  Suit
    Value int
}

func (c Card) String() string {
    value := ""
    switch c.Value {
    case 11:
        value = "J"
    case 12:
        value = "Q"
    case 13:
        value = "K"
    case 14:
        value = "A"
    default:
        value = fmt.Sprintf("%d", c.Value)
    }
    return fmt.Sprintf("%s%s", c.Suit, value)
} 