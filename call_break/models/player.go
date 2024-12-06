package models

type Player struct {
    Name       string
    Hand       []Card
    TricksBid  int
    TricksWon  int
    Score      int
}

func NewPlayer(name string) *Player {
    return &Player{
        Name:      name,
        Hand:      make([]Card, 0),
        TricksBid: 0,
        TricksWon: 0,
        Score:     0,
    }
}

func (p *Player) AddCards(cards []Card) {
    p.Hand = append(p.Hand, cards...)
}

func (p *Player) RemoveCard(index int) Card {
    if index < 0 || index >= len(p.Hand) {
        panic("Invalid card index")
    }
    card := p.Hand[index]
    p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
    return card
} 