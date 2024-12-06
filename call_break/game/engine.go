package game

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "bitbucket.org/junglee_games/call_break/models"
)

type GameEngine struct {
    Players       []*models.Player
    CurrentTrick  []models.Card
    LeadSuit      models.Suit
    TrumpSuit     models.Suit
    CurrentPlayer int
    Round         int
}

func NewGameEngine(playerNames []string) *GameEngine {
    if len(playerNames) != 4 {
        panic("Call Break requires exactly 4 players")
    }

    players := make([]*models.Player, len(playerNames))
    for i, name := range playerNames {
        players[i] = models.NewPlayer(name)
    }

    return &GameEngine{
        Players:       players,
        CurrentTrick:  make([]models.Card, 0),
        TrumpSuit:    models.Spades,
        CurrentPlayer: 0,
        Round:        1,
    }
}

func (g *GameEngine) StartGame() {
    fmt.Println("\nStarting new game of Call Break!")
    
    for g.Round <= 5 { // Play 5 rounds
        fmt.Printf("\n=== Round %d ===\n", g.Round)
        g.PlayRound()
        g.Round++
    }
    
    g.ShowFinalResults()
}

func (g *GameEngine) PlayRound() {
    // Deal cards
    deck := models.NewDeck()
    deck.Shuffle()
    
    for _, player := range g.Players {
        player.Hand = deck.Deal(13)
        player.TricksWon = 0
    }
    
    // Bidding phase
    g.BiddingPhase()
    
    // Play 13 tricks
    for trick := 0; trick < 13; trick++ {
        g.PlayTrick()
    }
    
    // Score the round
    g.ScoreRound()
}

func (g *GameEngine) BiddingPhase() {
    reader := bufio.NewReader(os.Stdin)
    
    for _, player := range g.Players {
        fmt.Printf("\n%s's hand: ", player.Name)
        for _, card := range player.Hand {
            fmt.Printf("%s ", card)
        }
        
        for {
            fmt.Printf("\n%s, enter your bid (1-13): ", player.Name)
            input, _ := reader.ReadString('\n')
            bid, err := strconv.Atoi(strings.TrimSpace(input))
            
            if err == nil && bid >= 1 && bid <= 13 {
                player.TricksBid = bid
                break
            }
            fmt.Println("Invalid bid. Please enter a number between 1 and 13.")
        }
    }
}

func (g *GameEngine) PlayTrick() {
    g.CurrentTrick = make([]models.Card, 0)
    fmt.Printf("\n\n=== New Trick ===\n")
    
    for i := 0; i < 4; i++ {
        player := g.Players[g.CurrentPlayer]
        
        fmt.Printf("\nCurrent trick: ")
        for _, card := range g.CurrentTrick {
            fmt.Printf("%s ", card)
        }
        
        fmt.Printf("\n%s's hand: ", player.Name)
        for j, card := range player.Hand {
            fmt.Printf("%d:%s ", j+1, card)
        }
        
        cardIndex := g.GetValidCardChoice(player)
        playedCard := player.RemoveCard(cardIndex)
        
        if len(g.CurrentTrick) == 0 {
            g.LeadSuit = playedCard.Suit
        }
        
        g.CurrentTrick = append(g.CurrentTrick, playedCard)
        g.CurrentPlayer = (g.CurrentPlayer + 1) % 4
    }
    
    // Determine trick winner
    winnerIndex := g.DetermineTrickWinner()
    g.Players[winnerIndex].TricksWon++
    g.CurrentPlayer = winnerIndex
    
    fmt.Printf("\nTrick won by %s\n", g.Players[winnerIndex].Name)
}

func (g *GameEngine) GetValidCardChoice(player *models.Player) int {
    reader := bufio.NewReader(os.Stdin)
    
    for {
        fmt.Printf("\n%s, choose a card (1-%d): ", player.Name, len(player.Hand))
        input, _ := reader.ReadString('\n')
        index, err := strconv.Atoi(strings.TrimSpace(input))
        index-- // Convert to 0-based index
        
        if err != nil || index < 0 || index >= len(player.Hand) {
            fmt.Println("Invalid card choice.")
            continue
        }
        
        // Validate following suit
        if len(g.CurrentTrick) > 0 && player.Hand[index].Suit != g.LeadSuit {
            canFollowSuit := false
            for _, card := range player.Hand {
                if card.Suit == g.LeadSuit {
                    canFollowSuit = true
                    break
                }
            }
            
            if canFollowSuit {
                fmt.Println("You must follow the lead suit if possible.")
                continue
            }
        }
        
        return index
    }
}

func (g *GameEngine) DetermineTrickWinner() int {
    winningCard := g.CurrentTrick[0]
    winnerIndex := 0
    
    for i := 1; i < len(g.CurrentTrick); i++ {
        card := g.CurrentTrick[i]
        
        if card.Suit == g.TrumpSuit && winningCard.Suit != g.TrumpSuit {
            winningCard = card
            winnerIndex = i
        } else if card.Suit == winningCard.Suit && card.Value > winningCard.Value {
            winningCard = card
            winnerIndex = i
        }
    }
    
    return (g.CurrentPlayer + winnerIndex) % 4
}

func (g *GameEngine) ScoreRound() {
    fmt.Println("\nRound Results:")
    for _, player := range g.Players {
        roundScore := 0
        if player.TricksWon >= player.TricksBid {
            roundScore = player.TricksBid
        } else {
            roundScore = -player.TricksBid
        }
        player.Score += roundScore
        
        fmt.Printf("%s: Bid %d, Won %d, Round Score %d, Total Score %d\n",
            player.Name, player.TricksBid, player.TricksWon, roundScore, player.Score)
    }
}

func (g *GameEngine) ShowFinalResults() {
    fmt.Println("\n=== Final Scores ===")
    for _, player := range g.Players {
        fmt.Printf("%s: %d points\n", player.Name, player.Score)
    }
} 